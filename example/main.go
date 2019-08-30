package main

import (
	"github.com/SimplyCodin/bilit"
)

func main() {
	str := bilit.Populate("Hello, I'm {{name}} from {{City}}, {{from_state}}", map[string]string{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	})
	print(str)
}