package websocket

import (
	"log"
	"net/http"
	"time"

	"github.com/Koenigseder/pixelart/internal/canvas"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is an instance for each connection
type Client struct {
	// Websocket connection
	conn *websocket.Conn
}

const (
	// Time allowed to read the next pong message from the peer
	pongWait = 60 * time.Second

	// Maximum message size allowed from peer
	maxMessageSize = 512
)

// ServeWebsocket handles websocket connections
func ServeWebsocket(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error establishing WebSocket connection: %v\n", err)
		return
	}

	client := &Client{conn: conn}

	go client.readPump()
	go client.writePump()
}

func (c *Client) writePump() {
	defer c.conn.Close()

	lastCanvasChange := canvas.Canvas.LastModified
	if err := c.conn.WriteJSON(canvas.Canvas); err != nil {
		log.Println(err)
		return
	}

	for {
		if lastCanvasChange != canvas.Canvas.LastModified {
			if err := c.conn.WriteJSON(canvas.Canvas); err != nil {
				log.Println(err)
				return
			}

			lastCanvasChange = canvas.Canvas.LastModified
		}
	}
}

func (c *Client) readPump() {
	defer c.conn.Close()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			return err
		}
		return nil
	})

	var requestMessage SetPixelRequestBody

	for {
		err := c.conn.ReadJSON(&requestMessage)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v\n", err)
			}
			break
		}

		if err := canvas.Canvas.UpdatePixel(requestMessage.X, requestMessage.Y, requestMessage.RGB); err != nil {
			log.Println(err.Error())
			return
		}
	}
}
