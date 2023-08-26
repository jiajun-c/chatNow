package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {
	app := gin.Default()
	app.GET("/chat", chat)
	app.LoadHTMLGlob("frontend/*")
	app.GET("/", home)
	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}

var connPool []*websocket.Conn

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func chat(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	connPool = append(connPool, ws)
	if err != nil {
		panic(err)
	}
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Errorf(err.Error())
		}
		for _, client := range connPool {
			client.WriteMessage(msgType, msg)
		}
	}
}

func home(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}
