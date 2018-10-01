package unhtml

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type (
	PartTypesStruct struct {
		Slice   []int    `html:"ul > li"`
		Struct  TestUser `html:"#test > div"`
		String  string   `html:"#test > p:nth-child(3)"`
		Int     int      `html:"#test > p:nth-child(4)"`
		Float64 float64  `html:"#test > p:nth-child(5)"`
		Bool    bool     `html:"#test > p:nth-child(6)"`
	}
)

func (PartTypesStruct) Root() string {
	return "#test"
}

func BenchmarkUnmarshalCourses(b *testing.B) {
	assert.NotNil(b, CourseHTML)
	for i := 0; i < b.N; i++ {
		courses := make(Courses, 0)
		assert.Nil(b, Unmarshal(CourseHTML, &courses))
	}
}

func BenchmarkUnmarshalPartTypes(b *testing.B) {
	assert.NotNil(b, AllTypeHTML)
	for i := 0; i < b.N; i++ {
		partTypes := PartTypesStruct{}
		assert.Nil(b, Unmarshal(AllTypeHTML, &partTypes))
	}
}

func BenchmarkParseCoursesLogically(b *testing.B) {
	assert.NotNil(b, CourseHTML)
	for i := 0; i < b.N; i++ {
		_, err := parseCoursesLogically()
		assert.Nil(b, err)
	}
}

func BenchmarkParsePartTypesLogically(b *testing.B) {
	assert.NotNil(b, AllTypeHTML)
	for i := 0; i < b.N; i++ {
		_, err := parsePartTypesLogically()
		assert.Nil(b, err)
	}
}

func getLink(selection *goquery.Selection) Link {
	link, _ := selection.Attr(AttrHref)
	return Link{Text: selection.Text(), Href: link}
}

func parseCoursesLogically() (Courses, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(CourseHTML))
	courses := make(Courses, 0)
	if err == nil {
		doc.Find(courses.Root()).Each(func(i int, selection *goquery.Selection) {
			course := Course{}
			course.Code = getLink(selection.Find(`td:nth-child(1) > a`))
			course.Name = getLink(selection.Find(`td:nth-child(2) > a`))
			course.Teacher = getLink(selection.Find(`td:nth-child(3) > a`))
			course.Semester = selection.Find(`td:nth-child(4)`).Text()
			course.Time = selection.Find(`td:nth-child(5)`).Text()
			course.Location = selection.Find(`td:nth-child(6)`).Text()
			courses = append(courses, course)
		})
	}

	return courses, err
}

func parsePartTypesLogically() (PartTypesStruct, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(AllTypeHTML))
	partTypes := PartTypesStruct{}
	if err == nil {
		selection := doc.Find(partTypes.Root())
		partTypes.Slice = make([]int, 0)
		selection.Find(`ul > li`).Each(func(i int, selection *goquery.Selection) {
			Int, parseErr := strconv.Atoi(selection.Text())
			if parseErr != nil {
				err = parseErr
			}
			partTypes.Slice = append(partTypes.Slice, Int)
		})
		if err == nil {
			partTypes.Struct.Name = selection.Find(`#test > div > p:nth-child(1)`).Text()
			Int, parseErr := strconv.Atoi(selection.Find(`#test > div > p:nth-child(2)`).Text())
			if err = parseErr; err == nil {
				partTypes.Struct.Age = uint(Int)
				Bool, parseErr := strconv.ParseBool(selection.Find(`#test > div > p:nth-child(3)`).Text())
				if err = parseErr; err == nil {
					partTypes.Struct.LikeLemon = Bool

					String := selection.Find(`#test > p:nth-child(3)`).Text()
					Int, parseErr := strconv.Atoi(selection.Find(`#test > p:nth-child(4)`).Text())
					if err = parseErr; err != nil {
						return partTypes, err
					}

					Float64, parseErr := strconv.ParseFloat(selection.Find(`#test > p:nth-child(5)`).Text(), 0)
					if err = parseErr; err != nil {
						return partTypes, err
					}

					Bool, parseErr := strconv.ParseBool(selection.Find(`#test > p:nth-child(6)`).Text())
					if err = parseErr; err != nil {
						return partTypes, err
					}
					partTypes.String = String
					partTypes.Int = Int
					partTypes.Float64 = Float64
					partTypes.Bool = Bool
				}
			}
		}
	}

	return partTypes, err
}
