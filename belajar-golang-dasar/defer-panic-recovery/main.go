package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function logging")
}

func runApplication(value int) {
	defer logging() // defer selalu diatas
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("result ", result)
	// logging()
}

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Applikasi error")
	}
	fmt.Println("Aplikasi berjalan")

}

func endAppRecovery() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error : ", message)
	}

	fmt.Println("Aplikasi selesai")
}

func runAppRecover(error bool) {
	defer endAppRecovery()

	if error {
		panic("error")
	}

	fmt.Println("aplikasi berjalan")
}

func main() {
	fmt.Println("defer panic recovery")
	// defer => menjalankan sebuah fungsi setelah fungsi walaupun fungsi tersebut error

	// runApplication(10)
	// runApplication(0)

	// panic => menghentikan program, biasanya ketika program error
	// saat panic selesai dieksekusi, defer tetap akan dieksekusi

	// runApp(false)
	// runApp(true)

	// recover => function yang kita gunakan untuk menangkap data panic
	// dengan recover, proses panic akan berhenti, sehingga program tetap berjalan
	// recover yang benar itu disimpan di defer

	runAppRecover(false)
	runAppRecover(true)
	fmt.Println("Eko")
}
