package belajargolangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

// insert
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO costumer(id, name) VALUES ('budi', 'BUDI')"
	_, err := db.ExecContext(ctx, query)
	fmt.Println(err)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses menambahkan customer baru")
}

// select, mengembalikan nilai
func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM costumer"
	rows, err := db.QueryContext(ctx, query)
	fmt.Println(err)

	if err != nil {
		panic(err)
	}

	defer rows.Close() // selalu ditutup rows nya setelah selesai

	// rows.Next() (Boolean) : iterasi rows
	// rows.Scan(column...) : membaca tiap data

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name) // harus sesuai dengan colom yang diselect tadi

		if err != nil {
			panic(err)
		}

		fmt.Println("id :", id)
		fmt.Println("name :", name)
	}
}

// select, mengembalikan nilai
func TestQueryComplexSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, created_at, married FROM costumer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close() // selalu ditutup rows nya setelah selesai

	// rows.Next() (Boolean) : iterasi rows
	// rows.Scan(column...) : membaca tiap data
	// kalau ada data NULL pake tipe data sql.NullString dll

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var created_at time.Time
		var birth_date sql.NullTime
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &created_at, &married) // harus sesuai dengan colom yang diselect tadi

		fmt.Println(err)
		if err != nil {
			panic(err)
		}

		fmt.Println("========================")
		fmt.Println("id :", id)
		fmt.Println("name :", name)
		if email.Valid {
			fmt.Println("email :", email.String)
		}
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		if birth_date.Valid {
			fmt.Println("birth_date :", birth_date.Time)
		}
		fmt.Println("created_at :", created_at)
		fmt.Println("married :", married)
	}
}

func TestQueryParameterSqlIncorrect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	user := "admin'; #"
	pass := "salah"

	ctx := context.Background()

	query := "SELECT username FROM users WHERE username = '" + user + "' AND password = '" + pass + "' LIMIT 1"
	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		panic(err)
	}

	defer rows.Close() // selalu ditutup rows nya setelah selesai

	if rows.Next() {
		var username string

		err := rows.Scan(&username) // harus sesuai dengan colom yang diselect tadi

		if err != nil {
			panic(err)
		}

		fmt.Println("sukses login")
	} else {
		fmt.Println("gagal login")
	}

	// sukses login
}

func TestQueryParameterSqlCorrect(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	user := "admin"
	pass := "admin"

	ctx := context.Background()

	query := "SELECT username FROM users WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(query)
	rows, err := db.QueryContext(ctx, query, user, pass)

	if err != nil {
		panic(err)
	}

	defer rows.Close() // selalu ditutup rows nya setelah selesai

	if rows.Next() {
		var username string

		err := rows.Scan(&username) // harus sesuai dengan colom yang diselect tadi

		if err != nil {
			panic(err)
		}

		fmt.Println("sukses login")
	} else {
		fmt.Println("gagal login")
	}

}

// insert
func TestExecParameterSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "dhany2"
	password := "dhany2"

	ctx := context.Background()

	query := "INSERT INTO users(username, password) VALUES (?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	fmt.Println(err)

	if err != nil {
		panic(err)
	}

	fmt.Println("Sukses menambahkan customer baru")
}

func TestExecParamAutoIncrementSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	email := "testlu@gmail.com"
	comment := "ini adalah comment"

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	insetId, err := result.LastInsertId() // mendapatkan last id auto increment

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Last insert id", insetId)

	fmt.Println("Sukses menambahkan customer baru")
}

func TestPrepareStatement(t *testing.T) {
	// statement => digunakan untuk mengeksekusi kuery yang berulang2
	// biasanya untuk keperluan bulk upload data, baik get maupun insert

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	statement, err := db.PrepareContext(ctx, query)

	if err != nil {
		panic(err)
	}
	defer statement.Close() // statement  nya selalu diclose

	for i := 0; i < 10; i++ {
		email := "person" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment -" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment) // query disini gak usah dibuat karena diatas sudah ada

		if err != nil {
			panic(err)
		}

		lastInserId, _ := result.LastInsertId()
		fmt.Println("Comment id :", lastInserId)

	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES (?, ?)"

	for i := 0; i < 10; i++ {
		email := "person" + strconv.Itoa(i) + "@gmail.com"
		comment := "Comment -" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)

		if err != nil {
			panic(err)
		}

		lastInserId, _ := result.LastInsertId()
		fmt.Println("Comment id :", lastInserId)

	}

	err = tx.Commit() // commit
	// err = tx.Rollback() // rollback

	if err != nil {
		panic(err)
	}
}
