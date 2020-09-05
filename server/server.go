package main

import (
	"net/http"
	"service/server/api"
)

func main() {
	// ルーティングを設定する
	http.HandleFunc("/api/v1/vision", api.CallVisionApi)
	// ポート8080番でサーバーを起動する
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
