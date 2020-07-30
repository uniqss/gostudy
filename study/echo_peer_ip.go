package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	listenAddr string
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":65432", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "http:", log.LstdFlags)

	logger.Println("server is ready to handle request at:", listenAddr)

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		//fmt.Fprintf(w, "hello world.  r.RemoteAddr:%s", r.RemoteAddr)
		strArr := strings.Split(r.RemoteAddr, ":")
		if len(strArr) >= 1 {
			fmt.Fprintf(w, strArr[0])
		}
	})

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	server.SetKeepAlivesEnabled(false)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("could not listen on %s: %v\n", listenAddr, err)
	}
	logger.Println("server stopped.")
}
