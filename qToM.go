package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
/*
	qToM takes the standard input and extracts floating point numbers delineated by white space, skipping the first text. Data is assumed to be quarterly and is converted to monthly using linear extrapolation with the trailing slope. Monthly data starts with in the quarter of the second row of data. Leading or trailing white space on each input line is ignored. Output is written to the standard output in a standard exponential format.
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	var v, v0 []float64
	for i := 0; ; i++ {
		line, err := in.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimRight(strings.TrimLeft(line, " \t\n"), " \t\n")
		sn := strings.Split(line, " ", -1)
		v = make([]float64, 0)
		for _, val := range sn {
			if val == "" {
				continue
			}
			n, err := strconv.Atof64(val)
			if err != nil {
				fmt.Println("problem converting line number", i, "to numbers", val, err)
				os.Exit(1)
			}
			v = append(v, n)
		}
		if i != 0 {
			var s0, s1, s2 string
			for j, v := range v {
				if j == 0 {
					continue
				}
				diff := v - v0[j]
				s0 += fmt.Sprintf("%15e ", v-diff/3.0)
				s1 += fmt.Sprintf("%15e ", v)
				s2 += fmt.Sprintf("%15e ", v+diff/3.0)
			}
			os.Stdout.WriteString(s0 + "\n")
			os.Stdout.WriteString(s1 + "\n")
			os.Stdout.WriteString(s2 + "\n")
		}
		v0 = make([]float64, len(v))
		copy(v0, v)
	}
}
