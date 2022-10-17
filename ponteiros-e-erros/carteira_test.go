package ponteiroseerros

import (
	"fmt"
	"testing"
)

func verificaMensagem(t *testing.T, resultado, esperado Bitcoin)  {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestCarteira(t *testing.T) {
	t.Run("Deposito", func(t *testing.T) {
		carteira := Carteira{}
	
		carteira.Depositar(Bitcoin(10))
	
		resultado := carteira.Saldo()
	
		fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)
	
		esperado := Bitcoin(10)
	
		verificaMensagem(t, resultado, esperado)
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: 20}

		carteira.Retirar(Bitcoin(10))

		resultado := carteira.saldo

		esperado := Bitcoin(10)

		verificaMensagem(t, esperado, resultado)
	})
}

