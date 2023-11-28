package main

type CCWC interface {
	bytesCount() (int, error)
	linesCount() (int, error)
	wordsCount() (int, error)
	charsCount() (int, error)
}
