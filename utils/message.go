package utils

type Message struct{
	Type string `json:"type"`
	Body map[string]interface{} `json:"body"`
}