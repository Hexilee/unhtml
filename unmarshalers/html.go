package unmarshalers

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strconv"
)

type (
	HTMLUnmarshaler struct {
		dto         reflect.Value
		kind        reflect.Kind
		dtoElemType reflect.Type
		selector    string
		attrKey     string
	}

	HTMLModel interface {
		Root() string
	}
)

const (
	SelectorKey    = "html"
	AttrKey        = "key"
	DefaultAttrKey = AttrText
)

const (
	AttrHref = "href"
	AttrText = "text"
	AttrSrc  = "src"
)

func NewHTMLUnmarshaler(tag reflect.StructTag) *HTMLUnmarshaler {
	selector := tag.Get(SelectorKey)
	attrKey := tag.Get(AttrKey)

	return &HTMLUnmarshaler{selector: selector, attrKey: attrKey}
}

func (marshaler *HTMLUnmarshaler) initDto(v reflect.Value) error {
	marshaler.dto = v
	return marshaler.parseType()
}

func (marshaler *HTMLUnmarshaler) setSelector(selector string) {
	marshaler.selector = selector
}

func (marshaler *HTMLUnmarshaler) setAttrKey(attrKey string) {
	marshaler.attrKey = attrKey
}

func (marshaler *HTMLUnmarshaler) getSelector() string {
	return marshaler.selector
}

func (marshaler *HTMLUnmarshaler) getAttrKey() string {
	result := marshaler.attrKey
	if result == "" {
		result = DefaultAttrKey
	}
	return result
}

func (marshaler *HTMLUnmarshaler) getDto() reflect.Value {
	return marshaler.dto
}

func (marshaler *HTMLUnmarshaler) getKind() reflect.Kind {
	return marshaler.kind
}

func (marshaler *HTMLUnmarshaler) getDtoElemType() reflect.Type {
	return marshaler.dtoElemType
}

func (marshaler *HTMLUnmarshaler) Unmarshal(data []byte, v interface{}) (err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err == nil {
		marshaler.initRoot(v)
		err = marshaler.unmarshal(doc.Selection, reflect.ValueOf(v))
	}
	return err
}

func (marshaler *HTMLUnmarshaler) unmarshal(selection *goquery.Selection, value reflect.Value) (err error) {
	err = marshaler.initDto(value)
	if err == nil {
		preSelection := selection
		if marshaler.getKind() == reflect.Slice {
			itemType := marshaler.getDtoElemType().Elem()
			isItemTypePtr := itemType.Kind() == reflect.Ptr
			sliceValue := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)

			if isItemTypePtr {
				itemType = itemType.Elem()
			}
			preSelection.Find(marshaler.getSelector()).Each(func(i int, selection *goquery.Selection) {
				newItem := reflect.New(itemType)
				parseErr := NewHTMLUnmarshaler("").unmarshal(selection, newItem)
				if parseErr != nil {
					err = parseErr
				} else {
					if !isItemTypePtr {
						newItem = newItem.Elem()
					}
					sliceValue = reflect.Append(sliceValue, newItem)
				}
			})
			marshaler.getDto().Elem().Set(sliceValue)
		}

		if err == nil {
			if marshaler.getSelector() != "" {
				selection.Find(marshaler.getSelector()).Each(func(i int, selection *goquery.Selection) {
					if i == 0 {
						preSelection = selection
					}
				})
			}

			switch marshaler.getKind() {
			case reflect.Struct:
				motherValue := marshaler.getDto().Elem()
				motherType := marshaler.getDtoElemType()
				for i := 0; i < motherValue.NumField(); i++ {
					fieldPtr := motherValue.Field(i).Addr()
					tag := motherType.Field(i).Tag
					err = NewHTMLUnmarshaler(tag).unmarshal(preSelection, fieldPtr)
					if err != nil {
						break
					}
				}
			case reflect.String:
				marshaler.getDto().Elem().SetString(marshaler.getAttrValue(preSelection))
			case reflect.Int:
				fallthrough
			case reflect.Int8:
				fallthrough
			case reflect.Int16:
				fallthrough
			case reflect.Int32:
				fallthrough
			case reflect.Int64:
				valueStr := marshaler.getAttrValue(preSelection)
				value, err := strconv.Atoi(valueStr)
				if err == nil {
					marshaler.getDto().Elem().SetInt(int64(value))
				}
			case reflect.Uintptr:
				fallthrough
			case reflect.Uint:
				fallthrough
			case reflect.Uint8:
				fallthrough
			case reflect.Uint16:
				fallthrough
			case reflect.Uint32:
				fallthrough
			case reflect.Uint64:
				valueStr := marshaler.getAttrValue(preSelection)
				value, err := strconv.ParseUint(valueStr, 0, 0)
				if err == nil {
					marshaler.getDto().Elem().SetUint(value)
				}
			case reflect.Float32:
				fallthrough
			case reflect.Float64:
				valueStr := marshaler.getAttrValue(preSelection)
				value, err := strconv.ParseFloat(valueStr, 0)
				if err == nil {
					marshaler.getDto().Elem().SetFloat(value)
				}
			case reflect.Bool:
				valueStr := marshaler.getAttrValue(preSelection)
				value, err := strconv.ParseBool(valueStr)
				if err == nil {
					marshaler.getDto().Elem().SetBool(value)
				}
			}
		}
	}

	return err
}

func (marshaler *HTMLUnmarshaler) initRoot(v interface{}) {
	if value, ok := v.(HTMLModel); ok {
		marshaler.setSelector(value.Root())
	}
}

func (marshaler *HTMLUnmarshaler) parseType() (err error) {
	dtoType := marshaler.getDto().Type()
	switch dtoType.Kind() {
	case reflect.Ptr:
		marshaler.dtoElemType = dtoType.Elem()
		marshaler.kind = marshaler.getDtoElemType().Kind()
		err = marshaler.checkItemKind()
	default:
		err = errors.New(UnmarshaledKindMustBePtr)
	}
	return
}

func (marshaler *HTMLUnmarshaler) checkItemKind() (err error) {
	switch marshaler.getKind() {
	case reflect.Ptr:
	case reflect.Interface:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Map:
	default:
		return
	}
	err = errors.New(UnmarshalerItemKindError)
	return
}

func (marshaler *HTMLUnmarshaler) getAttrValue(selection *goquery.Selection) (valueStr string) {
	if marshaler.getAttrKey() == AttrText {
		valueStr = selection.Text()
	} else {
		if str, exist := selection.Attr(marshaler.getAttrKey()); exist {
			valueStr = str
		}
	}
	return
}
