package entities

type ArgumentSchemaElement struct {
	Name, ArgumentType, Description, LongName string
	IsRequired                                bool
}

func NewArgumentSchemaElement(schemaElementFields ...func(*ArgumentSchemaElement)) *ArgumentSchemaElement {
	schemaElement := &ArgumentSchemaElement{IsRequired: true}
	for _, schemaElementField := range schemaElementFields {
		schemaElementField(schemaElement)
	}
	return schemaElement
}
