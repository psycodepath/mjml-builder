package body

import (
	"fmt"
	"github.com/psycodepath/mjml-builder/pkg/component"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const expectedTitle = "<mj-accordion-title border=\"1px solid black\" background-color=\"lime\" css-class=\"foo\" font-family=\"Comic Sans\" padding=\"40px 20px\" padding-top=\"40px\" padding-right=\"20px\" padding-bottom=\"40px\" padding-left=\"20px\">This is Sparta</mj-accordion-title>"
const expectedText = "<mj-accordion-text border=\"1px solid black\" background-color=\"brown\" css-class=\"foo\" color=\"#F00\" padding=\"40px 20px\" padding-top=\"40px\" padding-right=\"20px\" padding-bottom=\"40px\" padding-left=\"20px\" font-family=\"Comic Sans\" font-size=\"12px\" font-weight=\"700\" letter-spacing=\"37px\" line-height=\"100px\">tiny man makes tiny library</mj-accordion-text>"

func getExpectedTitleValues() map[string]string {
	return map[string]string{
		"border":           "1px solid black",
		"background-color": "lime",
		"css-class":        "foo",
		"font-family":      "Comic Sans",
		"padding":          "40px 20px",
		"padding-top":      "40px",
		"padding-right":    "20px",
		"padding-bottom":   "40px",
		"padding-left":     "20px",
	}
}

func createTitleWithValues(values map[string]string) MjAccordionTitle {
	return MjAccordionTitle{
		Border:     values["border"],
		Background: values["background-color"],
		CssClass:   values["css-class"],
		FontFamily: values["font-family"],
		Padding: &component.Padding{
			Padding: values["padding"],
			Top:     values["padding-top"],
			Right:   values["padding-right"],
			Bottom:  values["padding-bottom"],
			Left:    values["padding-left"],
		},
	}
}

func getExpectedTextValues() map[string]string {
	return map[string]string{
		"border":           "1px solid black",
		"background-color": "brown",
		"css-class":        "foo",
		"color":            "#F00",
		"padding":          "40px 20px",
		"padding-top":      "40px",
		"padding-right":    "20px",
		"padding-bottom":   "40px",
		"padding-left":     "20px",
		"font-family":      "Comic Sans",
		"font-size":        "12px",
		"font-weight":      "700",
		"letter-spacing":   "37px",
		"line-height":      "100px",
	}
}

func createTextWithValues(values map[string]string) MjAccordionText {
	return MjAccordionText{
		Border:     values["border"],
		Background: values["background-color"],
		CssClass:   values["css-class"],
		Color:      values["color"],
		Padding: &component.Padding{
			Padding: values["padding"],
			Top:     values["padding-top"],
			Right:   values["padding-right"],
			Bottom:  values["padding-bottom"],
			Left:    values["padding-left"],
		},
		Font: &component.Font{
			Family:        values["font-family"],
			Size:          values["font-size"],
			Weight:        values["font-weight"],
			LetterSpacing: values["letter-spacing"],
			LineHeight:    values["line-height"],
		},
	}
}

func getExpectedElementValues() map[string]string {
	return map[string]string{
		"border":             "1px solid black",
		"background-color":   "bordeaux",
		"css-class":          "foo",
		"font-family":        "Comic Sans",
		"padding":            "40px 20px",
		"padding-top":        "40px",
		"padding-right":      "20px",
		"padding-bottom":     "40px",
		"padding-left":       "20px",
		"icon-align":         "left",
		"icon-height":        "32px",
		"icon-position":      "left",
		"icon-unwrapped-alt": "An image",
		"icon-unwrapped-url": "http://image.example/turd.gif",
		"icon-width":         "100px",
		"icon-wrapped-alt":   "Another image",
		"icon-wrapped-url":   "http:///image.example/bag.jpg",
	}
}

func createElementWithValues(values map[string]string) MjAccordionElement {
	return MjAccordionElement{
		Border:     values["border"],
		Background: values["background-color"],
		CssClass:   values["css-class"],
		FontFamily: values["font-family"],
		Padding: &component.Padding{
			Padding: values["padding"],
			Top:     values["padding-top"],
			Right:   values["padding-right"],
			Bottom:  values["padding-bottom"],
			Left:    values["padding-left"],
		},
		Icon: createIconFromValues(values),
	}
}

func createIconFromValues(values map[string]string) *component.Icon {
	return &component.Icon{
		Align:        values["icon-align"],
		Height:       values["icon-height"],
		Position:     values["icon-position"],
		UnwrappedAlt: values["icon-unwrapped-alt"],
		UnwrappedUrl: values["icon-unwrapped-url"],
		Width:        values["icon-width"],
		WrappedAlt:   values["icon-wrapped-alt"],
		WrappedUrl:   values["icon-wrapped-url"],
	}
}

func getExpectedElementRender() string {
	var sb strings.Builder

	sb.WriteString("<mj-accordion-element border=\"1px solid black\" background-color=\"bordeaux\" css-class=\"foo\" font-family=\"Comic Sans\" padding=\"40px 20px\" padding-top=\"40px\" padding-right=\"20px\" padding-bottom=\"40px\" padding-left=\"20px\" icon-align=\"left\" icon-height=\"32px\" icon-position=\"left\" icon-unwrapped-alt=\"An image\" icon-unwrapped-url=\"http://image.example/turd.gif\" icon-width=\"100px\" icon-wrapped-alt=\"Another image\" icon-wrapped-url=\"http:///image.example/bag.jpg\">")
	sb.WriteString(expectedTitle)
	sb.WriteString(expectedText)
	sb.WriteString("</mj-accordion-element>")

	return sb.String()
}

// title
func TestMjAccordionTitle_AllowedAttributes(t *testing.T) {

	title := MjAccordionTitle{}
	actual := title.AllowedAttributes()

	message := func(n string) string {
		return fmt.Sprintf("MjAccordionTitle needs to to allow %s attribute", n)
	}

	assert.Exactly(t, 9, len(actual))
	assert.Contains(t, actual, "border", message("border"))
	assert.Contains(t, actual, "background-color", message("background-color"))
	assert.Contains(t, actual, "css-class", message("border"))
	assert.Contains(t, actual, "font-family", message("font-family"))
	assert.Contains(t, actual, "padding", message("padding"))
	assert.Contains(t, actual, "padding-top", message("padding-top"))
	assert.Contains(t, actual, "padding-right", message("padding-right"))
	assert.Contains(t, actual, "padding-bottom", message("padding-bottom"))
	assert.Contains(t, actual, "padding-left", message("padding-left"))

}

func TestMjAccordionTitle_Attributes(t *testing.T) {

	values := getExpectedTitleValues()
	title := createTitleWithValues(values)

	actual := title.Attributes()
	allowed := title.AllowedAttributes()

	// happy path with every attribute set
	assert.Exactly(t, 9, len(actual))
	for _, attr := range actual {
		assert.Contains(t, allowed, attr.Name, fmt.Sprintf("%s is not an allowed attribute", attr.Name))
		assert.Equal(t, values[attr.Name], attr.Content, fmt.Sprintf("Excpected \"%s\" but got \"%s\"", values[attr.Name], attr.Content))
	}

}

func TestMjAccordionTitle_Type(t *testing.T) {
	title := MjAccordionTitle{}
	assert.Equal(t, "mj-accordion-tile", title.Type())
}

func TestMjAccordionTitle_Render(t *testing.T) {
	title := createTitleWithValues(getExpectedTitleValues())
	title.Title = "This is Sparta"
	actual := title.Render()

	assert.Equal(t, expectedTitle, string(actual))
}

// text

func TestMjAccordionText_AllowedAttributes(t *testing.T) {

	title := MjAccordionText{}
	actual := title.AllowedAttributes()

	message := func(n string) string {
		return fmt.Sprintf("MjAccordionText needs to to allow %s attribute", n)
	}

	assert.Exactly(t, 13, len(actual))
	assert.Contains(t, actual, "border", message("border"))
	assert.Contains(t, actual, "background-color", message("background-color"))
	assert.Contains(t, actual, "css-class", message("border"))
	assert.Contains(t, actual, "padding", message("padding"))
	assert.Contains(t, actual, "padding-top", message("padding-top"))
	assert.Contains(t, actual, "padding-right", message("padding-right"))
	assert.Contains(t, actual, "padding-bottom", message("padding-bottom"))
	assert.Contains(t, actual, "padding-left", message("padding-left"))
}

func TestMjAccordionText_Attributes(t *testing.T) {
	values := getExpectedTextValues()
	text := createTextWithValues(values)

	actual := text.Attributes()
	allowed := text.AllowedAttributes()

	// happy path with every attribute set
	assert.Exactly(t, 9, len(actual))
	for _, attr := range actual {
		assert.Contains(t, allowed, attr.Name, fmt.Sprintf("%s is not an allowed attribute", attr.Name))
		assert.Equal(t, values[attr.Name], attr.Content, fmt.Sprintf("Excpected \"%s\" but got \"%s\"", values[attr.Name], attr.Content))
	}
}

func TestMjAccordionText_Type(t *testing.T) {
	text := MjAccordionText{}
	assert.Equal(t, "mj-accordion-text", text.Type())
}

func TestMjAccordionText_Render(t *testing.T) {
	values := getExpectedTextValues()
	text := createTextWithValues(values)
	text.Content = "tiny man makes tiny library"
	assert.Equal(t, expectedText, string(text.Render()))
}

// Element
func TestMjAccordionElement_AllowedAttributes(t *testing.T) {
	values := getExpectedElementValues()
	element := createElementWithValues(values)

	actual := element.Attributes()
	allowed := element.AllowedAttributes()

	// happy path with every attribute set
	assert.Exactly(t, 9, len(actual))
	for _, attr := range actual {
		assert.Contains(t, allowed, attr.Name, fmt.Sprintf("%s is not an allowed attribute", attr.Name))
		assert.Equal(t, values[attr.Name], attr.Content, fmt.Sprintf("Excpected \"%s\" but got \"%s\"", values[attr.Name], attr.Content))
	}
}

func TestMjAccordionElement_Attributes(t *testing.T) {
	values := getExpectedElementValues()
	text := createElementWithValues(values)

	actual := text.Attributes()
	allowed := text.AllowedAttributes()

	// happy path with every attribute set
	assert.Exactly(t, 9, len(actual))
	for _, attr := range actual {
		assert.Contains(t, allowed, attr.Name, fmt.Sprintf("%s is not an allowed attribute", attr.Name))
		assert.Equal(t, values[attr.Name], attr.Content, fmt.Sprintf("Excpected \"%s\" but got \"%s\"", values[attr.Name], attr.Content))
	}
}

func TestMjAccordionElement_Type(t *testing.T) {
	element := MjAccordionElement{}
	assert.Equal(t, "mj-accordion-element", element.Type())
}

func TestMjAccordionElement_Render(t *testing.T) {
	element := createElementWithValues(getExpectedElementValues())
	title := createTitleWithValues(getExpectedTitleValues())
	text := createTextWithValues(getExpectedTextValues())
	element.Title = &title
	element.Title.Title = "This is Sparta"
	element.Text = &text
	element.Text.Content = "tiny man makes tiny library"
	actual := element.Render()
	assert.Equal(t, getExpectedElementRender(), string(actual))
}

// accordion

func TestMjAccordion_AllowedAttributes(t *testing.T) {

}

func TestMjAccordion_Attributes(t *testing.T) {

}

func TestMjAccordion_Type(t *testing.T) {

}

func TestMjAccordion_Render(t *testing.T) {

}
