package main

import (
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Kou Ushio!"))
}

func main() {
	server := http.Server{
		Addr: ":3939",
	}

	http.HandleFunc("/", Hello)
	server.ListenAndServe()
}
