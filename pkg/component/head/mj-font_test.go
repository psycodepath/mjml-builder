package head

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMjFont_AllowedAttributes(t *testing.T) {
	c := MjFont{}
	allowed := c.AllowedAttributes()
	assert.Contains(t, allowed, "href", "MjFont should allow the attribute href")
	assert.Contains(t, allowed, "name", "MjFont should allow the attribute name")
	assert.Exactly(t, 2, len(allowed), "MjFont should allow exactly 2 attributes")
}

func TestMjFont_Type(t *testing.T) {
	c := MjFont{}
	assert.Equal(t, "mj-font", c.Type(), "MjFont.Type() should return \"mj-font\"")
}

func TestMjFont_Render(t *testing.T) {
	expected := "<mj-font href=\"https://fonts.googleapis.com/css2?family=Roboto\" name=\"Roboto\"></mj-font>"
	c := MjFont{href: "https://fonts.googleapis.com/css2?family=Roboto", name: "Roboto"}
	assert.Equal(t, expected, string(c.Render()))
}
