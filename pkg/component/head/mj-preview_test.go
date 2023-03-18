package head

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMjPreview_AllowedAttributes(t *testing.T) {
	preview := MjPreview{}
	assert.Nil(t, preview.AllowedAttributes())
}

func TestMjPreview_Type(t *testing.T) {
	preview := MjPreview{}
	assert.Equal(t, "mj-preview", preview.Type())
}

func TestMjPreview_Render(t *testing.T) {
	preview := MjPreview{content: "This is the content"}
	expected := "<mj-preview>This is the content</mj-preview>"

	assert.Equal(t, expected, string(preview.Render()))
}
