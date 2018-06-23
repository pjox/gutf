package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	eunicode "golang.org/x/text/encoding/unicode"
	"golang.org/x/text/encoding/unicode/utf32"
)

// StripSpaces removes the spaces from a string
func StripSpaces(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)
}

func main() {

	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "usage:\n\t%s [input] [encoding] [output] [encoding]\n", os.Args[0])
		os.Exit(1)
	}

	count := 0

	m := make(map[string]encoding.Encoding)
	for _, v := range charmap.All {
		s := StripSpaces(fmt.Sprint(v))
		count++
		m[s] = v
	}

	for _, v := range japanese.All {
		s := StripSpaces(fmt.Sprint(v))
		count++
		m[s] = v
	}

	for _, v := range korean.All {
		s := StripSpaces(fmt.Sprint(v))
		count++
		m[s] = v
	}

	for _, v := range simplifiedchinese.All {
		s := StripSpaces(fmt.Sprint(v))
		count++
		m[s] = v
	}

	for _, v := range traditionalchinese.All {
		s := StripSpaces(fmt.Sprint(v))
		count++
		m[s] = v
	}

	for _, v := range eunicode.All {
		s := StripSpaces(fmt.Sprint(v))
		count++
		m[s] = v
	}

	for _, v := range utf32.All {
		s := StripSpaces(fmt.Sprint(v))
		count++
		m[s] = v
	}

	input, iencoding, output, oencoding := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	f, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(output)

	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReaderSize(f, 1<<20)

	out := bufio.NewWriterSize(outFile, 1<<20)

	decode := m[iencoding].NewDecoder().Reader(in)

	encode := m[oencoding].NewEncoder().Writer(out)

	buf := make([]byte, 1<<20)

	io.CopyBuffer(encode, decode, buf)

	out.Flush()
	outFile.Close()
	f.Close()
}
