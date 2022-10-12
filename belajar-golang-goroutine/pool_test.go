package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// once = hanya bisa dieksekusi sekali
func TestPool(t *testing.T) {
	pool := sync.Pool{
		// membuat nilai default dari pool
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Dhany")
	pool.Put("Putra")
	pool.Put("Aritonang")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
