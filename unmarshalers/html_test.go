package unmarshalers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

var (
	htmlTestStruct = HtmlTestStruct{
		Interface: 0,
		Chan:      make(chan int, 0),
		Func: func() {

		},
		Map: make(map[string]string),
	}

	InterfaceAddr = reflect.ValueOf(&htmlTestStruct).Elem().FieldByName("Interface").Addr().Interface()
	ChanAddr      = reflect.ValueOf(&htmlTestStruct).Elem().FieldByName("Chan").Addr().Interface()
	FuncAddr      = reflect.ValueOf(&htmlTestStruct).Elem().FieldByName("Func").Addr().Interface()
	MapAddr       = reflect.ValueOf(&htmlTestStruct).Elem().FieldByName("Map").Addr().Interface()
)

type (
	HtmlTestStruct struct {
		Interface interface{}
		Chan      chan int
		Func      func()
		Map       map[string]string
	}

	Link struct {
		Text string `key:"text"`
		Href string `key:"href"`
	}

	Course struct {
		Code     Link   `html:"td:nth-child(1) > a"`
		Name     Link   `html:"td:nth-child(2) > a"`
		Teacher  Link   `html:"td:nth-child(3) > a"`
		Semester string `html:"td:nth-child(4)";key:"text"`
		Time     string `html:"td:nth-child(5)";key:"text"`
		Location string `html:"td:nth-child(6)";key:"text"`
	}

	Courses []Course
)

func (courses Courses) Root() string {
	return "#xsgrid > tbody > tr:nth-child(1n+2)"
}

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

func TestHTMLUnmarshaler_Unmarshal(t *testing.T) {
	TestHTML, err := ioutil.ReadFile("testFiles/courses.html")
	assert.Nil(t, err)
	courses := make(Courses, 0)
	assert.Nil(t, new(HTMLUnmarshaler).Unmarshal(TestHTML, &courses))
	result, err := json.Marshal(courses)
	assert.Nil(t, err)
	fmt.Println(string(result))
}
