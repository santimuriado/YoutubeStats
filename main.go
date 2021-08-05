package main

import (
	"fmt"
	"log"
	"net/http"

	websocket "github.com/santimuriado/YoutubeSubs/WebSocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Youtube Subscription Number")
}

func stats(w http.ResponseWriter, r *http.Request) {

	//Upgrade the connection to a ws.
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)
}

func setupRoutes() {

	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	fmt.Println("Youtube Subscriptions")
	setupRoutes()
}
