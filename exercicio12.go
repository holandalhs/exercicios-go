/* Para cada boa ideia de kata, parece haver algumas ruins!

Neste kata você precisa verificar o array fornecido (x) para boas ideias 'good'
e más ideias 'bad'. Se houver uma ou duas boas ideias, retorne 'Publish!', se
houver mais de 2, retorne 'I smell a series!'. Se não houver boas ideias, como
é frequentemente o caso, retorne 'Fail!'.
https://www.codewars.com/kata/57f222ce69e09c3630000212 */

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
