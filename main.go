package main

import "fmt"

type Shape interface {
	Area() float64
	Perimetro() float64
}

func main() {

	myCircle := Circle{Radius: 5}
	myRentangle := Rectangle{Ancho: 5, Alto: 5}
	myTriangle := Triangle{Altura: 5, Base: 3, Lado: 6}

	fmt.Println("Datos del círculo")
	PrintShapeProperties(myCircle)
	fmt.Println("Datos del rectángulo")
	PrintShapeProperties(myRentangle)
	fmt.Println("Datos del triangulo")
	PrintShapeProperties(myTriangle)

}

// Datos del triángulo
type Triangle struct {
	Altura float64
	Base   float64
	Lado   float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func (t Triangle) Perimetro() float64 {
	return t.Lado * 3
}

// Datos del Rectángulo
type Rectangle struct {
	Ancho float64
	Alto  float64
}

func (r Rectangle) Area() float64 {
	return r.Alto * r.Ancho
}

func (r Rectangle) Perimetro() float64 {
	return 2 * (r.Alto + r.Ancho)
}

// Datos del circulo
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * 3.14
}

func (c Circle) Perimetro() float64 {
	return 2 * 3.14 * c.Radius
}

func PrintShapeProperties(s Shape) {
	fmt.Printf("Area : %f\n", s.Area())
	fmt.Printf("Perímetro : %f\n", s.Perimetro())
}
