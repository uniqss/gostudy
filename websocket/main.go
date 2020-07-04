package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var addr = flag.String("addr", "0.0.0.0:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Error("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Error("write:", err)
			break
		}
	}
}


func main() {
	flag.Parse()
	//log.SetFlags(0)
	http.HandleFunc("/", echo)
	//log.Fatal(http.ListenAndServe(*addr, nil))
	certFile := ""
	keyFile := ""
	log.Fatal(http.ListenAndServeTLS(*addr, certFile, keyFile, nil))
}