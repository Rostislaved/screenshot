package main

import (
	"image/png"
	"log"
	"os"
	"github.com/rostislaved/screenshot"
)

func main() {

	screenshoter := screenshot.New()

	img, err := screenshoter.CaptureScreen()
	if err != nil {
		log.Fatal(err)
	}


	file, err := os.Create("./screenshot.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

}
