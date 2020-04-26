# Screenshot
Simple cross-platform pure Go screen shot library. (tested on linux&windows&osx)



## Basic Usage
Import the package
```go
import (
    "github.com/rostislaved/screenshot"
)
```

```go
func main() {

	screenshoter := screenshot.New()

	img, err := screenshot.CaptureScreen()
	if err != nil {
		log.Fatal(err)
	}
}
```

## Dependencies
* **Windows** - None
* **Linux/FreeBSD** - https://github.com/BurntSushi/xgb
* **OSX** - cgo (CoreGraphics,CoreFoundation, that should not be a problem)

## Examples
Look at `examples/` folder.
