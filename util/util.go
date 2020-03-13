package util

import (
	"fmt"
	"strconv"
	"strings"
)

func StringConv(s string, defaultValue interface{}, valueType string) interface{} {
	var err error
	var value interface{}
	switch valueType {
	case "bool":
		value, err = strconv.ParseBool(s)
	case "int":
		value, err = strconv.ParseInt(s, 10, 0)
		value = int(value.(int64))
	case "int8":
		value, err = strconv.ParseInt(s, 10, 8)
		value = int8(value.(int64))
	case "int16":
		value, err = strconv.ParseInt(s, 10, 16)
		value = int16(value.(int64))
	case "int32":
		value, err = strconv.ParseInt(s, 10, 32)
		value = int32(value.(int64))
	case "int64":
		value, err = strconv.ParseInt(s, 10, 64)
	case "float64":
		value, err = strconv.ParseFloat(s, 64)
	case "string":
		value = s
		if value == "" {
			value = defaultValue.(string)
		}
	case "[]string":
		value = strings.Fields(s)
	default:
		panic(fmt.Sprintf("nonexistent value type: %s", valueType))
	}
	if err != nil {
		value = defaultValue
	}
	return value
}
