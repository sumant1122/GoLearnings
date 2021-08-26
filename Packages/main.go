package main

import (
	"github.com/packages/email"
	"github.com/packages/invoice"
)

func main() {

	invoice.Create("ABC", 2, 4.5)
	email.Contents("Hey There", "ABC", 2)
	email.Send("Hmmm", "yaaa")
}
