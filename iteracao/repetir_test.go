package iteracao

import "testing"

const quantidade = 5
func TestRepetir(t *testing.T) {
    repeticoes := Repetir("a", quantidade)
    esperado := "aaaaa"

    if repeticoes != esperado {
        t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
    }
}

func BenchmarkRepetir(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repetir("a", quantidade)
    }
}