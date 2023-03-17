package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/models"
)

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

func GetExprScanInput(tableName string, filt expression.ConditionBuilder, proj expression.ProjectionBuilder) *dynamodb.ScanInput {
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		panic("got error building expression: " + err.Error())
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	return params
}
