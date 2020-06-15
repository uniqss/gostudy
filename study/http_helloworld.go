package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", defaultHandler)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("server error.")
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	//r.Header.Add("Connection", "Close")
	fmt.Fprintf(w, "hello world")
}
