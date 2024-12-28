// ​https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad
/*
Dado um conjunto de números, retorne o inverso aditivo de cada um.
Cada positivo se torna negativo, e os negativos se tornam positivos.

[1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
[1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
[] --> []

*/

package main

import "fmt"

func main() {
	//listaNumPrimos := [6]int{2, 3, 5, 7, 11, 13}
	listaNumNegativos := [5]int{-7, -9, -8, -3, -4}

	for _, x := range listaNumNegativos {
		w := x
		w = w * -1
		fmt.Println(w)
	}

}

/*
package kata


func Invert(arr []int) []int {
  var listainverte []int

  for _, numero := range arr {

      listainverte = append(listainverte, numero * (-1))
  }

  return listainverte
}
*/
