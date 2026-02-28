package entities

type ArgumentSchemaElement struct {
	Name, ArgumentType, Description, LongName string
	IsRequired                                bool `default:"true"`
}
