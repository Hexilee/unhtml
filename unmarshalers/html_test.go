package unmarshalers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type (
	HttpTestStruct struct {
		Interface interface{}
		Chan      chan int
		Func      func()
		Map       map[string]string
	}
)

var (
	httpTestStruct = HttpTestStruct{
		Interface: 0,
		Chan:      make(chan int, 0),
		Func: func() {

		},
		Map: make(map[string]string),
	}

	InterfaceAddr = reflect.ValueOf(&httpTestStruct).Elem().FieldByName("Interface").Addr().Interface()
	ChanAddr = reflect.ValueOf(&httpTestStruct).Elem().FieldByName("Chan").Addr().Interface()
	FuncAddr = reflect.ValueOf(&httpTestStruct).Elem().FieldByName("Func").Addr().Interface()
	MapAddr = reflect.ValueOf(&httpTestStruct).Elem().FieldByName("Map").Addr().Interface()
)

func TestHTMLMarshaler_parseType(t *testing.T) {
	var (
		Int      = 0
		IntPtr   = &Int
		IntSlice = make([]int, 0)
	)

	for _, testCase := range [] *struct {
		dto      interface{}
		kind     reflect.Kind
		itemType reflect.Type
		err      error
	}{
		{&IntSlice, reflect.Slice, reflect.TypeOf(IntSlice), nil},
		{Int, 0, nil, errors.New(UnmarshaledKindMustBePtr)},
		{IntPtr, reflect.Int, reflect.TypeOf(Int), nil},
		{&IntPtr, reflect.Ptr, reflect.TypeOf(IntPtr), errors.New(UnmarshalerItemKindError)},
		{InterfaceAddr, reflect.Interface, reflect.TypeOf(InterfaceAddr).Elem(), errors.New(UnmarshalerItemKindError)},
		{ChanAddr, reflect.Chan, reflect.TypeOf(ChanAddr).Elem(), errors.New(UnmarshalerItemKindError)},
		{FuncAddr, reflect.Func, reflect.TypeOf(FuncAddr).Elem(), errors.New(UnmarshalerItemKindError)},
		{MapAddr, reflect.Map, reflect.TypeOf(MapAddr).Elem(), errors.New(UnmarshalerItemKindError)},
	} {
		func() {
			result := new(HTMLUnmarshaler)
			parseErr := result.initDto(reflect.ValueOf(testCase.dto))
			if parseErr != nil && testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), parseErr.Error())
			} else {
				assert.Equal(t, testCase.err, parseErr)
			}
			assert.Equal(t, testCase.kind, result.kind)
			assert.Equal(t, testCase.itemType, result.dtoElemType)
		}()
	}
}
