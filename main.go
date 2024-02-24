package main

import (
	"encoding/json"
	"log"
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

type MessageStruct map[string]string

func KeyValueToJSON(key, value string) []byte {
	data := map[string]string{
		key: value,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return jsonData
}

func ByteToStruct(msgByte []byte) MessageStruct {
	var msgStruct MessageStruct

	err := json.Unmarshal(msgByte, &msgStruct)
	if err != nil {
		log.Println("Error:", err)
	}
	return msgStruct
}

func handleMessages(w http.ResponseWriter, r *http.Request) {

	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	for {
		// read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		log.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Unmarshal the JSON and extract value
		msgStruct := ByteToStruct(msg)
		inputText := msgStruct["text"]

		// logic
		returnText := ""
		if inputText == "wow" {
			returnText = "This is wow"
		} else {
			returnText = "This is NOT wow"
		}
		jsonByte := KeyValueToJSON("text", returnText)

		err = conn.WriteMessage(msgType, jsonByte)
		if err != nil {
			// Check if it's a broken pipe error
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				log.Printf("Error writing message to WebSocket: %v", err)
			} else {
				log.Println("WebSocket connection closed by client")
			}
			return
		}
	}
}

func main() {
	http.HandleFunc("/echo", handleMessages)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
