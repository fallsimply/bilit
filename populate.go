package bilit

func replacer(str []byte, data Data) []byte {
	var s = string(GroupRegexp.ReplaceAll(str, []byte("${groupName}")))
	if Debug {
		out.Printf("%s: %s\n", s, data[s])
	}
	return []byte(data[s])
}

//Populate populates a template with data
func (t Tmpl) Populate(data Data) string {
	return string(GroupRegexp.ReplaceAllFunc([]byte(t.template), func(s []byte) []byte {
		return replacer(s, data)
	}))
}

//Populate for old api compatibility
func Populate(template string, data map[string]string) string {
	return Template(template).Populate(data)
}
