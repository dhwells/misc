package main

import (
	"fmt"
	"crypto/openpgp/armor"
	"os"
)

func main() {
	hdr := map[string]string{}
	e, _ := armor.Encode(os.Stdout, "RADIX-64", hdr)

	const NBUF = 512
	var buf [NBUF]byte
	var err, erw os.Error
	var nr, nw int
	for {
		switch nr, err = os.Stdin.Read(buf[:]); true {
		case nr < 0:
			fmt.Fprintf(os.Stderr, "error reading from Stdin: %s\n", err.String())
			os.Exit(1)
		case nr == 0:
			break
		case nr > 0:
			if nw, erw = e.Write(buf[0:nr]); nw != nr {
				fmt.Fprintf(os.Stderr, "error writing encoded bytes to Stdout: %s\n", erw.String())
				os.Exit(1)
			}
		}
		if err == os.EOF { break }
	}
	e.Close()
}
