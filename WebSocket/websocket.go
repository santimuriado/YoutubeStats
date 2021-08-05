package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	youtube "github.com/santimuriado/YoutubeStats/Youtube"
)

//Set Read and Write buffer.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Upgrade recieves the request and upgrades the request to a websocket connection.
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	//Check host origin. Returns true to make it simpler.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	//Builds the ws connection.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}

func Writer(conn *websocket.Conn) {

	for {
		//New ticker that ticks every 5 seconds.
		ticker := time.NewTicker(5 * time.Second)

		//Every time the ticker ticks.
		for t := range ticker.C {

			fmt.Printf("Updating Stats: %v+v\n", t)

			items, err := youtube.GetChannel()
			if err != nil {
				fmt.Println(err)
			}
			//Marshall items into a JSON string.
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			//Write the JSON string to the ws connection and store any error that might come up.
			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
