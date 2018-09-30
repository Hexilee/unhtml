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
		dtoType reflect.Type
	}

	UnmarshalerItemKindError struct {
		dtoType reflect.Type
	}

	ConverterNotExistError struct {
		name string
	}

	ConverterTypeWrongError struct {
		name       string
		methodType reflect.Type
	}
)

func NewUnmarshaledKindMustBePtrError(dtoType reflect.Type) *UnmarshaledKindMustBePtrError {
	return &UnmarshaledKindMustBePtrError{dtoType}
}

func NewUnmarshalerItemKindError(dtoType reflect.Type) *UnmarshalerItemKindError {
	return &UnmarshalerItemKindError{dtoType}
}

func NewConverterNotExistError(name string) *ConverterNotExistError {
	return &ConverterNotExistError{name}
}

func NewConverterTypeWrongError(name string, methodType reflect.Type) *ConverterTypeWrongError {
	return &ConverterTypeWrongError{name, methodType}
}

func (err UnmarshaledKindMustBePtrError) Error() string {
	return UnmarshaledKindMustBePtr + ": " + err.dtoType.String()
}

func (err UnmarshalerItemKindError) Error() string {
	return UnmarshalerItemKind + ": " + err.dtoType.String()
}

func (err ConverterNotExistError) Error() string {
	return ConverterNotExist + ": " + err.name
}

func (err ConverterTypeWrongError) Error() string {
	return fmt.Sprintf(ConverterTypeWrong+"(%s): %s", err.name, err.methodType)
}
