package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type cli struct {
	checkBytes bool
	checkLines bool
	checkWords bool
	checkChars bool
}

func main() {

	checkBytes := flag.Bool("c", false, "count the number of bytes in a file")
	checkLines := flag.Bool("l", false, "count the number of lines in a file")
	checkWords := flag.Bool("w", false, "count the number of words in a file")
	checkChars := flag.Bool("m", false, "count the number of characters in a file")

	flag.Parse()

	app := cli{
		checkBytes: *checkBytes,
		checkLines: *checkLines,
		checkWords: *checkWords,
		checkChars: *checkChars,
	}

	var fileName string

	var ccwc CCWC

	if flag.NArg() == 0 {
		ccwc = stdinHandler{}
	} else {
		fileName = flag.Arg(0)
		ccwc = fileHandler{
			fileName: fileName,
		}
	}

	if err := app.print(ccwc, fileName); err != nil {
		log.Fatal(err)
	}

}

func (app *cli) isFlagSet() bool {
	set := false

	flag.Visit(func(f *flag.Flag) {
		set = true
	})
	return set
}

func (app *cli) print(wc CCWC, fileName string) error {
	var output int
	var err error

	if !app.isFlagSet() {
		val, err := app.printNoArgs(wc)
		if err != nil {
			return err
		}
		fmt.Println(val)
		return nil
	} else if app.checkBytes {
		output, err = wc.getBytesNumber()
	} else if app.checkLines {
		output, err = wc.getLinesNumber()
	} else if app.checkWords {
		output, err = wc.getWordsNumber()
	} else if app.checkChars {
		output, err = wc.getCharsNumber()
	}

	if err != nil {
		return err
	}

	fmt.Printf("%d \t%s\n", output, fileName)
	return nil
}

func (app *cli) printNoArgs(wc CCWC) (string, error) {
	args := os.Args[1:]
	fileName := args[0]

	bytesQt, err := wc.getBytesNumber()

	if err != nil {
		return "", err
	}

	linesQt, err := wc.getLinesNumber()

	if err != nil {
		return "", err
	}

	wordsQt, err := wc.getWordsNumber()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d\t%d\t%d\t%s", linesQt, wordsQt, bytesQt, fileName), nil
}
