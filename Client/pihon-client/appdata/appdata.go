package appdata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var DataPath string

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func SaveBookData(bookHash string, bookPosition int) {
	_ = os.Remove(DataPath + bookHash)
	f, err := os.Create(DataPath + bookHash)
	checkError(err)
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d", bookPosition))
	fmt.Printf("Saving current position : %d\n", bookPosition)
}

func ReadBookData(bookHash string) int {
	if _, err := os.Stat(DataPath + bookHash); errors.Is(err, os.ErrNotExist) {
		SaveBookData(bookHash, 0)
		return 0
	}
	content, err := ioutil.ReadFile(DataPath + bookHash)
	checkError(err)
	i, err := strconv.Atoi(string(content))
	checkError(err)
	return i
}
