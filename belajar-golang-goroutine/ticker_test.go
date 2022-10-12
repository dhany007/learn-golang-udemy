package belajargolanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

// kejadiannya hanya sekali dijalankan
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

// balikannya channel dan tidak bisa berhenti
// kejadian yang berulang
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}
}
