package main

import (
	"flag"
	"fmt"
	"strings"
)

const version = "1.0.0"

var asciiFont = map[rune][][]string{
	'A': {
		{"  █████  "},
		{" ██   ██ "},
		{"███████ "},
		{"██   ██ "},
		{"██   ██ "},
	},
	'B': {
		{"██████  "},
		{"██   ██ "},
		{"██████  "},
		{"██   ██ "},
		{"██████  "},
	},
	'C': {
		{" ██████ "},
		{"██      "},
		{"██      "},
		{"██      "},
		{" ██████ "},
	},
	'D': {
		{"██████  "},
		{"██   ██ "},
		{"██   ██ "},
		{"██   ██ "},
		{"██████  "},
	},
	'E': {
		{"███████ "},
		{"██      "},
		{"█████   "},
		{"██      "},
		{"███████ "},
	},
	'F': {
		{"███████ "},
		{"██      "},
		{"█████   "},
		{"██      "},
		{"██      "},
	},
	'G': {
		{" ██████  "},
		{"██       "},
		{"██   ███ "},
		{"██    ██ "},
		{" ██████  "},
	},
	'H': {
		{"██   ██ "},
		{"██   ██ "},
		{"███████ "},
		{"██   ██ "},
		{"██   ██ "},
	},
	'I': {
		{"██ "},
		{"██ "},
		{"██ "},
		{"██ "},
		{"██ "},
	},
	'J': {
		{"     ██ "},
		{"     ██ "},
		{"     ██ "},
		{"██   ██ "},
		{" █████  "},
	},
	'K': {
		{"██   ██ "},
		{"██  ██  "},
		{"█████   "},
		{"██  ██  "},
		{"██   ██ "},
	},
	'L': {
		{"██      "},
		{"██      "},
		{"██      "},
		{"██      "},
		{"███████ "},
	},
	'M': {
		{"███    ███ "},
		{"████  ████ "},
		{"██ ████ ██ "},
		{"██  ██  ██ "},
		{"██      ██ "},
	},
	'N': {
		{"███    ██ "},
		{"████   ██ "},
		{"██ ██  ██ "},
		{"██  ██ ██ "},
		{"██   ████ "},
	},
	'O': {
		{" ██████  "},
		{"██    ██ "},
		{"██    ██ "},
		{"██    ██ "},
		{" ██████  "},
	},
	'P': {
		{"██████  "},
		{"██   ██ "},
		{"██████  "},
		{"██      "},
		{"██      "},
	},
	'Q': {
		{" ██████  "},
		{"██    ██ "},
		{"██    ██ "},
		{"██ ▄▄ ██ "},
		{" ██████  "},
	},
	'R': {
		{"██████  "},
		{"██   ██ "},
		{"██████  "},
		{"██   ██ "},
		{"██   ██ "},
	},
	'S': {
		{" ███████ "},
		{"██       "},
		{" ███████ "},
		{"      ██ "},
		{" ███████ "},
	},
	'T': {
		{"████████ "},
		{"   ██    "},
		{"   ██    "},
		{"   ██    "},
		{"   ██    "},
	},
	'U': {
		{"██    ██ "},
		{"██    ██ "},
		{"██    ██ "},
		{"██    ██ "},
		{" ██████  "},
	},
	'V': {
		{"██    ██ "},
		{"██    ██ "},
		{"██    ██ "},
		{" ██  ██  "},
		{"  ████   "},
	},
	'W': {
		{"██     ██ "},
		{"██     ██ "},
		{"██  █  ██ "},
		{"██ ███ ██ "},
		{" ███ ███  "},
	},
	'X': {
		{"██   ██ "},
		{" ██ ██  "},
		{"  ███   "},
		{" ██ ██  "},
		{"██   ██ "},
	},
	'Y': {
		{"██    ██ "},
		{" ██  ██  "},
		{"  ████   "},
		{"   ██    "},
		{"   ██    "},
	},
	'Z': {
		{"███████ "},
		{"   ██   "},
		{"  ██    "},
		{" ██     "},
		{"███████ "},
	},
	' ': {
		{"    "},
		{"    "},
		{"    "},
		{"    "},
		{"    "},
	},
	'!': {
		{"██ "},
		{"██ "},
		{"██ "},
		{"   "},
		{"██ "},
	},
	'0': {
		{" ██████  "},
		{"██  ████ "},
		{"██ ██ ██ "},
		{"████  ██ "},
		{" ██████  "},
	},
	'1': {
		{" ██ "},
		{"███ "},
		{" ██ "},
		{" ██ "},
		{" ██ "},
	},
	'2': {
		{"██████  "},
		{"     ██ "},
		{"██████  "},
		{"██      "},
		{"███████ "},
	},
	'3': {
		{"██████  "},
		{"     ██ "},
		{"██████  "},
		{"     ██ "},
		{"██████  "},
	},
	'4': {
		{"██   ██ "},
		{"██   ██ "},
		{"███████ "},
		{"     ██ "},
		{"     ██ "},
	},
	'5': {
		{"███████ "},
		{"██      "},
		{"███████ "},
		{"     ██ "},
		{"███████ "},
	},
	'6': {
		{" ██████  "},
		{"██       "},
		{"███████  "},
		{"██    ██ "},
		{" ██████  "},
	},
	'7': {
		{"███████ "},
		{"     ██ "},
		{"    ██  "},
		{"   ██   "},
		{"   ██   "},
	},
	'8': {
		{" ██████  "},
		{"██    ██ "},
		{" ██████  "},
		{"██    ██ "},
		{" ██████  "},
	},
	'9': {
		{" ██████  "},
		{"██    ██ "},
		{" ███████ "},
		{"      ██ "},
		{" ██████  "},
	},
}

type Color int

const (
	Red Color = iota
	Green
	Blue
	Yellow
	Magenta
	Cyan
	Rainbow
)

var colorCodes = map[Color][]string{
	Red:     {"\033[91m"},
	Green:   {"\033[92m"},
	Blue:    {"\033[94m"},
	Yellow:  {"\033[93m"},
	Magenta: {"\033[95m"},
	Cyan:    {"\033[96m"},
	Rainbow: {"\033[91m", "\033[93m", "\033[92m", "\033[96m", "\033[94m", "\033[95m"},
}

const resetColor = "\033[0m"

func main() {
	text := flag.String("t", "HELLO", "Text to display")
	colorFlag := flag.String("c", "cyan", "Color: red, green, blue, yellow, magenta, cyan, rainbow")
	showVersion := flag.Bool("v", false, "Show version")
	border := flag.Bool("b", false, "Add border")

	flag.Parse()

	if *showVersion {
		fmt.Printf("bannify version %s\n", version)
		return
	}

	if flag.NArg() > 0 {
		*text = strings.Join(flag.Args(), " ")
	}

	color := parseColor(*colorFlag)
	banner := generateBanner(strings.ToUpper(*text), color, *border)
	fmt.Println(banner)
}

func parseColor(s string) Color {
	switch strings.ToLower(s) {
	case "red":
		return Red
	case "green":
		return Green
	case "blue":
		return Blue
	case "yellow":
		return Yellow
	case "magenta":
		return Magenta
	case "rainbow":
		return Rainbow
	default:
		return Cyan
	}
}

func generateBanner(text string, color Color, withBorder bool) string {
	if len(text) == 0 {
		return ""
	}

	var lines [5]strings.Builder

	for _, char := range text {
		charLines, exists := asciiFont[char]
		if !exists {
			charLines = asciiFont[' ']
		}

		for i := 0; i < 5; i++ {
			lines[i].WriteString(charLines[i][0])
		}
	}

	var result strings.Builder
	colors := colorCodes[color]

	if withBorder {
		borderLen := len(lines[0].String()) + 4
		result.WriteString(colors[0])
		result.WriteString(strings.Repeat("═", borderLen))
		result.WriteString(resetColor + "\n")
	}

	for i := 0; i < 5; i++ {
		colorCode := colors[i%len(colors)]
		if withBorder {
			result.WriteString(colorCode + "║ " + resetColor)
		}
		result.WriteString(colorCode)
		result.WriteString(lines[i].String())
		result.WriteString(resetColor)
		if withBorder {
			result.WriteString(colorCode + " ║" + resetColor)
		}
		result.WriteString("\n")
	}

	if withBorder {
		borderLen := len(lines[0].String()) + 4
		result.WriteString(colors[0])
		result.WriteString(strings.Repeat("═", borderLen))
		result.WriteString(resetColor + "\n")
	}

	return result.String()
}
