package entities

type ArgumentSet struct{}

type void struct{}

var (
	argumentSet map[string]void
	entry       void
)

func NewArgumentSet() ArgumentSet {
	argumentSet = make(map[string]void)
	return ArgumentSet{}
}

func (a ArgumentSet) Add(argument string) {
	argumentSet[argument] = entry
}

func (a ArgumentSet) Has(argument string) bool {
	_, found := argumentSet[argument]
	return found
}

func (a ArgumentSet) Length() int {
	return len(argumentSet)
}
