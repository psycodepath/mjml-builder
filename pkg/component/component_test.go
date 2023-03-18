package component

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttribute_Render(t *testing.T) {
	const expected = "href=\"https://www.example.com\""
	a := Attribute{Name: "href", Content: "https://www.example.com"}
	assert.Equal(t, expected, a.Render())
}

func TestRenderComponent(t *testing.T) {
	var attributes = make([]Attribute, 1)
	expected := []byte("<mj-style inline=\"inline\">.blue-text{color: red}</mj-style>")
	attributes[0] = Attribute{
		Name:    "inline",
		Content: "inline",
	}
	assert.Equal(t, expected, RenderComponent("mj-style", &attributes, ".blue-text{color: red}"))

}
