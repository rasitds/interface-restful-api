package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/backend/dynamo"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/db"
	"github.com/durmusrasit/sencha-restful-api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{SharedConfigState: session.SharedConfigEnable}))

	db := db.NewDynamo(sess)
	router := gin.Default()

	apiRouter := router.Group("/api")

	backend := server.NewThemeAPIServer(dynamo.NewDynamoBackend(*db))

	apiRouter.Handle("GET", "/themes", backend.GetThemes)
	apiRouter.Handle("GET", "/theme/:themeName/:userId", backend.ReadTheme)

	apiRouter.Handle("POST", "/theme", backend.CreateTheme)
	apiRouter.Handle("POST", "/theme/update", backend.UpdateTheme)
	apiRouter.Handle("POST", "/theme/delete", backend.DeleteTheme)

	router.Run()
}
