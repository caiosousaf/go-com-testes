package olamundo


func Ola(nome, idioma string) string {
    const (
        prefixoOlaPortugues = "Olá, "
        prefixoOlaEspannhol = "Hola, "
        prefixoOlaIngles = "Hello, "
    )

    if nome == "" {
        switch idioma {
        case "ingles":
            nome = "world"
        case "espanhol":
            nome = "mundo"
        case "portugues":
            nome = "mundo"
        }
    }
    switch idioma {
    case "portugues":
        return prefixoOlaPortugues + nome
    case "espanhol":
        return prefixoOlaEspannhol + nome
    case "ingles":
        return prefixoOlaIngles + nome
    }
    return nome + ",O idioma não foi encontrado"
}

