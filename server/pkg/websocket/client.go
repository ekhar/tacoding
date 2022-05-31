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
    //Type int    `json:"type"`
    Body string `json:"ops"`
    ID string `json:"id"`
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
        _, message, err := c.Conn.ReadMessage()
        //messageType, p, err := c.Conn.NextReader()
        if err != nil {
            log.Println(err)
            return
        }
        //make message
        //Broadcast
        c.Pool.Broadcast <- message
    }
}
