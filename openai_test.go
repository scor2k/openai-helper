package main

import (
	"io/fs"
	"os"
	"reflect"
	"testing"
)

func TestLoadPrompts(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "prompts.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write test data to the temporary file
	testData := `{"rewrite": "Hello"}`
	err = os.WriteFile(tempFile.Name(), []byte(testData), fs.FileMode(0644))
	if err != nil {
		t.Fatal(err)
	}

	// Call the loadPrompts function
	prompts, err := loadPrompts(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check if the prompts were loaded correctly
	expectedPrompts := Prompts{Rewrite: "Hello"}
	if !reflect.DeepEqual(prompts, expectedPrompts) {
		t.Errorf("Expected prompts %v, but got %v", expectedPrompts, prompts)
	}
}
