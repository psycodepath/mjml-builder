package head

import (
	"fmt"
	"github.com/psycodepath/mjml-builder/pkg/component"
)

type MjBreakPoint struct {
	width string
}

func (m *MjBreakPoint) attributes() *[]component.Attribute {
	return &[]component.Attribute{{Name: "width", Content: m.width}}
}

func (m *MjBreakPoint) Type() string {
	return "mj-breakpoint"
}

func (m *MjBreakPoint) Render() []byte {
	return []byte(fmt.Sprintf("<mj-breakpoint width=\"%s\"/>", m.width))
}

func (m *MjBreakPoint) AllowedAttributes() []string {
	return []string{"width"}
}
