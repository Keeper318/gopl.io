package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func convert(arg string) {
	if t, err := strconv.Atoi(arg); err == nil {
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convert(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			convert(input.Text())
		}
		if err := input.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
			os.Exit(1)
		}
	}
}
