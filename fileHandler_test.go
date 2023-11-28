package main

import (
	"os"
	"testing"
)

var testFileHandler fileHandler
var testContent = "test content\nwith multiple lines\nand words"

func fileHandlerSetupSuit(t *testing.T) func(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testing_file")
	if err != nil {
		t.Fatalf("unable to create temporary file: %s", err)
	}

	// Write some content to the temporary file
	if _, err := tmpFile.Write([]byte(testContent)); err != nil {
		t.Fatalf("unable to write to temporary file: %s", err)
	}
	tmpFile.Close()

	// Create a fileHandler with the temporary file
	testFileHandler = fileHandler{fileName: tmpFile.Name()}

	return func(t *testing.T) {
		//everything to run after test' suite
		os.Remove(tmpFile.Name())
	}
}

func TestFileHandler(t *testing.T) {
	tearDownSuite := fileHandlerSetupSuit(t)
	defer tearDownSuite(t)

	tests := []struct {
		id       int
		name     string
		expected bool
	}{
		{id: 1, name: "get bytes number", expected: true},
		{id: 2, name: "get lines number", expected: true},
		{id: 3, name: "get words number", expected: true},
		{id: 4, name: "get chars number", expected: true},
	}

	var count int
	var err error

	t.Log("Given the need to test fileHandler")
	{
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				switch tt.id {
				case 1:
					count, err = testFileHandler.bytesCount()
				case 2:
					count, err = testFileHandler.linesCount()
				case 3:
					count, err = testFileHandler.wordsCount()
				case 4:
					count, err = testFileHandler.charsCount()
				}
				if err != nil {
					t.Errorf("unexpected error: %s", err)
				}
				expected := count > 0
				if !expected {
					t.Errorf("expected number of bytes to be gt than zero but got %d", count)
				}

			})
		}
	}

}
