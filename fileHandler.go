package main

import (
	"bufio"
	"os"
)

type fileHandler struct {
	fileName string
	count    int
}

func (fh fileHandler) openFile() (*os.File, error) {
	file, err := os.Open(fh.fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (fh fileHandler) bytesCount() (int, error) {
	file, err := fh.openFile()
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

func (fh fileHandler) linesCount() (int, error) {
	file, err := fh.openFile()
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

func (fh fileHandler) wordsCount() (int, error) {
	file, err := fh.openFile()
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

func (fh fileHandler) charsCount() (int, error) {
	file, err := fh.openFile()
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
