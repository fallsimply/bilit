package bilit

import (
	"regexp"
)

var (
	//GroupRegexp extracts the name of a nunjucks style lookup
	GroupRegexp = regexp.MustCompile(`(?m){{(?P<groupName>[^{}]*)}}`)
)

//Data is a data map for bilit
type Data = map[string]string

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
