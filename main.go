package main

import (
	"fmt"

	"github.com/akgbytes/dsa/strings"
)

func main() {

	words := []string{"leet", "code"}
	fmt.Println(strings.FindWordsContaining(words, 'o'))
}
