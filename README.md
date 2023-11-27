# code challenge: wc command line tool

a simple version of the UNIX's command line tool "wc" written in golang.

## Installation

To use this, clone the repo:

```bash
git clone https://github.com/awmpietro/wc.git
```

## Usage:

Build and install

```bash
go install github.com/awmpietro/ccwc
```

Create an instance of NewLimiter and apply as a middleware to a handler:

```bash
// count the number of bytes
ccwc -c test.txt

// count the number of lines
ccwc -l test.txt

// count the number of words
ccwc -w test.txt

// count the number of characters
ccwc -m test.txt

// no options are provided is the equivalent to the -c, -l and -w options
ccwc test.txt

// read from standard input if no filename is specified
cat test.txt | ccwc -c
cat test.txt | ccwc -l
cat test.txt | ccwc -w
cat test.txt | ccwc -m
```

## Contributing

Feel free to open issues or PRs if you find any problems or have suggestions!
