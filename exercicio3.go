/* Você obtém uma matriz de números e retorna a soma de todos os positivos.
Exemplo [1,-4,7,12]=>1 + 7 + 12 = 20
Observação: se não houver nada para somar, a soma padrão será 0. */

package main

import "fmt"

func main() {
	//percorrendo uma lista de inteiros usando for-range
	numeros := []int{-6, -2, 0, 4, 8, 10}

	var positivo int

	for _, somapositivo := range numeros {
		if somapositivo > 0 {
			//positivo = positivo + somapositivo
			positivo += somapositivo
		} else {
			fmt.Println("Número não positivo", somapositivo)
		}
	}
	fmt.Println("A soma dos números positivos é:", positivo)

}

/*

//https://www.codewars.com/kata/5715eaedb436cf5606000381​
package kata

func PositiveSum(numbers []int) int {
  var positivo int

	for _, somapositivo := range numbers {
    if somapositivo > 0 {
			positivo += somapositivo
    }
 	}
  return positivo

}
*/
