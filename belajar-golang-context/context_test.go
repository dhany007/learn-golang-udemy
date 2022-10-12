package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {

	contextA := context.Background()

	// parent A,
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// value => mengambil value dari child ke parent tinggi nya
	// tidak nanyak kebawah, wajib nanyak keatas
	fmt.Println(contextF.Value("f")) // ada F
	fmt.Println(contextF.Value("c")) // ada C
	fmt.Println(contextF.Value("b")) // nill, beda parent
}

// membuat fungsi counter yang dijalankan terusmenerus
func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter += 1
		}
	}()

	return destination
}

func TestContextWIthCancelLeak(t *testing.T) {
	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 2

	destination := CreateCounterLeak()

	for n := range destination {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break // jika n = 0, maka kita break
		}
	}

	time.Sleep(3 * time.Second)
	// hasil yang diharapkan adalah : misal goroutine awal = 2, maka setelah break, seharusnya balik goroutine = 2
	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 3
	// tapi yang didapat kan hasil = 3, kenapa ?
	// karena channel nya itu tidak dibatalkan, tapi terus berjalan

}

// membuat fungsi counter yang dijalankan dan ada context cancel
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done(): // selama ctx done/cancel, maka goroutine nya berhenti
				return
			default: // akan dijalankan terusmenerus sampai channel mengirimkan cancel
				destination <- counter
				counter += 1
			}
		}
	}()

	return destination
}

func TestContextWIthCancel(t *testing.T) {
	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 2

	parent := context.Background() // membuat context

	ctx, cancel := context.WithCancel(parent) // selalu return 2, yaitu ctx dan cancel

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 3, ada proses goroutine

	for n := range destination {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break // jika n = 0, maka kita break
		}
	}
	cancel() // mengirimkan sinyal cancel

	time.Sleep(3 * time.Second)
	// hasil yang diharapkan adalah : misal goroutine awal = 2, maka setelah break, seharusnya balik goroutine = 2
	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 2

}

// membuat fungsi counter yang dijalankan dan ada context cancel
func CreateCounterTimeout(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done(): // selama ctx done/cancel, maka goroutine nya berhenti
				return
			default: // akan dijalankan terusmenerus sampai channel mengirimkan cancel
				destination <- counter
				counter += 1
				time.Sleep(1 * time.Second) // membuat slow
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 2

	parent := context.Background() // membuat context

	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // batas 5 second
	defer cancel()                                            // selalu dikirim cancel kalau proses selesai

	destination := CreateCounterTimeout(ctx)

	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 3, ada proses goroutine

	// ini akan berjalan terus menerus sampai batas timeout diatas
	for n := range destination {
		fmt.Println("Counter : ", n)
	}

	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 2

}

// menggunakan waktu deadline, yaitu pukul berapa
func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 2

	parent := context.Background() // membuat context

	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10*time.Second)) // batas 5 second diwaktu mendatang
	defer cancel()                                                              // selalu dikirim cancel kalau proses selesai

	destination := CreateCounterTimeout(ctx)

	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 3, ada proses goroutine

	// ini akan berjalan terus menerus sampai batas timeout diatas
	for n := range destination {
		fmt.Println("Counter : ", n)
	}

	fmt.Println("Total Goroutine =", runtime.NumGoroutine()) // 2

}
