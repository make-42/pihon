package screenfader

import (
	"math"
	"time"

	"github.com/fogleman/ease"
	"tokei.pi/screendriver"
	"tokei.pi/ui"
)

var ScreenFadeDelay = 5    //s Start fading screen after this duration
var ScreenFadeDuration = 1 //s
var LastClickTimestamp = time.Now().UnixMilli()

func ClickRegistered() {
	LastClickTimestamp = time.Now().UnixMilli()
	ui.UIScreenOff = false
}

func AdaptDisplayBrightness() {
	timeSinceLastClick := int(time.Now().UnixMilli() - LastClickTimestamp)
	if timeSinceLastClick < ScreenFadeDelay*1000 {
		screendriver.SetBrightness(255)
	} else if timeSinceLastClick < ScreenFadeDelay*1000+ScreenFadeDuration*1000 {
		t := 1.0 - float64(timeSinceLastClick-ScreenFadeDelay*1000)/float64(ScreenFadeDuration*1000)
		screendriver.SetBrightness(int(math.Round(ease.InOutQuart(t) * 255)))
	} else {
		screendriver.SetBrightness(0)
		ui.UIScreenOff = true
	}
}
