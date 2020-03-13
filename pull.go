package bilit

import (
	"fmt"
)

func makePullRegex(str string, pattern string) string {
	var substitution = "(?P<${groupName}>" + pattern + ")"
	return GroupRegexp.ReplaceAllString(str, substitution)
}

//Pull extracts data from a populated string using a template
func (t Tmpl) Pull(str string) map[string]string {
	if Debug {
		fmt.Println(t.regex)
	}
	matches := map[string]string{}
	for _, line := range t.RegEx().FindAllStringSubmatch(str, -1) {
		if Debug {
			println(line)
		}
		for n := 1; n < len(line); n++ {
			name := t.RegEx().SubexpNames()[n]
			if Debug {
				println(name + ": " + line[n])
			}
			matches[name] = string(line[n])
		}
	}
	return matches
}

//Pull for old api compatibility
func Pull(template, str string) map[string]string {
	return Template(template).Pull(str)
}
