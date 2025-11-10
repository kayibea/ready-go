package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

type TextInfo struct {
	linesEmpty uint64
	lines      uint64
	words      uint64
	chars      uint64
}

func (info TextInfo) printInfos() {
	var str string
	var builder strings.Builder

	str = fmt.Sprintf("%-20s: %d\n", "Total Line", info.lines)
	builder.WriteString(str)

	str = fmt.Sprintf("%-20s: %d\n", "Total Lines(Empty)", info.linesEmpty)
	builder.WriteString(str)

	str = fmt.Sprintf("%-20s: %d\n", "Total Words", info.words)
	builder.WriteString(str)

	str = fmt.Sprintf("%-20s: %d", "Total Chars", info.chars)
	builder.WriteString(str)

	fmt.Println(builder.String())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file>\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	var err error

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer file.Close()

	textInfo := TextInfo{}

	err = getTotalLines(file, &textInfo.lines, &textInfo.linesEmpty)
	if err != nil {
		fmt.Printf("Error: %q is not valid ASCII text file\n", path.Base(os.Args[1]))
		os.Exit(1)
	}

	if textInfo.lines == 0 {
		textInfo.printInfos()
		os.Exit(0)
	}

	getTotalWords(file, &textInfo.words)
	getTotalChars(file, &textInfo.chars)

	textInfo.printInfos()
}

func resetFileHead(f *os.File) {
	f.Seek(0, io.SeekStart)
}

func getTotalChars(f *os.File, count *uint64) {
	resetFileHead(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		*count += 1
	}
	resetFileHead(f)
}

func getTotalWords(f *os.File, count *uint64) {
	resetFileHead(f)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*count += 1
	}
	resetFileHead(f)
}

func getTotalLines(f *os.File, lines *uint64, empty *uint64) error {
	resetFileHead(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		if !utf8.ValidString(text) {
			return errors.New("not valid UTF-8 text")
		}

		if strings.TrimSpace(text) != "" {
			*lines += 1
		} else {
			*empty += 1
		}
	}

	resetFileHead(f)
	return nil
}
