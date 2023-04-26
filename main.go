package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kou!"))
}

func main() {
	http.HandleFunc("/", Hello)

	err := http.ListenAndServeTLS(":3939", "ncth-app.jp.pem", "ncth-app.jp.key", nil)
	if err != nil {
		fmt.Printf("ERROR : %s", err)
	}
}
