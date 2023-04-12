package routers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func RunAPIServer() *gin.Engine {

	// SOCKET
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})

	// Connect
	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("Connect List: ", server.Count())
		s.SetContext("")
		s.Join(s.ID())
		server.BroadcastToRoom("/", s.ID(), "joinRoom", nil)
		return nil
	})

	// Error
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("error:", e)
	})

	// Disconnect
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("closed", reason)
	})

	defer server.Close()

	// GIN
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Latency 응답 시간
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	socketGroup := r.Group("")
	socketGroup.Use(SocketIOJSFile)
	{
		socketGroup.GET("/socket.io/*any", gin.WrapH(server))
		socketGroup.POST("/socket.io/*any", gin.WrapH(server))
	}

	// log.Println("Serving at localhost:8080...")
	// _ = r.Run(":8080")

	return r
}
