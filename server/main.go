package main

import (
	"fmt"
	"net/http"

	"server_go/pkg/websocket"
)

//connect the websockets, register client to the correct pool
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    //create the websocket connection
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    //create the client
    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    //register the client to the pool
    pool.Register <- client

    //take in info that client sends through websocket
    client.Read()
}

//get the routes in order
func setupRoutes() {
    //create the pool in preperation for route
    pool := websocket.NewPool()
    go pool.Start()
    /*TODO
        add custom routes by name
        save/load these routes from mongodb
    */

    //handle when frontend connects to /ws
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

//startup
func main() {
    fmt.Println("Distributed Chat App v0.01")
    //set routes
    setupRoutes()
    http.ListenAndServe(":8000", nil)
}
