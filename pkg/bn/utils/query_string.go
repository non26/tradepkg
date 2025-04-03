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
			// if value.Kind() == reflect.String {
			// 	q.Add(field, value.String())
			// 	continue
			// }
			// if value.Kind() == reflect.Int64 {
			// 	q.Add(field, fmt.Sprintf("%d", value.Int()))
			// 	continue
			// }
			q.Add(field, fmt.Sprintf("%v", value.Interface()))

			// if value.Kind() == reflect.Slice {
			// 	q.Add(field, fmt.Sprintf("%v", value.Interface()))
			// 	continue
			// }
			// if value.Kind() == reflect.Slice {
			// 	for j := 0; j < value.Len(); j++ {
			// 		q.Add(field, fmt.Sprintf("%v", value.Index(j).Interface()))
			// 	}
			// 	continue
			// }
			// if value.Kind() == reflect.Struct {
			// 	q.Add(field, fmt.Sprintf("%v", value.Interface()))
			// 	continue
			// }
		}
	}
	return q.Encode()
}

func GetQueryStringBinanceSignature(query_string, bnsign string) string {
	return fmt.Sprintf("%v&signature=%v", query_string, bnsign)
}
