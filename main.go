package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]

	// Read the banner file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	banner := string(data)

	// Turn the \n from the command into real newlines
	input = convertNewLines(input)

	// Print the input
	printAscii(input, banner)
}

func convertNewLines(input string) string {
	result := ""

	for i := 0; i < len(input); i++ {
		if input[i] == '\\' && i+1 < len(input) && input[i+1] == 'n' {
			result += "\n"
			i++
		} else {
			result += string(input[i])
		}
	}

	return result
}

func printAscii(input string, banner string) {
	bannerLines := splitLines(banner)
	inputLines := splitLines(input)

	for _, line := range inputLines {

		// An empty line means print one empty line
		if line == "" {
			fmt.Println()
			continue
		}

		// Print the 8 rows of the ASCII characters
		for row := 0; row < 8; row++ {

			for i := 0; i < len(line); i++ {
				char := line[i]

				// ASCII characters start at 32 (space)
				charNumber := int(char) - 32

				// Each character takes 9 lines
				start := charNumber * 9

				fmt.Print(bannerLines[start+row])
			}

			fmt.Println()
		}
	}
}

func splitLines(text string) []string {
	lines := []string{}
	current := ""

	for i := 0; i < len(text); i++ {
		if text[i] == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(text[i])
		}
	}

	lines = append(lines, current)

	return lines
}
