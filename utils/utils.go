package utils

import "reflect"

func Construct(params interface{}, retval interface{}) {
	switch t := reflect.TypeOf(retval); t {
	case reflect.Type:
		// We want this

	default:
		panic("retval must be a pointer to a type")
	}
}
