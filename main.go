package main

import (
	"fmt"
	"net/http"
	"os"
)

func UsioOpen() {
	fmt.Println("ファイル読み取り処理を開始します")
	// ファイルをOpenする
	f, err := os.Open("test.txt")
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("error")
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
		fmt.Println(string(buf[:n]))
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kou kou!"))
}
func main() {
	http.HandleFunc("/", Hello) // /が来たときに func Hello を実行する

	err := http.ListenAndServeTLS(":3939", "ncth-app.jp.pem", "ncth-app.jp.key", nil)
	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler)　error
	// errにエラーメッセージを格納する

	if err != nil { // エラーメッセージがあるとき出力
		fmt.Printf("ERROR : %s", err)
	}
}
