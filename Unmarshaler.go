package gotten

type (
	Unmarshaler interface {
		Unmarshal(data []byte, v interface{}) error
	}
)
