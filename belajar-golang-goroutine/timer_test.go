package belajargolanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C // mengasign channelnya dari timer
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel // mengasign channelnya langsung
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)

	// fungsi yang dijalankan setelah waktu tertentu
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("setelah 5 detik :", time.Now())
		group.Done()
	})

	fmt.Println("Langsung dijalankan :", time.Now())

	group.Wait()
}
