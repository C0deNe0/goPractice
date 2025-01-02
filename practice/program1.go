/*
Design and implement ShapeFactory class that generates different types of Shape
objects (Circle, Square, Rectangle) based on input parameters using Factory Design Pattern.
*/
package main

import (
	"errors"
	"fmt"
)

// shape interface
type Shape interface {
	DrawShape() string
}

// circle struct
type Circle struct {
	radius float64
}

func (c Circle) DrawShape() string {
	return fmt.Sprintf(" Drawing a cicle with radius :%.2f", c.radius)
}

// square struct
type Square struct {
	side float64
}

func (s Square) DrawShape() string {
	return fmt.Sprintf(" Drawing a square with side: %.2f", s.side)

}

// shape Factory function
type shapeFactory struct{}

// CreateShape method to generate shapes based on type and parameters
func (sf shapeFactory) CreateShape(shapetype string, params ...float64) (Shape, error) {

	switch shapetype {
	case "circle":
		if len(params) != 1 {
			return nil, errors.New("Circle requires 1 parameter: radius")
		}
		return Circle{radius: params[0]}, nil

	case "square":
		if len(params) != 1 {
			return nil, errors.New("Square requires 1 parameter: side length")
		}
		return Square{side: params[0]}, nil

	default:
		return nil, fmt.Errorf("Unknown shape type: %v", shapetype)
	}

}

func main() {

	factory := shapeFactory{}

	circle, err := factory.CreateShape("circle", 3.5)
	if err == nil {
		fmt.Println(circle.DrawShape())
	} else {
		fmt.Println(err)
	}

	square, err := factory.CreateShape("square", 5.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(square.DrawShape())
	}
}
