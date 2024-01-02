package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/cagarcia2011/htmx-go-templ-tw-starter/templates"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	listenAddr := flag.String("listenAddr", ":3000", "the server address")
	flag.Parse()
	host := fmt.Sprintf("localhost%s", *listenAddr)
	startLog := fmt.Sprintf("Server starting at %s", host)

	e.GET("/", func(c echo.Context) error {
		component := templates.Home("Welcome to this HTMX, GO, Templ & Tailwind Starter")
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/test", func(c echo.Context) error {
		component := templates.Test()
		fmt.Println("Test received.")
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.Static("/static", "static")

	fmt.Println(startLog)
	e.Logger.Fatal(e.Start(host))
}
