package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]Message) // connected clients
var broadcast = make(chan Message)              // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message object
type Message struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Message     string `json:"message"`
	Information string `json:"information"`
}

func main() {
	go handleMessages()

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	http.Handle("/",
		http.FileServer(
			&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "public"}))

	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/sw.js", serviceWorkerHandler)

	// log.Fatal(http.ListenAndServeTLS(":8080", "certificate.crt", "private.key", nil))
	log.Fatal(http.ListenAndServeTLS(":443", "certificate.crt", "private.key", nil))
}

func redirect(w http.ResponseWriter, req *http.Request) {
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		http.StatusTemporaryRedirect)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error %v", err)
		return
	}

	ws.SetCloseHandler(func(code int, text string) error {
		broadcast <- Message{
			Email:       "anonymous@mail.com",
			Username:    clients[ws].Username,
			Message:     "#left#",
			Information: ""}
		return nil
	})
	defer ws.Close()

	clients[ws] = Message{
		Email:       "anonymous@mail.com",
		Username:    "unknown user",
		Message:     "#anonymous#",
		Information: ""}

	for client := range clients {
		if client != ws {
			msg := Message{
				Email:       "Manager@chat.com",
				Username:    "unknown user",
				Message:     "#userlistjoin#",
				Information: clients[client].Username}
			ws.WriteJSON(msg)
		}
	}

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		clients[ws] = msg
		broadcast <- msg
	}

}

func serviceWorkerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/javascript")
	http.ServeFile(w, r, "sw.js")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast

		if strings.TrimSpace(msg.Message) == "" {
			continue
		}

		if msg.Message == "#login#" {
			msg.Message = msg.Username + " joined this room."
			msg.Information = msg.Username
			msg.Email = "Manager@chat.com"
			msg.Username = "Room Manager"
		} else if msg.Message == "#join#" {
			msg.Message = "An unknown user set nickname as " + msg.Username + "."
			msg.Information = msg.Username
			msg.Email = "Manager@chat.com"
			msg.Username = "Room Manager"
		} else if msg.Message == "#left#" {
			msg.Message = msg.Username + " has left this room."
			msg.Information = msg.Username
			msg.Email = "Manager@chat.com"
			msg.Username = "Room Manager"
		} else if msg.Message == "#change#" {
			msg.Message = msg.Username + " change nickname as " + msg.Email + "."
			msg.Information = msg.Username + "," + msg.Email
			msg.Email = "Manager@chat.com"
			msg.Username = "Room Manager"
		}

		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
