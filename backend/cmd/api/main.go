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
	// ファイル名を元にファイル情報を取得
	file, err := getFileInfo(filename)
	if err != nil {
		log.Fatalln("ファイル情報が取得できませんでした")
	}
	defer file.Close()

	// DBの初期化および接続
	models.Init()

	err = insertUsersFromFile(file)
	if err != nil {
		log.Fatalln("ファイル情報をDBへ書き込む処理が失敗しました")
	}

	defer models.Db.Close()

}

func getFileInfo(filename string) (*os.File, error) {
	// ファイルをオープン
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ファイルを開けません:", err)
		return nil, err
	}
	return file, nil
}

func insertUsersFromFile(file *os.File) error {
	// ファイルの内容を一行ずつ読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var userData models.UserData
		err := json.Unmarshal(scanner.Bytes(), &userData)
		if err != nil {
			log.Println("JSONの解析エラー:", err)
			continue
		}
		user := models.User{
			Age:  userData.User.Age,
			Name: userData.User.Name,
			Role: userData.User.Role,
		}
		err = user.CreateUser()
		if err != nil {
			log.Println("ユーザー作成エラー:", err)
		}
		// ユーザーデータを出力
		fmt.Printf("User: %s, Age: %d, Role: %s\n", userData.User.Name, userData.User.Age, userData.User.Role)
	}

	// スキャン中にエラーが発生したか確認
	if err := scanner.Err(); err != nil {
		log.Println("ファイルの読み込みエラー:", err)
		return err
	}

	return nil
}
