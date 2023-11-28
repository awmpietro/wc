package main

import (
	"bufio"
	"os"
)

type stdinHandler struct {
	count int
}

func (sh stdinHandler) bytesCount() (int, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)

	if err := sh.scan(scanner); err != nil {
		return 0, err
	}
	return sh.count, nil
}

func (sh stdinHandler) linesCount() (int, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	if err := sh.scan(scanner); err != nil {
		return 0, err
	}
	return sh.count, nil
}

func (sh stdinHandler) wordsCount() (int, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if err := sh.scan(scanner); err != nil {
		return 0, err
	}
	return sh.count, nil
}

func (sh stdinHandler) charsCount() (int, error) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	if err := sh.scan(scanner); err != nil {
		return 0, err
	}
	return sh.count, nil
}

func (sh *stdinHandler) scan(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		sh.count++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
