/* https://www.codewars.com/kata/580755730b5a77650500010c
ÍNDICES
Dada uma string s, sua tarefa é retornar outra string de modo que caracteres indexados pares e ímpares ssejam agrupados e os grupos sejam separados por espaços. O grupo indexado par vem primeiro, seguido por um espaço e, então, pela parte indexada ímpar.

Exemplos
input:    "CodeWars" => "CdWr oeas"
           ||||||||      |||| ||||
indices:   01234567      0246 1357
Índices pares 0, 2, 4, 6, então temos "CdWr"como primeiro grupo.
Índices ímpares são 1, 3, 5, 7, então o segundo grupo é "oeas".
E a string final a retornar é "Cdwr oeas".

Notas
As strings testadas têm pelo menos 8 caracteres. */

package main

import "fmt"

func main() {
	palavra := "CaBeDiJu"
	var letraspares, letrasimpares string //letras dos índices pares

	for i := 0; i < len(palavra); i++ {
		if i%2 == 0 {
			//fmt.Println(string(palavra[i]))
			letraspares += string(palavra[i])
		} else {
			letrasimpares += string(palavra[i])
		}

	}
	fmt.Println(letraspares, letrasimpares)

	fmt.Println(fmt.Sprintf("%s %s", letraspares, letrasimpares))

}
