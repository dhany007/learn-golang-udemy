package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait() // menunggu cond supaya dilanjutkan kebawah
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestSyncCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		// perulangan sama diatas
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second) // sleep 1 sc untuk melihat proses
			cond.Signal()               // lanjut ke goroutine berikutnya
		}
	}()

	// go func() {
	// 	// perulangan sama diatas
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second) // sleep 1 sc untuk melihat semua proses
	// 		cond.Broadcast()            // ini melanjutkan semua goroutine
	// 	}
	// }()

	group.Wait()
}
