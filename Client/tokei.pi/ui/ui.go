package ui

import (
	"fmt"

	"tokei.pi/renderer"
	"tokei.pi/screendriver"
	"tokei.pi/tokei"
	"tokei.pi/weather"
)

var BuildVersion string

var UIDisplayedMenu = 0 // 0 - Boot ; 1-inf widgets
var UIWidgetCount int   // Count of widgets
var UITopBarTitle string
var UIConnectedToInternet = false

var UIMargin int
var UITopBarHeight int
var UILineHeight int
var UIFontHorizontalSpacing int

var UIMaxBookLines int

var UIFrameCount = 0

func RenderUI() {
	renderer.PerformanceTrackerStart()
	renderer.BlankScreen(0)
	switch UIDisplayedMenu {
	case 0:
		renderer.DrawPihonLogo()
		renderer.AddTextAllDisplays(UIMargin, renderer.ScreenHeight-UILineHeight-UIMargin+1, 1, BuildVersion)
	case 1:
		DrawTopBar(UIMargin, UITopBarHeight)
		// Left window
		renderer.DrawWidgetWindow(UIMargin, UITopBarHeight+UIMargin+1, renderer.ScreenWidth-UIMargin*2, renderer.ScreenHeight-UITopBarHeight-UIMargin*2-1)
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+2, 1, tokei.GetCurrentWeekdayString())
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight+2, 1, tokei.GetCurrentDateString())
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*2+2, 1, tokei.GetCurrentTimeString())
		renderer.DrawSunriseIcon(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*3+2)
		renderer.AddTextAllDisplays(UIMargin*2+9, UITopBarHeight+UIMargin*2+UILineHeight*3+2, 1, tokei.GetTimeStringFromTimeStampLowPrecision(int64(weather.Sunrise)))
		renderer.DrawSunsetIcon(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*4+2)
		renderer.AddTextAllDisplays(UIMargin*2+9, UITopBarHeight+UIMargin*2+UILineHeight*4+2, 1, tokei.GetTimeStringFromTimeStampLowPrecision(int64(weather.Sunset)))
		// right window
		renderer.DrawWidgetWindow(renderer.ScreenWidth+UIMargin, UITopBarHeight+UIMargin+1, renderer.ScreenWidth-UIMargin*2, renderer.ScreenHeight-UITopBarHeight-UIMargin*2-1)
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+2, 1, weather.CityOfObservation)
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight+2, 1, fmt.Sprintf("%0.1f°C", weather.Temperature))
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*2+2, 1, weather.Conditions)
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*3+2, 1, fmt.Sprintf("AQI %d", weather.AQI))
	case 2:
		DrawTopBar(UIMargin, UITopBarHeight)
		// Left window
		renderer.DrawWidgetWindow(UIMargin, UITopBarHeight+UIMargin+1, renderer.ScreenWidth-UIMargin*2, renderer.ScreenHeight-UITopBarHeight-UIMargin*2-1)
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+2, 1, fmt.Sprintf("AQI %d", weather.AQI))
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight+2, 1, fmt.Sprintf("CO  %0.1fμg/m3", weather.Components.CO))
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*2+2, 1, fmt.Sprintf("NO  %0.1fμg/m3", weather.Components.NO))
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*3+2, 1, fmt.Sprintf("NO2 %0.1fμg/m3", weather.Components.NO2))
		renderer.AddTextAllDisplays(UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*4+2, 1, fmt.Sprintf("O3  %0.1fμg/m3", weather.Components.CO))
		// right window
		renderer.DrawWidgetWindow(renderer.ScreenWidth+UIMargin, UITopBarHeight+UIMargin+1, renderer.ScreenWidth-UIMargin*2, renderer.ScreenHeight-UITopBarHeight-UIMargin*2-1)
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+2, 1, fmt.Sprintf("SO2 %0.1fμg/m3", weather.Components.SO2))
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight+2, 1, fmt.Sprintf("NH3 %0.1fμg/m3", weather.Components.NH3))
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*3+2, 1, fmt.Sprintf("PM2.5 %0.1fμg/m3", weather.Components.PM2_5))
		renderer.AddTextAllDisplays(renderer.ScreenWidth+UIMargin*2+1, UITopBarHeight+UIMargin*2+UILineHeight*4+2, 1, fmt.Sprintf("PM10  %0.1fμg/m3", weather.Components.PM10))
	case 3:
		DrawTopBar(UIMargin, UITopBarHeight)
	case 4:
		DrawTopBar(UIMargin, UITopBarHeight)
	default:
		fmt.Printf("Error! Menu does not exist!")
	}
	fmt.Printf("Frame render took %fms.\n", float64(renderer.PerformanceTrackerEnd())/float64(1000000))
	screendriver.UpdateDisplays()
	//renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[0], 0, UIFrameCount)
	//renderer.UpdateFrameBuffersToFiles(renderer.FrameBuffers[1], 1, UIFrameCount)
	UIFrameCount++
}

func DrawTopBar(UIMargin int, UITopBarHeight int) {
	// Top bar
	renderer.DrawHorizontalLine(0, UITopBarHeight, 256, 1)
	renderer.AddTextAllDisplays(UIMargin, UIMargin, 1, UITopBarTitle)
	renderer.DrawLightningIcon()
	if UIConnectedToInternet {
		renderer.DrawWifiEnabledIcon()
	} else {
		renderer.DrawWifiDisabledIcon()
	}
}
