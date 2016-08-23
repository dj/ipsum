package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Read flags
	setencesArg := flag.Int("sentences", 50, "Number of sentences to take from the input text")
	paragraphsArg := flag.Int("paragraphs", 1, "Number of paragraphs to take form the input text")
	inputArg := flag.String("input", "", "Path to input text file (otherwise reads from stdin)")
	outputArg := flag.String("output", "", "Path to output text file (otherwise prints to stdout)")
	outputFmtArg := flag.String("fmt", "html", "Output file format. One of: html, txt")
	flag.Parse()

	// Output to a file or stdout
	var outputWriter *bufio.Writer
	if *outputArg == "" {
		outputWriter = bufio.NewWriter(os.Stdout)
	} else {
		file, err := os.Create(*outputArg)
		if err != nil {
			panic(err)
		}
		outputWriter = bufio.NewWriter(file)
	}

	// Read input from a file or stdout
	var scanner *bufio.Scanner
	if *inputArg == "" {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(*inputArg)
		if err != nil {
			panic(err)
		}
		scanner = bufio.NewScanner(file)
	}
	scanner.Split(scanSentences)

	// Read from the input and write to the output
	for paragraphCount := 0; paragraphCount < *paragraphsArg; paragraphCount++ {
		if *outputFmtArg == "html" {
			outputWriter.WriteString("<p>")
		}

		for wordCount := 0; wordCount < *setencesArg; wordCount++ {
			scanner.Scan()
			outputWriter.WriteString(scanner.Text() + ".")

			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading input:", err)
			}
		}

		if *outputFmtArg == "html" {
			outputWriter.WriteString("</p> \n")
		}
	}

	err := outputWriter.Flush()
	if err != nil {
		fmt.Fprintln(os.Stderr, "writing output:", err)
	}
}

// scanSentences is a bufio.SplitFunc
func scanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Sentence delimiters. I have a feeling I'm missing a bunch
	var (
		i = bytes.IndexByte(data, '.')
		j = bytes.IndexByte(data, '?')
		k = bytes.IndexByte(data, '!')
	)
	if i >= 0 || j >= 0 || k >= 0 {
		// We have a full sentence.
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	// Read more data.
	return 0, nil, nil
}
