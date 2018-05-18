package entity

type Header struct {
	In              Location              `json:"-" builder:"required" default:"entity.InHeader"`
	Required        bool                  `json:"required,omitempty"`
	Description     string                `json:"description,omitempty"`
	Deprecated      bool                  `json:"deprecated,omitempty"`
	AllowEmptyValue bool                  `json:"allowEmptyValue,omitempty"`
	Explode         bool                  `json:"explode,omitempty"`
	AllowReserved   bool                  `json:"allowReserved,omitempty"`
	Schema          *Schema               `json:"schema,omitempty"`
	Example         interface{}           `json:"example,omitempty"`
	Examples        map[string]*Example   `json:"examples,omitempty"`
	Content         map[string]*MediaType `json:"content,omitempty"`
}
