package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"unicode/utf8"
)

type TextInfo struct {
	f          *os.File
	emptyLines uint64
	lines      uint64
	words      uint64
	chars      uint64
}

func (info *TextInfo) print() {
	fmt.Printf("%-20s: %d\n", "Total Lines", info.lines)
	fmt.Printf("%-20s: %d\n", "Total Empty Lines", info.emptyLines)
	fmt.Printf("%-20s: %d\n", "Total Words", info.words)
	fmt.Printf("%-20s: %d\n", "Total Characters", info.chars)
}

func main() {
	argc := len(os.Args) - 1

	if argc < 2 {
		fmt.Printf("Usage: %s <file> [file]...\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	files := make([]*os.File, 0, argc)
	defer closeFiles(&files)

	for _, name := range os.Args[1:] {
		file, err := os.Open(name)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		files = append(files, file)
	}

	var wg sync.WaitGroup
	errs := make(chan error, len(files))
	infos := make([]TextInfo, len(files))

	for i, file := range files {
		wg.Add(1)
		infos[i].f = file

		go func(i int, file *os.File) {
			defer wg.Done()

			err := getTextInfo(file, &infos[i])
			if err != nil {
				errs <- err
			}
		}(i, file)
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}

	sep := strings.Repeat("-", 15)
	for _, info := range infos {
		fmt.Printf("\nFile: %s\n%s\n", info.f.Name(), sep)
		info.print()
		fmt.Println()
	}

	os.Exit(0)
}

func closeFiles(files *[]*os.File) {
	for _, file := range *files {
		if file != nil {
			file.Close()
		}
	}
}

func getTextInfo(f *os.File, info *TextInfo) error {
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if !utf8.ValidString(line) {
			return fmt.Errorf("%q: invalid UTF-8 text file\n", f.Name())
		}

		if strings.TrimSpace(line) == "" {
			info.emptyLines++
		}

		info.lines++
		info.words += uint64(len(strings.Fields(line)))
		info.chars += uint64(len(line)) + 1
	}

	return nil
}
