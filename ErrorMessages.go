package unhtml

import (
	"reflect"
)

const (
	UnmarshaledKindMustBePtr = "unmarshaled kind must be Ptr"
	UnmarshalerItemKind      = "unmarshaled elem cannot be Ptr/Uintptr/Interface/Chan/Func/"
	DtoZero                  = "dto cannot be zero"
	SelectionNil             = "selection cannot be nil"
	ConverterNotExist        = "converter not exist"
)

type (
	UnmarshaledKindMustBePtrError struct {
		Type reflect.Type
	}

	UnmarshalerItemKindError struct {
		Type reflect.Type
	}

	ConverterNotExistError struct {
		Name string
	}
)

func (err UnmarshaledKindMustBePtrError) Error() string {
	return UnmarshaledKindMustBePtr + ": " + err.Type.String()
}

func (err UnmarshalerItemKindError) Error() string {
	return UnmarshalerItemKind + ": " + err.Type.String()
}

func (err ConverterNotExistError) Error() string {
	return UnmarshalerItemKind + ": " + err.Name
}
