package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func ok(w http.ResponseWriter, r *http.Request) {
	log.Println("WEB REQUEST ORIGIN:", r.Header.Get("Origin"))
	log.Println("WEB REQUEST HOST:", r.Host)
	w.WriteHeader(200)
	w.Write([]byte("ok"))
}

var upgrader = websocket.Upgrader{
	// Override the default with one that never rejects requests, and logs incoming headers
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ws(w http.ResponseWriter, r *http.Request) {
	log.Println("WS REQUEST ORIGIN:", r.Header.Get("Origin"))
	log.Println("WS REQUEST HOST:", r.Host)

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
	}

	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/testweb", ok)
	http.HandleFunc("/testws", ws)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
