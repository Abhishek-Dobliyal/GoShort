# GoShort
A web app to shorten URLs.

![](goshort.png)  
<img src="https://github.com/Abhishek-Dobliyal/GoShort/blob/main/goshort-mobile.png" width="600" height="1000">

### Installation & Requirements

- [NodeJS](https://nodejs.org/en/)
- [VueJS](https://www.npmjs.com/package/vue)
- [GoLang](https://go.dev/)
- [Fiber](https://docs.gofiber.io/)
- [Go-Redis](https://github.com/go-redis/redis)

### Usage

- Navigate to the project directory and inside the `api` folder open up the Terminal/Command Prompt and type in
```bash
go run main.go
```

- Once the `fiber` server starts running, navigate to the `frontend` folder and type in
```bash
npm run serve
```

> This will start a development server.

### To-Do

- [x] Send URL to backend for shortening and retrieving the response
- [ ] Add wait time animation
- [x] Add notification sound 
- [ ] Fix issues (if any)
