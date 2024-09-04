package elements

import (
	"fmt"
)

type Element struct {
	name       string
	attributes []fmt.Stringer
	children   []fmt.Stringer
}

func New(name string, attributes ...fmt.Stringer) func(children ...fmt.Stringer) Element {
	el := Element{name: name, attributes: attributes}

	return func(children ...fmt.Stringer) Element {
		el.children = children

		return el
	}
}

func (el Element) open() string {
	attributes := ""

	for _, attr := range el.attributes {
		attributes += fmt.Sprint(attr)
	}

	return "<" + el.name + attributes + ">"
}

func (el Element) close() string {
	return "</" + el.name + ">"
}

func (el Element) String() string {
	children := ""

	for _, child := range el.children {
		children += fmt.Sprint(child)
	}

	return el.open() + children + el.close()
}
