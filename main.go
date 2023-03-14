package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Theme struct {
	Background string `json:"background"`
	Foreground string `json:"foreground"`
}

var themes map[string]*Theme = map[string]*Theme{
	"default":      {Background: "black", Foreground: "white"},
	"cfi-blue":     {Background: "#1974D2", Foreground: "white"},
	"aqua":         {Background: "black", Foreground: "#33ffd0"},
	"white-orange": {Background: "#ff8000", Foreground: "white"},
	"light-blue":   {Background: "black", Foreground: "#33bbff"},
	"yellow":       {Background: "black", Foreground: "yellow"},
	"pinkish":      {Background: "black", Foreground: "#DE3163"},
	"dark":         {Background: "black", Foreground: "#f2f2f2"},
	"light":        {Background: "#f2f2f2", Foreground: "black"},
	"orange":       {Background: "black", Foreground: "#EA5B0C"},
	"cyan":         {Background: "black", Foreground: "#4CBEC5"},
	"green":        {Background: "black", Foreground: "#00CC11"},
	"pink":         {Background: "black", Foreground: "#FF6666"},
	"faint-orange": {Background: "black", Foreground: "#996633"},
	"neon-blue":    {Background: "black", Foreground: "#0033FF"},
	"ultra-green":  {Background: "black", Foreground: "#0AFF84"},
	"ultra-purple": {Background: "black", Foreground: "#8709F4"},
	"iron-gray":    {Background: "black", Foreground: "#52595D"},
	"bright-gray":  {Background: "black", Foreground: "#dcdcdc"},
	"bright-blue":  {Background: "black", Foreground: "#006EF0"},
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": themes,
		})
	})

	router.Run()
}
