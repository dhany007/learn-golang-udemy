package generic

import (
	"fmt"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

// jika menggunakan generic pada interface, maka pada struct nya juga wajib menggunakan generic
// mengikuti contract nya
func TestGenericInterface(t *testing.T) {
	data := MyData[string]{Value: "Dhany"}
	fmt.Println("before", data)
	result := ChangeValue[string](&data, "Kalai")
	fmt.Println("after", data)
	fmt.Println(result)
}
