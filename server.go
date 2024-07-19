package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = make([]*websocket.Conn, 0)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while connecting:", err)
		return
	}
	connections = append(connections, conn)
	fmt.Printf("User %s connected.\n", conn.RemoteAddr())

	for _, c := range connections {
		err := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("User %s connected.", conn.RemoteAddr())))
		if err != nil {
			fmt.Println("Error while sending message:", err)
		}
	}

	go func() {
		defer conn.Close()
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error while reading the message!")
				break
			}

			var msgData Message
			err = json.Unmarshal(msg, &msgData)
			if err != nil {
				fmt.Println("Error while unmarshalling message:", err)
				continue
			}

			fmt.Printf("User %s says: %s\n", msgData.Username, msgData.Message)

			for _, c := range connections {
				if err := c.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", msgData.Username, msgData.Message))); err != nil {
					fmt.Println("Error while sending message:", err)
				}
			}
		}
	}()
}

func main() {
	http.HandleFunc("/chat", handleConnection)
	serverAddr := "localhost:8080"
	fmt.Printf("Server is listening on port %s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
