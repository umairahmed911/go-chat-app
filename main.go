package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

type client struct {
	conn *websocket.Conn
}

var (
	clients   = make(map[*client]bool)
	broadcast = make(chan string)
	mutex     sync.Mutex
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	c := &client{conn: ws}
	mutex.Lock()
	clients[c] = true
	mutex.Unlock()
	defer func() {
		mutex.Lock()
		delete(clients, c)
		mutex.Unlock()
		ws.Close()
	}()
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		// Expect message in format 'name::message', broadcast as-is
		broadcast <- string(msg)
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for c := range clients {
			err := c.conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				c.conn.Close()
				delete(clients, c)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	go handleMessages()
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
