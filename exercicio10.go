package main

import "fmt"

/* Eu tenho um gato e um cachorro.

Eu os peguei na mesma época que o gatinho/cachorrinho. Isso foi humanYearsanos atrás.

Retorne suas respectivas idades agora como [ humanYears, catYears, dogYears]

NOTAS:

humanYears>= 1
	humanYearssão apenas números inteiros

Anos de gato
	15anos de gato para o primeiro ano
	+9anos de gato para o segundo ano
	+4anos de gato para cada ano depois disso

Anos de cão
	15anos de cachorro para o primeiro ano
	+9anos de cachorro para o segundo ano
	+5anos de cachorro para cada ano depois disso */

func main() {
	anos := 5
	humano := 20
	gato := 5
	cachorro := 7

	switch anos {
	case 1:
		humano = humano + 1
		gato = gato + 15
		cachorro += 15
		fmt.Println("Humano, Gato, Cachorro", humano, gato, cachorro)
	case 2:
		humano = humano + 2
		gato += 9
		cachorro += 9
		fmt.Println("Humano, Gato, Cachorro", humano, gato, cachorro)
	}

	if anos > 2 {
		humano = humano + anos
		gato = gato + (anos * 4)
		cachorro = cachorro + (anos * 5)
		fmt.Println("Humano, Gato, Cachorro", humano, gato, cachorro)
	}

}

/* package kata

func CalculateYears(anos int) (result [3]int) {

	if anos == 1 {
		return [3]int{1, 15, 15}
  } else if anos == 2 {
		return [3]int{2, 24, 24}
  }else{
		gato := 24 + (anos - 2) * 4
		cachorro := 24 + (anos - 2) * 5
		return [3]int{anos, gato, cachorro}
  }


}
*/

/* package kata

func CalculateYears(years int) (result [3]int) {
  switch years {
    case 1: result = [3]int{1, 15, 15}
    case 2: result = [3]int{2, 24, 24}
    default: result = [3]int{years, 24 + 4 * (years - 2), 24 + 5 * (years - 2)}
  }
  return
} */
