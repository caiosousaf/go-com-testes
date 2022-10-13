package estruturasmetodosinterfaces

import "testing"



func verificaMensagem(t *testing.T, resultado, esperado float64)  {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%.2f', esperado '%2.f'", resultado, esperado)
	}
}

func TestPerimetro(t *testing.T) {

	retangulo := Retangulo{12.0, 8.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	verificaMensagem(t, resultado, esperado)
}	

func TestArea(t *testing.T) {

	t.Run("retangulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		resultado := retangulo.Area()
		esperado := 72.0
	
		verificaMensagem(t, resultado, esperado)
	})

	t.Run("circulos", func(t *testing.T) {
		circulo := Circulo{10}
		resultado := circulo.Area()
		esperado := 314.1592653589793

		verificaMensagem(t, resultado, esperado)
	})
}