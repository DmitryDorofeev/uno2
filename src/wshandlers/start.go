package wshandlers

import (
	"github.com/gorilla/websocket"
	"utils"
	"room"
	"log"
)


func StartGame(ws *websocket.Conn, count int, userId string, rooms *map[int][]room.Room) {
	log.Println("Search for requested room. Requested: ", count)
	var selected room.Room
	var defined bool = false
	for _, r := range (*rooms)[count] {
		if len(r.Users) < count {
			log.Println("Found requested room. Users: ", len(r.Users), "| Requested: ", count)
			selected = r
			defined = true
			break
		}
	}

	if !defined {
		selected = room.Room{
			Broadcast:   make(chan *utils.Message),
			Register:    make(chan *room.User),
			Unregister:  make(chan *room.User),
			Users: make(map[string]*room.User),
			Size: count,
		}
		go selected.Run()
		(*rooms)[count] = append((*rooms)[count], selected)
		log.Println("Room created. Room size:", count)
	}
	selected.Register <- &room.User{
		Id: userId,
		Ws: ws,
		Login: userId,
	}
}
