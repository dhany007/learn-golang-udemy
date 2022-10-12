package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // group nya selesai

	group.Add(1) // menambahkan 1 proses goroutine ke group

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsynchronous(group) // menjalan kan group nya
	}

	group.Wait() // menunggu semua group selesai
	fmt.Println("Complete")
}
