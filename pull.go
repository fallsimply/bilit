package bilit

import (
	"fmt"
	"regexp"
)

func makePullRegex(str string, pattern ...string) string {
	if len(pattern) == 0 {
		pattern = append(pattern, ".+")
	}
	var substitution = "(?P<${groupName}>" + pattern[0] + ")"

	return GroupRegexp.ReplaceAllString(str, substitution)
}

//Pull extracts data from a populated string using a template
func Pull(template, str string) map[string]string {
	//Make Regular Expression from template
	var regex = regexp.MustCompile(makePullRegex(template))

	fmt.Println(regex)
	matches := map[string]string{}
	for _, line := range regex.FindAllStringSubmatch(str, -1) {
		for j := range line {
			n := 1
			if j != 0 {
				n = j
			}
			name := regex.SubexpNames()[n]
			// Debug Statement
			// println(name + ": " + line[n])
			matches[name] = string(line[n])
		}
	}
	return matches
}

//Special Thanks to Sundrique on Stack Overflow - https://stackoverflow.com/q/29937787/
