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

    t.Run("RetornarMundoQuandoNomeVazio", func(t *testing.T) {
        resultado := Ola("", "portugues")
        esperado := "Olá, mundo"

        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    })

    t.Run("Teste-Idioma-EN-Nome-Vazio", func(t *testing.T) {
        resultado := Ola("", "ingles")
        esperado := "Hello, world"

        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    })

    t.Run("Teste-Idioma-ES-Nome-Vazio", func(t *testing.T) {
        resultado := Ola("", "espanhol")
        esperado := "Hola, mundo"

        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    })

    t.Run("Teste-Idioma-Nao-Encontrado", func(t *testing.T) {
        resultado := Ola("Caio", "")
        esperado := "Caio,O idioma não foi encontrado"

        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    })
}