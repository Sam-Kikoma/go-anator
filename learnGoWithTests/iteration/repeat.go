package main

import (
	"fmt"
	"strings"
)

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func main(){
	fmt.Print(Repeat("a",5))
}
