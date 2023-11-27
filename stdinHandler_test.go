package main

import (
	"os"
	"testing"
)

var testStdinHandler stdinHandler
var testStdinContent = "test content\nwith multiple lines\nand words"

func stdinHandlerSetupSuit(t *testing.T) (*os.File, func(t *testing.T)) {
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatalf("unable to create temporary file: %s", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(testStdinContent)
	if err != nil {
		t.Fatalf("unable to write to temporary file: %s", err)
	}

	originalStdin := os.Stdin

	os.Stdin = tmpFile

	testStdinHandler = stdinHandler{}

	return tmpFile, func(t *testing.T) {
		os.Stdin = originalStdin
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
}

func TestStdinHandler(t *testing.T) {
	tmpFile, tearDownSuite := stdinHandlerSetupSuit(t)
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

	t.Log("Given the need to test stdinHandler")
	{
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tmpFile.Seek(0, 0)
				switch tt.id {
				case 1:
					count, err = testStdinHandler.getBytesNumber()
				case 2:
					count, err = testStdinHandler.getLinesNumber()
				case 3:
					count, err = testStdinHandler.getWordsNumber()
				case 4:
					count, err = testStdinHandler.getCharsNumber()
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
