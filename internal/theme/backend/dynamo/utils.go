package dynamo

import (
	"errors"

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

func CheckColorsAndUserIdGiven(theme models.Theme) error {
	if theme.BackgroundColor == "" {
		return errors.New("backgroundColor not given")
	}
	if theme.ForegroundColor == "" {
		return errors.New("foregroundColor not given")
	}
	if theme.UserID == "" {
		return errors.New("userId not given")
	}

	return nil
}
