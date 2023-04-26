package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kou Ushio!"))
}

func main() {
	http.HandleFunc("/", Hello)

	err := http.ListenAndServeTLS(":3939", "/etc/ssl/certs/ncth-app.jp.pem", "/etc/ssl/private/ncth-app.jp.key", nil)
	if err != nil {
		fmt.Printf("ERROR : %s", err)
	}
}
