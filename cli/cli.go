package cli

import (
	"flag"
)

// Cli represents the command line interface.
type Cli struct {
	CheckBytes bool
	CheckLines bool
	CheckWords bool
	CheckChars bool
	FileName   string
}

// New parses the command-line arguments and returns a new Cli.
func New() *Cli {
	checkBytes := flag.Bool("c", false, "count the number of bytes in a file")
	checkLines := flag.Bool("l", false, "count the number of lines in a file")
	checkWords := flag.Bool("w", false, "count the number of words in a file")
	checkChars := flag.Bool("m", false, "count the number of characters in a file")

	flag.Parse()

	return &Cli{
		CheckBytes: *checkBytes,
		CheckLines: *checkLines,
		CheckWords: *checkWords,
		CheckChars: *checkChars,
		FileName:   determineFileName(),
	}
}

// determineFileName extracts the file name from the command-line arguments.
func determineFileName() string {
	if flag.NArg() > 0 {
		return flag.Arg(0)
	}
	return ""
}
