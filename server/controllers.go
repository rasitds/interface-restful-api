package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *ThemeAPIServer) CreateTheme(c *gin.Context) {
	createdTheme, err := s.backend.CreateTheme(c)
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"createdTheme": createdTheme})
}

func (s *ThemeAPIServer) ReadTheme(c *gin.Context) {
	theme, err := s.backend.ReadTheme(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, theme)
}

func (s *ThemeAPIServer) UpdateTheme(c *gin.Context) {
	err := s.backend.UpdateTheme(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (s *ThemeAPIServer) DeleteTheme(c *gin.Context) {
	err := s.backend.DeleteTheme(c)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true})
}

func (s *ThemeAPIServer) GetThemes(c *gin.Context) {
	themeData := s.backend.GetThemes(c)
	c.JSON(http.StatusOK, gin.H{
		"data": themeData,
	})
}
