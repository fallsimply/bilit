# Bilit
[![Try it on Go Playground](https://img.shields.io/static/v1?label=Go%20Playground&message=Try%20It&color=00addc&style=for-the-badge&logo=go)](https://play.golang.org/p/n9MtRW_RIlg) [![MIT Licensed](https://img.shields.io/github/license/simplycodin/bilit?style=for-the-badge)](./LICENSE)

Bilit (pronounced /ˈbilit/) is a bidirectional template library for Go.
Inspired by Javascript's Template Literals and [Nunjucks](https://mozilla.github.io/nunjucks/templating.html)

Powered by Regular Expressions

## Examples
### Modern API - Basic:
``` go
template := bilit.Template("Hello, I'm {{name}} from ${City}, {{from_state}}")

dataSrc := map[string]string{
	"name":       "John",
	"City":       "Dallas",
	"from_state": "TX",
}

str := template.Populate(dataSrc)
fmt.Println(str, "end")
// "Hello, I'm John from Dallas, TX"

data := template.Pull(tr)
fmt.Println(data)
// map[City:Dallas from_state:TX name:John]

```
### Function API (Old API) - Basic:
``` go
const template = "Hello, I'm {{name}} from ${City}, {{from_state}}"

dataSrc := map[string]string{
	"name":       "John",
	"City":       "Dallas",
	"from_state": "TX",
}

str := bilit.Populate(template, dataSrc)
fmt.Println(str, "end")
// "Hello, I'm John from Dallas, TX"

data := bilit.Pull(template, str)
fmt.Println(data)
// map[City:Dallas from_state:TX name:John]
```
Open [Function API Example](./examples/function)

## TODO
- [ ] Escape RegExp Symbols
	- [x] Parenthesis - `()`
	- [x] Curly Brackets - `{}`
	- [x] Square Brackes - `[]`
	- [ ] Period - `.`
	- [ ] Plus Sign - `+`
	- [ ] Question Mark - `?`
	- [ ] Dollar Sign - `$`
	- [ ] Caret - `^`