package unhtml

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkUnmarshalCourses(b *testing.B) {
	assert.NotNil(b, CourseHTML)
	for i := 0; i < b.N; i++ {
		courses := make(Courses, 0)
		assert.Nil(b, Unmarshal(CourseHTML, &courses))
	}
}

func getLink(selection *goquery.Selection) Link {
	link, _ := selection.Attr(AttrHref)
	return Link{Text: selection.Text(), Href: link}
}

func parseHTMLLogically() (Courses, error) {
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

func BenchmarkParseCoursesLogically(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := parseHTMLLogically()
		assert.Nil(b, err)
	}
}
