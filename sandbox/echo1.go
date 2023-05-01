// Simple implementation of echo command
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// Print again but one char at a time as a string
	sep = " "
	for i := 1; i < len(os.Args); i++ {
		for j := 0; j < len(os.Args[i]); j++ {
			fmt.Print(string(os.Args[i][j]))
			// Wait x ms
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Print(string(sep))
	}
	fmt.Println()

	//  Going over arguments using range
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// Using Join Explicitly
	fmt.Print(string(strings.Join(os.Args[1:], " ")))
	fmt.Println()

	// Unpacking slice to fmt.Println as multiple arguments
	fmt.Println(interfaceSlice(os.Args[1:])...)
}

func interfaceSlice(slice []string) []interface{} {
	result := make([]interface{}, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}
