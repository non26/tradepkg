package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"slices"
)

// create tag, binance:"optional"
func CreateQueryStringFrom[T any](m *T, except_fields ...string) string {
	st := reflect.TypeOf(m).Elem()
	v := reflect.ValueOf(m).Elem()
	q := url.Values{}
	for i := 0; i < st.NumField(); i++ {
		json_tag_value := st.Field(i).Tag.Get("json")
		binance_tag_value := st.Field(i).Tag.Get("binance")
		if !slices.Contains(except_fields, json_tag_value) {
			json_value := v.FieldByIndex([]int{i})
			println(json_value.String())
			if json_value.Kind() == reflect.String && json_value.String() == "" && binance_tag_value == "optional" {
				continue
			}
			q.Add(json_tag_value, fmt.Sprintf("%v", json_value.Interface()))
		}
	}
	return q.Encode()
}

func GetQueryStringBinanceSignature(query_string, bnsign string) string {
	return fmt.Sprintf("%v&signature=%v", query_string, bnsign)
}
