package hw02unpackstring //hw02unpackstring

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
			fmt.Println("BooM1!")
			return false
		}

		if runeStr[i] == slash && (runeStr[i+1] != slash && !unicode.IsDigit(runeStr[i+1])) {
			fmt.Println("BooM2!")
			if len(runeStr) == 2 {
				return false
			}
			// Additional check
			if i-2 >= 0 { // `\\\a`
				if runeStr[i-1] == slash && runeStr[i-2] == slash {
					return false
				} else if runeStr[i-1] == slash && runeStr[i-2] != slash {
					return true
				} else if runeStr[i-1] != slash && runeStr[i-2] != slash { // `qw\ne`
					return false
				}
			} else { // `\\a`
				return true
			}
		}
	}

	return true
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
	var offset int = 0

	for pos := 0; pos <= len(runeStr)-1; {
		char := runeStr[pos]

		// Check last char
		if pos == len(runeStr)-1 {
			fmt.Fprintf(&output, "%c", char)
			break
		}

		// Main loop
		nextChar := runeStr[pos+1]
		//fmt.Printf("Cur = %c, Next = %c\n", char, nextChar)

		// Slash processing
		if char == slash { // `\\\\4`
			if runeStr[pos+1] == slash {
				if pos+2 < len(runeStr) {
					if unicode.IsDigit(runeStr[pos+2]) {
						repeat_num, _ := strconv.Atoi(string(runeStr[pos+2])) // Skip conversion error check
						tmp := strings.Repeat(string(slash), repeat_num)
						fmt.Fprintf(&output, "%s", tmp)
						// Skip slash (1) + next slash (1) + count (1)
						offset = 3
					} else {
						fmt.Fprintf(&output, "%c", slash)
						// Skip slash (1) + next slash (1)
						fmt.Println(output.String())
						offset = 2
					}
				}
			} else {
				// Skip slash (1)
				offset = 1
			}

		} else {
			// General execution flow
			if unicode.IsDigit(nextChar) {
				repeat_num, _ := strconv.Atoi(string(nextChar)) // Skip conversion error check
				tmp := strings.Repeat(string(char), repeat_num)
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
		fmt.Printf("Pos = %d, Offset = %d\n", pos, offset)
	}

	//fmt.Println(len(runeStr))
	fmt.Println(output.String())
	return output.String(), nil
}
