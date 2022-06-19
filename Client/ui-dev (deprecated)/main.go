package main

import (
	"log"

	"uidev/appdata"
	"uidev/extras"
	"uidev/phnparser"
	"uidev/renderer"
	"uidev/ui"
)

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func main() {
	// Config
	// Screen
	renderer.ScreenCount = 2
	renderer.ScreenWidth = 128
	renderer.ScreenHeight = 64
	// UI
	ui.BuildVersion = "V2022.05.04"
	ui.UIMargin = 1
	ui.UITopBarHeight = 10
	ui.UILineHeight = 8
	ui.UIFontHorizontalSpacing = 9
	ui.UIMaxBookLines = 6
	// .phn Parser
	phnparser.LibraryPath = "library/"
	phnparser.MaxLengthOfLine = 28
	// Appdata
	appdata.DataPath = "data/"
	// Renderer
	// Code
	// Renderer init
	renderer.InitFrameBuffers()
	// Boot screen show
	ui.DisplayedMenu = 0
	ui.RenderUI()
	// Virtual screens save
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, 0)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, 0)
	// Library show
	phnparser.ScanLibraryFolder()
	ui.TopBarTitle = extras.FitWithinCharacterLimits(phnparser.MaxLengthOfLine-2, extras.GetCWD()+phnparser.LibraryPath)
	ui.DisplayedMenu = 1
	ui.RenderUI()
	// Virtual screens save
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, 1)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, 1)
	// Read book
	phnparser.LoadBook(ui.UISelectionIndex)
	ui.DisplayedMenu = 2
	ui.UISelectionIndex = appdata.ReadBookData(phnparser.LoadedBookHash)
	ui.RenderUI()
	// Virtual screens save
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, 2)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, 2)
	ui.UISelectionIndex = 1
	ui.RenderUI()
	// Virtual screens save
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, 3)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, 3)
	ui.UISelectionIndex = 2
	ui.RenderUI()
	// Virtual screens save
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, 4)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, 4)
	ui.UISelectionIndex = 3
	ui.RenderUI()
	// Virtual screens save
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, 5)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, 5)
}
