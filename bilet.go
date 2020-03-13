package bilit

import (
	"regexp"
)

var (
	//GroupRegexp extracts the name of a nunjucks style lookup
	GroupRegexp = regexp.MustCompile(`(?m)[${]{2}(?P<groupName>[^{}$]*)[}]{1,2}`)
	//DefaultPattern is the default pattern for the regex
	DefaultPattern = ".+"
	//Debug enables printing of the map passed to a template
	Debug = false
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
func (t Tmpl) RegEx() *regexp.Regexp {
	t.regex = regexp.MustCompile(makePullRegex(t.template, t.pattern))
	return t.regex
}

//Template makes a template
func Template(str string, pattern ...string) Tmpl {
	if len(pattern) == 0 {
		pattern = append(pattern, DefaultPattern)
	}
	var obj = Tmpl{
		template: str,
		pattern:  pattern[0],
	}
	obj.RegEx()
	return obj
}
