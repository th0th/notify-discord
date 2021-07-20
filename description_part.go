package main

import "fmt"

type DescriptionPart struct {
	Name  string
	Value string
}

func (p DescriptionPart) ToString() string {
	return fmt.Sprintf("**%s**: %s", p.Name, p.Value)
}

func GetDescription(parts []DescriptionPart) string {
	d := ""

	for i, part := range parts {
		d += part.ToString()

		if i < len(parts) {
			d += "\n\n"
		}
	}

	return d
}
