step
1. clone project restful api 
2. go get github.com/google/wire => install google wire
3. install wire => go install github.com/google/wire/cmd/wire@latest
4. membuat provider (constructor function)
5. membuat injector


gunakan autogenerate

```
  wire gen namapackage
  wire gen dhany007/belajar-golang-restful-api/simple
```

atau masuk ke folder injector (simple) dan ketik "wire" di cmd.

tambahkan error pada struct, dan balikkan error pada paramater ke 2

tambahkan error pada injector parameter


provider tidak boleh memprovide lebih dari satu dengan tipe data yang sama,
maka dibuat alias, seperti pada file database.go

provider set => untuk menggrouping beberapa provider. contohnya ada di foo, bar, dan foobar

membuat provider dengan struct di file foo_bar.go

binding value
=> melakukan dependency injection terhadap value yang sudah ada tanpa harus membuat provider terlebih dahulu
=> langsung sebut object nya

struct field provider
=> menggunakan sebuah field dari struct untuk dijadikan dependency pada provider lain

kebanyakan yang diganti itu main.go nya

menjalankannya memakai :

go run .