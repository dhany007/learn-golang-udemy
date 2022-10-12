package belajargolangdatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql" // pastikan selalu manggil ini
)

func TestEmpty(t *testing.T) {
	fmt.Println("OK")
}

func TestOpenConnection(t *testing.T) {
	// sql.Open(driver, "username:password@tcp(server:port)/dbname")
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_mysql")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
