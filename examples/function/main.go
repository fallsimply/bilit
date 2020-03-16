package main

import (
	"fmt"

	"github.com/SimplyCodin/bilit"
)

func main() {
	bilit.Debug = true

	const template = "Hello, I'm [{{name}}] from (${City}), {{{from_state}}})"

	dataSrc := map[string]string{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	}

	str := bilit.Populate(template, dataSrc)
	fmt.Println(str)

	data := bilit.Pull(template, str)
	fmt.Println(data)
}
