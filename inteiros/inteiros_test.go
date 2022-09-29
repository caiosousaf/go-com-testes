package inteiros

import "testing"
	

func TestAdicionador(t *testing.T) {

    soma := Adiciona(2, 2)
    esperado := 4

    if soma != esperado {
        t.Errorf("esperado '%d', resultado '%d'", esperado, soma)
    }
}

func TestMultiplicador(t *testing.T) {
	multiplicar := Multiplicar(2, 2)
	esperado := 4

	if multiplicar != esperado {
        t.Errorf("esperado '%d', resultado '%d'", esperado, multiplicar)
    }
}