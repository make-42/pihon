package ui

import (
	"fmt"
	"uidev/phnparser"
	"uidev/renderer"
)

var BuildVersion string

var DisplayedMenu = 0 // 0 - Boot ; 1 - Book Select ; 2 - Book reading
var UISelectionIndex = 0
var TopBarTitle string

var UIMargin int
var UITopBarHeight int
var UILineHeight int
var UIFontHorizontalSpacing int

var UIMaxBookLines int

var UIFrameCount = 0

func RenderUI() {
	renderer.PerformanceTrackerStart()
	renderer.BlankScreen(0)
	switch DisplayedMenu {
	case 0:
		renderer.DrawPihonLogo()
		renderer.AddTextAllDisplays(UIMargin, renderer.ScreenHeight-UILineHeight-UIMargin+1, 1, BuildVersion)
	case 1:
		DrawTopBar(UIMargin, UITopBarHeight)
		// Book Info
		// Title
		renderer.AddTextAllDisplays(UIMargin, UIMargin*2+UITopBarHeight, 1, phnparser.LibraryTitles[UISelectionIndex])
		// Author
		renderer.AddTextAllDisplays(UIMargin, UIMargin*3+UILineHeight+UITopBarHeight, 1, phnparser.LibraryAuthors[UISelectionIndex])
		// Format / Size
		renderer.AddTextAllDisplays(UIMargin, renderer.ScreenHeight-UILineHeight-UIMargin+1, 1, phnparser.LibraryFormats[UISelectionIndex]+" "+fmt.Sprintf("%.2fMB", float64(phnparser.LibraryFileSizes[UISelectionIndex])/float64(1024*1024)))
		// Hash
		renderer.AddTextAllDisplays(UIMargin, renderer.ScreenHeight-2*UILineHeight-2*UIMargin+1, 1, phnparser.LibraryHashes[UISelectionIndex][:phnparser.MaxLengthOfLine])
		// Reading progress
		//renderer.AddTextAllDisplays(renderer.ScreenWidth*renderer.ScreenCount-len("1/123")*UIFontHorizontalSpacing-UIMargin+3, renderer.ScreenHeight-UIMargin-UILineHeight+1, 1, "1/123")
	case 2:
		TopBarTitle = phnparser.LoadedBookTitle
		DrawTopBar(UIMargin, UITopBarHeight)
		// Book text
		for i := 0; i < UIMaxBookLines; i++ {
			if UISelectionIndex*UIMaxBookLines+i < len(phnparser.LoadedBookLines) {
				renderer.AddTextAllDisplays(UIMargin, UIMargin*(2+i)+UITopBarHeight+i*UILineHeight, 1, phnparser.LoadedBookLines[UISelectionIndex*UIMaxBookLines+i])
			}
		}
	default:
		fmt.Printf("Error! Menu does not exist!")
	}
	fmt.Printf("Frame render took %fms.\n", float64(renderer.PerformanceTrackerEnd())/float64(1000000))
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, UIFrameCount)
	renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, UIFrameCount)
	UIFrameCount++
}

func DrawTopBar(UIMargin int, UITopBarHeight int) {
	// Top bar
	renderer.DrawHorizontalLine(0, UITopBarHeight, 256, 1)
	renderer.AddTextAllDisplays(UIMargin, UIMargin, 1, TopBarTitle)
	renderer.DrawLightningIcon()
	renderer.DrawWifiEnabledIcon()
}
