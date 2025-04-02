package utils

import "reflect"

func IsStructurePointer(arg any) bool {
	val := reflect.ValueOf(arg)
	return val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Struct
}
