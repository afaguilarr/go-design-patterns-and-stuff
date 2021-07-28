package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	ss []Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	for _, s := range spec.ss {
		if !(s.IsSatisfied(p)) {
			return false
		}
	}
	return true
}

type OrSpecification struct {
	ss []Specification
}

func (spec OrSpecification) IsSatisfied(p *Product) bool {
	for _, s := range spec.ss {
		if s.IsSatisfied(p) {
			return true
		}
	}
	return false
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}

	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}
	bf := BetterFilter{}

	fmt.Print("Green products (new):\n")
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeGreenSpec := AndSpecification{[]Specification{largeSpec, greenSpec}}
	largeOrGreenSpec := OrSpecification{[]Specification{largeSpec, greenSpec}}

	fmt.Print("Large green items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

	fmt.Print("Large or green items:\n")
	for _, v := range bf.Filter(products, largeOrGreenSpec) {
		fmt.Printf(" - %s is large or green\n", v.name)
	}
}
