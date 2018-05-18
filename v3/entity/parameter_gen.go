package entity

type Parameter struct {
	Name            string                `json:"name,omitempty" builder:"required"`
	In              Location              `json:"in" builder:"required"`
	Required        bool                  `json:"required,omitempty" default:"defaultParameterRequiredFromLocation(in)"`
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
