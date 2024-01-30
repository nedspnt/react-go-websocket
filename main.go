package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var clients []websocket.Conn

func handleMessages(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	clients = append(clients, *conn)

	for {
		// read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write the message so that frontend can consume
		for _, client := range clients {
			err := client.WriteMessage(msgType, msg)
			if err != nil {
				fmt.Println("Error sending message to client:", err)
			}
		}
	}
}

func main() {
	http.HandleFunc("/echo", handleMessages)

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
