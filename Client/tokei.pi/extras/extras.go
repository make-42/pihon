package extras

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-ping/ping"
)

var LeftButtonState = false
var RightButtonState = false
var LeftButtonPressedTime time.Time
var NavigationHoldDuration float64
var CloudflareIP string

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func GetCWD() string {
	path, err := os.Getwd()
	checkError(err)
	dirname, err := os.UserHomeDir()
	checkError(err)
	return strings.Replace(path+"/", dirname, "~", 1)
}

func FitWithinCharacterLimits(characterLimit int, stringToFit string) string {
	if characterLimit > len(stringToFit) {
		return stringToFit
	}
	return string(stringToFit[:characterLimit-1])
}

func PingCloudflare() bool {
	pinger, err := ping.NewPinger(CloudflareIP)
	if err != nil {
		panic(err)
	}
	pinger.Count = 1
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	if stats.PacketLoss == 0 {
		return true
	} else {
		return false
	}
}
