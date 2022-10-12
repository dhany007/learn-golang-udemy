package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

// once = hanya bisa dieksekusi sekali
func TestOnce(t *testing.T) {
	var group sync.WaitGroup
	var once sync.Once

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			once.Do(OnlyOnce) // hanya akan dieksekusi sekali
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
