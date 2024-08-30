package main

import (
	"fmt"
	"strings"
)

func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

func main() {
	input := "Aaslema , Baba!"
	result := toUpperCase(input)
	fmt.Println(result)
}
