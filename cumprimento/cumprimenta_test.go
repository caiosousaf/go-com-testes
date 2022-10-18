package cumprimento

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}

	Cumprimenta(&buffer, "Caio")

	resultado := buffer.String()
	esperado := "Olá, Caio"

	if resultado != esperado {
        t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
    }
}