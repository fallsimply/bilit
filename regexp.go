package bilit

import (
	"regexp"
)

var (
	// GroupRegexp extracts the name of a nunjucks style lookup
	GroupRegexp = regexp.MustCompile(
		`(?m)[${]{2}(?P<groupName>[^{}$]*)[}]{1,2}`,
	)

	// BeforeEscape escapes characters - Before Groups
	BeforeEscape = regexp.MustCompile(
		`(?m)[\[\]().+?^]`,
	)

	// QuantityEscape escape quantity specifiers - After Groups
	QuantityEscape = regexp.MustCompile(
		`{(.*)}`,
	)
	// AfterEscape escapes characters - After Groups
	AfterEscape = regexp.MustCompile(`[$]`)
)

func pullRegex(str string, pattern string) string {
	var substitution = "(?P<${groupName}>" + pattern + ")"

	if !Unsafe {
		str = BeforeEscape.ReplaceAllString(str, `\$0`)
	}

	str = GroupRegexp.ReplaceAllString(str, substitution)

	if !Unsafe {
		// Quantity is after group replacement to avoid collision with template groups
		str = QuantityEscape.ReplaceAllString(str, `\{$1\}`)
		str = AfterEscape.ReplaceAllString(str, `\$0`)
	}

	return str
}
