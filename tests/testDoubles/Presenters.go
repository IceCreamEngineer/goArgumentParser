package testDoubles

type PresenterSpy struct{}

var presented string

func NewPresenterSpy() PresenterSpy {
	presented = ""
	return PresenterSpy{}
}

func (p PresenterSpy) Present(message string) {
	presented += message
}

func (p PresenterSpy) GetPresented() string {
	return presented
}
