package main

import "testing"

func TestOla(t *testing.T) {

    t.Run("Teste-Idioma-BR", func (t *testing.T) {
        resultado := Ola("Caio", "portugues")
        esperado := "Olá, Caio"
    
        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }

    })

    t.Run("Teste-Idioma-ES", func(t *testing.T) {
        resultado := Ola("Bruno", "espanhol")
        esperado := "Hola, Bruno"

        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    })

    
}