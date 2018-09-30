package unhtml

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func BenchmarkUnmarshalCourses(b *testing.B) {
	assert.NotNil(b, CourseHTML)
	for i := 0; i < b.N; i++ {
		courses := make(Courses, 0)
		assert.Nil(b, Unmarshal(CourseHTML, &courses))
	}
}

func BenchmarkUnmarshalAllTypes(b *testing.B) {
	assert.NotNil(b, AllTypeHTML)
	for i := 0; i < b.N; i++ {
		allTypes := AllTypeTest{}
		assert.Nil(b, Unmarshal(AllTypeHTML, &allTypes))
	}
}

func BenchmarkParseCoursesLogically(b *testing.B) {
	assert.NotNil(b, CourseHTML)
	for i := 0; i < b.N; i++ {
		_, err := parseCoursesLogically()
		assert.Nil(b, err)
	}
}

func BenchmarkParseAllTypesLogically(b *testing.B) {
	assert.NotNil(b, AllTypeHTML)
	for i := 0; i < b.N; i++ {
		_, err := parseAllTypesLogically()
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

func parseAllTypesLogically() (AllTypeTest, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(AllTypeHTML))
	allTypes := AllTypeTest{}
	if err == nil {
		selection := doc.Find(allTypes.Root())
		allTypes.Slice = make([]int, 0)
		selection.Find(`ul > li`).Each(func(i int, selection *goquery.Selection) {
			Int, parseErr := strconv.Atoi(selection.Text())
			if parseErr != nil {
				err = parseErr
			}
			allTypes.Slice = append(allTypes.Slice, Int)
		})
		if err == nil {
			allTypes.Struct.Name = selection.Find(`#test > div > p:nth-child(1)`).Text()
			Int, parseErr := strconv.Atoi(selection.Find(`#test > div > p:nth-child(2)`).Text())
			if err = parseErr; err == nil {
				allTypes.Struct.Age = uint(Int)
				Bool, parseErr := strconv.ParseBool(selection.Find(`#test > div > p:nth-child(3)`).Text())
				if err = parseErr; err == nil {
					allTypes.Struct.LikeLemon = Bool

					String := selection.Find(`#test > p:nth-child(3)"`).Text()
					Int, parseErr := strconv.Atoi(selection.Find(`#test > p:nth-child(4)`).Text())

					if err = parseErr; err != nil {
						return allTypes, err
					}
					Int8, parseErr := strconv.Atoi(selection.Find(`#test > p:nth-child(4)`).Text())
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Int16, parseErr := strconv.Atoi(selection.Find(`#test > p:nth-child(4)`).Text())
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Int32, parseErr := strconv.Atoi(selection.Find(`#test > p:nth-child(4)`).Text())
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Int64, parseErr := strconv.Atoi(selection.Find(`#test > p:nth-child(4)`).Text())
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Uint, parseErr := strconv.ParseUint(selection.Find(`#test > p:nth-child(4)`).Text(), 0, 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Uint8, parseErr := strconv.ParseUint(selection.Find(`#test > p:nth-child(4)`).Text(), 0, 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Uint16, parseErr := strconv.ParseUint(selection.Find(`#test > p:nth-child(4)`).Text(), 0, 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Uint32, parseErr := strconv.ParseUint(selection.Find(`#test > p:nth-child(4)`).Text(), 0, 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}
					Uint64, parseErr := strconv.ParseUint(selection.Find(`#test > p:nth-child(4)`).Text(), 0, 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}

					Float32, parseErr := strconv.ParseFloat(selection.Find(`#test > p:nth-child(5)`).Text(), 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}

					Float64, parseErr := strconv.ParseFloat(selection.Find(`#test > p:nth-child(5)`).Text(), 0)
					if err = parseErr; err != nil {
						return allTypes, err
					}

					Bool, parseErr := strconv.ParseBool(selection.Find(`#test > p:nth-child(6)`).Text())
					if err = parseErr; err != nil {
						return allTypes, err
					}
					allTypes.String = String
					allTypes.Int = Int
					allTypes.Int8 = int8(Int8)
					allTypes.Int16 = int16(Int16)
					allTypes.Int32 = int32(Int32)
					allTypes.Int64 = int64(Int64)
					allTypes.Uint = uint(Uint)
					allTypes.Uint8 = uint8(Uint8)
					allTypes.Uint16 = uint16(Uint16)
					allTypes.Uint32 = uint32(Uint32)
					allTypes.Uint64 = uint64(Uint64)
					allTypes.Float32 = float32(Float32)
					allTypes.Float64 = Float64
					allTypes.Bool = Bool

				}
			}
		}
	}

	return allTypes, err
}
