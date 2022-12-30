package engagebay

const (
	PropertyTypeText = "TEXT"
)

type Contact struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Tag struct {
	Value     string `json:"tag"`
	CreatedAt int    `json:"assigned_time"`
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type CreateContactParams struct {
	Properties []Property `json:"properties,omitempty"`
	Tags       []string   `json:"tags,omitempty"`
}

func (p *CreateContactParams) SetName(name string) {
	p.AddProperty("name", name, "SYSTEM")
}

func (p *CreateContactParams) SetEmail(email string) {
	p.AddProperty("email", email, "SYSTEM")
}

func (p *CreateContactParams) AddTags(tags ...string) {
	p.Tags = append(p.Tags, tags...)
}

func (p *CreateContactParams) AddProperty(name, value, propType string) {
	p.Properties = append(p.Properties, Property{
		Name:  name,
		Value: value,
		Type:  propType,
	})
}
