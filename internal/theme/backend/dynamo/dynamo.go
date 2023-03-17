package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/backend"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/backend/utils"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DynamoBackend struct {
	DB dynamodb.DynamoDB
}

func NewDynamoBackend(db dynamodb.DynamoDB) backend.Backend {
	return &DynamoBackend{
		DB: db,
	}
}

var loggedUserId = "0"

func (b *DynamoBackend) CreateTheme(c *gin.Context) (*models.Theme, error) {
	var newTheme models.Theme

	if c.Bind(&newTheme) == nil {
		newTheme.ID = uuid.NewString()
		av, err := dynamodbattribute.MarshalMap(newTheme)
		if err != nil {
			return nil, errors.New("Got error marshalling new theme: " + err.Error())
		}

		tableName := "themes"

		filt := expression.Name("themeName").Equal(expression.Value(newTheme.ThemeName)).And(expression.Name("userId").Equal(expression.Value(loggedUserId)))

		proj := expression.NamesList(expression.Name("id"), expression.Name("themeName"), expression.Name("backgroundColor"), expression.Name("foregroundColor"), expression.Name("userId"))

		params := utils.GetExprScanInput(tableName, filt, proj)

		result, err := b.DB.Scan(params)
		if err != nil {
			return nil, errors.New("Query API call failed." + err.Error())
		}

		if len(result.Items) > 0 {
			return nil, errors.New("theme name exists")
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(tableName),
		}

		_, err = b.DB.PutItem(input)
		if err != nil {
			return nil, errors.New("Got error calling PutItem: " + err.Error())
		}

	}

	return &newTheme, nil
}

func (b *DynamoBackend) ReadTheme(c *gin.Context) (*models.Theme, error) {
	themeName := c.Params.ByName("name")

	tableName := "themes"

	filt := expression.Name("themeName").Equal(expression.Value(themeName))

	proj := expression.NamesList(expression.Name("id"), expression.Name("themeName"), expression.Name("backgroundColor"), expression.Name("foregroundColor"), expression.Name("userId"))

	params := utils.GetExprScanInput(tableName, filt, proj)

	result, err := b.DB.Scan(params)
	if err != nil {
		return nil, errors.New("Query API call failed: " + err.Error())
	}

	getTheme := models.Theme{}

	for _, i := range result.Items {
		theme := models.Theme{}

		err = dynamodbattribute.UnmarshalMap(i, &theme)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}

		if theme.UserID == loggedUserId {
			getTheme.ID = theme.ID
			getTheme.ThemeName = theme.ThemeName
			getTheme.BackgroundColor = theme.BackgroundColor
			getTheme.ForegroundColor = theme.ForegroundColor
			getTheme.UserID = theme.UserID
			return &getTheme, nil
		}
	}

	return nil, errors.New("theme not found")
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
