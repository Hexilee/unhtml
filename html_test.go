package unhtml

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const (
	CoursesJSON  = `[{"code":{"text":"061B0020","href":"#"},"name":{"text":"复变函数与积分变换","href":"#"},"teacher":{"text":"王伟","href":"#"},"semester":"秋","time":"周一第1,2节周四第1,2节","location":"紫金港西2-205(多)紫金港西2-205(多)"},{"code":{"text":"101C0350","href":"#"},"name":{"text":"电路与模拟电子技术","href":"#"},"teacher":{"text":"孙盾","href":"#"},"semester":"秋冬","time":"周二第6,7节周二第8节{单周}周五第3,4,5节","location":"紫金港西1-417(多)紫金港西1-417(多)紫金港西1-417(多)"},{"code":{"text":"101C0360","href":"#"},"name":{"text":"电路与模拟电子技术实验","href":"#"},"teacher":{"text":"干于","href":"#"},"semester":"秋冬","time":"周四第3,4,5节","location":"紫金港东3-202"},{"code":{"text":"241L0020","href":"#"},"name":{"text":"博弈论基础","href":"#"},"teacher":{"text":"蒋文华","href":"#"},"semester":"冬","time":"周三第6,7,8节","location":"紫金港西1-316(多)*"},{"code":{"text":"261C0070","href":"#"},"name":{"text":"工程力学","href":"#"},"teacher":{"text":"吴禹季葆华","href":"#"},"semester":"秋冬","time":"周二第1,2节{单周}周四第6,7节周四第8节{双周}","location":"紫金港西1-404(多)紫金港西1-404(多)紫金港西1-404(多)"},{"code":{"text":"74188020","href":"#"},"name":{"text":"专业实习","href":"#"},"teacher":{"text":"陈家旺黄豪彩","href":"#"},"semester":"短","time":" ","location":" "},{"code":{"text":"761T0010","href":"#"},"name":{"text":"大学物理（甲）Ⅰ","href":"#"},"teacher":{"text":"潘国卫","href":"#"},"semester":"秋冬","time":"周六第6,7,8,9节","location":"紫金港西2-101(多)"},{"code":{"text":"761T0020","href":"#"},"name":{"text":"大学物理（甲）Ⅱ","href":"#"},"teacher":{"text":"郑大方","href":"#"},"semester":"秋冬","time":"周一第3,4节周三第1,2节","location":"紫金港西2-202(多)#"},{"code":{"text":"821T0020","href":"#"},"name":{"text":"微积分（甲）Ⅱ","href":"#"},"teacher":{"text":"薛儒英","href":"#"},"semester":"秋冬","time":"周六第1,2,3,4节{单周}周六第1,2,3,4,5节{双周}","location":"紫金港西2-105(多)"}]`
	AllTypesJSON = `{"Slice":[0,1,2,3],"Struct":{"Name":"Hexilee","Age":20,"LikeLemon":true},"String":"Hello World!","Int":10,"Int8":10,"Int16":10,"Int32":10,"Int64":10,"Uint":10,"Uint8":10,"Uint16":10,"Uint32":10,"Uint64":10,"Float32":3.14,"Float64":3.14,"Bool":true}`
)

var (
	CourseHTML, _  = ioutil.ReadFile("testHTML/courses.html")
	AllTypeHTML, _ = ioutil.ReadFile("testHTML/all-type.html")
)

type (
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
		String  string   `html:"#test > p:nth-child(3)"`
		Int     int      `html:"#test > p:nth-child(4)"`
		Int8    int8     `html:"#test > p:nth-child(4)"`
		Int16   int16    `html:"#test > p:nth-child(4)"`
		Int32   int32    `html:"#test > p:nth-child(4)"`
		Int64   int64    `html:"#test > p:nth-child(4)"`
		Uint    uint     `html:"#test > p:nth-child(4)"`
		Uint8   uint8    `html:"#test > p:nth-child(4)"`
		Uint16  uint16   `html:"#test > p:nth-child(4)"`
		Uint32  uint32   `html:"#test > p:nth-child(4)"`
		Uint64  uint64   `html:"#test > p:nth-child(4)"`
		Float32 float32  `html:"#test > p:nth-child(5)"`
		Float64 float64  `html:"#test > p:nth-child(5)"`
		Bool    bool     `html:"#test > p:nth-child(6)"`
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

func TestHTMLUnmarshaler_Unmarshal(t *testing.T) {
	assert.NotNil(t, CourseHTML)
	courses := make(Courses, 0)
	assert.Nil(t, new(HTMLUnmarshaler).Unmarshal(CourseHTML, &courses))
	result, err := json.Marshal(courses)
	assert.Nil(t, err)
	assert.Equal(t, CoursesJSON, string(result))

	assert.NotNil(t, AllTypeHTML)
	allTypes := AllTypeTest{}
	assert.Nil(t, new(HTMLUnmarshaler).Unmarshal(AllTypeHTML, &allTypes))
	result, err = json.Marshal(&allTypes)
	assert.Nil(t, err)
	assert.Equal(t, AllTypesJSON, string(result))
}

func BenchmarkUnmarshalCourse(b *testing.B) {
	assert.NotNil(b, CourseHTML)
	for i := 0; i < b.N; i++ {
		courses := make(Courses, 0)
		assert.Nil(b, new(HTMLUnmarshaler).Unmarshal(CourseHTML, &courses))
	}
}

func getLink(t assert.TestingT, selection *goquery.Selection) Link {
	link, _ := selection.Attr(AttrHref)
	return Link{Text: selection.Text(), Href: link}
}

func parseHTMLLogically(t assert.TestingT) (Courses, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(CourseHTML))
	courses := make(Courses, 0)
	if err == nil {
		doc.Find(courses.Root()).Each(func(i int, selection *goquery.Selection) {
			course := Course{}
			course.Code = getLink(t, selection.Find(`td:nth-child(1) > a`))
			course.Name = getLink(t, selection.Find(`td:nth-child(2) > a`))
			course.Teacher = getLink(t, selection.Find(`td:nth-child(3) > a`))
			course.Semester = selection.Find(`td:nth-child(4)`).Text()
			course.Time = selection.Find(`td:nth-child(5)`).Text()
			course.Location = selection.Find(`td:nth-child(6)`).Text()
			courses = append(courses, course)
		})
	}

	return courses, err
}

func BenchmarkParseHTMLLogically(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := parseHTMLLogically(b)
		assert.Nil(b, err)
	}
}
