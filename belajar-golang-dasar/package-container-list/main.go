package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("double linked list")
	// membuat double linked list baru
	data := list.New()

	// menambah item di double linked list
	data.PushBack("dhany")
	data.PushBack("putra")
	data.PushBack("aritonang")

	// list.Front => depan
	// list.Back => belakang
	// list.Next => selanjutnya
	// list.Prev => sebelumnya

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
