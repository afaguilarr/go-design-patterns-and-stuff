package main

import "fmt"

type color int

const (
	red color = iota
	green
	blue
)

type size int

const (
	small size = iota
	medium
	large
)

type product struct {
	name  string
	color color
	size  size
}

type specification interface {
	isSatisfied(p *product) bool
}

type colorSpecification struct {
	color color
}

func (spec colorSpecification) isSatisfied(p *product) bool {
	return p.color == spec.color
}

type sizeSpecification struct {
	size size
}

func (spec sizeSpecification) isSatisfied(p *product) bool {
	return p.size == spec.size
}

type andSpecification struct {
	ss []specification
}

func (spec andSpecification) isSatisfied(p *product) bool {
	for _, s := range spec.ss {
		if !(s.isSatisfied(p)) {
			return false
		}
	}
	return true
}

type orSpecification struct {
	ss []specification
}

func (spec orSpecification) isSatisfied(p *product) bool {
	for _, s := range spec.ss {
		if s.isSatisfied(p) {
			return true
		}
	}
	return false
}

func filter(products []product, spec specification) []*product {
	result := []*product{}
	for i, v := range products {
		if spec.isSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := product{"Apple", green, small}
	tree := product{"Tree", green, large}
	house := product{"House", blue, large}

	products := []product{apple, tree, house}

	greenSpec := colorSpecification{green}
	largeSpec := sizeSpecification{large}

	fmt.Print("Green products (new):\n")
	for _, v := range filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeGreenSpec := andSpecification{[]specification{largeSpec, greenSpec}}
	largeOrGreenSpec := orSpecification{[]specification{largeSpec, greenSpec}}

	fmt.Print("Large green items:\n")
	for _, v := range filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}

	fmt.Print("Large or green items:\n")
	for _, v := range filter(products, largeOrGreenSpec) {
		fmt.Printf(" - %s is large or green\n", v.name)
	}
}
