package renderer

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"

	"github.com/pbnjay/pixfont"
	/*
		"github.com/golang/freetype/truetype"
		"golang.org/x/image/font"
		"golang.org/x/image/math/fixed"
	*/)

var FrameBuffers []*image.Gray

//var FontBuffer *truetype.Font

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func InitFrameBuffer(bufferWidth int, bufferHeight int) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{bufferWidth, bufferHeight}
	img := image.NewGray(image.Rectangle{upLeft, lowRight})
	FrameBuffers = append(FrameBuffers, img)
}

func DrawPixel(frameBuffer *image.Gray, x int, y int, v int) {
	frameBuffer.Set(x, y, color.Gray{uint8(255 * v)})
}

func DrawHorizontalLine(frameBuffer *image.Gray, startX int, startY, length int, v int) {
	for i := 0; i < length; i++ {
		DrawPixel(frameBuffer, startX+i, startY, v)
	}
}

func DrawVerticalLine(frameBuffer *image.Gray, startX int, startY, length int, v int) {
	for i := 0; i < length; i++ {
		DrawPixel(frameBuffer, startX, startY+i, v)
	}
}

/*
func LoadFontFile(fontFileLocation string) {
	// Read the font data.
	fontBytes, err := ioutil.ReadFile(fontFileLocation)
	checkError(err)
	FontBuffer, err = truetype.Parse(fontBytes)
	checkError(err)
}
*/

func AddText(frameBuffer *image.Gray, x int, y int, v int, label string) {
	col := color.Gray{uint8(255 * v)}
	pixfont.DrawString(frameBuffer, x, y, label, col)
	/*
		point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

		d := &font.Drawer{
			Dst: frameBuffer,
			Src: image.NewUniform(col),
			Face: truetype.NewFace(FontBuffer, &truetype.Options{
				Size:    float64(fontSize),
				DPI:     float64(dpi),
				Hinting: font.HintingNone,
			}),
			Dot: point,
		}
		d.DrawString(label)
	*/
}

func UpdateFrameBuffersToFiles(frameBuffer *image.Gray, bufferIndex int) {
	f, _ := os.Create(strconv.Itoa(bufferIndex) + "img.png")
	png.Encode(f, frameBuffer)
}
