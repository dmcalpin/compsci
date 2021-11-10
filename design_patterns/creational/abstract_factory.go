package creational

import "math"

// Abstract factory creates object
// factories which can create objects
// without specifying their classes
//
// A great explaination:
// https://www.tutorialspoint.com/design_pattern/abstract_factory_pattern.htm

// FactoryProducer creates different types
// of factories
type FactoryProducer struct{}

func (p *FactoryProducer) GetFactory(kind string) AbstractFactory {
	switch kind {
	case "round":
		return &RoundShapeFactory{}
	default:
		return &ShapeFactory{}
	}
}

type AbstractFactory interface {
	GetShape(string) Shape
}

// ShapeFactory implements AbstractFactory
// It can make any kind of Shape
type ShapeFactory struct{}

func (f *ShapeFactory) GetShape(name string) Shape {
	switch name {
	case "square":
		return &Square{}
	case "circle":
		return &Circle{}
	default:
		return nil
	}
}

// RoundShapeFactory implements AbstractFactory
// It only produces round shapes
type RoundShapeFactory struct{}

func (f *RoundShapeFactory) GetShape(name string) Shape {
	switch name {
	case "circle":
		return &Circle{}
	default:
		return nil
	}
}

// Shape is the type of object returned
// from the factories the abstract
// factory creates
type Shape interface {
	GetName() string
	GetArea() float64
}

// Square implements Shape
type Square struct {
	length float64
	width  float64
}

func (s *Square) GetName() string {
	return "Square"
}

func (s *Square) GetArea() float64 {
	return s.width * s.length
}

// Circle implements Shape
type Circle struct {
	radius float64
}

func (s *Circle) GetName() string {
	return "Circle"
}

func (s *Circle) GetArea() float64 {
	return math.Pi * s.radius * s.radius
}
