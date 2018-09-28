package unmarshalers

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"reflect"
)

type (
	HTMLMarshaler struct {
		dto      interface{}
		isList   bool
		isPtr    bool
		itemType reflect.Type
	}

	HTMLModel interface {
		Root() string
	}
)

const (
	HTMLRootNode = "html"
	SelectorKey  = "html"
	RangeKey     = "range"
	AttrKey      = "key"
)

func (marshaler *HTMLMarshaler) setDto(v interface{}) {
	marshaler.dto = v
}

func (marshaler *HTMLMarshaler) getDto() interface{} {
	return marshaler.dto
}

func (marshaler *HTMLMarshaler) Unmarshal(data []byte, v interface{}) (err error) {
	marshaler.setDto(v)
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err == nil {
		doc.Find(marshaler.getRoot()).Each(func(i int, selection *goquery.Selection) {
			_ = selection
		})
	}
	return err
}

func (marshaler *HTMLMarshaler) getRoot() string {
	root := HTMLRootNode
	if value, ok := marshaler.dto.(HTMLModel); ok {
		root = value.Root()
	}
	return root
}

func (marshaler *HTMLMarshaler) parseType() (err error) {
	dtoType := reflect.TypeOf(marshaler.dto)
	switch dtoType.Kind() {
	case reflect.Slice:
		marshaler.isList = true
		marshaler.itemType = dtoType.Elem()
		if marshaler.itemType.Kind() == reflect.Ptr {
			marshaler.itemType = marshaler.itemType.Elem()
			marshaler.isPtr = true
		}
	case reflect.Ptr:
		marshaler.itemType = dtoType.Elem()
		marshaler.isPtr = true
	default:
		err = errors.New(UnmarshaledKindMustBePtrOrSlice)
	}
	if err == nil {
		err = marshaler.checkItemType()
	}
	return
}

func (marshaler *HTMLMarshaler) checkItemType() (err error) {
	if marshaler.itemType.Kind() == reflect.Ptr {
		err = errors.New(UnmarshalerItemTypeCannotBePtr)
	}
	return
}
