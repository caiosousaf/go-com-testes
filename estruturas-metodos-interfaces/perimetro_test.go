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
    testesArea := []struct {
		nome	string
        forma    Forma
        esperado float64
    }{
		{nome:"Retangulo", forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
        {nome:"Circulo", forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
        {nome:"Triangulo", forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
    }

    for _, tt := range testesArea {
        resultado := tt.forma.Area()
        if resultado != tt.esperado {
            t.Errorf("%#v resultado %.2f, esperado %.2f",tt.forma, resultado, tt.esperado)
        }
    }
}

// func TestArea(t *testing.T) {

// 	verificaArea := func(t *testing.T, forma Forma, esperado float64)  {
// 		t.Helper()
// 		resultado := forma.Area()
// 		if resultado != esperado {
// 			t.Errorf("resultado '%.2f', esperado '%2.f'", resultado, esperado)
// 		}
// 	}

// 	t.Run("retangulos", func(t *testing.T) {
// 		retangulo := Retangulo{12.0, 6.0}
// 		esperado := 72.0
// 		verificaArea(t, retangulo, esperado)
// 	})

// 	t.Run("circulos", func(t *testing.T) {
// 		circulo := Circulo{10}
// 		esperado := 314.1592653589793

// 		verificaArea(t, circulo, esperado)
// 	})
// }