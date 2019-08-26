package bilit

func makePullRegex(str string, pattern ...string) string {
	if pattern[0] == "" {
		pattern[0] = ".*"
	}
	var substitution = "(?P<${groupName}>" + pattern[0] + ")"

	return GroupRegexp.ReplaceAllString(str, substitution)
}
