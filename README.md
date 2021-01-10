# The Fibonacci API 

This API exposes three endpoints that can be called via HTTP requests:
* /current - returns the current number in the sequence
* /next - returns the next number in the sequence
* /previous - returns the previous number in the sequence

## How it Works
`go get -u github.com/smith4040/fibAPI`

Run the fibAPI in the terminal or build and run with Docker.

Using Postman or your favorite API testing tool, send "GET" request to "http://localhost:8080/current", "http://localhost:8080/next" or "http://localhost:8080/previous".

## About the API  - Why choose...
`gin-gonic/gin` web framework was chosen due to its great performance, able to handle high throughput and minimalistic design. I went with the `gin.Default` implementation which handles the logging and recovery from `panic`.

`stretchr/testify` testing package was chosen to aid in testing the handlers with the `assert` function.

`t-pwk/go-fibonacci` package contains a benchmarked and tested Fibonacci function that is ideal for this project.

`fvbock/endles` router handles unexpected crashes of the router and automatically restarts.

A `Dockerfile` has been add to quickly run this application in a Docker container.