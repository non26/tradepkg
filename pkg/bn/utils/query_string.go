package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"slices"
)

func CreateQueryStringFrom[T any](m *T, except_fields ...string) string {
	st := reflect.TypeOf(m).Elem()
	v := reflect.ValueOf(m).Elem()
	q := url.Values{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Tag.Get("json")
		if !slices.Contains(except_fields, field) {
			value := v.FieldByIndex([]int{i})
			q.Add(field, fmt.Sprintf("%v", value.Interface()))
		}
	}
	return q.Encode()
}

func GetQueryStringBinanceSignature(query_string, bnsign string) string {
	return fmt.Sprintf("%v&signature=%v", query_string, bnsign)
}
