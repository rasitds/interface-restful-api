package memory

import (
	"errors"

	"github.com/durmusrasit/sencha-restful-api/internal/theme/backend"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/backend/utils"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/models"
	"github.com/gin-gonic/gin"
)

type MemoryBackend struct {
	DB []models.Theme
}

func NewMemoryBackend(db []models.Theme) backend.Backend {
	return &MemoryBackend{
		DB: db,
	}
}

func (b *MemoryBackend) CreateTheme(c *gin.Context) (*models.Theme, error) {
	var newTheme models.Theme

	if c.Bind(&newTheme) == nil {
		index := utils.IsThemeExistsByName(b.DB, newTheme.ThemeName)
		if index != nil {
			return nil, errors.New("theme name exists")
		}

		b.DB = append(b.DB, newTheme)
	}

	return &newTheme, nil
}

func (b *MemoryBackend) ReadTheme(c *gin.Context) (*models.Theme, error) {
	themeName := c.Params.ByName("name")

	for _, t := range b.DB {
		if t.ThemeName == themeName {
			return &t, nil
		}
	}

	return nil, errors.New("theme not found")
}

func (b *MemoryBackend) UpdateTheme(c *gin.Context) error {
	themeId := c.Param("id")

	index := utils.IsThemeExistsById(b.DB, themeId)
	if index == nil {
		return errors.New("theme not found")
	}

	var json struct {
		BackgroundColor string `json:"backgroundColor"`
		ForegroundColor string `json:"foregroundColor"`
	}

	var i = *index
	if c.Bind(&json) == nil {
		b.DB[i].BackgroundColor = json.BackgroundColor
		b.DB[i].ForegroundColor = json.ForegroundColor
	}

	return nil
}

func (b *MemoryBackend) DeleteTheme(c *gin.Context) error {
	themeId := c.Param("id")

	index := utils.IsThemeExistsById(b.DB, themeId)
	if index == nil {
		return errors.New("theme not found")
	}

	var i = *index
	b.DB = append(b.DB[:i], b.DB[i+1:]...)

	return nil
}

func (b *MemoryBackend) GetThemes(c *gin.Context) []models.Theme {
	return b.DB
}
