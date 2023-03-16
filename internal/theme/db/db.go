package db

import "github.com/durmusrasit/sencha-restful-api/internal/theme/models"

func NewMemory() *[]models.Theme {
	var themes = []models.Theme{
		{ID: "1", ThemeName: "red", BackgroundColor: "#000", ForegroundColor: "#FF0000"}, {ID: "2", ThemeName: "yellow", BackgroundColor: "#000", ForegroundColor: "#FFFF00"},
	}

	return &themes
}
