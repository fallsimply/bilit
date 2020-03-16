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
	CharClassEscape = regexp.MustCompile(
		`\[(?P<Content>[^?].*)\]`,
	)

	// QuantityEscape escape quantity specifiers - After
	QuantityEscape = regexp.MustCompile(
		`{(?P<Content>.*)}`,
	)
)

func pullRegex(str string, pattern string) string {
	var substitution = "(?P<${groupName}>" + pattern + ")"

	if !Unsafe {
		// Group & Character class are before group replacement to avoid collision with patterns
		str = GroupEscape.ReplaceAllString(str, `\(${Content}\)`)
		str = CharClassEscape.ReplaceAllString(str, `\[${Content}\]`)
	}

	str = GroupRegexp.ReplaceAllString(str, substitution)

	if !Unsafe {
		// Quantity is after group replacement to avoid collision with template groups
		str = QuantityEscape.ReplaceAllString(str, `\{${Content}\}`)
	}

	return str
}
