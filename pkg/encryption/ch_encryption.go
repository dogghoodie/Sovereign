package encryption

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Declare filepaths
var enList, _ = extractEnglish("resources/english.csv")
var chList, _ = extractChinese("resources/hanziDB.csv")

// Function to extract the second column of a CSV file and return it as a list of strings
func extractChinese(filePath string) ([]string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}

	// Initialize reader
	reader := csv.NewReader(file)

	// Read the CSV file
	// Records are columns in csv
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV: %v", err)
	}

	// Create a list to store the second column
	var secondColumn []string

	// Loop through the records and extract the second column
	for _, record := range records {
		if len(record) > 1 {
			secondColumn = append(secondColumn, record[1]) // Record[1] is the second column
		}
	}

	file.Close()
	return secondColumn, nil
}

// Function to extract english alphabet
func extractEnglish(filePath string) ([]string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}

	// Initialize reader
	reader := csv.NewReader(file)

	// Read the CSV file
	// Records are columns in csv
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV: %v", err)
	}

	// Create a list to store the second column
	var firstColumn []string

	// Loop through the records and extract the second column
	for _, record := range records {
		if len(record) > 0 {
			firstColumn = append(firstColumn, record[0]) // Record[1] is the second column
		}
	}

	file.Close()
	return firstColumn, nil
}

// Cipher struct as before
type Cipher struct {
	enToChMap map[string]string
	chToEnMap map[string]string
}

// SetSeed initializes the random number generator with a specific seed and shuffles the character list
func SetSeed(seed int64) *Cipher {
	r := rand.New(rand.NewSource(seed))

	// Shuffle chList based on the seed
	r.Shuffle(len(chList), func(i, j int) {
		chList[i], chList[j] = chList[j], chList[i]
	})

	// Create the mapping between English and Chinese characters
	enToChMap := make(map[string]string)
	chToEnMap := make(map[string]string)

	for i := range enList {
		enToChMap[enList[i]] = chList[i]
		chToEnMap[chList[i]] = enList[i]
	}

	// Return the cipher with the mappings
	return &Cipher{enToChMap, chToEnMap}
}

// EncryptMessage encrypts an English message into a string of Chinese characters using the mapping
func (c *Cipher) EncryptMessage(message string) string {
	var encryptedString strings.Builder
	for _, char := range message {
		letter := string(char)
		if val, ok := c.enToChMap[letter]; ok {
			encryptedString.WriteString(val)
		} else {
			encryptedString.WriteString(letter) // if not in the map, retain the original character
		}
	}
	return encryptedString.String()
}

// DecryptMessage decrypts a string of Chinese characters into an English message using the mapping
func (c *Cipher) DecryptMessage(encryptedMessage string) string {
	var decryptedString strings.Builder
	for _, char := range encryptedMessage {
		letter := string(char)
		if val, ok := c.chToEnMap[letter]; ok {
			decryptedString.WriteString(val)
		} else {
			decryptedString.WriteString(letter) // if not in the map, retain the original character
		}
	}
	return decryptedString.String()
}

// lil main
func ch_encryption_test() { // add comments, folds,
	// and design inputs from user, and use functions directly from main
	// Example usage
	seed := int64(12)
	message := "callme"

	// Set the seed and create the cipher
	cipher := SetSeed(seed)

	// Encrypt the message
	encryptedMessage := cipher.EncryptMessage(message)
	fmt.Println("Encrypted Message:", encryptedMessage)

	// Decrypt the message
	decryptedMessage := cipher.DecryptMessage(encryptedMessage)
	fmt.Println("Decrypted Message:", decryptedMessage)
}