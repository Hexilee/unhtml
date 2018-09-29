package unmarshalers

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

const (
	CoursesJSON = `[{"code":{"text":"061B0020","href":"#"},"name":{"text":"复变函数与积分变换","href":"#"},"teacher":{"text":"王伟","href":"#"},"semester":"秋","time":"周一第1,2节周四第1,2节","location":"紫金港西2-205(多)紫金港西2-205(多)"},{"code":{"text":"101C0350","href":"#"},"name":{"text":"电路与模拟电子技术","href":"#"},"teacher":{"text":"孙盾","href":"#"},"semester":"秋冬","time":"周二第6,7节周二第8节{单周}周五第3,4,5节","location":"紫金港西1-417(多)紫金港西1-417(多)紫金港西1-417(多)"},{"code":{"text":"101C0360","href":"#"},"name":{"text":"电路与模拟电子技术实验","href":"#"},"teacher":{"text":"干于","href":"#"},"semester":"秋冬","time":"周四第3,4,5节","location":"紫金港东3-202"},{"code":{"text":"241L0020","href":"#"},"name":{"text":"博弈论基础","href":"#"},"teacher":{"text":"蒋文华","href":"#"},"semester":"冬","time":"周三第6,7,8节","location":"紫金港西1-316(多)*"},{"code":{"text":"261C0070","href":"#"},"name":{"text":"工程力学","href":"#"},"teacher":{"text":"吴禹季葆华","href":"#"},"semester":"秋冬","time":"周二第1,2节{单周}周四第6,7节周四第8节{双周}","location":"紫金港西1-404(多)紫金港西1-404(多)紫金港西1-404(多)"},{"code":{"text":"74188020","href":"#"},"name":{"text":"专业实习","href":"#"},"teacher":{"text":"陈家旺黄豪彩","href":"#"},"semester":"短","time":" ","location":" "},{"code":{"text":"761T0010","href":"#"},"name":{"text":"大学物理（甲）Ⅰ","href":"#"},"teacher":{"text":"潘国卫","href":"#"},"semester":"秋冬","time":"周六第6,7,8,9节","location":"紫金港西2-101(多)"},{"code":{"text":"761T0020","href":"#"},"name":{"text":"大学物理（甲）Ⅱ","href":"#"},"teacher":{"text":"郑大方","href":"#"},"semester":"秋冬","time":"周一第3,4节周三第1,2节","location":"紫金港西2-202(多)#"},{"code":{"text":"821T0020","href":"#"},"name":{"text":"微积分（甲）Ⅱ","href":"#"},"teacher":{"text":"薛儒英","href":"#"},"semester":"秋冬","time":"周六第1,2,3,4节{单周}周六第1,2,3,4,5节{双周}","location":"紫金港西2-105(多)"}]`
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
		Text string `json:"text"`
		Href string `key:"href" json:"href"`
	}

	Course struct {
		Code     Link   `html:"td:nth-child(1) > a" json:"code"`
		Name     Link   `html:"td:nth-child(2) > a" json:"name"`
		Teacher  Link   `html:"td:nth-child(3) > a" json:"teacher"`
		Semester string `html:"td:nth-child(4)" json:"semester"`
		Time     string `html:"td:nth-child(5)" json:"time"`
		Location string `html:"td:nth-child(6)" json:"location"`
	}

	Courses []Course

	AllTypeTest struct {
		Slice   []int    `html:"ul > li"`
		Struct  TestUser `html:"div"`
		String  string   `html:"p:nth-child(1)"`
		Int     int      `html:"p:nth-child(2)"`
		Int8    int8     `html:"p:nth-child(2)"`
		Int16   int16    `html:"p:nth-child(2)"`
		Int32   int32    `html:"p:nth-child(2)"`
		Int64   int64    `html:"p:nth-child(2)"`
		Uint    uint     `html:"p:nth-child(2)"`
		Uint8   uint8    `html:"p:nth-child(2)"`
		Uint16  uint16   `html:"p:nth-child(2)"`
		Uint32  uint32   `html:"p:nth-child(2)"`
		Uint64  uint64   `html:"p:nth-child(2)"`
		Float32 float32  `html:"p:nth-child(3)"`
		Float64 float64  `html:"p:nth-child(3)"`
		Bool    bool     `html:"p:nth-child(4)"`
	}

	TestUser struct {
		Name      string `html:"p:nth-child(1)"`
		Age       uint   `html:"p:nth-child(2)"`
		LikeLemon bool   `html:"p:nth-child(3)"`
	}
)

func (courses Courses) Root() string {
	return "#xsgrid > tbody > tr:nth-child(1n+2)"
}

func (AllTypeTest) Root() string {
	return "#test"
}

//func TestHTMLMarshaler_parseType(t *testing.T) {
//	var (
//		Int      = 0
//		IntPtr   = &Int
//		IntSlice = make([]int, 0)
//	)
//
//	for _, testCase := range []*struct {
//		dto      interface{}
//		kind     reflect.Kind
//		itemType reflect.Type
//		err      error
//	}{
//		{&IntSlice, reflect.Slice, reflect.TypeOf(IntSlice), nil},
//		{Int, 0, nil, errors.New(UnmarshaledKindMustBePtr)},
//		{IntPtr, reflect.Int, reflect.TypeOf(Int), nil},
//		{&IntPtr, reflect.Ptr, reflect.TypeOf(IntPtr), errors.New(UnmarshalerItemKindError)},
//		{InterfaceAddr, reflect.Interface, reflect.TypeOf(InterfaceAddr).Elem(), errors.New(UnmarshalerItemKindError)},
//		{ChanAddr, reflect.Chan, reflect.TypeOf(ChanAddr).Elem(), errors.New(UnmarshalerItemKindError)},
//		{FuncAddr, reflect.Func, reflect.TypeOf(FuncAddr).Elem(), errors.New(UnmarshalerItemKindError)},
//		{MapAddr, reflect.Map, reflect.TypeOf(MapAddr).Elem(), errors.New(UnmarshalerItemKindError)},
//	} {
//		func() {
//			builder := new(HTMLUnmarshalerBuilder).setDto(reflect.ValueOf(testCase.dto))
//
//			parseErr := builder.
//			if parseErr != nil && testCase.err != nil {
//				assert.Equal(t, testCase.err.Error(), parseErr.Error())
//			} else {
//				assert.Equal(t, testCase.err, parseErr)
//			}
//			assert.Equal(t, testCase.kind, result.kind)
//			assert.Equal(t, testCase.itemType, result.dtoElemType)
//		}()
//	}
//}

func TestHTMLUnmarshaler_Unmarshal(t *testing.T) {
	TestHTML, err := ioutil.ReadFile("testFiles/courses.html")
	assert.Nil(t, err)
	courses := make(Courses, 0)
	assert.Nil(t, new(HTMLFooUnmarshaler).Unmarshal(TestHTML, &courses))
	result, err := json.Marshal(courses)
	assert.Nil(t, err)
	assert.Equal(t, CoursesJSON, string(result))

	AllTypeHTML, err := ioutil.ReadFile("testFiles/all-type.html")
	assert.Nil(t, err)
	allTypes := AllTypeTest{}
	assert.Nil(t, new(HTMLFooUnmarshaler).Unmarshal(AllTypeHTML, &allTypes))
	result, err = json.Marshal(&allTypes)
	assert.Nil(t, err)
	fmt.Println(string(result))
}
