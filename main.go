package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("you need just one input")
		return
	}
	Ascii(args[0])
}

/* Ascii(args[0]) */

func Ascii(word string) string {

	if word == "" {
		return ""
	}

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("File read error", err)
		return word
	} // checking file reading

	content := string(data)
	lines := strings.Split(content, "\n")
	// transforming and splitting the text to the strings
	// all the strings are in the "lines" now

	// somehow i need to connect ascii symbols with file line by line
	/* char := 'B'
	   index := int(char) - 32 */

	for i := 0; i < 8; i++ {
		row := "" // making the horizontal output line

		for _, char := range row {

			// adding those spaaaces
			if char == ' ' {
				row += "      "
				continue
			}

			index := int(char) - 33

			line := index*9 + 1 + i // number of the line in file
			row += lines[line]      // taking out one line from the file from its number
			// useful stuff: row += is row = row + ...

		}

		fmt.Println(row)
	}
	return word
}
