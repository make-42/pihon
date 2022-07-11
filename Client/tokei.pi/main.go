package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"tokei.pi/extras"
	"tokei.pi/renderer"
	"tokei.pi/screendriver"
	"tokei.pi/screenfader"
	"tokei.pi/ui"
	"tokei.pi/weather"

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

/*
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
}*/

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
			fmt.Println("Back nav placeholder")
		}
	}
	if edge == "falling" {
		/*
			if checkIfBothButtonsArePressed("left") {
				return
			}*/
		fmt.Println("Left button pressed!")
		screenfader.ClickRegistered()
		ui.UIDisplayedMenu--
		if ui.UIDisplayedMenu < 1 {
			ui.UIDisplayedMenu = ui.UIWidgetCount
		}
		ui.UITopBarTitle = fmt.Sprintf("Widget set %d", ui.UIDisplayedMenu)
		ui.RenderUI()
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
		/*
			if checkIfBothButtonsArePressed("right") {
				return
			}*/
		fmt.Println("Right button pressed!")
		screenfader.ClickRegistered()
		ui.UIDisplayedMenu++
		if ui.UIDisplayedMenu > ui.UIWidgetCount {
			ui.UIDisplayedMenu = 1
		}
		ui.UITopBarTitle = fmt.Sprintf("Widget set %d", ui.UIDisplayedMenu)
		ui.RenderUI()
	}
}

func UIMainloop() {
	for true {
		// Rerender UI every second
		time.Sleep(time.Second)
		ui.RenderUI()
	}
}

func UIWeatherMainloop() {
	for true {
		weather.UpdateCurrentWeather()
		time.Sleep(240 * time.Second)
	}
}

func ScreenFaderMainloop() {
	for true {
		screenfader.AdaptDisplayBrightness()
		timeSinceLastClick := int(time.Now().UnixMilli() - screenfader.LastClickTimestamp)
		if timeSinceLastClick < screenfader.ScreenFadeDelay*1000 {
			time.Sleep(100 * time.Millisecond)
		} else if timeSinceLastClick < screenfader.ScreenFadeDelay*1000+screenfader.ScreenFadeDuration*1000 {
			time.Sleep(17 * time.Millisecond)
		} else {
			time.Sleep(1000 * time.Millisecond)
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
	ui.BuildVersion = "V2022.07.11"
	ui.UIMargin = 1
	ui.UITopBarHeight = 10
	ui.UILineHeight = 8
	ui.UIFontHorizontalSpacing = 9
	ui.UIMaxBookLines = 6
	ui.UIWidgetCount = 4
	// Display Driver
	screendriver.I2CBuses = []string{"1", "4"}
	// Extras
	extras.CloudflareIP = "1.1.1.1"
	// Weather
	weather.OpenWeatherMapAPIKey = "2ee79601c807b825fa622addcb4d9b57"
	weather.CityOfObservation = "Paris, FR"
	weather.Latitude = 48.834590
	weather.Longitude = 2.398260
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
	go UIMainloop()
	go UIWeatherMainloop()
	go ScreenFaderMainloop()
	// Ping every 60s
	for true {
		time.Sleep(time.Minute)
		ui.UIConnectedToInternet = extras.PingCloudflare()
	}

}
