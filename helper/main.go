package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// reverseString reverses the input string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Define the source and output directories
	srcDir := "/src"
	outputDir := "/output"

	// Ensure the output directory exists
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create output directory: %v\n", err)
		return
	}

	// Read all files in the source directory
	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		fmt.Printf("Failed to read directory: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		// Construct full file path for source and output
		srcFilePath := filepath.Join(srcDir, file.Name())
		outputFilePath := filepath.Join(outputDir, file.Name())

		// Read file content
		content, err := ioutil.ReadFile(srcFilePath)
		if err != nil {
			fmt.Printf("Failed to read file %s: %v\n", srcFilePath, err)
			continue
		}

		// Reverse file content
		reversedContent := reverseString(string(content))

		// Write reversed content to output file
		err = ioutil.WriteFile(outputFilePath, []byte(reversedContent), 0644)
		if err != nil {
			fmt.Printf("Failed to write to file %s: %v\n", outputFilePath, err)
			continue
		}

		fmt.Printf("Reversed content of file %s and wrote to %s\n", srcFilePath, outputFilePath)
	}

	fmt.Println("All files processed.")
}
