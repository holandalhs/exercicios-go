/* ​https://www.codewars.com/kata/544675c6f971f7399a000e79
Descrição
Precisamos de uma função que possa transformar uma string em um número.
Quais maneiras de conseguir isso você conhece?

Observação: não se preocupe, todas as entradas serão strings e cada string
é uma representação perfeitamente válida de um número integral.

Exemplos
"1234" --> 1234
"605"  --> 605
"1405" --> 1405
"-7" --> -7 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	salario := "1000"
	aumento := "20"

	salarioConvertido, err := strconv.Atoi(salario)
	if err != nil {
		fmt.Println(err)
	}

	aumentoConvertido, err := strconv.Atoi(aumento)
	if err != nil {
		panic(err)
	}

	fmt.Println("O salário é", salario, "e o aumento de", aumento+"%")

	novoSalario := ((salarioConvertido * aumentoConvertido) / 100)
	novo := novoSalario + salarioConvertido
	fmt.Println("O novo salário é", novo)

}

/* package kata

import "strconv"

func StringToNumber(str string) int {
  conversao, err := strconv.Atoi(str)
  if err != nil{
    panic(err)
  }

  return conversao

} */
