package visuals

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Define character of interest
var charsOfInterest = []rune{'$', '/', '\\', '|', '_', ' '}

// Define blocks and their coordinates
var blockCoordinates = map[string][2][2]int{
	"S": {{1, 1}, {10, 12}},
	"o": {{3, 13}, {10, 22}},
	"v": {{3, 23}, {10, 33}},
	"e1": {{3, 33}, {10, 42}},
	"r": {{3, 42}, {10, 52}},
	"e2": {{3, 53}, {10, 62}},
	"i": {{1, 63}, {10, 66}},
	"g": {{3, 67}, {13, 76}},
	"n": {{3, 77}, {10, 86}},
}

// Fill characters
var fillChars = []rune("qwertyughlodkrasd")

// Function to read the text file and populate a 2D array
func readFileInto2DArray(filePath string) [][]rune {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var array2D [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		array2D = append(array2D, line)
	}
	return array2D
}

// Get coordinates for characters of interest
func getCoordinates(array2D [][]rune, charsOfInterest []rune) map[rune][][]int {
	coordMap := make(map[rune][][]int)
	for row := range array2D {
		for col := range array2D[row] {
			char := array2D[row][col]
			for _, targetChar := range charsOfInterest {
				if char == targetChar {
					coordMap[targetChar] = append(coordMap[targetChar], []int{row + 1, col + 1}) // 1-indexed
				}
			}
		}
	}
	return coordMap
}

// Get the coordinates for each block's characters of interest
func getBlockCoordinates(blockLetter string, coordMap map[rune][][]int) map[rune][][]int {
	blockCoord := blockCoordinates[blockLetter]
	blockCoordsMap := make(map[rune][][]int)

	for char, coords := range coordMap {
		for _, coord := range coords {
			if coord[0] >= blockCoord[0][0] && coord[0] <= blockCoord[1][0] && // Check row
				coord[1] >= blockCoord[0][1] && coord[1] <= blockCoord[1][1] { // Check column
				blockCoordsMap[char] = append(blockCoordsMap[char], coord)
			}
		}
	}
	return blockCoordsMap
}

// Print character to terminal at specified coordinates
func printAtCoordinate(row, col int, char rune) {
	fmt.Printf("\033[%d;%dH%s", row, col, string(char)) // Move cursor and print
}

// Function to animate a block positively by filling chars randomly
func positiveLetter(blockLetter string, coordMap map[rune][][]int, delay time.Duration) {
	timeTotal := 2*time.Second - delay
	time.Sleep(delay)

	startTime := time.Now()
	for time.Since(startTime) < timeTotal {
		for _, coord := range append(coordMap['$'], coordMap['/']...) { // Combine coord lists of interest
			row, col := coord[0], coord[1]
			randomChar := fillChars[rand.Intn(len(fillChars))]
			printAtCoordinate(row, col, randomChar)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Function to animate a block negatively by filling space coordinates
func negativeLetter(blockLetter string, coordMap map[rune][][]int, delay time.Duration) {
	timeTotal := 2*time.Second - delay
	time.Sleep(delay)

	startTime := time.Now()
	for time.Since(startTime) < timeTotal {
		for _, coord := range coordMap[' '] { // Use space coordinates
			row, col := coord[0], coord[1]
			randomChar := fillChars[rand.Intn(len(fillChars))]
			printAtCoordinate(row, col, randomChar)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// Main function to process blocks and animate them
func animateBlocks(blockList []string, coordMap map[rune][][]int) {
	for _, blockLetter := range blockList {
		blockCoords := getBlockCoordinates(blockLetter, coordMap)
		delay := time.Duration(rand.Intn(450)+50) * time.Millisecond
		if rand.Float64() > 0.3 {
			positiveLetter(blockLetter, blockCoords, delay)
		} else {
			negativeLetter(blockLetter, blockCoords, delay)
		}
	}
}

func Animate_call() {
	rand.Seed(time.Now().UnixNano())
	// Read the ASCII art file into a 2D array
	array2D := readFileInto2DArray("resources/logo.txt")

	// Get all coordinates of characters of interest
	coordMap := getCoordinates(array2D, charsOfInterest)

	// List of blocks to animate
	blockList := []string{"S", "o", "v", "e1", "r", "e2", "i", "g", "n"}

	// Start the animation
	animateBlocks(blockList, coordMap)
}
