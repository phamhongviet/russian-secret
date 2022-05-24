package main

func GetCapacity(text string) int {
	bits := 0
	for _, c := range text {
		_, keyExist := Codebook[c]
		if keyExist {
			bits += 1
		}
	}
	bytes := bits/8 - 1
	return(bytes)
}
