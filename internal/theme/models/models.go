package models

type Theme struct {
	ID              string `json:"id"`
	ThemeName       string `json:"themeName"`
	BackgroundColor string `json:"backgroundColor"`
	ForegroundColor string `json:"foregroundColor"`
	UserID          string `json:"userId"`
}
