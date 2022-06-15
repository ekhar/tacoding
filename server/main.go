package main

import (
	"fmt"
	"net/http"
	"server_go/pkg/websocket"

	"github.com/gorilla/mux"
)

//connect the websockets, register client to the correct pool
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("THIS IS POOL SERVE", pool)
    fmt.Println(pool)
    fmt.Println(pool.ID)
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
func setupRoutes(r *mux.Router, url_cache map[string]*websocket.Pool) {
    /*TODO
        save/load these routes from mongodb
    */

    //handle when frontend connects to /ws
    r.HandleFunc("/ws/{key}", func(w http.ResponseWriter, r *http.Request) {
        //create the pool in preperation for route
        //var pool *websocket.Pool
        id := mux.Vars(r)["key"]
        //id is in the cache
        pool, hit := url_cache[id]
        if !hit{
            fmt.Println("WE MAKING A NEW POOL")
            pool = websocket.NewPool(id)
            url_cache[id]=pool
            go pool.Start()
        }
        serveWs(pool, w, r)
    })
}

//startup
func main() {
    fmt.Println("Distributed Chat App v0.01")
    r := mux.NewRouter()
    url_cache := make(map[string]*websocket.Pool)
    //set routes
    //maps are passed by ref default golang
    setupRoutes(r,url_cache)
    http.ListenAndServe(":8000", r)
}
