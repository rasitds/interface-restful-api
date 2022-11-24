package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	app = fiber.New()
)

func main() {
	fmt.Println("INIT")

	app.Use(cors.New())

	app.Static("/", "../sencha-web-ui/build")

	app.Get("/fiber", func(c *fiber.Ctx) error {
		return c.SendString("Fiber")
	}).Name("fiberapi")

	apiRouteHandler()
	apiThemeRoute()

	app.Listen(":1010")
}

func apiRouteHandler() {
	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("API Route Handler")
		return c.Next()
	})
}

func apiThemeRoute() {
	app.Get("/api/themes/:theme", func(c *fiber.Ctx) error {
		fmt.Println("api themes handler")

		var themeData = map[string]interface{}{
			"default": map[string]interface{}{
				"background": "black",
				"foreground": "white",
			},
			"cfi-blue": map[string]interface{}{
				"background": "#1974D2",
				"foreground": "white",
			},
			"aqua": map[string]interface{}{"background": "black", "foreground": "#33ffd0"},
			"white-orange": map[string]interface{}{
				"background": "#ff8000",
				"foreground": "white",
			},
			"light-blue": map[string]interface{}{
				"background": "black",
				"foreground": "#33bbff",
			},
			"yellow":       map[string]interface{}{"background": "black", "foreground": "yellow"},
			"pinkish":      map[string]interface{}{"background": "black", "foreground": "#DE3163"},
			"dark":         map[string]interface{}{"background": "black", "foreground": "#f2f2f2"},
			"light":        map[string]interface{}{"background": "#f2f2f2", "foreground": "black"},
			"orange":       map[string]interface{}{"background": "black", "foreground": "#EA5B0C"},
			"cyan":         map[string]interface{}{"background": "black", "foreground": "#4CBEC5"},
			"green":        map[string]interface{}{"background": "black", "foreground": "#00CC11"},
			"pink":         map[string]interface{}{"background": "black", "foreground": "#FF6666"},
			"faint-orange": map[string]interface{}{"background": "black", "foreground": "#996633"},
			"neon-blue":    map[string]interface{}{"background": "black", "foreground": "#0033FF"},
			"ultra-green":  map[string]interface{}{"background": "black", "foreground": "#0aff84"},
			"ultra-purple": map[string]interface{}{"background": "black", "foreground": "#8709f4"},
			"iron-gray":    map[string]interface{}{"background": "black", "foreground": "#52595D"},
			"bright-gray":  map[string]interface{}{"background": "black", "foreground": "#DCDCDC"},
			"bright-blue":  map[string]interface{}{"background": "black", "foreground": "#006EF0"},
		}

		themeInfo := themeData[c.Params("theme")]

		if themeInfo == nil {
			themeInfo = map[string]interface{}{"message": "Theme not found."}
		}

		jsonData, err := json.MarshalIndent(themeInfo, "", " ")

		if err != nil {
			fmt.Printf("/api/themes Error: %s\n", err)
		}

		return c.Send(jsonData)
	})
}
