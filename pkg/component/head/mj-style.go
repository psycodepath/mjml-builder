package head

import "github.com/psycodepath/mjml-builder/pkg/component"

type MjStyle struct {
	inline  bool
	Content string
}

func (m *MjStyle) Type() string {
	return "mj-style"
}

func (m *MjStyle) Render() []byte {
	var attr []component.Attribute
	if m.inline == true {
		attr = append(attr, component.Attribute{Name: "inline", Content: "inline"})
	}
	return component.RenderComponent(m.Type(), &attr, m.Content)
}

func (m *MjStyle) AllowedAttributes() []string {
	return []string{"inline"}
}
