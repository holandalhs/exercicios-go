/* ​https://www.codewars.com/kata/57a0e5c372292dd76d000d7e

Escreva uma função que aceite um inteiro n e uma string s como parâmetros e
retorne uma string de vezes sexatas repetidas n.

Exemplos (entrada -> saída)
6, "I"     -> "IIIIII"
5, "Hello" -> "HelloHelloHelloHelloHello" */

package main

import "fmt"

func main() {
	numero := 3
	nome := "Lua"

	for i := 1; i <= numero; i++ {
		fmt.Println(nome)
	}

}

/*
package kata

func RepeatStr(repetitions int, value string) string {
  var repetir string

  for i := 1; i <= repetitions; i++ {
    repetir = repetir + value

  }
  return repetir      //NÃO PODE USAR O RETURN DENTRO DO FOR PORQUE já sai do loop
					  //e não faz a contagem

} */

//OUTRA FORMA:
/* package kata

import "strings"

func RepeatStr(repititions int, value string) string {
  return strings.Repeat(value, repititions)
} */
