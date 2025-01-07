/*
	https://www.codewars.com/kata/56f3a1e899b386da78000732/train/go

Escreva uma função partlistque forneça todas as maneiras de dividir uma lista (uma matriz) de pelo menos dois elementos em duas partes não vazias.

Cada duas partes não vazias estarão em um par (ou uma matriz para linguagens sem tuplas ou structem C - C: veja Exemplos Casos de teste - )
Cada parte estará em uma string
Os elementos de um par devem estar na mesma ordem que na matriz original.

não pode ter partes vazias
*/
package main

import "fmt"

func main() {
	palavras := []string{"I", "wish", "I", "hadn't", "come"}
	fmt.Println(len(palavras))

	for i := 0; i < len((palavras)); i++ {
		fmt.Println(palavras[:1]) //pega tudo antes da posição 1
		fmt.Println(palavras[1:]) //pega da posição 1 até o final
	}

}
