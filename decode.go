package main

func Decode(text string, stop byte) []byte {
	data := make([]byte, 0)
	i := 0
	for _, c := range text {
		_, e1 := ReverseCodebook[c]
		if e1 {
			if (i % 8) == 0 {
				data = append(data, 0)
			}
			data[i/8] = data[i/8] + (1 << (7 - i%8))
			i += 1
			if i%8 == 7 {
				if data[i/8] == stop {
					break
				}
			}
		}
		_, e2 := Codebook[c]
		if e2 {
			if i%8 == 0 {
				data = append(data, 0)
			}
			i += 1
			if i%8 == 7 {
				if data[i/8] == stop {
					break
				}
			}
		}
	}
	return (data)
}
