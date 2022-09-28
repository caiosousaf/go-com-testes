package main

import "testing"

func TestOla(t *testing.T) {
    verificaMensagem := func(t *testing.T, resultado, esperado string)  {
        t.Helper()
        if resultado != esperado {
            t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
        }
    }

    t.Run("Teste-Idioma-BR", func (t *testing.T) {
        resultado := Ola("Caio", "portugues")
        esperado := "Olá, Caio"
    
        verificaMensagem(t, resultado, esperado)

    })

    t.Run("Teste-Idioma-ES", func(t *testing.T) {
        resultado := Ola("Bruno", "espanhol")
        esperado := "Hola, Bruno"

        verificaMensagem(t, resultado, esperado)
    })

    t.Run("RetornarMundoQuandoNomeVazio", func(t *testing.T) {
        resultado := Ola("", "portugues")
        esperado := "Olá, mundo"

        verificaMensagem(t, resultado, esperado)
    })

    t.Run("Teste-Idioma-EN-Nome-Vazio", func(t *testing.T) {
        resultado := Ola("", "ingles")
        esperado := "Hello, world"

        verificaMensagem(t, resultado, esperado)
    })

    t.Run("Teste-Idioma-ES-Nome-Vazio", func(t *testing.T) {
        resultado := Ola("", "espanhol")
        esperado := "Hola, mundo"

        verificaMensagem(t, resultado, esperado)
    })

    t.Run("Teste-Idioma-Nao-Encontrado", func(t *testing.T) {
        resultado := Ola("Caio", "")
        esperado := "Caio,O idioma não foi encontrado"

        verificaMensagem(t, resultado, esperado)
    })
}