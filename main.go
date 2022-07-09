package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var line = strings.Join(os.Args[1:], " ")

	var words = strings.Fields(line)
	fmt.Println(len(words))
}
