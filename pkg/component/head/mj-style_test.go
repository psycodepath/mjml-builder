package head

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMjStyle_AllowedAttributes(t *testing.T) {
	c := MjStyle{}
	allowed := c.AllowedAttributes()
	assert.Contains(t, allowed, "inline", "MjStyle should allow the attribute inline")
	assert.Exactly(t, 1, len(allowed), "MjStyle should only allow 1 attribute")
}

func TestMjStyle_Type(t *testing.T) {
	c := MjStyle{}
	assert.Equal(t, "mj-style", c.Type(), "MjStyle.Type() expected to return MjStyle")
}

func TestMjStyle_Render(t *testing.T) {
	e1 := "<mj-style inline=\"inline\">#x{background: lime;}</mj-style>"
	e2 := "<mj-style>#x{background: pink;}</mj-style>"

	c1 := MjStyle{inline: true, Content: "#x{background: lime;}"}
	c2 := MjStyle{inline: false, Content: "#x{background: pink;}"}

	assert.Equal(t, e1, string(c1.Render()))
	assert.Equal(t, e2, string(c2.Render()))
}
