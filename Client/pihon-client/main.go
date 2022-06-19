package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"pihonclient/appdata"
	"pihonclient/extras"
	"pihonclient/phnparser"
	"pihonclient/renderer"
	"pihonclient/screendriver"
	"pihonclient/ui"

	"github.com/warthog618/gpiod"
	"github.com/warthog618/gpiod/device/rpi"
)

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func checkGPIOInitError(error error) {
	if error != nil {
		log.Fatal(fmt.Sprintf("RequestLine returned error: %s\n", error))
		if error == syscall.Errno(22) {
			log.Fatal("Note that the WithPullUp option requires kernel V5.5 or later - check your kernel version.")
		}
	}
}

func selectBook() {
	// Read book
	phnparser.LoadBook(ui.UISelectionIndex)
	ui.UIDisplayedMenu = 2
	ui.UISelectionIndex = appdata.ReadBookData(phnparser.LoadedBookHash)
	ui.RenderUI()
}

func checkIfBothButtonsArePressed(currentButton string) bool {
	if extras.LeftButtonState == true {
		if extras.RightButtonState == true {
			fmt.Println("Both buttons pressed!")
			ui.UISelectionIndex = ui.UISelectionIndexBuffer
			if ui.UIDisplayedMenu == 1 {
				selectBook()
				fmt.Println("Loaded a book!")
				return true
			} else if ui.UIDisplayedMenu == 2 {
				appdata.SaveBookData(phnparser.LoadedBookHash, ui.UISelectionIndex)
				ui.RenderUI()
				fmt.Println("Saved the position!")
				return true
			}
		}
	}
	return false
}

func leftEventHandler(evt gpiod.LineEvent) {
	edge := "rising"
	extras.LeftButtonState = false
	if evt.Type == gpiod.LineEventFallingEdge {
		edge = "falling"
		extras.LeftButtonState = true
		extras.LeftButtonPressedTime = time.Now()
	}
	if edge == "rising" {
		if time.Now().Sub(extras.LeftButtonPressedTime).Seconds() >= extras.NavigationHoldDuration {
			if ui.UIDisplayedMenu == 2 {
				ui.UITopBarTitle = extras.FitWithinCharacterLimits(phnparser.MaxLengthOfLine-2, extras.GetCWD()+phnparser.LibraryPath)
				ui.UIDisplayedMenu = 1
				ui.UISelectionIndex = 0
				ui.UISelectionIndexBuffer = ui.UISelectionIndex
				ui.RenderUI()
			}
		}
	}
	if edge == "falling" {
		if checkIfBothButtonsArePressed("left") {
			return
		}
		fmt.Println("Left button pressed!")
		if ui.UIDisplayedMenu == 1 {
			if ui.UISelectionIndex > 0 {
				ui.UISelectionIndexBuffer = ui.UISelectionIndex
				ui.UISelectionIndex--
			}
			ui.RenderUI()
		} else if ui.UIDisplayedMenu == 2 {
			if ui.UISelectionIndex > 0 {
				ui.UISelectionIndexBuffer = ui.UISelectionIndex
				ui.UISelectionIndex--
			}
			ui.RenderUI()
		}
	}
}

func rightEventHandler(evt gpiod.LineEvent) {
	edge := "rising"
	extras.RightButtonState = false
	if evt.Type == gpiod.LineEventFallingEdge {
		edge = "falling"
		extras.RightButtonState = true
	}
	if edge == "falling" {
		if checkIfBothButtonsArePressed("right") {
			return
		}
		fmt.Println("Right button pressed!")
		if ui.UIDisplayedMenu == 0 {
			fmt.Println("Displaying library!")
			// Library show
			phnparser.ScanLibraryFolder()
			ui.UITopBarTitle = extras.FitWithinCharacterLimits(phnparser.MaxLengthOfLine-2, extras.GetCWD()+phnparser.LibraryPath)
			ui.UIDisplayedMenu = 1
			ui.RenderUI()
		} else if ui.UIDisplayedMenu == 1 {
			if ui.UISelectionIndex < len(phnparser.LibraryHashes)-1 {
				ui.UISelectionIndexBuffer = ui.UISelectionIndex
				ui.UISelectionIndex++
			}
			ui.RenderUI()
		} else if ui.UIDisplayedMenu == 2 {
			if ui.UISelectionIndex*ui.UIMaxBookLines < len(phnparser.LoadedBookLines)-1 {
				ui.UISelectionIndexBuffer = ui.UISelectionIndex
				ui.UISelectionIndex++
			}
			ui.RenderUI()
		}

	}
}

func main() {
	// Config
	// Input
	leftButtonPinRef := rpi.J8p11
	rightButtonPinRef := rpi.J8p13
	debouncingPeriod := 10 * time.Millisecond
	extras.NavigationHoldDuration = 3
	// Screen
	renderer.ScreenCount = 2
	renderer.ScreenWidth = 128
	renderer.ScreenHeight = 64
	// UI
	ui.BuildVersion = "V2022.05.05"
	ui.UIMargin = 1
	ui.UITopBarHeight = 10
	ui.UILineHeight = 8
	ui.UIFontHorizontalSpacing = 9
	ui.UIMaxBookLines = 6
	// Display Driver
	screendriver.I2CBuses = []string{"1", "4"}
	// .phn Parser
	phnparser.LibraryPath = "library/"
	phnparser.MaxLengthOfLine = 28
	// Appdata
	appdata.DataPath = "data/"
	// Extras
	extras.CloudflareIP = "1.1.1.1"
	// Code
	// GPIO init
	l, err := gpiod.RequestLine("gpiochip0", leftButtonPinRef, gpiod.WithPullUp, gpiod.WithBothEdges, gpiod.WithEventHandler(leftEventHandler), gpiod.WithDebounce(debouncingPeriod))
	checkGPIOInitError(err)
	defer l.Close()
	r, err := gpiod.RequestLine("gpiochip0", rightButtonPinRef, gpiod.WithPullUp, gpiod.WithBothEdges, gpiod.WithEventHandler(rightEventHandler), gpiod.WithDebounce(debouncingPeriod))
	checkGPIOInitError(err)
	defer r.Close()
	// Displays init
	screendriver.InitializeDisplays()
	// Renderer init
	renderer.InitFrameBuffers()
	// Boot screen show
	ui.UIDisplayedMenu = 0
	ui.RenderUI()
	ui.UIConnectedToInternet = extras.PingCloudflare()
	// Ping every 60s
	for true {
		time.Sleep(time.Minute)
		ui.UIConnectedToInternet = extras.PingCloudflare()
	}
}
