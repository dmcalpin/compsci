package structural

// Decorators add functionality to existing
// Classes. In this example we have two typical
// fast food items: Burger, French Fries, and
// two typical add-ons: Katchup, Cheese.
// All types implement FoodItem, but the
// decorators (Cheese and Katchup) have knowledge
// about the base food item, and they add their
// own description and cost to that base item
// i.e. decoration.

// FoodItem is something you can buy at a
// restaurant. It has a cost and a description
type FoodItem interface {
	GetCost() float64
	GetDescription() string
}

// BaseFoodItem implements FoodItem and
// can be embedded by FootItem structs
type BaseFoodItem struct {
	cost        float64
	description string
}

func (i *BaseFoodItem) GetCost() float64 {
	return i.cost
}

func (i *BaseFoodItem) GetDescription() string {
	return i.description
}

// Burger implements FoodItem
type Burger struct {
	BaseFoodItem
}

func NewBurger() *Burger {
	return &Burger{BaseFoodItem{2.99, "A Hamburger"}}
}

// Fries implements FoodItem
type Fries struct {
	BaseFoodItem
}

func NewFries() *Fries {
	return &Fries{BaseFoodItem{1.99, "French Fries"}}
}

// Katchup is a decorator which
// adds katchup to any FoodItem
type Katchup struct {
	BaseFoodItem
	rootItem FoodItem
}

func AddKatchup(f FoodItem) FoodItem {
	return &Katchup{
		rootItem: f,
	}
}

// Katchup adds $0.25 to the cost
func (k *Katchup) GetCost() float64 {
	return k.rootItem.GetCost() + .25
}

func (k *Katchup) GetDescription() string {
	return k.rootItem.GetDescription() + ", with Katchup"
}

type Cheese struct {
	BaseFoodItem
	rootItem FoodItem
}

func AddCheese(f FoodItem) FoodItem {
	return &Cheese{
		rootItem: f,
	}
}

// Cheese adds $0.50 to the cost
func (k *Cheese) GetCost() float64 {
	return k.rootItem.GetCost() + .50
}

func (k *Cheese) GetDescription() string {
	return k.rootItem.GetDescription() + ", with Cheese"
}
