package unmarshalers

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestHTMLMarshaler_parseType(t *testing.T) {
	for _, testCase := range [] *struct {
		dto      interface{}
		isList   bool
		isPtr    bool
		itemType reflect.Type
		err      error
	}{
		{make([]int, 0), true, false, reflect.TypeOf(int(1)), nil},
		{make([]*int, 0), true, true, reflect.TypeOf(int(1)), nil},
		{make([]**int, 0), true, true, reflect.TypeOf(new(int)), errors.New(UnmarshalerItemTypeCannotBePtr)},
		{int(0), false, false, nil, errors.New(UnmarshaledKindMustBePtrOrSlice)},
		{new(int), false, true, reflect.TypeOf(int(1)), nil},
		{make([]**int, 1)[0], false, true, reflect.TypeOf(new(int)), errors.New(UnmarshalerItemTypeCannotBePtr)},
	} {
		func() {
			result := new(HTMLMarshaler)
			result.setDto(testCase.dto)
			parseErr := result.parseType()
			if parseErr != nil && testCase.err != nil {
				assert.Equal(t, testCase.err.Error(), parseErr.Error())
			} else {
				assert.Equal(t, testCase.err, parseErr)
			}
			assert.Equal(t, testCase.isList, result.isList)
			assert.Equal(t, testCase.itemType, result.itemType)
		}()
	}
}
