package utils

import "reflect"

func Construct[P any, T any, PT interface{ *T }](params interface{}) PT {
	p := PT(new(T))
	Construct0(params, p)
	return p
}

func Construct0(params interface{}, retval interface{}) {
	// Check if retval is a pointer
	rv := reflect.ValueOf(retval)
	if rv.Kind() != reflect.Ptr {
		panic("retval is not a pointer")
	}

	// Dereference the pointer to get the underlying value
	rv = rv.Elem()

	// Check if the dereferenced value is a struct
	if rv.Kind() != reflect.Struct {
		panic("retval is not a pointer to a struct")
	}

	// Now, get the value of params
	rp := reflect.ValueOf(params)
	if rp.Kind() != reflect.Struct {
		panic("params is not a struct")
	}

	// Iterate over the fields of params and copy to retval
	for i := 0; i < rp.NumField(); i++ {
		name := rp.Type().Field(i).Name
		field, ok := rv.Type().FieldByName(name)
		if ok && field.Type == rp.Field(i).Type() {
			rv.FieldByName(name).Set(rp.Field(i))
		}
	}

}
