package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// ini emang dikomentari
//go:embed version.txt
var version string // variabel nya harus diluar function

//go:embed version2.txt
var version2 string

func TestEmbedString(t *testing.T) {
	fmt.Println("tes embed")
	fmt.Println(version)
	fmt.Println(version2)
	fmt.Println("selesai")
}

//go:embed tes.png
var logo []byte

func TestEmbedByte(t *testing.T) {
	err := ioutil.WriteFile("logo_baru.png", logo, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt

var files embed.FS

func TestEmbedFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

// digunakan ketika patern/pola nya sama
// baca dokumentasi patch matcher
// regex

// semua yang ekstension nya txt
//go:embed files/*.txt
var path embed.FS

func TestEmbedPatchMathcer(t *testing.T) {
	dir, _ := path.ReadDir("files")

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

// package embed, ini sudah dimasukkan kedalam binary nya secara permanent,
// jadi jika sudah dikompile, maka kita bisa menghapus file external
// walaupun diubah isinya, tetap yang hasil compile yang dipakai
