package entities

type ArgumentSchemaElement struct {
	Name, ArgumentType, Description, LongName string
	Required                                  *bool
}

// IsRequired Default Required to true if unset
func (a *ArgumentSchemaElement) IsRequired() bool { return a.Required == nil || *a.Required }
