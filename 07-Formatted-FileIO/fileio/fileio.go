package fileio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Cat() {
	for _, fname := range os.Args[1:] {

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		file.Close()
	}
}

func FileSize() {
	for _, fname := range os.Args[1:] {

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		data, err := io.ReadAll(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Printf("The file size of %s is %d bytes.\n", fname, len(data))

		file.Close()
	}
}

func WordCount() {
	for _, fname := range os.Args[1:] {

		var lc, wc, cc int

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}
		fmt.Printf("%s has %d lines, %d words and %d characters\n", fname, lc, wc, cc)
		file.Close()
	}
}
