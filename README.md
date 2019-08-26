# Bilit
Bilit (pronounced /ˈbilit/) is a bidirectional template library for golang.
Inspired by Javascript's Template Literals and [Nunjucks](https://mozilla.github.io/nunjucks/templating.html)

Powered by Regular Expressions

### Examples
Simple (map):
``` go
func main() {
	const template = "Hello, I'm {{name}} from {{City}}, {{from_state}}"
	var dataSrc = map[string]string{
		"name": "John",
		"City": "Dallas",
		"from_state": "TX"
	}

	var str = bilit.Populate(template, dataSrc) //"Hello, I'm John from Dallas, TX"

	var data = billit.Pull(template, str) //{name: John, City: Dallas, from_state: TX}

}
```