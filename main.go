/*
package main

import (
	"log"
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>Hello</title></head>
		<body><h1>Hello, world!</h1></body>
	</html>
	`), "", 1250, 770)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
*/
package main

import (
	"log"
	"net/http"
)

func main() {
	//ディレクトリを指定する
	fs := http.FileServer(http.Dir("www"))
	//ルーティング設定。"/"というアクセスがきたらwwwディレクトリのコンテンツを表示させる
	http.Handle("/", fs)

	log.Println("Listening...")
	// 3000ポートでサーバーを立ち上げる
	http.ListenAndServe(":8080", nil)
}
