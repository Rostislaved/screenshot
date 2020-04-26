package main

import (
	"log"
	"my-projects/screenshot/screenshoters"
	"net/http"

	_ "net/http/pprof"
)

func main() {

	screenshoter := screenshoters.New()

	go http.ListenAndServe(":8080", nil)

	for {
		img, err := screenshoter.CaptureScreen()
		if err != nil {
			log.Fatal(err)
		}
		_ = img
	}

	//fileName := "screen"
	//file, _ := os.Create(fileName)
	//defer file.Close()
	//png.Encode(file, img)

}
