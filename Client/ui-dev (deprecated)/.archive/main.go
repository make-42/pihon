package main

import (
	"fmt"
	"log"

	"uidev/renderer"
)

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func leftButtonPressed() {
	fmt.Println("Left button pressed!")
}

func rightButtonPressed() {
	fmt.Println("Right button pressed!")
}

func main() {
	// Config
	// Screen
	screenCount := 2
	screenWidth := 128
	screenHeight := 64
	// UI
	UIMargin := 1
	// Renderer
	//fontLocation := "./assets/Sauce Code Pro Nerd Font Complete.ttf"
	// Code
	// Renderer init
	for i := 0; i < screenCount; i++ {
		renderer.InitFrameBuffer(screenWidth, screenHeight)
	}
	// Font init
	//renderer.LoadFontFile(fontLocation)
	// Top bar
	renderer.DrawHorizontalLine(renderer.FrameBuffers[0], 0, 10, 128, 1)
	renderer.DrawHorizontalLine(renderer.FrameBuffers[1], 0, 10, 128, 1)
	renderer.AddText(renderer.FrameBuffers[0], UIMargin, UIMargin, 1, "Library @ ~/ontake/Books/")
	renderer.AddText(renderer.FrameBuffers[1], UIMargin-screenWidth, UIMargin, 1, "Library @ ~/ontake/Books/")
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1)
}
