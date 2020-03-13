# Bilit
[![Try it on Go Playground](https://img.shields.io/static/v1?label=Go%20Playground&message=Try%20It&color=00addc&style=for-the-badge&logo=go)]()  

Bilit (pronounced /ˈbilit/) is a bidirectional template library for golang.
Inspired by Javascript's Template Literals and [Nunjucks](https://mozilla.github.io/nunjucks/templating.html)

Powered by Regular Expressions

## Examples
### OLD API - Basic:
``` go
func main() {
	const template = "Hello, I'm {{name}} from ${City}, {{from_state}}"
	var dataSrc = map[string]string{
		"name": "John",
		"City": "Dallas",
		"from_state": "TX",
	}

	var str = bilit.Populate(template, dataSrc)
	fmt.Println(str)
	// "Hello, I'm John from Dallas, TX"
	var data = bilit.Pull(template, str)
	fmt.Println(data)
	// map[City:Dallas from_state:TX name:John]
}
```
Open [Example Folder](./example)

## Special Thanks
