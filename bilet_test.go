package bilit

import (
	"fmt"
	"testing"
)

const (
	testTmplStr = `Hello, I'm {{name}} from ${City}, {{from_state}}`
)

var (
	testTemplate = Template(testTmplStr)
	// %[1]s syntax = fmt explicit arguments
	testRegex = fmt.Sprintf("Hello, I'm (?P<name>%[1]s) from (?P<City>%[1]s), (?P<from_state>%[1]s)", DefaultPattern)
)

var testData = map[string]string{
	"name":       "John",
	"City":       "Dallas",
	"from_state": "TX",
}

func TestPopulate(t *testing.T) {
	Debug = true
	str := testTemplate.Populate(testData)
	if str != "Hello, I'm John from Dallas, TX" {
		t.Error(str)
	}
}

func ExamplePopulate() {
	Debug = true
	tmpl := Template("Hello, I'm {{name}} from ${City}, {{from_state}}")
	str := tmpl.Populate(map[string]string{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	})
	fmt.Print(str)
	// Output: Hello, I'm John from Dallas, TX
}

func TestMakePullRegex(t *testing.T) {
	Debug = true
	fmt.Println(testTemplate.template)
	str := pullRegex(testTemplate.template, ".+")
	if str != testRegex {
		t.Error("Hardcoded:", testRegex)
		t.Error("Generated:", str)
	}
}

func TestFunctionalPopulate(t *testing.T) {
	Debug = true
	str := Populate(testTmplStr, testData)
	if str != "Hello, I'm John from Dallas, TX" {
		t.Error(str)
	}
}
