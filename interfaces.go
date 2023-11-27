package main

type CCWC interface {
	getBytesNumber() (int, error)
	getLinesNumber() (int, error)
	getWordsNumber() (int, error)
	getCharsNumber() (int, error)
}
