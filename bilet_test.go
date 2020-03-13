package bilit

import (
	"testing"
)

const (
	testTemplate = "Hello, I'm {{name}} from {{City}}, {{from_state}}"
	testRegex    = "Hello, I'm (?P<name>.*) from (?P<City>.*), (?P<from_state>.*)"
)

var testData = map[string]string{
	"name":       "John",
	"City":       "Dallas",
	"from_state": "TX",
}

func TestPopulate(t *testing.T) {
	str := Populate(testTemplate, testData)
	if str != "Hello, I'm John from Dallas, TX" {
		t.Error(str)
	}
}

func ExamplePopulate() {
	str := Populate("Hello, I'm {{name}} from {{City}}, {{from_state}}", Data{
		"name":       "John",
		"City":       "Dallas",
		"from_state": "TX",
	})
	print(str)
	//Output: Hello, I'm John from Dallas, TX
}

func TestMakePullRegex(t *testing.T) {
	str := makePullRegex(testTemplate, ".+")
	if str != testRegex {
		t.Error(testTemplate)
		t.Error(str)
	}
}
