package bilit

import (
	"regexp"
)

var (
	// DefaultPattern is the default pattern for the regex
	DefaultPattern = ".+"

	// Unsafe disables RegExp escaping
	Unsafe = false
)

// Tmpl is a bilit template
type Tmpl struct {
	template string
	pattern  string
	regex    *regexp.Regexp
}

// RegEx makes the pull regular expression
func (t Tmpl) RegEx() *regexp.Regexp {
	t.regex = regexp.MustCompile(pullRegex(t.template, t.pattern))
	return t.regex
}

// Template makes a template
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
