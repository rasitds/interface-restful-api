package db

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/models"
)

func NewMemory() *[]models.Theme {
	var themes = []models.Theme{
		{ID: "1", ThemeName: "red", BackgroundColor: "#000", ForegroundColor: "#FF0000"}, {ID: "2", ThemeName: "yellow", BackgroundColor: "#000", ForegroundColor: "#FFFF00"},
	}

	return &themes
}

func NewDynamo(sess *session.Session) *dynamodb.DynamoDB {
	svc := dynamodb.New(sess)

	return svc
}
