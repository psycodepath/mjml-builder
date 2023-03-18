package head

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMjBreakPoint_AllowedAttributes(t *testing.T) {
	c := MjBreakPoint{}
	attr := c.AllowedAttributes()
	assert.Contains(t, attr, "width", "MjBreakpoint should allow the attribute width")
	assert.Equal(t, 1, len(attr), "MjBreakpoint should only allow 1 attribute")
}

func TestMjBreakPoint_Render(t *testing.T) {
	c := MjBreakPoint{width: "200px"}
	expected := "<mj-breakpoint width=\"200px\"/>"
	assert.Equal(t, expected, string(c.Render()))
}

func TestMjBreakPoint_Type(t *testing.T) {
	c := MjBreakPoint{}
	assert.Equal(t, "mj-breakpoint", c.Type(), "MJBreakpoint.Type() should return \"mj-breakpoint\"")
}
