package generic

import (
	"fmt"
	"testing"
)

type Data2[T any] struct {
	First  T
	Second T
}

// didalam constraint method struct nya wajib dibuat _ jika constraint tidak dipakai
func (d *Data2[_]) SayHello(name string) string {
	return "Hello " + name
}

// didalam constraint nya wajib karena akan dipakai di param dan return method nya
func (d *Data2[T]) Changefirst(first T) T {
	d.First = first
	return first
}

func TestGenericMethodStruct(t *testing.T) {
	data := Data2[string]{First: "Dhany", Second: "Aritonang"}
	fmt.Println("before", data)

	newFirst := "Kalai"
	fmt.Println(data.SayHello(newFirst))
	data.Changefirst(newFirst)
	fmt.Println("after", data)
}
