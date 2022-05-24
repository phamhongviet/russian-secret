package main

import (
	"fmt"
)

func main() {
	data := []byte("Steganography is fun")
	data = append(data, 0x4)
	fmt.Printf("[Original data]:\n%v\n", data)
	fmt.Printf("[Decoded data, ascii]:\n%s\n", string(data))
	text := `The Internet is the global system of interconnected computer networks that uses the Internet protocol suite (TCP/IP) to communicate between networks and devices. It is a network of networks that consists of private, public, academic, business, and government networks of local to global scope, linked by a broad array of electronic, wireless, and optical networking technologies. The Internet carries a vast range of information resources and services, such as the inter-linked hypertext documents and applications of the World Wide Web (WWW), electronic mail, telephony, and file sharing. `
	fmt.Printf("[Original text]:\n%s\n", text)

	c := GetCapacity(text)
	fmt.Printf("[Capacity]: %d bytes\n", c)
	fmt.Printf("[Data length]: %d bytes\n", len(data))
	encodedText := Encode(text, data)
	fmt.Printf("[Encoded text]:\n%s\n", encodedText)

	decodedData := Decode(encodedText, 0x4)
	fmt.Printf("[Decoded data]:\n%v\n", decodedData)
	fmt.Printf("[Decoded data, ascii]:\n%s\n", string(decodedData))
}
