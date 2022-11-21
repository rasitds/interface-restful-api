package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var (
	app = fiber.New()
)

func main() {
	fmt.Println("INIT")

	app.Static("/", "./public")

	app.Get("/fiber", func(c *fiber.Ctx) error {
		return c.SendString("Fiber")
	}).Name("fiberapi")

	anotherRoute()
	anotherRoute2()

	app.Listen(":3000")
}

func anotherRoute() {
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("API Handler")
		return c.Next()
	})

	app.Get("/api/info/fiberapi", func(c *fiber.Ctx) error {
		fmt.Println("info fiberapi api handler")
		data, _ := json.MarshalIndent(app.GetRoute("fiberapi"), "", " ")
		return c.Send(data)
	})
}

func anotherRoute2() {
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("%s is %s years old. %s", c.Params("name"), c.Params("age"), c.Params("gender"))

		return c.SendString(msg)
	})
}
