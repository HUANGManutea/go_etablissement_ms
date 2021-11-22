package search

import (
	"bufio"
	"os"
	"sort"
)

func BinarySearchFile(filename string, target string, offset int, length int) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// extract all substrings
	substrings := make([]string, 0)

	// we need to store the length of the line for File.Seek
	lineLength := -1

	for scanner.Scan() {
		line := scanner.Text()

		// init length
		if lineLength < 0 {
			lineLength = len(line)
		}

		// extract
		substring := line[offset : offset+length]
		substrings = append(substrings, substring)
	}

	//use builtin binary search on substrings, position is our line number
	position := int64(sort.SearchStrings(substrings, target))

	// add 1 because of "\n"
	lineByteLength := lineLength + 1

	// we can now point at the line and read the line
	_, err = file.Seek(position*int64(lineByteLength), 0)
	if err != nil {
		return "", err
	}

	// lineByteLength is only used for Seek, we don't want to store "\n", so we use lineLength
	lineBuffer := make([]byte, lineLength)

	_, err = file.Read(lineBuffer)
	if err != nil {
		return "", err
	}

	return string(lineBuffer), nil

}
