package creational

import (
	"fmt"
)

// A builder creates complex objects from simple
// objects. In this example I use the concept
// of building a computer. We have 2 types of
// computers, PCs and Macs, each has their own
// way of being configured and built into a list
// of parts. But the parts adhere to a common
// interface of Part which has the GetName and
// GetPrice methods.

// ComputerBuilder is configured to make either
// a PC or a Mac. Each contains a list of parts
// which have a name and a price, and a special
// field unique to the type of part
type ComputerBuilder struct{}

func (b *ComputerBuilder) makePC() *Computer {
	pc := &Computer{
		operatingSystem: "Windows 10",
	}

	pc.AddPart(&CPU{
		BasePart: BasePart{
			name:  "Intel i7",
			price: 299.85,
		},
		clockSpeed: 4.5,
	})

	pc.AddPart(&Case{
		BasePart: BasePart{
			name:  "Coolmax Pro",
			price: 99.23,
		},
		size: "Full",
	})

	pc.AddPart(&Motherboard{
		BasePart: BasePart{
			name:  "MSI Gamer Pro",
			price: 199.99,
		},
		form: "Mini ITX",
	})

	pc.AddPart(&Memory{
		BasePart: BasePart{
			name:  "Crucial 8x4",
			price: 179.99,
		},
		capacity: 32,
	})

	pc.AddPart(&HardDrive{
		BasePart: BasePart{
			name:  "Samsung Evo 970",
			price: 179.99,
		},
		capacity: 500,
	})

	return pc
}

func (b *ComputerBuilder) makeMac() *Computer {
	pc := &Computer{
		operatingSystem: "macOS",
	}

	pc.AddPart(&CPU{
		BasePart: BasePart{
			name:  "8-Core Processor",
			price: 200.00,
		},
		clockSpeed: 4.5,
	})

	pc.AddPart(&Case{
		BasePart: BasePart{
			name:  "Macbook Pro",
			price: 1000.00,
		},
		size: "14\"",
	})

	pc.AddPart(&Memory{
		BasePart: BasePart{
			name:  "Apple Memory",
			price: 200.00,
		},
		capacity: 16,
	})

	pc.AddPart(&HardDrive{
		BasePart: BasePart{
			name:  "Apple SSD",
			price: 300.00,
		},
		capacity: 500,
	})

	return pc
}

type Computer struct {
	parts           []Part
	operatingSystem string
}

func (c *Computer) AddPart(p Part) {
	c.parts = append(c.parts, p)
}

func (c *Computer) ListParts() string {
	partsList := ""

	for _, p := range c.parts {
		partsList += "- " + p.GetName() + "\n"
	}

	return partsList
}

func (c *Computer) GetTotalPrice() float64 {
	price := 0.0

	for _, p := range c.parts {
		price += p.GetPrice()
	}

	return price
}

type Part interface {
	GetName() string
	GetPrice() float64
}

type BasePart struct {
	name  string
	price float64
}

func (p *BasePart) GetName() string {
	return p.name
}
func (p *BasePart) GetPrice() float64 {
	return p.price
}

type Motherboard struct {
	BasePart
	// ITX, ATX, etc.
	form string
}

func (mb *Motherboard) GetName() string {
	return fmt.Sprintf("%s: %s", mb.name, mb.form)
}

type CPU struct {
	BasePart
	// in htz
	clockSpeed float64
}

func (c *CPU) GetName() string {
	return fmt.Sprintf("%s: %.02f", c.name, c.clockSpeed)
}

type Memory struct {
	BasePart
	// in GB
	capacity int
}

func (m *Memory) GetName() string {
	return fmt.Sprintf("%s: %dGB", m.name, m.capacity)
}

type HardDrive struct {
	BasePart
	// in GB
	capacity int
}

func (d *HardDrive) GetName() string {
	return fmt.Sprintf("%s: %dGB", d.name, d.capacity)
}

type Case struct {
	BasePart
	// Full, Mid, Mini
	size string
}

func (c *Case) GetName() string {
	return fmt.Sprintf("%s: %s", c.name, c.size)
}
