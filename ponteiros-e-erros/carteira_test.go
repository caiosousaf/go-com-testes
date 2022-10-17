package ponteiroseerros

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(20)}
		carteira.Retirar(10)
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)

		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}

func confirmaErro(t *testing.T, resultado error, esperado error) {
    t.Helper()
    if resultado == nil {
        t.Fatal("esperava um erro, mas nenhum ocorreu")
    }

    if resultado != esperado {
        t.Errorf("erro resultado %s, erro esperado %s", resultado, esperado)
    }
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}