perintah eksekusi di root folder : 
go test -v ./...

assertion ambil dari framework lain :
go get github.com/stretchr/testify

run benchmark
go test -v -bench=. (didalam folder berisi file test)

run benchmark di root
go test -v -bench=../...

hanya menjalankan benchmark tanpa fungsi test lain
go test -v -run=NamaFunsiLain -bench=BenchmarkHelloWorldSub