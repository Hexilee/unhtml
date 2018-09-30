package unhtml

import (
	"reflect"
)

const (
	ErrorMethodName = "Error"
)

// cannot use it for reference kind (Ptr, Interface, Func, Map, Slice)
//func isZero(v interface{}) bool {
//	return v == reflect.Zero(reflect.TypeOf(v)).Interface()
//}

// Converter: Func (inputType) -> (resultType, error)
func checkConverter(methodName string, methodType reflect.Type, expectResultType reflect.Type) (inputValuePtr reflect.Value, err error) {
	err = NewConverterTypeWrongError(methodName, methodType)
	if methodType.NumIn() == 1 &&
		methodType.NumOut() == 2 &&
		methodType.Out(0) == expectResultType {
		if _, exist := methodType.Out(1).MethodByName(ErrorMethodName); exist {
			inputValuePtr = reflect.New(methodType.In(0))
			err = nil
		}
	}
	return
}