package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage:\n\t%s [input] [output]\n", os.Args[0])
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	in := bufio.NewReaderSize(f, 1<<20)

	out := bufio.NewWriterSize(outFile, 1<<20)

	r := charmap.ISO8859_1.NewDecoder().Reader(in)

	buf := make([]byte, 1048576)

	io.CopyBuffer(out, r, buf)

	out.Flush()
	outFile.Close()
	f.Close()
}
