package head

import "github.com/psycodepath/mjml-builder/pkg/component"

type MjTitle struct {
	content string
}

func (m *MjTitle) Type() string {
	return "mj-title"
}

func (m *MjTitle) Render() []byte {
	return component.RenderComponent(m.Type(), &[]component.Attribute{}, m.content)
}

func (m *MjTitle) AllowedAttributes() []string {
	return []string{}
}
