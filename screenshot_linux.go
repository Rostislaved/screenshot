// +build linux freebsd

package main

import (
	"image"
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

type screenshoter struct {
	conn       *xgb.Conn
	screenInfo *xproto.ScreenInfo
}

func New() *screenshoter {
	conn, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	screenInfo := xproto.Setup(conn).DefaultScreen(conn)

	return &screenshoter{
		conn:       conn,
		screenInfo: screenInfo,
	}
}

func (s *screenshoter) CaptureScreen() (img *image.RGBA, err error) {
	screenRectangle := s.getScreenRectangle()

	img, err  = s.CaptureRectangle(screenRectangle)

	return
}

func (s *screenshoter) CaptureRectangle(rectangle image.Rectangle) (*image.RGBA, error) {
	x, y := rectangle.Dx(), rectangle.Dy()

	xImg, err := xproto.GetImage(s.conn, xproto.ImageFormatZPixmap, xproto.Drawable(s.screenInfo.Root), int16(rectangle.Min.X), int16(rectangle.Min.Y), uint16(x), uint16(y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{
		Pix:    data,
		Stride: 4 * x,
		Rect:   image.Rect(0, 0, x, y),
	}

	return img, nil
}

func (s *screenshoter) getScreenRectangle() image.Rectangle {
	x := s.screenInfo.WidthInPixels
	y := s.screenInfo.HeightInPixels

	return image.Rect(0, 0, int(x), int(y))
}
