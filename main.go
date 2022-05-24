package main

import (
	"fmt"
)

func main() {
	/*
	fmt.Printf("Codebook:\n")
	for k, v := range Codebook {
		fmt.Printf("Character: %c \t Value: %d \t Character: %c \t Value: %d \n", k, k, v, v)
	}
	fmt.Printf("Reverse Codebook:\n")
	for k, v := range ReverseCodebook {
		fmt.Printf("Character: %c \t Value: %d \t Character: %c \t Value: %d \n", k, k, v, v)
	}
	*/
	c := GetCapacity(`

	EOT is one of a number of control characters used by seerial devices. There are a number of other control characters which are related to transmission of data over serial lines or storage of files on a serial source like paper tape. These include characters such as SOH, STX, ETX, FS, RS, GS, and, US. Additional control characters are used for transmission control and error correction.

	On a serial connection an EOT (End Of Transmission) character indicates a desire to end the transmission. Serial connections are usually accessed using a file driver. When the serial transmission ends, the file driver reports this as an EOF (End Of File) condition.

	EOF is not a character. getchar() returns an integer. A valid character while will have a value in the range 0 to 255. The value of -1 is often used as false/invalid/fail indicator on Unix/Linux. (Actually a non 0 value, as there are any number of reasons not to succeed, but usually only one success case.) When getchar() returns -1 it is clearly not returning a character. However, if you store the output in a byte, you won't be able to distinguish EOF from the DEL (Delete) character.
	`)
	fmt.Printf("Capacity: %d bytes\n", c)
}
