package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-docker-practice/models"
	"log"
	"os"
)

// ユーザーデータを表す構造体

func main() {

	// コマンドライン引数の数をチェック
	if len(os.Args) != 2 {
		log.Fatalln("使用方法: go run main.go [ファイル名]")
		os.Exit(1)
	}

	// コマンドライン引数からファイル名を取得
	filename := os.Args[1]

	// ファイルをオープン
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ファイルを開けません:", err)
		os.Exit(1)
	}
	defer file.Close()

	models.Init()

	// ファイルの内容を一行ずつ読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var userData models.UserData
		err := json.Unmarshal(scanner.Bytes(), &userData)
		if err != nil {
			fmt.Println("JSONの解析エラー:", err)
			continue
		}
		user := models.User{
			Age:  userData.User.Age,
			Name: userData.User.Name,
			Role: userData.User.Role,
		}
		err = user.CreateUser()

		if err != nil {
			// エラーの処理
			log.Println("ユーザー作成エラー:", err)
		}
		// ユーザーデータを出力
		fmt.Printf("User: %s, Age: %d, Role: %s\n", userData.User.Name, userData.User.Age, userData.User.Role)
	}

	// スキャン中にエラーが発生したか確認
	if err := scanner.Err(); err != nil {
		fmt.Println("ファイルの読み込みエラー:", err)
		os.Exit(1)
	}
}
