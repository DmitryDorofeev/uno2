package handlers

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"utils"
	"wshandlers"
	"room"
)

var rooms = map[int][]room.Room{
	2: []room.Room{},
	3: []room.Room{},
	4: []room.Room{},
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	userId := cookie.Value

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	jsonMessage := &utils.Message{}
	err = ws.ReadJSON(jsonMessage)

	switch jsonMessage.Type {
	case "gameInfo":
		playersCount := int(jsonMessage.Body["players"].(float64))
		wshandlers.StartGame(ws, playersCount, userId, &rooms)
	default:
		panic("ba")
	}
}
