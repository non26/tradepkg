package dynamodbconfig

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UpdateTableAt[T any] struct {
	data                      *T
	updateExpression          []string
	expressionAttributeValues map[string]types.AttributeValue
}

// T = table type
func NewUpdateTable[T any](with_data *T) *UpdateTableAt[T] {
	return &UpdateTableAt[T]{
		updateExpression:          []string{},
		expressionAttributeValues: map[string]types.AttributeValue{},
		data:                      with_data,
	}
}

func (u *UpdateTableAt[T]) setExpression(field string) {
	u.updateExpression = append(u.updateExpression, fmt.Sprintf("%v = :%v", field, field))
}

func (u *UpdateTableAt[T]) GetExpressionAttributeValues() map[string]types.AttributeValue {
	return u.expressionAttributeValues
}

func (u *UpdateTableAt[T]) setExpressionAttributeValues(field string, to_value any, value_type reflect.Type) {
	key := fmt.Sprintf(":%v", field)
	switch value_type.Kind() {
	case reflect.String:
		u.expressionAttributeValues[key] = &types.AttributeValueMemberS{Value: to_value.(string)}
	case reflect.Int64:
		v := to_value.(int64)
		u.expressionAttributeValues[key] = &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", v)}
	case reflect.Bool:
		v := to_value.(bool)
		u.expressionAttributeValues[key] = &types.AttributeValueMemberBOOL{Value: v}
	}
}

func (u *UpdateTableAt[T]) BuildExpression() *string {
	return aws.String("set " + strings.Join(u.updateExpression, ", "))
}

// don't set key
func (u *UpdateTableAt[T]) Set(field func() (string, reflect.Type), to_value any) *UpdateTableAt[T] {
	_field, of_type := field()
	u.setExpression(_field)
	u.setExpressionAttributeValues(_field, to_value, of_type)
	return u
}
