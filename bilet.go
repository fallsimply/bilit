package bilit

import (
	"regexp"
)

var (
	//GroupRegexp extracts the name of a nunjucks style lookup
	GroupRegexp = regexp.MustCompile(`(?m){{(?P<groupName>[^{}]*)}}`)
)

// Subgroup Helper Function
// func makeSubgroupNameMap(re regexp.Regexp) map[int]string {
// 	result := make(map[int]string)
// 	for i, name := range re.SubexpNames() {
// 		if i != 0 && name != "" {
// 			result[i] = name
// 		}
// 	}
// 	return result
// }

func main() {
	str := Populate("Hello, I'm {{name}} from {{City}}, {{from_state}}", map[string]string{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	})
	print(str)
}
