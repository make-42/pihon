package screendriver

import (
	"image"
	"log"
	"pihonclient/renderer"

	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ssd1306"
	"periph.io/x/host/v3"
)

var DisplaysList []*ssd1306.Dev
var I2CBuses []string

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}
func InitializeDisplays() {
	// Load all the drivers:
	_, err := host.Init()
	checkError(err)
	for _, busString := range I2CBuses {
		// Open a handle to the I²C bus:
		bus, err := i2creg.Open(busString)
		checkError(err)

		// Open a handle to a ssd1306 connected on the I²C bus:
		dev, err := ssd1306.NewI2C(bus, &ssd1306.DefaultOpts)
		checkError(err)
		DisplaysList = append(DisplaysList, dev)
	}
}

func UpdateDisplays() {
	for i := 0; i < renderer.ScreenCount; i++ {
		go DisplaysList[i].Draw(renderer.FrameBuffers[i].Bounds(), renderer.FrameBuffers[i], image.Point{})
	}
}
