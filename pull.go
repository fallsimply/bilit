package bilit

func makePullRegex(str string, pattern ...string) string {
	if len(pattern) == 0 {
		pattern = append(pattern, ".*")
	}
	var substitution = "(?P<${groupName}>" + pattern[0] + ")"

	return GroupRegexp.ReplaceAllString(str, substitution)
}
