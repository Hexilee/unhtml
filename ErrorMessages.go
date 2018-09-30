package unhtml

import (
	"reflect"
)

const (
	UnmarshaledKindMustBePtr = "unmarshaled kind must be Ptr"
	UnmarshalerItemKind      = "unmarshaled elem cannot be Ptr/Interface/Chan/Func/"
	DtoZero                  = "dto cannot be zero"
	SelectionNil             = "selection cannot be nil"
)

type (
	UnmarshaledKindMustBePtrError struct {
		Type reflect.Type
	}

	UnmarshalerItemKindError struct {
		Type reflect.Type
	}
)

func (err UnmarshaledKindMustBePtrError) Error() string {
	return UnmarshaledKindMustBePtr + ": " + err.Type.Name()
}

func (err UnmarshalerItemKindError) Error() string {
	return UnmarshalerItemKind + ": " + err.Type.Name()
}
