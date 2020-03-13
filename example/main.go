package main

import (
	//"encoding/json"
	"fmt"

	"github.com/SimplyCodin/bilit"
)

func main() {
	bilit.Debug = true
	const template = "Hello, I'm {{name}} from ${City}, {{from_state}}"
	var dataSrc = map[string]string{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	}

	var str = bilit.Populate(template, dataSrc)
	fmt.Println(str, "end")
	var data = bilit.Pull(template, str)
	fmt.Println(data)
	// bts, _ := json.Marshal(data)
	// fmt.Println(string(bts))
}
