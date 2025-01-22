package main

import (
	"fmt"
	"strings"
	)

func main() {
	var input string
	fmt.Println("Enter a string : ")
	fmt.Scanln(&input)
	
	res := stringPattern(input)	
	
	if res {
	fmt.Println("Valid String")	
	} else {
	fmt.Println("Invalid String")
	}
}


func stringPattern(input string) bool {
 if len(input) < 3 {
	return false
 }

 if (strings.HasPrefix(input, "i") || strings.HasPrefix(input, "I")) && (strings.Contains(input, "a") || strings.Contains(input, "A")) && (strings.HasSuffix(input, "n") || strings.HasSuffix(input, "N")) {
	return true
} else  {
	return false 
}
}
