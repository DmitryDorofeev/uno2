package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/satori/go.uuid"
	"fmt"
)

type ProfileResponse struct {
	Login string `json:"login"`
	Email string `json:"email"`
}

func HandleProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		userId := uuid.NewV4()

		cookie = &http.Cookie{
			Name: "user",
			Value: userId.String(),
			Path: "/",
			HttpOnly: false,
		}

		fmt.Println("Cookie", cookie.Value, r.Host)
		http.SetCookie(w, cookie)
	}

	resp, _ := json.Marshal(ProfileResponse{
		Login: "DmitryDorofeev",
		Email: cookie.Value,
	})
	w.WriteHeader(400)
	w.Write([]byte(resp))
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	params := make([]byte, 1024)
	r.Body.Read(params)

	w.WriteHeader(200)
	w.Write(params)
}
