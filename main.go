package main

import (
	"fmt"

	cb "russian-secret/codebook"
)

func main() {
	fmt.Printf("Codebook:\n")
	for k, v := range cb.Codebook {
		fmt.Printf("Character: %c \t Value: %d \t Character: %c \t Value: %d \n", k, k, v, v)
	}
	fmt.Printf("Reverse Codebook:\n")
	for k, v := range cb.ReverseCodebook {
		fmt.Printf("Character: %c \t Value: %d \t Character: %c \t Value: %d \n", k, k, v, v)
	}
}
