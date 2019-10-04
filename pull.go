package bilit

//Try it on Go Playground - https://play.golang.org/p/qp5QvkZ-RRZ
//Special Thanks to Sundrique on Stack Overflow - https://stackoverflow.com/q/29937787/

import (
	"fmt"
)

func makePullRegex(str string, pattern ...string) string {
	if len(pattern) == 0 {
		pattern = append(pattern, ".+")
	}
	var substitution = "(?P<${groupName}>" + pattern[0] + ")"

	return GroupRegexp.ReplaceAllString(str, substitution)
}

//Pull extracts data from a populated string using a template
func (t Tmpl) Pull(str string) map[string]string {
	fmt.Println(t.regex)
	matches := map[string]string{}
	for _, line := range t.RegEx().FindAllStringSubmatch(str, -1) {
		println(line)
		for j := range line {
			n := 1
			if j != 0 {
				n = j
			}
			name := t.RegEx().SubexpNames()[n]
			// Debug Statement
			// println(name + ": " + line[n])
			matches[name] = string(line[n])
		}
	}
	return matches
}

//Pull for old api compatibility
func Pull(template, str string) map[string]string {
	return Template(template).Pull(str)
}
