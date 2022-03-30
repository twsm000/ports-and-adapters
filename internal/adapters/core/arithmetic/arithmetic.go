package arithmetic

import (
	"fmt"
)

func NewAdapter() *Adapter {
	return new(Adapter)
}

type Adapter struct{}

func (adp Adapter) Add(a int32, b int32) (int32, error)   { return a + b, nil }
func (adp Adapter) Sub(a int32, b int32) (int32, error)   { return a - b, nil }
func (adp Adapter) Multi(a int32, b int32) (int32, error) { return a * b, nil }
func (adp Adapter) Div(a int32, b int32) (int32, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
