package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

const slash rune = 92

func IsStringValid(runeStr []rune) bool {
	if unicode.IsDigit(runeStr[0]) {
		return false
	}

	for i := 0; i < len(runeStr)-1; i++ {
		if unicode.IsDigit(runeStr[i]) && unicode.IsDigit(runeStr[i+1]) && runeStr[i-1] != slash {
			return false
		}

		if runeStr[i] == slash && runeStr[i+1] != slash && !unicode.IsDigit(runeStr[i+1]) {
			if len(runeStr) == 2 {
				return false
			}
			if i-2 >= 0 && (runeStr[i-1] == slash || runeStr[i-1] != slash) && (runeStr[i-2] == slash || runeStr[i-2] != slash) {
				return false
			}
		}
	}

	return true
}

func processSlash(runeStr []rune, pos int, output *strings.Builder) int {
	// Skip step -- default case = 1
	offset := 1
	if runeStr[pos+1] == slash && pos+2 < len(runeStr) {
		if unicode.IsDigit(runeStr[pos+2]) {
			repeatNum, _ := strconv.Atoi(string(runeStr[pos+2])) // Skip conversion error check
			tmp := strings.Repeat(string(slash), repeatNum)
			fmt.Fprintf(output, "%s", tmp)
			// Skip slash (1) + next slash (1) + count (1)
			offset = 3
		} else {
			fmt.Fprintf(output, "%c", slash)
			// Skip slash (1) + next slash (1)
			offset = 2
		}
	}
	return offset
}

func Unpack(inputString string) (string, error) {
	if len(inputString) == 0 {
		return "", nil
	}

	runeStr := []rune(inputString)

	isValidInStr := IsStringValid(runeStr)
	if !isValidInStr {
		return "", ErrInvalidString
	}

	var output strings.Builder
	var offset int

	for pos := 0; pos <= len(runeStr)-1; {
		char := runeStr[pos]

		// Check last char
		if pos == len(runeStr)-1 {
			fmt.Fprintf(&output, "%c", char)
			break
		}

		// Processing flow

		nextChar := runeStr[pos+1]

		switch char {
		// Slash processing
		case slash:
			offset = processSlash(runeStr, pos, &output)
		// General execution flow
		default:
			if unicode.IsDigit(nextChar) {
				repeatNum, _ := strconv.Atoi(string(nextChar)) // Skip conversion error check
				tmp := strings.Repeat(string(char), repeatNum)
				fmt.Fprintf(&output, "%s", tmp)
				// Skip char (1) + count (1)
				offset = 2
			} else {
				fmt.Fprintf(&output, "%c", char)
				// Skip char (1)
				offset = 1
			}
		}

		// Update pos
		pos += offset
		// fmt.Printf("Pos = %d, Offset = %d\n", pos, offset)
	}

	// fmt.Println(len(runeStr))
	fmt.Println(output.String())
	return output.String(), nil
}
