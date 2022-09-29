package arrays

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	numeros := []int{1, 2, 3, 4, 5}

	resultado := Soma(numeros)
	esperado := 15

	if esperado != resultado {
		t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
	}

	t.Run("soma-2", func(t *testing.T) {
		resultado := Soma(numeros)
		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
        numeros := []int{1, 2, 3}

        resultado := Soma(numeros)
        esperado := 6

        if resultado != esperado {
            t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
        }
    })
}

func TestSomaTudo(t *testing.T) {

    recebido := SomaTudo([]int{1,2}, []int{0,9})
    esperado := []int{3, 9}

    if !reflect.DeepEqual(recebido, esperado) {
        t.Errorf("recebido %v esperado %v", recebido, esperado)
    }
}

func TestSomaTodoOResto(t *testing.T) {
	verificaSomas := func(t *testing.T, resultado, esperado []int) {
        t.Helper()
        if !reflect.DeepEqual(resultado, esperado) {
            t.Errorf("resultado %v, esperado %v", resultado, esperado)
        }
    }

	t.Run("Faz-Soma-Slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1,2}, []int{0,9})
		esperado := []int{2, 9}
	
		verificaSomas(t,resultado, esperado)
	})

	t.Run("soma-slices-vazio", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
        esperado := []int{0, 9}

        verificaSomas(t, resultado, esperado)
	})
}