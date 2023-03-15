package server

import "github.com/durmusrasit/sencha-gin-api/internal/theme/backend"

type ThemeAPIServer struct{ backend backend.Backend }

func NewThemeAPIServer(backend backend.Backend) *ThemeAPIServer {
	return &ThemeAPIServer{backend: backend}
}
