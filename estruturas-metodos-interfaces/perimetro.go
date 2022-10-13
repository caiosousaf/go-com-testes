package estruturasmetodosinterfaces

import "math"

type Forma interface {
    Area() float64
}

type Retangulo struct {
	Largura float64
	Altura	float64
}
func (r Retangulo) Area() float64  {
    return r.Largura * r.Altura
}

type Circulo struct {
	Raio	float64
}
func (c Circulo) Area() float64  {
    return math.Pi * c.Raio * c.Raio
}


func Perimetro(retangulo Retangulo) float64{
	
	return 2 * (retangulo.Largura + retangulo.Altura)
}

