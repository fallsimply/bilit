package bilit

import (
	"fmt"
	"reflect"
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

func BenchmarkPopulate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testTemplate.Populate(testData)
	}
}

func BenchmarkPull(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testTemplate.Pull("Hello, I'm John from Dallas, TX")
	}
}

func BenchmarkTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Template(testTmplStr)
	}
}

func TestPopulate(t *testing.T) {
	str := testTemplate.Populate(testData)
	if str != "Hello, I'm John from Dallas, TX" {
		t.Error(str)
	}
}

func TestFunctionalPopulate(t *testing.T) {
	str := Populate(testTmplStr, testData)
	if str != "Hello, I'm John from Dallas, TX" {
		t.Error(str)
	}
}

func TestPull(t *testing.T) {
	data := testTemplate.Pull("Hello, I'm John from Dallas, TX")
	if !reflect.DeepEqual(data, testData) {
		t.Error(data)
	}
}

func TestFunctionalPull(t *testing.T) {
	data := Pull(testTmplStr, "Hello, I'm John from Dallas, TX")
	if !reflect.DeepEqual(data, testData) {
		t.Error(data)
	}
}

func TestMakePullRegex(t *testing.T) {
	str := pullRegex(testTemplate.template, ".+")
	if str != testRegex {
		t.Error("Hardcoded:", testRegex)
		t.Error("Generated:", str)
	}
}
func ExampleTmpl_Populate() {
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

func ExampleTmpl_Pull() {
	Debug = true
	tmpl := Template("Hello, I'm {{name}} from ${City}, {{from_state}}")
	str := tmpl.Pull("Hello, I'm John from Dallas, TX")
	fmt.Print(str)
	// Output: map[City:Dallas from_state:TX name:John]
}
