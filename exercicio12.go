package main

import "fmt"

func percorre(ideias []string) string {
	var numideiasboas int
	for _, ideia := range ideias {
		if ideia == "boa" {
			numideiasboas += 1
		}
	}

	if numideiasboas == 0 {
		return "falha"
	}

	if numideiasboas <= 2 {
		return "publica"
	}

	return "Sinto cheiro de séries!"
}

func main() {
	entrada := []string{"boa", "boa", "boa", "ruim", "ruim"} //slice
	fmt.Println(percorre(entrada))
}
