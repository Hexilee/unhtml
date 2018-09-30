package unhtml

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strconv"
)

type (
	// HTMLUnmarshalerBuilder: all methods is private
	HTMLUnmarshalerBuilder struct {
		dto         reflect.Value
		kind        reflect.Kind
		dtoElemType reflect.Type
		selection   *goquery.Selection
		selector    string
		attrKey     string
	}

	// HTMLUnmarshaler: all methods is private
	HTMLUnmarshaler struct {
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
	SelectorKey  = "html"
	AttrKey      = "key"
	ConverterKey = "converter"
	ZeroInt      = 0
	ZeroStr      = ""
)

const (
	AttrHref = "href"
	AttrSrc  = "src"
)

func Unmarshal(data []byte, v interface{}) error {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err == nil {
		err = unmarshal(reflect.ValueOf(v), *doc.Selection, "")
	}
	return err
}

func unmarshal(ptr reflect.Value, selection goquery.Selection, tag reflect.StructTag) (err error) {
	realUnmarshal, buildErr := new(HTMLUnmarshalerBuilder).
		setDto(ptr).
		setSelection(&selection).
		setSelector(tag.Get(SelectorKey)).
		setAttrKey(tag.Get(AttrKey)).
		build()

	if err = buildErr; err == nil {
		err = realUnmarshal.unmarshal()
	}
	return err
}

func (builder *HTMLUnmarshalerBuilder) build() (unmarshaler *HTMLUnmarshaler, err error) {
	if err = builder.initRoot(); err == nil {
		if err = builder.parseType(); err == nil {
			if err = builder.checkBeforeReturn(); err == nil {
				unmarshaler = &HTMLUnmarshaler{
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

func (builder *HTMLUnmarshalerBuilder) setDto(v reflect.Value) *HTMLUnmarshalerBuilder {
	builder.dto = v
	return builder
}

func (builder *HTMLUnmarshalerBuilder) setSelector(selector string) *HTMLUnmarshalerBuilder {
	builder.selector = selector
	return builder
}

func (builder *HTMLUnmarshalerBuilder) setAttrKey(attrKey string) *HTMLUnmarshalerBuilder {
	builder.attrKey = attrKey
	return builder
}

func (builder *HTMLUnmarshalerBuilder) setSelection(selection *goquery.Selection) *HTMLUnmarshalerBuilder {
	builder.selection = selection
	return builder
}

func (builder *HTMLUnmarshalerBuilder) initRoot() (err error) {
	//if err = builder.checkDtoZero(); err == nil {
	if value, ok := builder.dto.Interface().(HTMLModel); ok {
		builder.selector = value.Root()
	}
	//}
	return
}

func (builder *HTMLUnmarshalerBuilder) parseType() (err error) {
	//if err = builder.checkDtoZero(); err == nil {
	dtoType := builder.dto.Type()
	switch dtoType.Kind() {
	case reflect.Ptr:
		builder.dtoElemType = dtoType.Elem()
		builder.kind = builder.dtoElemType.Kind()
	default:
		err = UnmarshaledKindMustBePtrError{dtoType}
	}
	//}

	return
}

func (builder *HTMLUnmarshalerBuilder) checkItemKind() (err error) {
	switch builder.kind {
	case reflect.Ptr:
		fallthrough
	case reflect.Uintptr:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Map:
		err = UnmarshalerItemKindError{builder.dtoElemType}
	default:
	}
	return
}

func (builder *HTMLUnmarshalerBuilder) checkBeforeReturn() (err error) {
	//if err = builder.checkDtoZero(); err == nil {
	//	if err = builder.checkSelectionNil(); err == nil {
	err = builder.checkItemKind()
	//}
	//}
	return
}

// never return err in production env
//func (builder *HTMLUnmarshalerBuilder) checkDtoZero() (err error) {
//	// Zero reflect.Value
//	if isZero(builder.dto) {
//		err = errors.New(DtoZero)
//	}
//	return
//}

// never return err in production env
//func (builder *HTMLUnmarshalerBuilder) checkSelectionNil() (err error) {
//	if builder.selection == nil {
//		err = errors.New(SelectionNil)
//	}
//	return
//}

func (marshaler HTMLUnmarshaler) getSelection() goquery.Selection {
	return marshaler.selection
}

func (marshaler HTMLUnmarshaler) getSelector() string {
	return marshaler.selector
}

func (marshaler HTMLUnmarshaler) getAttrKey() string {
	return marshaler.attrKey
}

func (marshaler HTMLUnmarshaler) getDto() reflect.Value {
	return marshaler.dto
}

func (marshaler HTMLUnmarshaler) getKind() reflect.Kind {
	return marshaler.kind
}

func (marshaler HTMLUnmarshaler) getDtoElemType() reflect.Type {
	return marshaler.dtoElemType
}

func (marshaler HTMLUnmarshaler) unmarshalSlice(preSelection goquery.Selection) (err error) {
	itemType := marshaler.getDtoElemType().Elem()
	sliceValue := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)
	preSelection.Each(func(i int, selection *goquery.Selection) {
		newItem := reflect.New(itemType)
		if err = unmarshal(newItem, *selection, ""); err == nil {
			sliceValue = reflect.Append(sliceValue, newItem.Elem())
		}
	})
	marshaler.getDto().Elem().Set(sliceValue)
	return err
}

func (marshaler HTMLUnmarshaler) callConverter(converter string, fieldIndex int, preSelection goquery.Selection) (result reflect.Value, err error) {
	motherValue := marshaler.getDto().Elem()
	motherType := marshaler.getDtoElemType()
	tag := motherType.Field(fieldIndex).Tag
	resultType := motherType.Field(fieldIndex).Type
	method, exist := motherType.MethodByName(converter)
	if !exist {
		err = &ConverterNotExistError{converter}
	}
	if err == nil {
		methodValue := motherValue.MethodByName(converter)
		inputValuePtr, converterTypeErr := checkConverter(method.Name, methodValue.Type(), resultType)
		if err = converterTypeErr; err == nil {
			if err = unmarshal(inputValuePtr, preSelection, tag); err == nil {
				results := methodValue.Call([]reflect.Value{inputValuePtr.Elem()})
				if errInterface := results[1].Interface(); errInterface != nil {
					err = errInterface.(error)
				}
				if err == nil {
					result = results[0]
				}
			}
		}
	}
	return
}

func (marshaler HTMLUnmarshaler) unmarshalStruct(preSelection goquery.Selection) (err error) {
	motherValue := marshaler.getDto().Elem()
	motherType := marshaler.getDtoElemType()
	for i := 0; i < motherValue.NumField(); i++ {
		fieldPtr := motherValue.Field(i).Addr()
		tag := motherType.Field(i).Tag
		if converter := tag.Get(ConverterKey); converter != ZeroStr {
			result, callConverterErr := marshaler.callConverter(converter, i, preSelection)
			if err = callConverterErr; err == nil {
				fieldPtr.Elem().Set(result)
			}
		} else {
			err = unmarshal(fieldPtr, preSelection, tag)
		}

		if err != nil {
			break
		}
	}
	return
}

func (marshaler HTMLUnmarshaler) unmarshal() (err error) {
	preSelection := marshaler.getSelection()
	if marshaler.getSelector() != ZeroStr {
		preSelection = *preSelection.Find(marshaler.getSelector())
	}
	switch marshaler.getKind() {
	case reflect.Slice:
		err = marshaler.unmarshalSlice(preSelection)
	case reflect.Struct:
		err = marshaler.unmarshalStruct(preSelection)
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

	return err
}

func (marshaler HTMLUnmarshaler) getAttrValue(selection goquery.Selection) (valueStr string) {
	if marshaler.getAttrKey() == ZeroStr {
		valueStr = selection.Text()
	} else {
		if str, exist := selection.Attr(marshaler.getAttrKey()); exist {
			valueStr = str
		}
	}
	return
}
