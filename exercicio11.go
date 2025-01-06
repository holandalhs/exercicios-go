/* Pedra, papel, tesoura Vamos jogar!
Você tem que retornar qual jogador ganhou! Em caso de empate, retorne Draw!.

Exemplos(Entrada1, Entrada2 --> Saída):

"scissors", "paper" --> "Player 1 won!"
"scissors", "rock" --> "Player 2 won!"
"paper", "paper" --> "Draw!" */

package main

import "fmt"

func main() {
	jogador1 := "tesoura"
	jogador2 := "pedra"

	if jogador1 == jogador2 {
		fmt.Println("Deu empate!")
		return
	}
	//e/AND -> &&   ou/OR -> ||
	if jogador1 == "pedra" && jogador2 == "tesoura" {
		fmt.Println("Jogador 1 ganhou.")
		return
	}

	if jogador1 == "tesoura" && jogador2 == "papel" {
		fmt.Println("Jogador 1 ganhou.")
		return
	}

	if jogador1 == "papel" && jogador2 == "tesoura" {
		fmt.Println("Jogador 2 ganhou.")
		return
	}

	fmt.Println("Jogador 2 ganhou.")

}
