package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	bytesArg := flag.Bool("c", false, "bytes arguement")
	linesArg := flag.Bool("l", false, "lines arguement")
	wordsArg := flag.Bool("w", false, "words arguement")
	charsArg := flag.Bool("m", false, "characters arguement")
	flag.Parse()
	fileArg := flag.Args()

	var scanner *bufio.Scanner
	var filename string
	var lineCount, wordCount, charCount int
	var byteCount int64

	if len(fileArg) > 0 {
		filename = fileArg[0]

		fileInfo, err := os.Stat(filename)
		if err != nil {
			log.Fatal(err)
		}
		byteCount = fileInfo.Size()

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		//no filename read from stdin
		// bufio.ScanBytes() <---- try this
		scanner = bufio.NewScanner(os.Stdin)
		fileInfo, err := os.Stdin.Stat()
		if err != nil {
			log.Fatal(err)
		}
		byteCount = fileInfo.Size()
	}

	for scanner.Scan() {
		line := scanner.Text()
		wordCount += len(strings.Fields(line))
		charCount += utf8.RuneCountInString(line)
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if *charsArg {
		fmt.Println(charCount)
	} else if *bytesArg {
		fmt.Println(byteCount)
	} else if *linesArg {
		fmt.Println(lineCount)
	} else if *wordsArg {
		fmt.Println(wordCount)
	} else {
		fmt.Println(lineCount, wordCount, byteCount)
	}
}
