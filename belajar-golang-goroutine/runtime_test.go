package belajargolanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// kejadiannya hanya sekali dijalankan
func TestGomaxprocs(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("total cpu: ", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread: ", totalThread)
	// mengubah, tapi jarang digunakan
	runtime.GOMAXPROCS(20)
	totalThread2 := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread ubah: ", totalThread2)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine: ", totalGoroutine)
}
