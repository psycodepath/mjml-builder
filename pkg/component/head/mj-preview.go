package head

import "github.com/psycodepath/mjml-builder/pkg/component"

type MjPreview struct {
	content string
}

func (m *MjPreview) Type() string {
	return "mj-preview"
}

func (m *MjPreview) Render() []byte {
	return component.RenderComponent(m.Type(), &[]component.Attribute{}, m.content)
}

func (m *MjPreview) AllowedAttributes() []string {
	return []string{}
}
