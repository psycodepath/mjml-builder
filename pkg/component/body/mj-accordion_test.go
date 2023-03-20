package body

import (
	"fmt"
	"github.com/psycodepath/mjml-builder/pkg/component"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	expected := "<mj-accordion-title border=\"1px solid black\" background-color=\"lime\" css-class=\"foo\" font-family=\"Comic Sans\" padding=\"40px 20px\" padding-top=\"40px\" padding-right=\"20px\" padding-bottom=\"40px\" padding-left=\"20px\">This is Sparta</mj-accordion-title>"

	assert.Equal(t, expected, string(actual))
}

// text

func TestMjAccordionText_AllowedAttributes(t *testing.T) {

}

func TestMjAccordionText_Attributes(t *testing.T) {

}

func TestMjAccordionText_Type(t *testing.T) {

}

func TestMjAccordionText_Render(t *testing.T) {

}

// Element
func TestMjAccordionElement_AllowedAttributes(t *testing.T) {

}

func TestMjAccordionElement_Attributes(t *testing.T) {

}

func TestMjAccordionElement_Type(t *testing.T) {

}

func TestMjAccordionElement_Render(t *testing.T) {

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
