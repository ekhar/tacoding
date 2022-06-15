package websocket

import "fmt"

//Pools have channels for registering and unregistering clients
//Pools have boolean value for if a client exists
//Pools have channel for broadcasting a message
//Lastly pools have ID's
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan []byte
    ID         string
}

//create a pool struct and return pointer to the pool
func NewPool(id string) *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte),
        ID: id,
	}
}

//Pool begins
func (pool *Pool) Start() {
    //infinite loop
	for {
        //handle client channel cases
		select {
        //if a new client registers
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
			}
        //if client exits
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))

        //If we get a message, lets broadcast it
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
            //for all clients in pool
			for client := range pool.Clients {
                //if client is not the sender of the message
                if err := client.Conn.WriteMessage(1,message); err != nil {
                    fmt.Println(err)
                    return
                }
			}

		}
	}
}
