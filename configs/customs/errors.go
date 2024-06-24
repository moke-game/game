package customs

import "fmt"

// ErrInvalidKeyType .
type ErrInvalidKeyType struct {
	Key      string
	Expected string
	Actual   string
}

// NewErrInvalidKeyType .
func NewErrInvalidKeyType(key, expected, actual string) *ErrInvalidKeyType {
	return &ErrInvalidKeyType{
		Key:      key,
		Expected: expected,
		Actual:   actual,
	}
}

func (e *ErrInvalidKeyType) Error() string {
	return fmt.Sprintf("invalid key type! key: \"%s\", expected: %s, actual: %s", e.Key, e.Expected, e.Actual)
}
