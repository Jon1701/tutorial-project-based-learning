package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define an Upgrader. This will require a Read and Write buffer size.
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// Checks the Origin of the connection. For now, just allow any connection.
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

// Define a reader which will listed for new messages sent to our websocket
// endpoint.
func reader(conn *websocket.Conn) {
	for {
		// Read message.
		messageType, p, err := conn.ReadMessage()

		// If there was a problem reading the message, log it.
		if err != nil {
			log.Println(err)
			return;
		}

		// Print out message for clarity.
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return;
		}
	}
}

// Define WebSocket endpoint.
func HandlerServeWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// Upgrade endpoint to a WebSocket.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// Listen indefinitely for new messages coming through on our WebSocket
	// connection.
	reader(ws)
}

func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple Server")
}

func SetUpRoutes() {
	http.HandleFunc("/", HandlerRoot)
	http.HandleFunc("/ws", HandlerServeWebSocket)
}

func main() {
	fmt.Println("Chat App v0.01")
	SetUpRoutes()
	http.ListenAndServe(":8080", nil)
}