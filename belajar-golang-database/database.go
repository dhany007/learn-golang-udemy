package belajargolangdatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// parseTime = true => golang konversi otomatis dari time ke string
	db, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/belajar_database_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)                  // minimal jumlah koneksi
	db.SetMaxOpenConns(100)                 // maksimal jumlah koneksi
	db.SetConnMaxIdleTime(5 * time.Minute)  // maksimal waktu diam koneksi
	db.SetConnMaxLifetime(60 * time.Minute) // maksimal waktu perbaharui diganti koneksi baru

	return db
}
