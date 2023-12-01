package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	bytesFlag := flag.Bool("c", false, "The number of bytes in each input file is written to the standard output.")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Printf("Usage: mycli <parameter>\n")
		os.Exit(1)
	}
	filePath := flag.Arg(0)

	if *bytesFlag {
		fi, err := os.Stat(filePath)
		if err != nil {
			os.Exit(1)
		}
		size := fi.Size()
		fmt.Printf("Bytes: %d\n", size)
	}

}
