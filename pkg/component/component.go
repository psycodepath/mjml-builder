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

type Padding struct {
	Padding string
	Top     string
	Right   string
	Bottom  string
	Left    string
}

func (p *Padding) Attributes() []Attribute {
	var attributes []Attribute

	if p.Padding != "" {
		attributes = append(attributes, Attribute{Name: "padding", Content: p.Padding})
	}

	if p.Top != "" {
		attributes = append(attributes, Attribute{Name: "padding-top", Content: p.Top})
	}

	if p.Right != "" {
		attributes = append(attributes, Attribute{Name: "padding-right", Content: p.Right})
	}

	if p.Bottom != "" {
		attributes = append(attributes, Attribute{Name: "padding-bottom", Content: p.Bottom})
	}

	if p.Left != "" {
		attributes = append(attributes, Attribute{Name: "padding-left", Content: p.Left})
	}

	return attributes
}

func (p *Padding) AllowedAttributes() []string {
	return []string{
		"padding",
		"padding-top",
		"padding-right",
		"padding-bottom",
		"padding-left",
	}
}

type Icon struct {
	Align        string
	Height       string
	Position     string
	UnwrappedAlt string
	UnwrappedUrl string
	Width        string
	WrappedAlt   string
	WrappedUrl   string
}

func (i *Icon) Attributes() []Attribute {
	var attributes []Attribute

	if i.Align != "" {
		attributes = append(attributes, Attribute{Name: "icon-align", Content: i.Align})
	}

	if i.Height != "" {
		attributes = append(attributes, Attribute{Name: "icon-height", Content: i.Height})
	}

	if i.Position != "" {
		attributes = append(attributes, Attribute{Name: "icon-position", Content: i.Position})
	}

	if i.UnwrappedAlt != "" {
		attributes = append(attributes, Attribute{Name: "icon-unwrapped-alt", Content: i.UnwrappedAlt})
	}

	if i.UnwrappedUrl != "" {
		attributes = append(attributes, Attribute{Name: "icon-unwrapped-url", Content: i.UnwrappedUrl})
	}

	if i.Width != "" {
		attributes = append(attributes, Attribute{Name: "icon-width", Content: i.Width})
	}

	if i.WrappedAlt != "" {
		attributes = append(attributes, Attribute{Name: "icon-wrapped-alt", Content: i.WrappedAlt})
	}

	if i.WrappedUrl != "" {
		attributes = append(attributes, Attribute{Name: "icon-wrapped-url", Content: i.WrappedUrl})
	}

	return attributes
}

func (i *Icon) AllowedAttributes() []string {
	return []string{
		"icon-align",
		"icon-top",
		"icon-height",
		"icon-position",
		"icon-unwrapped-alt",
		"icon-unwrapped-url",
		"icon-width",
		"icon-wrapped-alt",
		"icon-wrapped-url",
	}
}

type Font struct {
	Family        string
	Size          string
	Weight        string
	LetterSpacing string
	LineHeight    string
}

func (f *Font) Attributes() []Attribute {
	var attributes []Attribute

	if f.Family != "" {
		attributes = append(attributes, Attribute{Name: "font-family", Content: f.Family})
	}

	if f.Size != "" {
		attributes = append(attributes, Attribute{Name: "font-size", Content: f.Size})
	}

	if f.Weight != "" {
		attributes = append(attributes, Attribute{Name: "font-weight", Content: f.Weight})
	}

	if f.LetterSpacing != "" {
		attributes = append(attributes, Attribute{Name: "letter-spacing", Content: f.LetterSpacing})
	}

	if f.LineHeight != "" {
		attributes = append(attributes, Attribute{Name: "line-height", Content: f.LineHeight})
	}

	return attributes
}

func (f *Font) AllowedAttributes() []string {
	return []string{
		"font-family",
		"font-size",
		"font-weight",
		"letter-spacing",
		"line-height",
	}
}
