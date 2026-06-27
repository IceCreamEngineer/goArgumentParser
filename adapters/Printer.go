package adapters

import "fmt"

type Printer struct{}

func (p Printer) Present(message string) {
	fmt.Println(message)
}
