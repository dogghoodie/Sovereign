package visuals

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

//TODO: Figure out both unused parameter warnings.
// ( unusedparams: unused paramter: blockLetter ).

// Define character of interest
var charsOfInterest = []rune{'$', '/', '\\', '|', '_', ' '}

// Define blocks and their coordinates
var blockCoordinates = map[string][2][2]int{
	"S":  {{1, 1}, {10, 12}},
	"o":  {{3, 13}, {10, 22}},
	"v":  {{3, 23}, {10, 33}},
	"e1": {{3, 33}, {10, 42}},
	"r":  {{3, 42}, {10, 52}},
	"e2": {{3, 53}, {10, 62}},
	"i":  {{1, 63}, {10, 66}},
	"g":  {{3, 67}, {13, 76}},
	"n":  {{3, 77}, {10, 86}},
}

// Fill characters
var fillChars = []rune("qwertyughlodkrasd:,.'")

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
	defined_origin := []int{0, 0}
	// Define an array or slice of the available white colors
	colorOptions := []string{
		//	Just fucking around with the colors, the matrix
		//	theme is actually pretty sick looking.
		//	Colors.WHITE, Colors.WHITE2,
		//	Colors.WHITE3, Colors.WHITE4}
		Colors.MATRIX1, Colors.MATRIX2,
		Colors.MATRIX3, Colors.GREEN}

	// Pick a random color from the array
	chosenColor := colorOptions[rand.Intn(len(colorOptions))]

	// Print with the randomly chosen color
	fmt.Printf(chosenColor+"\033[%d;%dH%s", defined_origin[0]+row, defined_origin[1]+col, string(char)+Colors.ANSI_RESET)

}

// Function to animate a block positively by filling chars randomly (with separate fillChars for each character type)
func positiveLetter(blockLetter string, coordMap map[rune][][]int, delay time.Duration) {
	timeTotal := 2*time.Second - delay
	time.Sleep(delay)

	startTime := time.Now()

	// Define different fillChars strings for each charType
	fillCharsMap := map[rune]string{
		'$':  "$qwertyughlodkrasd", // Fill characters for '$'
		'/':  "7/",                 // Fill characters for '/'
		'\\': "\\>",                // Fill characters for '\\'
		'|':  "lI|[}{[",            // Fill characters for '|'
		'_':  "=-~",                // Fill characters for '_'
	}

	// Function to animate a specific character type's coordinates
	animateCharType := func(charCoords [][]int, fillChars string) {
		for time.Since(startTime) < timeTotal {
			for _, coord := range charCoords {
				row, col := coord[0], coord[1]
				randomChar := rune(fillChars[rand.Intn(len(fillChars))])
				printAtCoordinate(row, col, randomChar)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}

	// Create a WaitGroup to wait for all Goroutines
	var wg sync.WaitGroup

	// Launch separate Goroutines for each character type of interest
	for _, charType := range []rune{'$', '/', '\\', '|', '_'} {
		if coords, ok := coordMap[charType]; ok {
			wg.Add(1) // Increment the WaitGroup counter
			go func(charCoords [][]int, fillChars string) {
				defer wg.Done() // Decrement the counter when the Goroutine finishes
				animateCharType(charCoords, fillChars)
			}(coords, fillCharsMap[charType])
		}
	}

	// Wait for all Goroutines to finish
	wg.Wait()
}

// Function to animate a block negatively by filling space coordinates
// TODO: Figure out this unused blockLetter
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

// Helper function to shuffle an array
func shuffle_array(arr []int) {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(arr), func(i, j int) {
        arr[i], arr[j] = arr[j], arr[i]
    })
}

// Function that generates the dictionary
func generateDict(numNegative int) map[string]int {
    // Create an array of length 3 with numNegative zeros and the rest as ones
    arr := make([]int, 3)
    for i := 0; i < numNegative; i++ {
        arr[i] = 0
    }
    for i := numNegative; i < 3; i++ {
        arr[i] = 1
    }

    // Shuffle the array
    shuffle_array(arr)

    // Assign shuffled array values to "r", "v", and "n"
    dictFinal := map[string]int{
        "r": arr[0],
        "v": arr[1],
        "n": arr[2],
    }

    return dictFinal
}

// Main function to process blocks and animate them concurrently
func animateBlocks(blockList []string, coordMap map[rune][][]int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure this finishes when done

	for _, blockLetter := range blockList {
		wg.Add(1) // Increment the WaitGroup counter
		go func(blockLetter string) {
			defer wg.Done() // Decrement the counter when the goroutine finishes
			blockCoords := getBlockCoordinates(blockLetter, coordMap)
			delay := time.Duration(rand.Intn(950)+50) * time.Millisecond

			// Seed the random number generator
			rand.Seed(time.Now().UnixNano())

			// Generate numNegative (random number between 1 and 3 inclusive)
			numNegative := rand.Intn(3) + 1

			// Get the dictionary
			dict := generateDict(numNegative)

			// check if letter is in dict; r, v, n
			if value, exists := dict[blockLetter]; exists {
				if value == 1 { // 1 for positive letters
					positiveLetter(blockLetter, blockCoords, delay)
				} else { // 0 for negative letters
					negativeLetter(blockLetter, blockCoords, delay)
				}
			} else {
				positiveLetter(blockLetter, blockCoords, delay)
			}
		}(blockLetter) // Pass blockLetter to the Goroutine
	}
}

// Function to shuffle a slice of coordinates
func shuffle(coords [][]int) [][]int {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(coords), func(i, j int) {
		coords[i], coords[j] = coords[j], coords[i]
	})
	return coords
}

// Function to print characters from a given range in a random order with adjustable speed
func printInRandomOrder(array2D [][]rune, coordRange [2][2]int, speed time.Duration, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when done

	// Store the coordinates from [row1,col1] to [row2,col2] in a slice
	var coords [][]int
	for col := coordRange[0][1]; col <= coordRange[1][1]; col++ {
		coords = append(coords, []int{coordRange[0][0], col})
	}

	// Shuffle the coordinates for random printing order
	shuffledCoords := shuffle(coords)

	// Print each character at a shuffled coordinate with the specified speed
	for _, coord := range shuffledCoords {
		row, col := coord[0], coord[1]
		char := array2D[row-1][col-1]     // Fetch character from 2D array, adjust for 0-indexing
		printAtCoordinate(row, col, char) // Print the actual character at its original position
		time.Sleep(speed)
	}
}

// Start the animation.
func Draw_Launch_Animation() {
	// TODO: Update this to New(NewSource(seed))
	rand.Seed(time.Now().UnixNano())

	// Read the ASCII art file into a 2D array
	array2D := readFileInto2DArray("resources/logo.txt")

	// Get all coordinates of characters of interest
	coordMap := getCoordinates(array2D, charsOfInterest)

	// List of blocks to animate
	blockList := []string{"S", "o", "v", "e1", "r", "e2", "i", "g", "n"}

	var wg sync.WaitGroup

	// Start the animation of blocks
	wg.Add(1)
	go animateBlocks(blockList, coordMap, &wg)

	// Start printing characters from [11,32] to [11,56] in random order
	wg.Add(1)
	go printInRandomOrder(array2D, [2][2]int{{11, 32}, {11, 56}}, 50*time.Millisecond, &wg)

	// Wait for both animations to finish
	wg.Wait()
}
