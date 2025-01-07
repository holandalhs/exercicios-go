/* Nosso time de futebol terminou o campeonato.
https://www.codewars.com/kata/5bb904724c47249b10000131
Os resultados das partidas do nosso time são registrados em uma coleção de strings. Cada partida é representada por uma string no formato "x:y", onde xé a pontuação do nosso time e yé a pontuação do nosso oponente.

Por exemplo: ["3:1", "2:2", "0:1", ...]

Os pontos são atribuídos para cada partida da seguinte forma:

se x > y: 3 pontos (vitória)
se x < y: 0 pontos (perda)
se x = y: 1 ponto (empate)
Precisamos escrever uma função que pegue essa coleção e retorne o número de pontos
que nosso time ( x) obteve no campeonato pelas regras fornecidas acima.

Notas:

nosso time sempre joga 10 partidas no campeonato
0 <= x <= 4
0 <= e <= 4 */

//Testing for Points['1:0','2:0','3:0','4:0','2:1','3:1','4:1','3:2','4:2','4:3']
//percorrer a lista e depois separa os pontos do time x e do time adversário
//entender se ganhou ou perdeu ou empatou
//somar os pontos de cada situação

package main

import (
	"fmt"
	"strings"
)

func Pontos(games []string) int {
	var pontoscorridos int
	for _, pontuacao := range games {
		placar := strings.Split(pontuacao, ":")
		fmt.Println(placar[0])
		fmt.Println(placar[1])

		if placar[0] == placar[1] {
			pontoscorridos += 1
			fmt.Println("empate")
		}
		if placar[0] > placar[1] {
			pontoscorridos += 3
			fmt.Println("vitória")
		}
		if placar[0] < placar[1] {
			fmt.Println("o time não pontua, derrota")
		}
	}
	return pontoscorridos
}

func main() { //declaração de um slice de strings
	jogos := []string{"1:1", "3:2", "3:4", "4:4", "5:2", "3:1", "1:4", "3:3", "4:1", "4:4"}
	fmt.Println("O time finalizou o campeonato com", Pontos(jogos), "pontos")

}
