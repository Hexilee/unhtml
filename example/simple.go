package example

import (
	"github.com/Hexilee/gotten"
	"net/http"
	"time"
)

type (
	SimpleParams struct {
		Id   gotten.PathInt
		Page gotten.QueryInt
	}

	SimpleResult struct {
		Status       gotten.StatusCode `expect:"200"`
		ExpectResult ExpectResult      `expect:"200"`
		BadResult    ObjectNotFound    `expect:"404"`
	}

	Item struct {
		TypeId      int
		IId         int
		Name        string
		Description string
	}

	ExpectResult []*Item
	ObjectNotFound struct {
		Key         string
		Reason      string
		Description string
	}

	SimpleService struct {
		GetItems func(SimpleParams) (SimpleResult, error) `method:"GET";path:"itemType/{id}"`
	}
)

var (
	creator = gotten.NewBuilder().
		SetBaseUrl("https://api.sample.com").
		AddCookie(&http.Cookie{Name: "clientcookieid", Value: "121", Expires: time.Now().Add(111 * time.Second)}).
		Build()

	simpleServiceImpl = new(SimpleService)
)

func init() {
	err := creator.Impl(simpleServiceImpl)
	if err != nil {
		panic(err)
	}
}
