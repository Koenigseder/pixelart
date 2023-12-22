package canvas

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type CanvasStruct struct {
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Pixels       [][][]int `json:"pixels"`
	LastModified int       `json:"lastModified"`
}

var Canvas CanvasStruct

// InitializeEmpty - Initialize an empty canvas with a fixed size
func (can *CanvasStruct) InitializeEmpty(width, height int) {
	can.Width = width
	can.Height = height
	can.Pixels = make([][][]int, width)
	can.LastModified = int(time.Now().Unix())

	for i := 0; i < can.Width; i++ {
		can.Pixels[i] = make([][]int, height)

		for j := 0; j < can.Height; j++ {
			can.Pixels[i][j] = []int{0, 0, 0}
		}
	}
}

// InitializeFromLatestBackup - Initialize a canvas from the latest backup file.
// `disasterWidth` and `disasterHeight` are used to have a fallback canvas
func (can *CanvasStruct) InitializeFromLatestBackup(disasterWidth, disasterHeight int) error {
	backupFilePath := os.Getenv("BACKUP_DIR_PATH")

	files, err := os.ReadDir(backupFilePath)

	if len(files) == 0 {
		can.InitializeEmpty(disasterWidth, disasterHeight)
		return err
	}

	log.Printf("Using %s as backup file\n", files[len(files)-1].Name())
	data, err := os.ReadFile(fmt.Sprintf("%s/%s", backupFilePath, files[len(files)-1].Name()))
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Canvas)
	if err != nil {
		return err
	}

	LastSaveTimestamp = Canvas.LastModified

	return nil
}

// UpdatePixel - Update a specific pixel with
func (can *CanvasStruct) UpdatePixel(x, y int, rgb []int) error {
	xValid := x < 0 || x > Canvas.Width-1
	yValid := y < 0 || y > Canvas.Height-1

	if xValid {
		return fmt.Errorf("'x' cannot be less than 0 and greater than %d", Canvas.Width-1)
	}

	if yValid {
		return fmt.Errorf("'y' cannot be less than 0 and greater than %d", Canvas.Height-1)
	}

	if len(rgb) < 3 {
		return errors.New("'rgb' has to have exactly three values between 0 and 255")
	}

	rValid := rgb[0] < 0 || rgb[0] > 255
	gValid := rgb[1] < 0 || rgb[1] > 255
	bValid := rgb[2] < 0 || rgb[2] > 255

	if rValid || gValid || bValid {
		return errors.New("RGB values are invalid")
	}

	can.Pixels[y][x] = rgb
	can.LastModified = int(time.Now().Unix())

	return nil
}
