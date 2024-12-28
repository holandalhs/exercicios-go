/* ​https://www.codewars.com/kata/5265326f5fda8eb1160004c8
Convertendo número para string */

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	salario := 7800

	nome := strconv.Itoa(salario)
	fmt.Println(nome)

	fmt.Println("%d para número", salario)
	fmt.Println(fmt.Sprintf("%s para string", nome))
	//o Sprintf transforma qualquer coisa em string

	fmt.Println(reflect.TypeOf(salario))
	fmt.Println(reflect.TypeOf(nome))

}

/* package kata

import "strconv"

func NumberToString(n int) string {
  nome := strconv.Itoa(n)
  return nome
}
*/

/* package kata

import "strconv"

func NumberToString(n int) string {
  return strconv.Itoa(n)  //SÓ UMA LINHA DE CÓDIGO, RETORNO DIRETO SEM CRIAR String
}
*/
