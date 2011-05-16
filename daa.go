package main

import (
	"fmt"
	"crypto/openpgp/armor"
	"os"
)

func main() {
	b, erd := armor.Decode(os.Stdin)
	if erd != nil {
		fmt.Println("unable to decode the Stdin")
		os.Exit(2)
	}
	const NBUF = 512
	var buf [NBUF]byte
	var err, erw os.Error
	var nr, nw int
	for {
		switch nr, err = b.Body.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "error reading from Stdin: %s\n", err.String())
			os.Exit(1)
		case nr == 0:
			break
		case nr > 0:
			if nw, erw = os.Stdout.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "error writing decoded bytes to Stdout: %s\n", erw.String())
				os.Exit(1)
			}
		}
		if err == os.EOF {
			break
		}
	}
}
