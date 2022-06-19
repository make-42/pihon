package renderer

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/pbnjay/pixfont"
)

var FrameBuffers []*image.Gray
var PerformanceTrackerTimer int64
var ScreenCount int
var ScreenWidth int
var ScreenHeight int

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func InitFrameBuffers() {
	for i := 0; i < ScreenCount; i++ {
		upLeft := image.Point{0, 0}
		lowRight := image.Point{ScreenWidth, ScreenHeight}
		img := image.NewGray(image.Rectangle{upLeft, lowRight})
		FrameBuffers = append(FrameBuffers, img)
	}
}

func DrawPixel(x int, y int, v int) {
	FrameBuffers[int(math.Floor(float64(x)/float64(ScreenWidth)))].Set(x-int(math.Floor(float64(x)/float64(ScreenWidth)))*ScreenWidth, y, color.Gray{uint8(255 * v)})
}

func DrawHorizontalLine(startX int, startY, length int, v int) {
	for i := 0; i < length; i++ {
		DrawPixel(startX+i, startY, v)
	}
}

func DrawVerticalLine(frameBuffer *image.Gray, startX int, startY, length int, v int) {
	for i := 0; i < length; i++ {
		DrawPixel(startX, startY+i, v)
	}
}

func AddTextAllDisplays(x int, y int, v int, label string) {
	col := color.Gray{uint8(255 * v)}
	for i := 0; i < ScreenCount; i++ {
		pixfont.DrawString(FrameBuffers[i], x-i*ScreenWidth, y, label, col)
	}
}

func BlankScreen(v int) {
	for i := 0; i < ScreenCount; i++ {
		draw.Draw(FrameBuffers[i], FrameBuffers[i].Bounds(), &image.Uniform{color.Gray{uint8(255 * v)}}, image.ZP, draw.Src)
	}
}

func PerformanceTrackerStart() {
	PerformanceTrackerTimer = time.Now().UTC().UnixNano()
}

func PerformanceTrackerEnd() int64 {
	return time.Now().UTC().UnixNano() - PerformanceTrackerTimer
}

func DrawPihonLogo() {
	DrawHorizontalLine(113, 4, 8, 1)
	DrawHorizontalLine(135, 4, 8, 1)
	DrawHorizontalLine(110, 5, 14, 1)
	DrawHorizontalLine(132, 5, 14, 1)
	DrawHorizontalLine(108, 6, 3, 1)
	DrawHorizontalLine(113, 6, 1, 1)
	DrawHorizontalLine(121, 6, 1, 1)
	DrawHorizontalLine(123, 6, 3, 1)
	DrawHorizontalLine(130, 6, 3, 1)
	DrawHorizontalLine(134, 6, 1, 1)
	DrawHorizontalLine(142, 6, 1, 1)
	DrawHorizontalLine(145, 6, 3, 1)
	DrawHorizontalLine(108, 7, 1, 1)
	DrawHorizontalLine(125, 7, 2, 1)
	DrawHorizontalLine(129, 7, 2, 1)
	DrawHorizontalLine(147, 7, 1, 1)
	DrawHorizontalLine(108, 8, 2, 1)
	DrawHorizontalLine(125, 8, 2, 1)
	DrawHorizontalLine(129, 8, 2, 1)
	DrawHorizontalLine(146, 8, 2, 1)
	DrawHorizontalLine(108, 9, 1, 1)
	DrawHorizontalLine(126, 9, 4, 1)
	DrawHorizontalLine(141, 9, 1, 1)
	DrawHorizontalLine(147, 9, 1, 1)
	DrawHorizontalLine(108, 10, 2, 1)
	DrawHorizontalLine(116, 10, 1, 1)
	DrawHorizontalLine(126, 10, 4, 1)
	DrawHorizontalLine(139, 10, 1, 1)
	DrawHorizontalLine(146, 10, 2, 1)
	DrawHorizontalLine(108, 11, 2, 1)
	DrawHorizontalLine(118, 11, 2, 1)
	DrawHorizontalLine(127, 11, 2, 1)
	DrawHorizontalLine(136, 11, 2, 1)
	DrawHorizontalLine(146, 11, 2, 1)
	DrawHorizontalLine(109, 12, 2, 1)
	DrawHorizontalLine(120, 12, 1, 1)
	DrawHorizontalLine(127, 12, 2, 1)
	DrawHorizontalLine(135, 12, 1, 1)
	DrawHorizontalLine(145, 12, 2, 1)
	DrawHorizontalLine(109, 13, 2, 1)
	DrawHorizontalLine(121, 13, 2, 1)
	DrawHorizontalLine(126, 13, 4, 1)
	DrawHorizontalLine(133, 13, 2, 1)
	DrawHorizontalLine(145, 13, 2, 1)
	DrawHorizontalLine(110, 14, 2, 1)
	DrawHorizontalLine(122, 14, 2, 1)
	DrawHorizontalLine(126, 14, 4, 1)
	DrawHorizontalLine(132, 14, 2, 1)
	DrawHorizontalLine(144, 14, 2, 1)
	DrawHorizontalLine(110, 15, 2, 1)
	DrawHorizontalLine(124, 15, 8, 1)
	DrawHorizontalLine(144, 15, 2, 1)
	DrawHorizontalLine(111, 16, 3, 1)
	DrawHorizontalLine(124, 16, 8, 1)
	DrawHorizontalLine(142, 16, 3, 1)
	DrawHorizontalLine(111, 17, 3, 1)
	DrawHorizontalLine(123, 17, 10, 1)
	DrawHorizontalLine(142, 17, 2, 1)
	DrawHorizontalLine(113, 18, 2, 1)
	DrawHorizontalLine(122, 18, 12, 1)
	DrawHorizontalLine(141, 18, 2, 1)
	DrawHorizontalLine(114, 19, 28, 1)
	DrawHorizontalLine(114, 20, 11, 1)
	DrawHorizontalLine(131, 20, 11, 1)
	DrawHorizontalLine(113, 21, 3, 1)
	DrawHorizontalLine(121, 21, 2, 1)
	DrawHorizontalLine(133, 21, 2, 1)
	DrawHorizontalLine(140, 21, 3, 1)
	DrawHorizontalLine(112, 22, 2, 1)
	DrawHorizontalLine(120, 22, 2, 1)
	DrawHorizontalLine(134, 22, 2, 1)
	DrawHorizontalLine(142, 22, 2, 1)
	DrawHorizontalLine(111, 23, 3, 1)
	DrawHorizontalLine(120, 23, 2, 1)
	DrawHorizontalLine(134, 23, 2, 1)
	DrawHorizontalLine(143, 23, 2, 1)
	DrawHorizontalLine(111, 24, 2, 1)
	DrawHorizontalLine(118, 24, 5, 1)
	DrawHorizontalLine(134, 24, 3, 1)
	DrawHorizontalLine(143, 24, 2, 1)
	DrawHorizontalLine(110, 25, 2, 1)
	DrawHorizontalLine(117, 25, 7, 1)
	DrawHorizontalLine(132, 25, 7, 1)
	DrawHorizontalLine(144, 25, 2, 1)
	DrawHorizontalLine(110, 26, 2, 1)
	DrawHorizontalLine(116, 26, 24, 1)
	DrawHorizontalLine(144, 26, 2, 1)
	DrawHorizontalLine(110, 27, 2, 1)
	DrawHorizontalLine(115, 27, 4, 1)
	DrawHorizontalLine(124, 27, 8, 1)
	DrawHorizontalLine(138, 27, 3, 1)
	DrawHorizontalLine(144, 27, 2, 1)
	DrawHorizontalLine(110, 28, 2, 1)
	DrawHorizontalLine(113, 28, 4, 1)
	DrawHorizontalLine(125, 28, 6, 1)
	DrawHorizontalLine(139, 28, 4, 1)
	DrawHorizontalLine(144, 28, 2, 1)
	DrawHorizontalLine(109, 29, 7, 1)
	DrawHorizontalLine(140, 29, 7, 1)
	DrawHorizontalLine(108, 30, 8, 1)
	DrawHorizontalLine(127, 30, 2, 1)
	DrawHorizontalLine(141, 30, 3, 1)
	DrawHorizontalLine(145, 30, 3, 1)
	DrawHorizontalLine(107, 31, 3, 1)
	DrawHorizontalLine(112, 31, 3, 1)
	DrawHorizontalLine(127, 31, 2, 1)
	DrawHorizontalLine(141, 31, 2, 1)
	DrawHorizontalLine(146, 31, 3, 1)
	DrawHorizontalLine(107, 32, 2, 1)
	DrawHorizontalLine(113, 32, 2, 1)
	DrawHorizontalLine(127, 32, 2, 1)
	DrawHorizontalLine(141, 32, 2, 1)
	DrawHorizontalLine(147, 32, 2, 1)
	DrawHorizontalLine(106, 33, 2, 1)
	DrawHorizontalLine(113, 33, 1, 1)
	DrawHorizontalLine(122, 33, 12, 1)
	DrawHorizontalLine(142, 33, 1, 1)
	DrawHorizontalLine(148, 33, 2, 1)
	DrawHorizontalLine(106, 34, 2, 1)
	DrawHorizontalLine(113, 34, 1, 1)
	DrawHorizontalLine(127, 34, 2, 1)
	DrawHorizontalLine(142, 34, 1, 1)
	DrawHorizontalLine(148, 34, 2, 1)
	DrawHorizontalLine(106, 35, 2, 1)
	DrawHorizontalLine(113, 35, 2, 1)
	DrawHorizontalLine(126, 35, 4, 1)
	DrawHorizontalLine(141, 35, 2, 1)
	DrawHorizontalLine(148, 35, 2, 1)
	DrawHorizontalLine(106, 36, 2, 1)
	DrawHorizontalLine(112, 36, 3, 1)
	DrawHorizontalLine(125, 36, 6, 1)
	DrawHorizontalLine(141, 36, 3, 1)
	DrawHorizontalLine(148, 36, 2, 1)
	DrawHorizontalLine(106, 37, 2, 1)
	DrawHorizontalLine(112, 37, 3, 1)
	DrawHorizontalLine(124, 37, 2, 1)
	DrawHorizontalLine(127, 37, 2, 1)
	DrawHorizontalLine(130, 37, 2, 1)
	DrawHorizontalLine(141, 37, 3, 1)
	DrawHorizontalLine(148, 37, 2, 1)
	DrawHorizontalLine(106, 38, 2, 1)
	DrawHorizontalLine(112, 38, 4, 1)
	DrawHorizontalLine(123, 38, 2, 1)
	DrawHorizontalLine(127, 38, 2, 1)
	DrawHorizontalLine(131, 38, 2, 1)
	DrawHorizontalLine(140, 38, 4, 1)
	DrawHorizontalLine(148, 38, 2, 1)
	DrawHorizontalLine(106, 39, 2, 1)
	DrawHorizontalLine(112, 39, 6, 1)
	DrawHorizontalLine(122, 39, 2, 1)
	DrawHorizontalLine(125, 39, 6, 1)
	DrawHorizontalLine(132, 39, 2, 1)
	DrawHorizontalLine(137, 39, 7, 1)
	DrawHorizontalLine(147, 39, 3, 1)
	DrawHorizontalLine(107, 40, 2, 1)
	DrawHorizontalLine(111, 40, 10, 1)
	DrawHorizontalLine(122, 40, 1, 1)
	DrawHorizontalLine(127, 40, 2, 1)
	DrawHorizontalLine(133, 40, 1, 1)
	DrawHorizontalLine(135, 40, 10, 1)
	DrawHorizontalLine(147, 40, 2, 1)
	DrawHorizontalLine(107, 41, 14, 1)
	DrawHorizontalLine(127, 41, 2, 1)
	DrawHorizontalLine(135, 41, 5, 1)
	DrawHorizontalLine(144, 41, 5, 1)
	DrawHorizontalLine(108, 42, 4, 1)
	DrawHorizontalLine(116, 42, 6, 1)
	DrawHorizontalLine(127, 42, 2, 1)
	DrawHorizontalLine(134, 42, 5, 1)
	DrawHorizontalLine(144, 42, 4, 1)
	DrawHorizontalLine(108, 43, 3, 1)
	DrawHorizontalLine(117, 43, 5, 1)
	DrawHorizontalLine(134, 43, 4, 1)
	DrawHorizontalLine(145, 43, 3, 1)
	DrawHorizontalLine(108, 44, 3, 1)
	DrawHorizontalLine(118, 44, 4, 1)
	DrawHorizontalLine(134, 44, 3, 1)
	DrawHorizontalLine(145, 44, 2, 1)
	DrawHorizontalLine(109, 45, 2, 1)
	DrawHorizontalLine(119, 45, 3, 1)
	DrawHorizontalLine(134, 45, 2, 1)
	DrawHorizontalLine(145, 45, 2, 1)
	DrawHorizontalLine(109, 46, 2, 1)
	DrawHorizontalLine(120, 46, 2, 1)
	DrawHorizontalLine(134, 46, 2, 1)
	DrawHorizontalLine(145, 46, 2, 1)
	DrawHorizontalLine(109, 47, 2, 1)
	DrawHorizontalLine(120, 47, 2, 1)
	DrawHorizontalLine(134, 47, 1, 1)
	DrawHorizontalLine(145, 47, 2, 1)
	DrawHorizontalLine(110, 48, 2, 1)
	DrawHorizontalLine(120, 48, 3, 1)
	DrawHorizontalLine(133, 48, 2, 1)
	DrawHorizontalLine(144, 48, 2, 1)
	DrawHorizontalLine(110, 49, 3, 1)
	DrawHorizontalLine(121, 49, 3, 1)
	DrawHorizontalLine(132, 49, 3, 1)
	DrawHorizontalLine(143, 49, 3, 1)
	DrawHorizontalLine(111, 50, 2, 1)
	DrawHorizontalLine(121, 50, 14, 1)
	DrawHorizontalLine(143, 50, 2, 1)
	DrawHorizontalLine(112, 51, 3, 1)
	DrawHorizontalLine(120, 51, 15, 1)
	DrawHorizontalLine(141, 51, 3, 1)
	DrawHorizontalLine(113, 52, 4, 1)
	DrawHorizontalLine(119, 52, 7, 1)
	DrawHorizontalLine(131, 52, 5, 1)
	DrawHorizontalLine(139, 52, 4, 1)
	DrawHorizontalLine(114, 53, 9, 1)
	DrawHorizontalLine(134, 53, 7, 1)
	DrawHorizontalLine(116, 54, 6, 1)
	DrawHorizontalLine(134, 54, 6, 1)
	DrawHorizontalLine(118, 55, 4, 1)
	DrawHorizontalLine(134, 55, 4, 1)
	DrawHorizontalLine(120, 56, 3, 1)
	DrawHorizontalLine(133, 56, 3, 1)
	DrawHorizontalLine(121, 57, 4, 1)
	DrawHorizontalLine(131, 57, 3, 1)
	DrawHorizontalLine(123, 58, 10, 1)
	DrawHorizontalLine(125, 59, 6, 1)
}

func DrawLightningIcon() {
	DrawHorizontalLine(253, 1, 1, 1)
	DrawHorizontalLine(252, 2, 1, 1)
	DrawHorizontalLine(251, 3, 2, 1)
	DrawHorizontalLine(250, 4, 5, 1)
	DrawHorizontalLine(249, 5, 5, 1)
	DrawHorizontalLine(251, 6, 2, 1)
	DrawHorizontalLine(251, 7, 1, 1)
	DrawHorizontalLine(250, 8, 1, 1)
}

func DrawWifiEnabledIcon() {
	DrawHorizontalLine(239, 1, 8, 1)
	DrawHorizontalLine(238, 2, 1, 1)
	DrawHorizontalLine(247, 2, 1, 1)
	DrawHorizontalLine(240, 3, 6, 1)
	DrawHorizontalLine(239, 4, 1, 1)
	DrawHorizontalLine(246, 4, 1, 1)
	DrawHorizontalLine(241, 5, 4, 1)
	DrawHorizontalLine(240, 6, 1, 1)
	DrawHorizontalLine(245, 6, 1, 1)
	DrawHorizontalLine(242, 7, 2, 1)
	DrawHorizontalLine(242, 8, 2, 1)
}

func DrawWifiDisabledIcon() {
	DrawHorizontalLine(239, 1, 8, 1)
	DrawHorizontalLine(238, 2, 1, 1)
	DrawHorizontalLine(245, 2, 1, 1)
	DrawHorizontalLine(247, 2, 1, 1)
	DrawHorizontalLine(240, 3, 6, 1)
	DrawHorizontalLine(239, 4, 1, 1)
	DrawHorizontalLine(243, 4, 1, 1)
	DrawHorizontalLine(246, 4, 1, 1)
	DrawHorizontalLine(241, 5, 4, 1)
	DrawHorizontalLine(240, 6, 2, 1)
	DrawHorizontalLine(245, 6, 1, 1)
	DrawHorizontalLine(240, 7, 1, 1)
	DrawHorizontalLine(242, 7, 2, 1)
	DrawHorizontalLine(239, 8, 1, 1)
	DrawHorizontalLine(242, 8, 2, 1)
}

func UpdateFrameBuffersToFiles(frameBuffer *image.Gray, bufferIndex int, sequenceIndex int) {
	f, _ := os.Create("./virtual-screens/" + strconv.Itoa(sequenceIndex) + "." + strconv.Itoa(bufferIndex) + "img.png")
	png.Encode(f, frameBuffer)
}
