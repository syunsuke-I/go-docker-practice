package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQLへの接続情報
	const (
		host     = "postgres"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "postgres"
	)

	// 接続文字列を作成
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// PostgreSQLに接続
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続を確認
	err = db.Ping()
	if err != nil {
		log.Fatal("接続失敗:", err)
	}

	fmt.Println("接続成功!")
}
