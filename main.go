package main

import (
	"fmt"
	"net/http"
	"os"
)

func UsioOpen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ファイル読み取り処理を開始します"))

	// ファイルをOpenする
	f, err := os.Open("logs/log.txt")
	// 読み取り時の例外処理
	if err != nil {
		w.Write([]byte("error\n"))
	}
	// 関数が終了した際に確実に閉じるようにする
	defer f.Close()

	// バイト型スライスの作成
	buf := make([]byte, 1024)
	for {
		// nはバイト数を示す
		n, err := f.Read(buf)
		// バイト数が0になることは、読み取り終了を示す
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		// バイト型スライスを文字列型に変換してファイルの内容を出力
		w.Write(buf[:n])
	}
}

// w.Write([]byte("Kou kou!"))

func main() {
	http.HandleFunc("/", UsioOpen) // /が来たときに func Hello を実行する

	/* HandleFunc の第一引数のパス指定の最後には "/" を付けるべきでは? */
	/* (前回は "/end" と記述されていた https://github.com/dokimiki/ncth-scfests-backend/commit/8321e5efddf4b86cfcede430f056f894e8a241ac ) */
	http.HandleFunc("/webhook/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is Webhook."))
	})

	err := http.ListenAndServeTLS(":3939", "ncth-app.jp.pem", "ncth-app.jp.key", nil)
	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler)　error
	// errにエラーメッセージを格納する

	if err != nil { // エラーメッセージがあるとき出力
		fmt.Printf("ERROR : %s", err)
	}
}
