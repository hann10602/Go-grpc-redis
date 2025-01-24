package main

import (
	"context"
	"go-rGPC/config"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	e := echo.New()

	var RedisDB *redis.Client = config.ExampleClient()

	err := RedisDB.Set(ctx, "name", "ngocha and dieuanh", 0).Err()

    if err != nil {
        panic(err)
    }

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}