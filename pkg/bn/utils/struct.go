package utils

import (
	"fmt"
	"reflect"
)

// TODO Duplicate with okx
func GetStructTagValueByField(st reflect.Type, field string, tag_name string) (string, error) {
	_field, found := st.FieldByName(field)
	if !found {
		return "", fmt.Errorf("field not found under %s field", field)
	}
	return _field.Tag.Get(tag_name), nil
}

func GetStructTagValueByIndex(st reflect.Type, tag string, index int) string {
	return st.FieldByIndex([]int{index}).Tag.Get(tag)
}
