package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	echo2_start := time.Now()
	echo2()
	fmt.Println("echo2 duration:", time.Since(echo2_start))

	echo3_start := time.Now()
	echo3()
	fmt.Println("echo3 duration:", time.Since(echo3_start))
}
