package main

import (
	"bufio"
	"strings"
	"testing"
	"fmt"
)

// Function to be tested
func getUserInput(reader *bufio.Reader) (string, error) {
	fmt.Print("Enter something: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func TestGetUserInput(t *testing.T) {
	// Create a mocked input source using a string reader
	input := "test input\n"
	reader := bufio.NewReader(strings.NewReader(input))

	// Call the function with the mocked input source
	result, err := getUserInput(reader)

	// Check if the result and error are as expected
	if err != nil {
		t.Errorf("Expected no error, got this error%v", err)
	}
	expectedResult := "test input"
	if result != expectedResult {
		t.Errorf("Expected '%s', got '%s'", expectedResult, result)
	}
}

func main(){
	TestGetUserInput()
}
