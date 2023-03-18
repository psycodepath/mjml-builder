package component

import (
	"bytes"
	"fmt"
)

type Component interface {
	Type() string
	Render() []byte
	AllowedAttributes() []string
}

type Attribute struct {
	Name    string
	Content string
}

func (a *Attribute) Render() string {
	return fmt.Sprintf("%s=\"%s\"", a.Name, a.Content)
}

func RenderComponent(componentType string, attributes *[]Attribute, content string) []byte {
	var sb bytes.Buffer
	sb.WriteString(fmt.Sprintf("<%s", componentType))

	for _, a := range *attributes {
		sb.WriteString(" ")
		sb.WriteString(a.Render())
	}
	sb.WriteString(fmt.Sprintf(">%s</%s>", content, componentType))
	return sb.Bytes()
}
