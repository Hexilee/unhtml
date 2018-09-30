package unhtml

import "reflect"

// cannot use it for reference kind (Ptr, Interface, Func, Map, Slice)
func isZero(v interface{}) bool {
	return v == reflect.Zero(reflect.TypeOf(v)).Interface()
}

// Converter: Func (inputType) -> (resultType, error)
func checkConverter(methodType reflect.Method, expectResultType reflect.Type) (inputType reflect.Type, err error) {

	return
}
