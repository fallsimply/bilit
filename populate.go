package bilit

import (
	"fmt"
)

func replacer(str []byte, data map[string]string) []byte {
	var s = GroupRegexp.ReplaceAll(str, []byte("${groupName}"))
	fmt.Println(data[string(s)])
	return []byte(data[string(s)])
}

//Populate populates a template with data
func Populate(template string, data map[string]string) string {
	return string(GroupRegexp.ReplaceAllFunc([]byte(template), func(s []byte) []byte {
		return replacer(s, data)
	}))
}
