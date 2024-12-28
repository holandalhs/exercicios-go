/* exercício 4:
https://www.codewars.com/kata/576bb71bbbcf0951d5000044
*/

/*  Dado um array de inteiros.
Retorna uma matriz, onde o primeiro elemento é a contagem de números positivos e
o segundo elemento é a soma de números negativos. 0 não é positivo nem negativo.
 Se a entrada for uma matriz vazia ou for nula, retorne uma matriz vazia.
 Exemplo
Para entrada [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], você deve retornar [10, -65].
*/

package main

import "fmt"

func main() {
	//percorrendo uma matriz usando for-range, exemplo:
	//coral := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
	//fmt.Println(coral)
	numeros := []int{-6, -2, 0, 4, 8, 10, 0, -4, 90}

	var positivo, negativo int

	for _, soma := range numeros {
		if soma > 0 {
			//ele não quer a soma (USARIA: positivo += soma) e sim a CONTAGEM dos números positivos
			//positivo = positivo + 1
			//positivo += 1
			positivo++
		} else {
			negativo += soma
		}
	}
	fmt.Println("Soma dos números positivos e negativos", positivo, negativo)
	//fmt.Println("Soma dos números negativos", negativo)

}

/* package kata

func CountPositivesSumNegatives(numbers []int) []int {
  var positivo, negativo int

	for _, soma := range numbers {
		if soma > 0 {
			positivo++
		} else {
			negativo += soma
		}
	}
	return []int{positivo, negativo}

} */
