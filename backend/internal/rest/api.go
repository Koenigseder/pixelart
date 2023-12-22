package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Koenigseder/pixelart/internal/canvas"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// GetCanvas (GET) - Get canvas as JSON
func GetCanvas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, canvas.Canvas)
}

// SetPixel (POST) - Set a specific pixel to a RGB value. Returns the request body at success
func SetPixel(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Cannot convert body")
		return
	}

	if len(jsonData) == 0 {
		c.IndentedJSON(http.StatusBadRequest, "Please provide a request body")
		return
	}

	var requestBody SetPixelRequestBody

	err = json.Unmarshal(jsonData, &requestBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid body content")
		return
	}

	if err := canvas.Canvas.UpdatePixel(requestBody.X, requestBody.Y, requestBody.RGB); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(200, requestBody)
}

// WebSocketEndpoint - Endpoint for the WebSocket
func WebSocketEndpoint(c *gin.Context) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("Error establishing WebSocket connection: %v\n", err)
		return
	}
	defer ws.Close()

	lastCanvasChange := canvas.Canvas.LastModified
	if err := ws.WriteJSON(canvas.Canvas); err != nil {
		log.Println(err)
		return
	}

	for {
		if lastCanvasChange != canvas.Canvas.LastModified {
			if err := ws.WriteJSON(canvas.Canvas); err != nil {
				log.Println(err)
				return
			}

			lastCanvasChange = canvas.Canvas.LastModified
		}
	}
}
