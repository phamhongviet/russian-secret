package main

import (
	"bytes"
	"fmt"
)

func GetCapacity(text string) int {
	bits := 0
	for _, c := range text {
		_, keyExist := Codebook[c]
		if keyExist {
			bits += 1
		}
	}
	bytes := bits/8 - 1
	return (bytes)
}

func bytesToBits(data []byte) string {
	var bits bytes.Buffer
	for _, n := range data {
		s := fmt.Sprintf("%08b", n)
		bits.WriteString(s)
	}
	return (fmt.Sprintf("%s", bits))
}

func Encode(text string, data []byte) string {
	// TODO check capacity
	bits := bytesToBits(data)
	textBytes := []int32(text)
	i := 0
	for t, c := range textBytes {
		v, keyExist := Codebook[c]
		if keyExist {
			if bits[i] == '1' {
				textBytes[t] = v
			}
			i += 1
		}
	}
	return (string(textBytes))
}
