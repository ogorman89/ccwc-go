package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	var reader *bufio.Reader
	var filename string
	var lineCount, wordCount, charCount, byteCount int

	bytesArg := flag.Bool("c", false, "bytes arguement")
	linesArg := flag.Bool("l", false, "lines arguement")
	wordsArg := flag.Bool("w", false, "words arguement")
	charsArg := flag.Bool("m", false, "characters arguement")
	flag.Parse()
	fileArg := flag.Args()

	if len(fileArg) > 0 {
		filename = fileArg[0]

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal()
		}
		byteCount += len(line)
		wordCount += len(strings.Fields(string(line)))
		charCount += utf8.RuneCountInString(string(line))
		lineCount++
	}

	if *charsArg {
		fmt.Println(charCount, filename)
	} else if *bytesArg {
		fmt.Println(byteCount, filename)
	} else if *linesArg {
		fmt.Println(lineCount, filename)
	} else if *wordsArg {
		fmt.Println(wordCount, filename)
	} else {
		fmt.Println(lineCount, wordCount, byteCount, filename)
	}
}
