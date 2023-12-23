package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Koenigseder/pixelart/internal/canvas"
	"github.com/Koenigseder/pixelart/internal/websocket"
	"github.com/gin-gonic/gin"
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

	var requestBody websocket.SetPixelRequestBody

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
