package bilit

import (
	"regexp"
)

var (
	//GroupRegexp extracts the name of a nunjucks style lookup
	GroupRegexp = regexp.MustCompile(`(?m)[${]{2}(?P<groupName>[^{}]*)[}]{1,2}`)
)

//Data is a data map for bilit
type Data map[string]string

//Tmpl is a bilit template
type Tmpl struct {
	template string
	pattern  string
	regex    *regexp.Regexp
}

//RegEx makes the pull regular expression
func (t Tmpl) RegEx() {
	t.regex = regexp.MustCompile(makePullRegex(t.template))
}

//Template makes a template
func Template(str string) Tmpl {
	var obj = Tmpl{
		template: str,
	}
	obj.RegEx()
	return obj
}

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
