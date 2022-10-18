package dicionario


type Dicionario map[string]string

type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
} 

var (
	ErrNaoEncontrado = ErrDicionario("não foi possivel encontrar a palavra que voce procura")
	ErrPalavraExistente = ErrDicionario("não é possivel adicionar a palavra pois ela já existe")
)


func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}

    return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
    switch err {
    case ErrNaoEncontrado:
        d[palavra] = definicao
    case nil:
        return ErrPalavraExistente
    default:
        return err

    }

    return nil
}