package phnparser

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var LibraryPath string

var LibraryTitles []string
var LibraryAuthors []string
var LibraryFileLocations []string
var LibraryFormats []string
var LibraryFileSizes []int
var LibraryHashes []string

var LoadedBookLines []string
var LoadedBookTitle string
var LoadedBookAuthor string
var LoadedBookHash string

var MaxLengthOfLine int

func checkError(error error) {
	if error != nil {
		log.Fatal(error)
	}
}

func ScanLibraryFolder() {
	files, err := ioutil.ReadDir(LibraryPath)
	checkError(err)
	var authors []string
	var titles []string
	var fileLocations []string
	var formats []string
	var fileSizes []int
	var hashes []string
	for _, file := range files {
		if !file.IsDir() {
			content, err := ioutil.ReadFile(LibraryPath + file.Name())
			checkError(err)
			fileSizes = append(fileSizes, len(content))
			formats = append(formats, "PHN")
			fileLocations = append(fileLocations, LibraryPath+file.Name())
			authors = append(authors, strings.Split(string(strings.Split(string(content), "\n-+---+-BOOK CONTENT-+---+-")[0]), "-+---+-BOOK AUTHOR-+---+-\n")[1])
			titles = append(titles, strings.Split(string(content), "\n-+---+-BOOK AUTHOR-+---+-")[0][25:])
			hashes = append(hashes, fmt.Sprintf("%x", md5.Sum(content)))
		}
	}
	LibraryTitles = titles
	LibraryAuthors = authors
	LibraryFileLocations = fileLocations
	LibraryFormats = formats
	LibraryFileSizes = fileSizes
	LibraryHashes = hashes
}

func LoadBook(bookIndex int) {
	LoadedBookLines = []string{}
	LoadedBookTitle = LibraryTitles[bookIndex]
	LoadedBookAuthor = LibraryAuthors[bookIndex]
	LoadedBookHash = LibraryHashes[bookIndex]
	content, err := ioutil.ReadFile(LibraryFileLocations[bookIndex])
	checkError(err)
	lines := strings.Split(strings.Split(string(content), "-+---+-BOOK CONTENT-+---+-\n")[1], "\n")
	for _, line := range lines {
		lineWithReturn := line + "\n"
		currentBuffer := []rune{}
		for _, character := range lineWithReturn {
			if (len(currentBuffer) >= MaxLengthOfLine) || (string(character) == "\n") {
				LoadedBookLines = append(LoadedBookLines, string(currentBuffer))
				currentBuffer = []rune{}
			}
			currentBuffer = append(currentBuffer, character)
		}
	}
}
