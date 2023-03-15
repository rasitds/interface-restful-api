package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Theme struct {
	ID              string `json:"id"`
	ThemeName       string `json:"themeName"`
	BackgroundColor string `json:"backgroundColor"`
	ForegroundColor string `json:"foregroundColor"`
}

var themes = []Theme{
	{ID: "1", ThemeName: "red", BackgroundColor: "#000", ForegroundColor: "#FF0000"}, {ID: "2", ThemeName: "yellow", BackgroundColor: "#000", ForegroundColor: "#FFFF00"},
}

func IsThemeExistsById(themes []Theme, themeId string) *int {
	for i, t := range themes {
		if t.ID == themeId {
			return &i
		}
	}
	return nil
}

func IsThemeExistsByName(themes []Theme, themeName string) *int {
	for i, t := range themes {
		if t.ThemeName == themeName {
			return &i
		}
	}
	return nil
}

func main() {
	router := gin.Default()

	apiRouter := router.Group("/api")

	apiRouter.Handle("GET", "/themes", getThemes)
	apiRouter.Handle("GET", "/theme/:name", readTheme)

	apiRouter.Handle("POST", "/theme", createTheme)
	apiRouter.Handle("POST", "/theme/update/:id", updateTheme)
	apiRouter.Handle("POST", "/theme/delete/:id", deleteTheme)

	router.Run()
}

func createTheme(c *gin.Context) {
	var newTheme Theme

	if c.Bind(&newTheme) == nil {
		index := IsThemeExistsByName(themes, newTheme.ThemeName)
		if index != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"data": "theme name exists"})
			return
		}

		themes = append(themes, newTheme)
		c.JSON(http.StatusOK, gin.H{"createdTheme": newTheme})
	}
}

func readTheme(c *gin.Context) {
	themeName := c.Params.ByName("name")

	for _, t := range themes {
		if t.ThemeName == themeName {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"data": "theme not found"})
}

func updateTheme(c *gin.Context) {
	themeId := c.Param("id")

	index := IsThemeExistsById(themes, themeId)
	if index == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"data": "theme not found"})
		return
	}

	var json struct {
		BackgroundColor string `json:"backgroundColor"`
		ForegroundColor string `json:"foregroundColor"`
	}

	var i = *index
	if c.Bind(&json) == nil {
		themes[i].BackgroundColor = json.BackgroundColor
		themes[i].ForegroundColor = json.ForegroundColor
	}

	c.JSON(http.StatusOK, gin.H{"status": true})
}

func deleteTheme(c *gin.Context) {
	themeId := c.Param("id")

	index := IsThemeExistsById(themes, themeId)
	if index == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"data": "theme not found"})
		return
	}

	var i = *index
	themes = append(themes[:i], themes[i+1:]...)

	c.JSON(http.StatusOK, gin.H{"status": true})
}

func getThemes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": themes,
	})
}
