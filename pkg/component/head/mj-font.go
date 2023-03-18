package head

import "github.com/psycodepath/mjml-builder/pkg/component"

type MjFont struct {
	href string
	name string
}

func (f *MjFont) attributes() *[]component.Attribute {
	a := make([]component.Attribute, 2)
	a[0] = component.Attribute{Name: "href", Content: f.href}
	a[1] = component.Attribute{Name: "name", Content: f.name}
	return &a
}

func (m *MjFont) Type() string {
	return "mj-font"
}

func (m *MjFont) Render() []byte {
	return component.RenderComponent(m.Type(), m.attributes(), "")
}

func (m *MjFont) AllowedAttributes() []string {
	return []string{"href", "name"}
}
