package utils

import "github.com/durmusrasit/sencha-gin-api/internal/theme/models"

func IsThemeExistsById(themes []models.Theme, themeId string) *int {
	for i, t := range themes {
		if t.ID == themeId {
			return &i
		}
	}
	return nil
}

func IsThemeExistsByName(themes []models.Theme, themeName string) *int {
	for i, t := range themes {
		if t.ThemeName == themeName {
			return &i
		}
	}
	return nil
}
