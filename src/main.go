package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// UniqueIntFile holds a map of unique integers found in the file to avoid
// duplicates
type UniqueIntFile struct {
	uniqueValues map[int]bool
}

// Creates a new UniqueIntFile
func NewUniqueIntFile() *UniqueIntFile {
	return &UniqueIntFile{
		uniqueValues: make(map[int]bool),
	}
}

// LoadDataFromFile reads data from the file
func (u *UniqueIntFile) LoadDataFromFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Loop through each line in the file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			continue
		}

		// check has only one integer else skip
		parts := strings.Fields(line)
		if len(parts) != 1 {
			continue
		}

		// Convert to integer
		val, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		// Check if value is within range asked
		if val >= -1023 && val <= 1023 {
			u.uniqueValues[val] = true
		}
	}

	return scanner.Err()
}

func (u *UniqueIntFile) GetSortedUniqueValues() []int {
	// Extract unique values from the map into a slice
	// This is done to sort the values later
	values := make([]int, 0, len(u.uniqueValues))
	for val := range u.uniqueValues {
		values = append(values, val)
	}

	// Sort values in ascending order
	sort.Ints(values)
	return values
}

// ProcessFile processes an input file and generates the output file
func ProcessFile(inputFilePath, outputFilePath string) error {
	// Create new UniqueIntFile
	uniqueInt := NewUniqueIntFile()

	// Load data from input file
	err := uniqueInt.LoadDataFromFile(inputFilePath)
	if err != nil {
		return err
	}

	// Get sorted unique values
	sortedUniqueValues := uniqueInt.GetSortedUniqueValues()

	// Create output file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Write each unique value to the output file
	for _, val := range sortedUniqueValues {
		fmt.Fprintln(outputFile, val)
	}

	return nil
}

const (
	file1        = "sample_data/sample_01.txt"
	file1_output = "sample_results/sample_01_output.txt"
	file2        = "sample_data/sample_04.txt"
	file2_output = "sample_results/sample_04_output.txt"
)

func main() {

	// Process the file 1
	err := ProcessFile(file1, file1_output)
	if err != nil {
		fmt.Printf("Error processing file1: %v\n", err)
		return
	}

	// Process the file2
	err = ProcessFile(file2, file2_output)
	if err != nil {
		fmt.Printf("Error processing file12: %v\n", err)
		return
	}

	fmt.Println("Successful processed files")
}
