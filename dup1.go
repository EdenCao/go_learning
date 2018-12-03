package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		fmt.Printf("%s\t%d\n", line, n)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
