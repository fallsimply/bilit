package test

const (
	template = "Hello, I'm {{name}} from {{City}}, {{from_state}}"
)

var exdata = map[string]string{
	"name":       "John",
	"City":       "Dallas",
	"from_state": "TX",
}

func main() {
	fmt.Println(Populate(template, exdata))
}

