package main

import (
	"github.com/durmusrasit/sencha-gin-api/internal/theme/backend/memory"
	"github.com/durmusrasit/sencha-gin-api/internal/theme/db"
	"github.com/durmusrasit/sencha-gin-api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiRouter := router.Group("/api")

	db := db.NewMemory()
	backend := server.NewThemeAPIServer(memory.NewMemoryBackend(*db))

	apiRouter.Handle("GET", "/themes", backend.GetThemes)
	apiRouter.Handle("GET", "/theme/:name", backend.ReadTheme)

	apiRouter.Handle("POST", "/theme", backend.CreateTheme)
	apiRouter.Handle("POST", "/theme/update/:id", backend.UpdateTheme)
	apiRouter.Handle("POST", "/theme/delete/:id", backend.DeleteTheme)

	router.Run()
}
