package ponteiroseerros

import "fmt"

type Bitcoin int

type Carteira struct {
    saldo  Bitcoin
}

type Stringer interface {
    String() string
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Depositar(quantidade Bitcoin)   {
    fmt.Printf("O endereço do saldo no Depositar é %v \n", &c.saldo)
    c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin  {
    return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) {
    c.saldo -= quantidade
}
