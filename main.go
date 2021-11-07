package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ユーザーを取得
	tokenFlag := flag.String("u", "private", "変えたいユーザー")
	flag.Parse()
	var token string
	switch *tokenFlag {
	case "private":
		token = os.Getenv("PRIVATE_ACCESS_TOKEN")
	case "enterprise":
		token = os.Getenv("ENTERPRISE_ACCESS_TOKEN")
	}

	cmd := exec.Command("gh", "auth", "login", "--with-token")

	// このpipeで何してるのか謎出力をwith-tokenに渡せてるっぽいので成功するんやけどここら辺イミフ。
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	io.WriteString(stdin, token)
	stdin.Close()

	// 結果を取得するのとエラーを標準エラーで表示してくれるように書き込み先を変更
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		panic(err)
	}

	fmt.Println("Successfully change user!!")
}
