package unhtml

import (
	"fmt"
	"reflect"
)

const (
	UnmarshaledKindMustBePtr = "unmarshaled kind must be Ptr"
	UnmarshalerItemKind      = "unmarshaled elem cannot be Ptr/Uintptr/Interface/Chan/Func/"
	DtoZero                  = "dto cannot be zero"
	SelectionNil             = "selection cannot be nil"
	ConverterNotExist        = "converter not exist"
	ConverterTypeWrong       = "type of converter is wrong"
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

	ConverterTypeWrongError struct {
		name       string
		methodType reflect.Type
	}
)

func NewConverterTypeWrongError(name string, methodType reflect.Type) *ConverterTypeWrongError {
	return &ConverterTypeWrongError{name, methodType}
}

func (err UnmarshaledKindMustBePtrError) Error() string {
	return UnmarshaledKindMustBePtr + ": " + err.Type.String()
}

func (err UnmarshalerItemKindError) Error() string {
	return UnmarshalerItemKind + ": " + err.Type.String()
}

func (err ConverterNotExistError) Error() string {
	return ConverterNotExist + ": " + err.Name
}

func (err ConverterTypeWrongError) Error() string {
	return fmt.Sprintf(ConverterTypeWrong+"(%s): %s", err.name, err.methodType)
}
