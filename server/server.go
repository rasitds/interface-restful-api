package server

import "github.com/durmusrasit/sencha-restful-api/internal/theme/backend"

type ThemeAPIServer struct{ backend backend.Backend }

func NewThemeAPIServer(backend backend.Backend) *ThemeAPIServer {
	return &ThemeAPIServer{backend: backend}
}
