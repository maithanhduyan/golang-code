# GO Echo


~~~~
    go mod example.com/go-echo
~~~~

### Add echo package
~~~~
    go get github.com/labstack/echo/v4
~~~~

#### Source code of main.go file
~~~~
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
~~~~

#### Run GO app

~~~~
    go run main.go
~~~~



#### Build app on Docker

- Create a Dockerfile

- Build GO image
~~~~
    docker build --tag go-echo .
~~~~

#### Start the container and expose port 8080 to port 8080 on the host.
~~~~
    docker run --publish 8080:8080 go-echo
~~~~


