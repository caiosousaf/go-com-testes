package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste", "sla": "Opa"}
	t.Run("Palavra conhecida", func(t *testing.T) {

		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"

		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, ErrNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {

	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario.Adiciona(palavra, definicao)

		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "teste novo")

		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("deveria ter encontrado palavra adicionada:", err)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s',  esperado '%s'", resultado, definicao)
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "teste")
	}
}

func TestUpdate(t *testing.T) {
	t.Run("palavra existente", func(t *testing.T) {
        palavra := "teste"
        definicao := "isso é apenas um teste"
        novaDefinicao := "nova definição"
        dicionario := Dicionario{palavra: definicao}
        err := dicionario.Atualiza(palavra, novaDefinicao)

        comparaErro(t, err, nil)
        comparaDefinicao(t, dicionario, palavra, novaDefinicao)
    })

    t.Run("palavra nova", func(t *testing.T) {
        palavra := "teste"
        definicao := "isso é apenas um teste"
        dicionario := Dicionario{}

        err := dicionario.Atualiza(palavra, definicao)

        comparaErro(t, err, ErrPalavraInexistente)
    })
}