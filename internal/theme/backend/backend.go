package backend

import (
	"github.com/durmusrasit/sencha-gin-api/internal/theme/models"
	"github.com/gin-gonic/gin"
)

type Backend interface {
	CreateTheme(c *gin.Context) (*models.Theme, error)
	ReadTheme(c *gin.Context) (*models.Theme, error)
	UpdateTheme(c *gin.Context) error
	DeleteTheme(c *gin.Context) error
	GetThemes(c *gin.Context) []models.Theme
}
