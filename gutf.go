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

	r := charmap.Windows1254.NewDecoder().Reader(f)

	io.Copy(out, r)

	out.Close()
	f.Close()
}
