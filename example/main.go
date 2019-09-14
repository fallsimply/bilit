package main

import (
	"github.com/SimplyCodin/bilit"
)

func main() {
	str := bilit.Populate("Hello, I'm {{name}} from {{City}}, {{from_state}}", bilit.Data{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	})
	print(str)
	popstr := "Hello, I'm John from Dallas, TX"
	test := bilit.Pull("Hello, I'm {{name}} from {{City}}, {{from_state}}", popstr)
	print(test)
}
