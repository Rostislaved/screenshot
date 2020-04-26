package main

import (
	"image/png"
	"log"
	"my-projects/screenshot/screenshoters"
	"os"
)

func main() {
	img, err := screenshoters.CaptureRect()
	if err != nil {
		log.Fatal(err)
	}

	fileName := "screen"
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)


}