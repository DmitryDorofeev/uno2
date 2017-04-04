package room

import (
	"github.com/gorilla/websocket"
	"utils"
	"log"
)

type Player struct{
	Id string `json:"id"`
	Login string `json:"login"`
}

type Card struct{
	Id int `json:"id"`
	Width int `json:"width"`
	Height int `json:"height"`
	X int `json:"x"`
	Y int `json:"y"`
	Color string `json:"color"`
	Type string `json:"type"`
}

type User struct {
	Ws *websocket.Conn
	Login string
	Id string
}

type Room struct {
	Users map[string]*User
	Broadcast chan *utils.Message
	Register chan *User
	Unregister chan *User
	Size int
}


func (r *Room) SendBroadcast(message *utils.Message) {
	log.Println("Sending messages to group...")
	for _, user := range r.Users {
		log.Println("Sending message to", user.Login)
		err := user.Ws.WriteJSON(message)
		if err != nil {
			log.Fatal("Unable to send message to", user.Login)
		}
	}
}


func (r *Room) Run() {
	for {
		select {
		case user := <-r.Register:
			r.Users[user.Id] = user
			log.Println("Register user", user.Id)

			if (len(r.Users) == r.Size) {
				log.Println("Group is full. Try to send messages")
				players := []Player{}

				for _, user := range r.Users {
					players = append(players, Player{
						Id: user.Id,
						Login: user.Login,
					})
				}

				r.SendBroadcast(&utils.Message{
					Type: "start",
					Body: map[string]interface{}{
						"players": players,
					},
				})

				r.SendBroadcast(&utils.Message{
					Type: "cards",
					Body: map[string]interface{}{
						"cards": []Card{
							{
								Id: 1,
								Width: 120,
								Height: 180,
								Color: "red",
								Type: "number",
								X: 0,
								Y: 0,
							},
							{
								Id: 105,
								Width: 120,
								Height: 180,
								Color: "blue",
								Type: "skip",
								X: 1200,
								Y: 1260,
							},
						},
					},
				})
			}
		case message := <- r.Broadcast:
			r.SendBroadcast(message)
		}
	}
}

func (r *Room) getUsers() map[string]*User {
	return r.Users
}

