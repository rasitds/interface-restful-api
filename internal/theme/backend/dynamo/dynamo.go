package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/backend"
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

func (b *DynamoBackend) CreateTheme(c *gin.Context) (*models.Theme, error) {
	var newTheme models.Theme
	tableName := "themes"

	if c.Bind(&newTheme) == nil {
		newTheme.ID = uuid.NewString()
		av, err := dynamodbattribute.MarshalMap(newTheme)
		if err != nil {
			return nil, errors.New("Got error marshalling new theme: " + err.Error())
		}

		if err := CheckColorsAndUserIdGiven(newTheme); err != nil {
			return nil, err
		}

		params := GetThemeNameAndUserIdFiltExprScanInput(tableName, newTheme)

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
	themeName := c.Params.ByName("themeName")
	userId := c.Params.ByName("userId")

	tableName := "themes"

	data := models.Theme{
		ThemeName: themeName,
		UserID:    userId,
	}

	params := GetThemeNameAndUserIdFiltExprScanInput(tableName, data)

	result, err := b.DB.Scan(params)
	if err != nil {
		return nil, errors.New("Query API call failed: " + err.Error())
	}

	if len(result.Items) < 1 {
		return nil, errors.New("theme not found")
	}

	var theme models.Theme
	for _, i := range result.Items {

		err = dynamodbattribute.UnmarshalMap(i, &theme)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}

	}

	return &theme, nil
}

func (b *DynamoBackend) UpdateTheme(c *gin.Context) error {
	tableName := "themes"

	var updatedTheme models.Theme

	if c.Bind(&updatedTheme) == nil {
		if err := CheckColorsAndUserIdGiven(updatedTheme); err != nil {
			return err
		}

		params := GetThemeNameAndUserIdFiltExprScanInput(tableName, updatedTheme)

		result, resultErr := b.DB.Scan(params)
		if resultErr != nil {
			return errors.New("Query API call failed." + resultErr.Error())
		}

		if len(result.Items) < 1 {
			return errors.New("theme does not exists")
		}

		input := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":bg": {
					S: aws.String(updatedTheme.BackgroundColor),
				},
				":fg": {
					S: aws.String(updatedTheme.ForegroundColor),
				},
			},
			TableName: aws.String(tableName),
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(updatedTheme.ID),
				},
				"themeName": {
					S: aws.String(updatedTheme.ThemeName),
				},
			},
			ReturnValues:     aws.String("UPDATED_NEW"),
			UpdateExpression: aws.String("set backgroundColor = :bg, foregroundColor = :fg"),
		}

		_, err := b.DB.UpdateItem(input)
		if err != nil {
			return errors.New("error while updateitem" + err.Error())
		}
	}

	return nil
}

func (b *DynamoBackend) DeleteTheme(c *gin.Context) error {
	tableName := "themes"

	var theme models.Theme

	if c.Bind(&theme) == nil {
		if theme.ThemeName == "" {
			return errors.New("themeName not given")
		}
		if theme.UserID == "" {
			return errors.New("userId not given")
		}

		params := GetThemeNameAndUserIdFiltExprScanInput(tableName, theme)

		result, resultErr := b.DB.Scan(params)
		if resultErr != nil {
			return errors.New("Query API call failed." + resultErr.Error())
		}

		if len(result.Items) < 1 {
			return errors.New("theme does not exists")
		}

		input := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(theme.ID),
				},
				"themeName": {
					S: aws.String(theme.ThemeName),
				},
			},
			TableName: aws.String(tableName),
		}

		_, err := b.DB.DeleteItem(input)
		if err != nil {
			return errors.New("An error accurred while calling DeleteItem: " + err.Error())
		}
	}

	return nil
}

func (b *DynamoBackend) GetThemes(c *gin.Context) []models.Theme {
	return nil
}
