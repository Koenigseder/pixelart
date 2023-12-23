package websocket

// SetPixelRequestBody - Request body for SetPixel
type SetPixelRequestBody struct {
	X   int   `json:"x"`
	Y   int   `json:"y"`
	RGB []int `json:"rgb"`
}
