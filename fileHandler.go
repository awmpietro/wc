package main

import (
	"bufio"
	"os"
)

type fileHandler struct {
	fileName string
	count    int
}

func (fh fileHandler) getBytesNumber() (int, error) {
	file, err := os.Open(fh.fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	if err := fh.scan(scanner); err != nil {
		return 0, err
	}
	return fh.count, nil
}

func (fh fileHandler) getLinesNumber() (int, error) {
	file, err := os.Open(fh.fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	if err := fh.scan(scanner); err != nil {
		return 0, err
	}
	return fh.count, nil
}

func (fh fileHandler) getWordsNumber() (int, error) {
	file, err := os.Open(fh.fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	if err := fh.scan(scanner); err != nil {
		return 0, err
	}
	return fh.count, nil
}

func (fh fileHandler) getCharsNumber() (int, error) {
	file, err := os.Open(fh.fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	if err := fh.scan(scanner); err != nil {
		return 0, err
	}
	return fh.count, nil
}

func (fh *fileHandler) scan(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		fh.count++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
