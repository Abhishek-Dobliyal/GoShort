package routes

import (
	"goshort/database"
	"goshort/helpers"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/asaskevich/govalidator"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL           string        `json:"url"`
	CustomShort   string        `json:"custom_short"`
	Expiry        time.Duration `json:"expiry"`
	RateRemaining int           `json:"rate_rem"`
	RateLimitRest time.Duration `json:"rate_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "could not parse JSON"})
	}

	r2 := database.CreateClient(1)
	defer r2.Close()

	value, err := r2.Get(database.Ctx, c.IP()).Result()
	if err == redis.Nil {
		_ = r2.Set(database.Ctx, c.IP(), helpers.ApiQuota, 60 * time.Second).Err()
	} else {
		valueToInt, _ := strconv.Atoi(value)
		if valueToInt <= 0 {
			limit, _ := r2.TTL(database.Ctx, c.IP()).Result()
			return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
				"error": "rate limit reached",
				"rate_limit": limit,
			 })
		}
	}

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid URL"})
	}

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "invalid URL"})
	}

	body.URL = helpers.EnforceHTTP(body.URL)
	
	var id string
	if body.CustomShort == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	value, _ = r.Get(database.Ctx, id).Result()
	if value != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "url already in use"})
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	status := r.Set(database.Ctx, id, body.URL, body.Expiry * 3600 * time.Second)
	if status.Err() != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to connect"})
	}

	resp := response{
		URL: body.URL,
		CustomShort: "",
		Expiry: body.Expiry,
		RateRemaining: 10,
		RateLimitRest: 30,
	}

	r2.Decr(database.Ctx, c.IP())

	value, _ = r2.Get(database.Ctx, c.IP()).Result()
	resp.RateRemaining, _ = strconv.Atoi(value)

	ttl, _ := r2.TTL(database.Ctx, c.IP()).Result()
	resp.RateLimitRest = ttl / time.Nanosecond / time.Minute
	resp.CustomShort = helpers.Domain + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)
}
