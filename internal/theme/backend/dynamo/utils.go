package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/durmusrasit/sencha-restful-api/internal/theme/models"
)

func GetThemeNameAndUserIdFiltExprScanInput(tableName string, theme models.Theme) *dynamodb.ScanInput {
	filt := expression.Name("themeName").Equal(expression.Value(theme.ThemeName)).And(expression.Name("userId").Equal(expression.Value(theme.UserID)))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		panic("got error building expression: " + err.Error())
	}

	scanInput := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	}

	return scanInput
}
