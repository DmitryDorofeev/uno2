package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

upgrager := websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Conn struct {
	ws *websocket.Conn
	send chan []byte
}

type Room struct {
	connections map[*Conn]bool
	broadcast chan []byte
	register chan *Conn
	unregister chan *Conn
}

var room = Room{
	broadcast:   make(chan []byte),
	register:    make(chan *Conn),
	unregister:  make(chan *Conn),
	connections: make(map[*Conn]bool),
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn := &Conn{send: make(chan []byte, 256), ws: ws}
	room.register <- conn
}

func main() {
	// http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	err := http.ListenAndServe(":8008", nil)
}
