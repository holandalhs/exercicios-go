/* Escreva um programa que encontre a soma de cada número de 1 a num.
O número sempre será um inteiro positivo maior que 0. Sua função só precisa
retornar o resultado, o que é mostrado entre parênteses no exemplo abaixo é
como você chega a esse resultado e não faz parte dele, veja os testes de exemplo.
https://www.codewars.com/kata/55d24f55d7dd296eb9000030
*/
//forma de incrementar o número de 1 até n/ ele

package kata

/* package main

import "fmt" */

func Summation(n int) int {
	// the sleeper must awaken!
	var somatorio int
	for i := 1; i <= n; i++ {
		somatorio = somatorio + i
	}
	//fmt.Println("o valor final é", somatorio)
	return somatorio

}
