package gotten

import "io"

type (
	PathVar interface {
		String() string
	}
	Query interface {
		String() string
	}
	Part interface {
		String() string
	}

	PathStr string

	PathInt int

	QueryStr string

	QueryInt int

	PartStr string

	PartInt int

	PartReader io.Reader

	PartFile string
)
