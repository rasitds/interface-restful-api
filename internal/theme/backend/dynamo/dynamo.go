package dynamo

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/backend"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/models"
	"github.com/gin-gonic/gin"
)

type DynamoBackend struct {
	DB dynamodb.DynamoDB
}

func NewDynamoBackend(db dynamodb.DynamoDB) backend.Backend {
	return &DynamoBackend{
		DB: db,
	}
}

func (b *DynamoBackend) CreateTheme(c *gin.Context) (*models.Theme, error) {
	return nil, nil
}

func (b *DynamoBackend) ReadTheme(c *gin.Context) (*models.Theme, error) {
	return nil, nil
}

func (b *DynamoBackend) UpdateTheme(c *gin.Context) error {
	return nil
}

func (b *DynamoBackend) DeleteTheme(c *gin.Context) error {
	return nil
}

func (b *DynamoBackend) GetThemes(c *gin.Context) []models.Theme {
	return nil
}
