package unmarshalers

import "reflect"

// cannot use it for reference kind (Ptr, Interface, Func, Map, Slice)
func isZero(v interface{}) bool {
	return v == reflect.Zero(reflect.TypeOf(v)).Interface()
}
