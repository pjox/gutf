package main

import (
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

	out, err := os.Create(os.Args[2])

	if err != nil {
		log.Fatal(err)
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(f)

	buf := make([]byte, 1048576)

	io.CopyBuffer(out, r, buf)

	out.Close()
	f.Close()
}
