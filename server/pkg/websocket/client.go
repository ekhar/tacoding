package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

//Clients have ID's, a designated pool, and a ws connection
type Client struct {
    ID   string
    Conn *websocket.Conn
    Pool *Pool
}


//Messages definition
type Message struct {
    Type int    `json:"type"`
    Body string `json:"body"`
    ID string
}

//Server reading what comes from client
func (c *Client) Read() {

    //logoff defered to end
    defer func() {
        c.Pool.Unregister <- c
        c.Conn.Close()
    }()

    //infinite loop of reading incoming messages
    for {
        //read
        messageType, p, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        //make message
        message := Message{Type: messageType, Body: string(p), ID: c.ID}
        //Broadcast
        c.Pool.Broadcast <- message
    }
}
