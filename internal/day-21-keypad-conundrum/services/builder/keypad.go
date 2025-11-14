package day21keypadconundrum_services_builder

import (
	"image"

	common "github.com/ewoutquax/advent-of-code-2024/internal/day-21-keypad-conundrum/common"
)

type (
	FuncKeypadOption func(*common.Keypad)
)

func BuildKeypad(opts ...FuncKeypadOption) common.Keypad {
	var defaultKeypad = common.Keypad{}
	for _, fnOpt := range opts {
		fnOpt(&defaultKeypad)
	}

	return defaultKeypad
}

func WithLayout(layout common.KeypadLayout) FuncKeypadOption {
	return func(keypad *common.Keypad) {
		var keypadLayout []string

		switch layout {
		case common.KeypadLayoutNumeric:
			keypadLayout = keypadNumericLayout()
		case common.KeypadLayoutDirectional:
			keypadLayout = keypadDirectionalLayout()
		default:
			panic("No valid case found")
		}
		parsedKeypad := parseKeypad(keypadLayout)

		keypad.Keys = parsedKeypad.Keys
		keypad.Start = parsedKeypad.Start
	}
}

func parseKeypad(lines []string) common.Keypad {
	var keypad = common.Keypad{
		Keys: make(map[common.Location]string, 10),
	}

	for y := 1; y < len(lines); y += 2 {
		for x := 2; x < len(lines[0]); x += 4 {
			char := string(lines[y][x])
			if char != " " {
				convertedLocation := image.Pt((x-2)/4, (y-1)/2)
				keypad.Keys[convertedLocation] = char
				if char == "A" {
					keypad.Start = convertedLocation
				}
			}
		}
	}

	return keypad
}

func keypadNumericLayout() []string {
	return []string{
		"+---+---+---+",
		"| 7 | 8 | 9 |",
		"+---+---+---+",
		"| 4 | 5 | 6 |",
		"+---+---+---+",
		"| 1 | 2 | 3 |",
		"+---+---+---+",
		"    | 0 | A |",
		"    +---+---+",
	}
}

func keypadDirectionalLayout() []string {
	return []string{
		"    +---+---+",
		"    | ^ | A |",
		"+---+---+---+",
		"| < | v | > |",
		"+---+---+---+",
	}
}
