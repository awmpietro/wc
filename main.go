package main

import (
	"fmt"
	"log"
	"os"

	"github.com/awmpietro/ccwc/cli"
)

func main() {

	app := cli.New()

	var ccwc CCWC

	if app.FileName == "" {
		isEmpty, err := isEmptyStdin()
		if err != nil {
			log.Fatal(err)
		}
		if isEmpty {
			log.Fatal("no filename and no named pipeline for stdin")
		}
		ccwc = stdinHandler{}
	} else {
		ccwc = fileHandler{
			fileName: app.FileName,
		}
	}

	if err := print(app, ccwc); err != nil {
		log.Fatal(err)
	}

}

func isFlagSet(app *cli.Cli) bool {
	return app.CheckBytes || app.CheckLines || app.CheckWords || app.CheckChars
}

func isEmptyStdin() (bool, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return fi.Mode()&os.ModeNamedPipe == 0, nil

}

func print(app *cli.Cli, wc CCWC) error {
	var output int
	var err error

	if !isFlagSet(app) {
		val, err := printNoArgs(app, wc)
		if err != nil {
			return err
		}
		fmt.Println(val)
		return nil
	} else if app.CheckBytes {
		output, err = wc.bytesCount()
	} else if app.CheckLines {
		output, err = wc.linesCount()
	} else if app.CheckWords {
		output, err = wc.wordsCount()
	} else if app.CheckChars {
		output, err = wc.charsCount()
	}

	if err != nil {
		return err
	}

	fmt.Printf("%d \t%s\n", output, app.FileName)
	return nil
}

func printNoArgs(app *cli.Cli, wc CCWC) (string, error) {
	bytesQt, err := wc.bytesCount()

	if err != nil {
		return "", err
	}

	linesQt, err := wc.linesCount()

	if err != nil {
		return "", err
	}

	wordsQt, err := wc.wordsCount()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d\t%d\t%d\t%s", linesQt, wordsQt, bytesQt, app.FileName), nil
}
