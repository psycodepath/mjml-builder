package head

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMjTitle_Type(t *testing.T) {
	m := MjTitle{}
	assert.Equal(t, "mj-title", m.Type())
}

func TestMjTitle_Render(t *testing.T) {
	m := MjTitle{content: "This is expected"}
	expected := "<mj-title>This is expected</mj-title>"
	assert.Equal(t, expected, string(m.Render()))
}

func TestMjTitle_AllowedAttributes(t *testing.T) {
	m := MjTitle{}
	assert.Equal(t, []string{}, m.AllowedAttributes())
}
