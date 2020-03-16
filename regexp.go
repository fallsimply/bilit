package bilit

import "regexp"

var (
	// GroupRegexp extracts the name of a nunjucks style lookup
	GroupRegexp = regexp.MustCompile(
		`(?m)[${]{2}(?P<groupName>[^{}$]*)[}]{1,2}`,
	)

	// GroupEscape escapes capture groups - Before
	GroupEscape = regexp.MustCompile(
		`(?m)\((?P<Content>[^?].*)\)`,
	)

	// CharClassEscape escapes character classes - Before
	CharClassEscape = ``

	// QuantityEscape escape quantity specifiers - After
	QuantityEscape = `{(?P<Content>[^0-9].*)}`
)

func pullRegex(str string, pattern string) string {
	var substitution = "(?P<${groupName}>" + pattern + ")"

	str = GroupEscape.ReplaceAllString(str, `\(${Content}\)`)

	return GroupRegexp.ReplaceAllString(str, substitution)
}
