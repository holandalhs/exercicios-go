/*
	exercício 1:

https://www.codewars.com/kata/55685cd7ad70
*/
package main

import "fmt"

func main() {

	//percorrendo uma lista de inteiros usando for-range
	numeros := []int{-6, -2, 0, 4, 8, 10}

	for _, numeroescolhido := range numeros {

		if numeroescolhido < 0 {
			fmt.Println("número escolhido é negativo", numeroescolhido)

		} else if numeroescolhido > 0 {
			fmt.Println("número escolhido é maior que zero", numeroescolhido)
			numeroescolhido = numeroescolhido * -1
			fmt.Println("valor transformado", numeroescolhido)
		} else {
			fmt.Println("valor igual a zero")
		}
	}

}

/* //forma mais simples
package kata

func MakeNegative(x int) int {
  if x <= 0 {
    return x
  } else {
    return -x
  }
} */
