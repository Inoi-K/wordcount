package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var in = bufio.NewReader(os.Stdin)
	var out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var line string
	line, _ = in.ReadString('\n')

	var words = strings.Split(line, " ")
	var wordCount int
	for _, w := range words {
		if len(w) > 0 {
			wordCount++
		}
	}
	fmt.Fprint(out, wordCount)
}
