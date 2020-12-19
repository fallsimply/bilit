package bilit

func replacer(str []byte, data map[string]string) []byte {
	var s = string(GroupRegexp.ReplaceAll(str, []byte("${groupName}")))

	return []byte(data[s])
}

// Populate populates a template with data
func (t Tmpl) Populate(data map[string]string) string {
	return string(GroupRegexp.ReplaceAllFunc([]byte(t.template), func(s []byte) []byte {
		return replacer(s, data)
	}))
}

// Populate for old api compatibility
func Populate(template string, data map[string]string) string {
	return Template(template).Populate(data)
}
