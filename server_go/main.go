package main

import (
	"fmt"
	"net/http"
	"server_go/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}

func corsHandler(h http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if (r.Method == "OPTIONS") {
      //handle preflight in here
    } else {
      h.ServeHTTP(w,r)
    }
  }
}

func testFunc(){
    fmt.Println("I am working. notice me senpai")
}

func setupRoutes() {
    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })

    http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        fmt.Print("It worked. cors is stupid")
        })
}

func main() {
    fmt.Println("Distributed Chat App v0.01")
    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })

    http.HandleFunc("/execute", func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        fmt.Print("It worked. cors is stupid")
        })
    http.ListenAndServe(":8000", nil)
}