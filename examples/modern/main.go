package main

import (
	"fmt"

	"github.com/SimplyCodin/bilit"
)

func main() {
	bilit.Debug = true

	template := bilit.Template("Hello, I'm [{{name}}] from (${City}), {{{from_state}}}")

	dataSrc := map[string]string{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	}

	str := template.Populate(dataSrc)
	fmt.Println("[Modern API]", str)

	data := template.Pull(str)
	fmt.Println("[Modern API]", data)
}
