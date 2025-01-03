/* Crie uma função que receba um inteiro como argumento e retorne "Even"para
números pares ou "Odd"ímpares.
*/
//https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/solutions/go?filter=me&sort=best_practice&invalids=false

package kata

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}
