package body

import (
	"github.com/psycodepath/mjml-builder/pkg/component"
	"strings"
)

type MjAccordion struct {
	Border     string
	Background string
	CssClass   string
	FontFamily string
	Padding    *component.Padding
	Icon       *component.Icon
	Elements   []*MjAccordionElement
}

func (m *MjAccordion) Type() string {
	return "mj-accordion"
}

func (m *MjAccordion) Attributes() []component.Attribute {
	var attributes []component.Attribute

	if m.Border != "" {
		attributes = append(attributes, component.Attribute{Name: "border", Content: m.Border})
	}

	if m.Background != "" {
		attributes = append(attributes, component.Attribute{Name: "container-background-color", Content: m.Background})
	}

	if m.FontFamily != "" {
		attributes = append(attributes, component.Attribute{Name: "font-family", Content: m.FontFamily})
	}

	if m.CssClass != "" {
		attributes = append(attributes, component.Attribute{Name: "css-class", Content: m.CssClass})
	}

	attributes = append(attributes, m.Padding.Attributes()...)
	attributes = append(attributes, m.Icon.Attributes()...)
	return attributes
}

func (m *MjAccordion) Render() []byte {
	var attributes []component.Attribute = m.Attributes()
	var sb = strings.Builder{}

	for _, element := range m.Elements {
		sb.Write(element.Render())
	}

	return component.RenderComponent(m.Type(), &attributes, "")
}

func (m MjAccordion) AllowedAttributes() []string {
	var allowed []string

	allowed = append(allowed, m.Padding.AllowedAttributes()...)
	allowed = append(allowed, m.Icon.AllowedAttributes()...)

	return allowed
}

type MjAccordionElement struct {
	Title      *MjAccordionTitle
	Text       *MjAccordionText
	Border     string
	Background string
	CssClass   string
	FontFamily string
	Padding    *component.Padding
	Icon       *component.Icon
}

func (m *MjAccordionElement) Attributes() []component.Attribute {
	var attributes []component.Attribute

	if m.Border != "" {
		attributes = append(attributes, component.Attribute{Name: "border", Content: m.Border})
	}

	if m.Background != "" {
		attributes = append(attributes, component.Attribute{Name: "background-color", Content: m.Background})
	}

	if m.CssClass != "" {
		attributes = append(attributes, component.Attribute{Name: "css-class", Content: m.CssClass})
	}

	if m.FontFamily != "" {
		attributes = append(attributes, component.Attribute{Name: "font-family", Content: m.FontFamily})
	}

	attributes = append(attributes, m.Padding.Attributes()...)
	attributes = append(attributes, m.Icon.Attributes()...)

	return attributes
}

func (a *MjAccordionElement) Type() string {
	return "mj-accordion-element"
}

func (m *MjAccordionElement) Render() []byte {
	var sb strings.Builder

	sb.WriteString(string(m.Title.Render()))
	sb.WriteString(string(m.Text.Render()))

	attributes := m.Attributes()
	return component.RenderComponent(m.Type(), &attributes, sb.String())
}

func (m *MjAccordionElement) AllowedAttributes() []string {

	panic("implement me")
}

type MjAccordionTitle struct {
	Title      string
	Border     string
	Background string
	CssClass   string
	FontFamily string
	Padding    *component.Padding
}

func (m *MjAccordionTitle) Attributes() []component.Attribute {
	var attributes []component.Attribute

	if m.Border != "" {
		attributes = append(attributes, component.Attribute{Name: "border", Content: m.Border})
	}

	if m.Background != "" {
		attributes = append(attributes, component.Attribute{Name: "background-color", Content: m.Background})
	}
	if m.CssClass != "" {
		attributes = append(attributes, component.Attribute{Name: "css-class", Content: m.CssClass})
	}

	if m.FontFamily != "" {
		attributes = append(attributes, component.Attribute{Name: "font-family", Content: m.FontFamily})
	}

	attributes = append(attributes, m.Padding.Attributes()...)

	return attributes
}

func (m *MjAccordionTitle) Type() string {
	return "mj-accordion-title"
}

func (m *MjAccordionTitle) Render() []byte {
	attributes := m.Attributes()
	return component.RenderComponent(m.Type(), &attributes, m.Title)
}

func (m *MjAccordionTitle) AllowedAttributes() []string {
	return append(m.Padding.AllowedAttributes(), []string{
		"border",
		"background-color",
		"css-class",
		"font-family",
	}...)
}

type MjAccordionText struct {
	Content    string
	Border     string
	Background string
	CssClass   string
	Color      string
	Font       *component.Font
	Padding    *component.Padding
}

func (m *MjAccordionText) Attributes() []component.Attribute {
	var attributes []component.Attribute

	if m.Border != "" {
		attributes = append(attributes, component.Attribute{Name: "border", Content: m.Border})
	}

	if m.Background != "" {
		attributes = append(attributes, component.Attribute{Name: "background-color", Content: m.Background})
	}
	if m.CssClass != "" {
		attributes = append(attributes, component.Attribute{Name: "css-class", Content: m.CssClass})
	}

	if m.Color != "" {
		attributes = append(attributes, component.Attribute{Name: "color", Content: m.Color})
	}

	attributes = append(attributes, m.Padding.Attributes()...)
	attributes = append(attributes, m.Font.Attributes()...)

	return attributes
}

func (m *MjAccordionText) Type() string {
	return "mj-accordion-text"
}

func (m *MjAccordionText) Render() []byte {
	attributes := m.Attributes()
	return component.RenderComponent(m.Type(), &attributes, m.Content)
}

func (m *MjAccordionText) AllowedAttributes() []string {

	attributes := []string{"border", "background-color", "css-class"}
	attributes = append(attributes, m.Padding.AllowedAttributes()...)
	attributes = append(attributes, m.Font.AllowedAttributes()...)

	return attributes
}
