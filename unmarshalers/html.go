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
	}

	RealRealHTMLUnmarshalerBuilder struct {
		dto         reflect.Value
		kind        reflect.Kind
		dtoElemType reflect.Type
		selection   *goquery.Selection
		selector    string
		attrKey     string
	}

	RealHTMLUnmarshaler struct {
		dto         reflect.Value
		kind        reflect.Kind
		dtoElemType reflect.Type
		selection   goquery.Selection
		selector    string
		attrKey     string
	}

	HTMLModel interface {
		Root() string
	}
)

const (
	SelectorKey = "html"
	AttrKey     = "key"
	ZeroInt     = 0
	ZeroStr     = ""
)

const (
	AttrHref = "href"
	AttrSrc  = "src"
)

func (builder *RealRealHTMLUnmarshalerBuilder) Build() (unmarshaler *RealHTMLUnmarshaler, err error) {
	if err = builder.initRoot(); err == nil {
		if err = builder.parseType(); err == nil {
			if err = builder.checkBeforeReturn(); err == nil {
				unmarshaler = &RealHTMLUnmarshaler{
					dto:         builder.dto,
					kind:        builder.kind,
					dtoElemType: builder.dtoElemType,
					selection:   *builder.selection,
					selector:    builder.selector,
					attrKey:     builder.attrKey,
				}
			}
		}
	}
	return
}

func (builder *RealRealHTMLUnmarshalerBuilder) setDto(v reflect.Value) *RealRealHTMLUnmarshalerBuilder {
	builder.dto = v
	return builder
}

func (builder *RealRealHTMLUnmarshalerBuilder) setSelector(selector string) *RealRealHTMLUnmarshalerBuilder {
	builder.selector = selector
	return builder
}

func (builder *RealRealHTMLUnmarshalerBuilder) setAttrKey(attrKey string) *RealRealHTMLUnmarshalerBuilder {
	builder.attrKey = attrKey
	return builder
}

func (builder *RealRealHTMLUnmarshalerBuilder) setSelection(selection *goquery.Selection) *RealRealHTMLUnmarshalerBuilder {
	builder.selection = selection
	return builder
}

func (builder *RealRealHTMLUnmarshalerBuilder) initRoot() (err error) {
	if err = builder.checkDtoZero(); err == nil {
		if value, ok := builder.dto.Interface().(HTMLModel); ok {
			builder.selector = value.Root()
		}
	}
	return
}

func (builder *RealRealHTMLUnmarshalerBuilder) parseType() (err error) {
	if err = builder.checkDtoZero(); err == nil {
		dtoType := builder.dto.Type()
		switch dtoType.Kind() {
		case reflect.Ptr:
			builder.dtoElemType = dtoType.Elem()
			builder.kind = builder.dtoElemType.Kind()
		default:
			err = errors.New(UnmarshaledKindMustBePtr)
		}
	}

	return
}

func (builder *RealRealHTMLUnmarshalerBuilder) checkItemKind() error {
	err := errors.New(UnmarshalerItemKindError)
	switch builder.kind {
	case reflect.Ptr:
	case reflect.Uintptr:
	case reflect.Interface:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Map:
	default:
		err = nil
	}
	return err
}

func (builder *RealRealHTMLUnmarshalerBuilder) checkBeforeReturn() (err error) {
	if err = builder.checkDtoZero(); err == nil {
		if err = builder.checkItemKind(); err == nil {
			err = builder.checkSelectionZero()
		}
	}
	return
}

func (builder *RealRealHTMLUnmarshalerBuilder) checkDtoZero() (err error) {
	if IsZero(builder.dto) {
		err = errors.New(DtoZero)
	}
	return
}

func (builder *RealRealHTMLUnmarshalerBuilder) checkSelectionZero() (err error) {
	if IsZero(builder.selection) {
		err = errors.New(SelectionZero)
	}
	return
}

func (marshaler *RealHTMLUnmarshaler) getSelection() goquery.Selection {
	return marshaler.selection
}

func (marshaler *RealHTMLUnmarshaler) getSelector() string {
	return marshaler.selector
}

func (marshaler *RealHTMLUnmarshaler) getAttrKey() string {
	return marshaler.attrKey
}

func (marshaler *RealHTMLUnmarshaler) getDto() reflect.Value {
	return marshaler.dto
}

func (marshaler *RealHTMLUnmarshaler) getKind() reflect.Kind {
	return marshaler.kind
}

func (marshaler *RealHTMLUnmarshaler) getDtoElemType() reflect.Type {
	return marshaler.dtoElemType
}

func (HTMLUnmarshaler) Unmarshal(data []byte, v interface{}) error {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err == nil {
		realUnmarshal, buildErr := new(RealRealHTMLUnmarshalerBuilder).
			setDto(reflect.ValueOf(v)).
			setSelection(doc.Selection).
			Build()
		if err = buildErr; err == nil {
			err = realUnmarshal.unmarshal()
		}
	}
	return err
}

func (marshaler *RealHTMLUnmarshaler) unmarshal() (err error) {
	preSelection := *marshaler.selection.Clone()
	if marshaler.getKind() == reflect.Slice {
		itemType := marshaler.getDtoElemType().Elem()
		isItemTypePtr := itemType.Kind() == reflect.Ptr
		sliceValue := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)

		if isItemTypePtr {
			itemType = itemType.Elem()
		}
		preSelection.Find(marshaler.getSelector()).Each(func(i int, selection *goquery.Selection) {
			newItem := reflect.New(itemType)
			newUnmarshaler, buildErr := new(RealRealHTMLUnmarshalerBuilder).
				setDto(newItem).
				setSelection(selection).
				Build()
			if err = buildErr; err == nil {
				if err = newUnmarshaler.unmarshal(); err == nil {
					if !isItemTypePtr {
						newItem = newItem.Elem()
					}
					sliceValue = reflect.Append(sliceValue, newItem)
				}
			}
		})
		marshaler.getDto().Elem().Set(sliceValue)
	}

	if err == nil {
		if marshaler.getSelector() != ZeroStr {
			marshaler.selection.Find(marshaler.getSelector()).Each(func(i int, newSelection *goquery.Selection) {
				if i == ZeroInt {
					preSelection = *newSelection
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
				newUnmarshal, buildErr := new(RealRealHTMLUnmarshalerBuilder).
					setDto(fieldPtr).
					setSelection(&preSelection).
					setSelector(tag.Get(SelectorKey)).
					setAttrKey(tag.Get(AttrKey)).
					Build()
				if err = buildErr; err != nil {
					break
				}
				if err = newUnmarshal.unmarshal(); err != nil {
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

	return err
}

func (marshaler *RealHTMLUnmarshaler) getAttrValue(selection goquery.Selection) (valueStr string) {
	if marshaler.getAttrKey() == ZeroStr {
		valueStr = selection.Text()
	} else {
		if str, exist := selection.Attr(marshaler.getAttrKey()); exist {
			valueStr = str
		}
	}
	return
}
