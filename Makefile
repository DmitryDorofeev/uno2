all: get build
	

run:
	go run main/main.go

build:
	go build -o build/uno main/main.go

get:
	go get github.com/gorilla/websocket
	go get github.com/satori/go.uuid
	go get github.com/go-martini/martini
	go get github.com/mattn/go-sqlite3
