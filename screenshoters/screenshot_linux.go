package screenshoters

import (
	"image"
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type screenshoter struct {
	conn *xgb.Conn
}

func New() *screenshoter {
	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
		//return image.Rectangle{}, err
	}
	defer conn.Close()

	return &screenshoter{conn: conn}
}

func (s *screenshoter) CaptureScreen() (*image.RGBA, error) {
	screenRectangle, err := s.getScreenRectangle()
	if err != nil {
		return nil, err
	}

	return s.CaptureRectangle(screenRectangle)
}

func (s *screenshoter) CaptureRectangle(rectangle image.Rectangle) (*image.RGBA, error) {

	screenInfo := xproto.Setup(s.conn).DefaultScreen(s.conn)
	x, y := rectangle.Dx(), rectangle.Dy()

	xImg, err := xproto.GetImage(s.conn, xproto.ImageFormatZPixmap, xproto.Drawable(screenInfo.Root), int16(rectangle.Min.X), int16(rectangle.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{
		Pix: data,
		Stride: 4 * x,
		Rect: image.Rect(0, 0, x, y),
	}

	return img, nil
}

func (s *screenshoter) getScreenRectangle() (image.Rectangle, error) {
	screen := xproto.Setup(s.conn).DefaultScreen(s.conn)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	return image.Rect(0, 0, int(x), int(y)), nil
}
