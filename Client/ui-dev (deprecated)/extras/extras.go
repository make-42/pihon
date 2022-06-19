package extras

import (
	"log"
	"os"
	"strings"
)

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
	return strings.Replace(path, dirname, "~", 1)
}

func FitWithinCharacterLimits(characterLimit int, stringToFit string) string {
	if characterLimit > len(stringToFit) {
		return stringToFit
	}
	return string(stringToFit[:characterLimit-1])
}
