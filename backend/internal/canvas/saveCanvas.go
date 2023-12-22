package canvas

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// SaveCanvasAsFile - Save the canvas as a backup file
func SaveCanvasAsFile() {
	backupFilePath := os.Getenv("BACKUP_DIR_PATH")
	if err := os.MkdirAll(backupFilePath, os.ModePerm); err != nil {
		log.Fatalf("Cannot create folder; no backups will be made: %v\n", err)
	}

	for {
		time.Sleep(10 * time.Second)

		data, err := json.Marshal(Canvas)
		if err != nil {
			log.Printf("Cannot convert struct to byte array: %v\n", err)
			continue
		}

		filePath := fmt.Sprintf("%s/%d", backupFilePath, time.Now().Unix())
		err = os.WriteFile(filePath, data, 0664)
		if err != nil {
			log.Printf("Cannot write to location %s\n", filePath)
			continue
		}

		log.Printf("Successfully wrote to %s\n", filePath)
	}
}
