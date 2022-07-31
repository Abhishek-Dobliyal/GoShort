package helpers

import "strings"

const (
	DbAddress = "db:6379"
	DbPass    = ""
	AppPort   = ":3000"
	Domain    = "localhost:3000"
	ApiQuota  = 10
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}

	return url
}

func RemoveDomainErr(url string) bool {
	if url == Domain {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == Domain {
		return false
	}

	return true
}

func RemoveDomainError(string) bool {

}