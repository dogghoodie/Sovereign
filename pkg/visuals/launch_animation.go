package visuals

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Speed adjusters for each group of characters
var (
	switchSpeedSlash    = 100 * time.Millisecond // Speed for / and 7
	switchSpeedBackslash = 150 * time.Millisecond // Speed for \ and >
	switchSpeedPipe     = 200 * time.Millisecond // Speed for |, [, ], I, l, etc.
	switchSpeedDollar   = 250 * time.Millisecond // Speed for $, s, h, etc.
	letterAppearance    = 0.5  // 50% chance of a letter block appearing
	negativeChance      = 0.2  // 20% chance of a letter block coming in as negative
)

// The Sovereign logo, with each letter represented by an array of strings (blocks)
var logoLetters = [][]string{
	{ // "S"
		"  /$$$$$$  ",
		" /$$__  $$ ",
		"| $$  \\ $$ ",
		"| $$$$$$$$ ",
		"| $$__  $$ ",
		"| $$  | $$ ",
		"|  $$$$$$/ ",
		" \\______/  ",
	},
	{ // "o"
		"   /$$$$$$  ",
		"  /$$__  $$ ",
		" | $$  \\__/ ",
		" |  $$$$$$  ",
		"  \\____  $$ ",
		"  /$$  \\ $$ ",
		" |  $$$$$$/ ",
		"  \\______/  ",
	},
	{ // "v"
		"  /$$    /$$ ",
		" |  $$  /$$/ ",
		"  \\  $$/$$/  ",
		"   \\  $$$/   ",
		"    \\  $/    ",
		"     | $$    ",
		"     | $$    ",
		"     |__/    ",
	},
	{ // "e"
		"  /$$$$$$$$ ",
		" | $$_____/ ",
		" | $$       ",
		" | $$$$$$   ",
		" | $$__  $$ ",
		" | $$  \\ $$ ",
		" |  $$$$$$/ ",
		"  \\______/  ",
	},
	{ // "r"
		"  /$$$$$$   ",
		" /$$__  $$  ",
		"| $$  \\ $$  ",
		"| $$$$$$$$  ",
		"| $$__  $$  ",
		"| $$  | $$  ",
		"|  $$$$$$$  ",
		" \\_______/  ",
	},
	{ // "e"
		"  /$$$$$$$$ ",
		" | $$_____/ ",
		" | $$       ",
		" | $$$$$$   ",
		" | $$__  $$ ",
		" | $$  \\ $$ ",
		" |  $$$$$$/ ",
		"  \\______/  ",
	},
	{ // "i"
		"    /$$   ",
		"   | $$   ",
		"   | $$   ",
		"   | $$   ",
		"   | $$   ",
		"   | $$   ",
		"   | $$$$$",
		"   \\_____/ ",
	},
	{ // "g"
		"   /$$$$$$  ",
		"  /$$__  $$ ",
		" | $$  \\__/ ",
		" |  $$$$$$  ",
		"  \\____  $$ ",
		"  /$$  \\ $$ ",
		" |  $$$$$$/ ",
		"  \\______/  ",
	},
	{ // "n"
		"  /$$   /$$ ",
		" | $$  / $$ ",
		" |  $$/ $$  ",
		"  \\  $$$$   ",
		"   \\  $$/   ",
		"    | $$    ",
		"    | $$    ",
		"    |__/    ",
	},
}

// Substitution mappings for switching characters
var switchMap = map[rune][]rune{
	'/':  {'7'},
	'\\': {'>'},
	'|':  {'[', ']', 'I', 'l', ';', ':', '!', 'i'},
	'$':  {'s', 'h', 'b', 'V', 'D', 'g', 'w', '#', 'k', '*', 'M', 'm', '`', '.', '%', '~', '"', '\''},
}

// Substitution timing for each group of characters
func substituteWithDelay(c rune, delay time.Duration) rune {
	time.Sleep(delay)
	return substituteChar(c)
}

// Substitution logic for characters
func substituteChar(c rune) rune {
	if subs, exists := switchMap[c]; exists {
		return subs[rand.Intn(len(subs))]
	}
	return c
}

// Randomize appearance of individual letter blocks with optional negative effects
func randomizeLetter(letter []string) []string {
	result := make([]string, len(letter))
	for i, line := range letter {
		var newLine strings.Builder
		for _, char := range line {
			// Apply different delays based on character type
			switch char {
			case '/', '7':
				char = substituteWithDelay(char, switchSpeedSlash)
			case '\\', '>':
				char = substituteWithDelay(char, switchSpeedBackslash)
			case '|', '[', ']', 'I', 'l', ';', ':', '!', 'i':
				char = substituteWithDelay(char, switchSpeedPipe)
			case '$', 's', 'h', 'b', 'V', 'D', 'g', 'w', '#', 'k', '*', 'M', 'm', '`', '.', '%', '~', '"', '\'':
				char = substituteWithDelay(char, switchSpeedDollar)
			}

			// Random chance for negative effect
			if rand.Float64() < negativeChance {
				newLine.WriteRune(' ')
			} else {
				newLine.WriteRune(char)
			}
		}
		result[i] = newLine.String()
	}
	return result
}

// Render the ASCII logo with randomized effects for each letter block
func renderLogo() {
	for row := 0; row < len(logoLetters[0]); row++ { // Loop through each row of the letters
		for _, letter := range logoLetters {
			if rand.Float64() < letterAppearance {
				randomizedLetter := randomizeLetter(letter)
				fmt.Print(randomizedLetter[row])
			} else {
				fmt.Print(strings.Repeat(" ", len(letter[row]))) // Empty space for missing letter
			}
			fmt.Print("  ") // Extra spacing between letters
		}
		fmt.Println() // Move to next line
	}
}

// Continuously animate the logo in the terminal
func animateLogo() {
	for {
		renderLogo()
		time.Sleep(100 * time.Millisecond) // Adjust frame rate here
	}
}

func animate_call() {
	rand.Seed(time.Now().UnixNano())

	// Customize animation parameters here
	switchSpeedSlash = 80 * time.Millisecond
	switchSpeedBackslash = 120 * time.Millisecond
	switchSpeedPipe = 180 * time.Millisecond
	switchSpeedDollar = 220 * time.Millisecond
	letterAppearance = 0.7  // 70% chance of letters appearing
	negativeChance = 0.3    // 30% chance of negative effect

	// Start the ASCII animation
	animateLogo()
}


//TODO: Settle on an animation idea then implement with python frames script.
