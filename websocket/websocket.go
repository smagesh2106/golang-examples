package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/echo", DoEcho)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websocket.html")
	})

	http.ListenAndServe(":8080", nil)
}

func DoEcho(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)

	if err := conn.WriteMessage(websocket.TextMessage, []byte("Welcome to echo chat...")); err != nil {
		return
	}
	stopCh := make(chan interface{})
	go loop(conn, stopCh)

	fmt.Println("conn : ", conn.UnderlyingConn())
label:
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Unable to upgrade to websocket. ", err)
			return
		}
		if string(msg) == "bye" {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("goodbye client..")); err != nil {
				//ignore
			}
			stopCh <- struct{}{}
			break label

		}
		sendMsg := fmt.Sprintf("From Server ++> :%s", msg)

		// Print the message to the console
		fmt.Printf("Msg Type %v \n", msgType)
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), fmt.Sprintf("--->%s", string(msg)))

		// Write message back to browser
		//if err = conn.WriteMessage(msgType, msg); err != nil {
		if err = conn.WriteMessage(msgType, []byte(sendMsg)); err != nil {
			return
		}
	}
	fmt.Println("Ending chat...")
}

func loop(conn *websocket.Conn, stop chan interface{}) {

	timer1 := time.NewTimer(1 * time.Second)
	for {
		select {
		case t := <-timer1.C:
			if err := conn.WriteMessage(websocket.TextMessage, []byte("wake up..."+t.String())); err != nil {
				return
			}
			fmt.Println(".....")
			timer1.Reset(time.Second * 1)

		case <-stop:
			timer1.Stop()
			conn.Close()
		}

	}
}
