package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kou kou!"))
}
func main() {
	http.HandleFunc("/", Hello) // /が来たときに func Hello を実行する

	err := http.ListenAndServeTLS(":3939", "ncth-app.jp.pem", "ncth-app.jp.key", nil)
	//func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler)

	if err != nil {
		fmt.Printf("ERROR : %s", err)
	}
}
