![Bilit Logo](./assets/readme.svg)

[![MIT Licensed][License Badge]][LICENSE] [![Latest Version][Release Badge]][Latest Version]
[![pkg.go.dev docs][Docs Badge]][Doc Site]

Bilit (pronounced /ˈbilit/) is a bidirectional template library for Go.
Syntax inspired by Javascript's Template Literals and [Nunjucks][Nujucks Site]

Powered by Regular Expressions

## Examples
### Modern API - Basic: 
[![Open Example Folder][Open Badge]](./examples/modern) [![Try it on Go Playground][Play Badge]][Modern API]
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
[![Open Example Folder][Open Badge]](./examples/function) [![Try it on Go Playground][Play Badge]][Function API]
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

[Play Badge]: https://img.shields.io/static/v1?label=Go%20Playground&message=Try%20It&style=for-the-badge&logo=go&color=4E98A6&labelColor=333A47
[License Badge]: https://img.shields.io/github/license/fallsimply/bilit?style=for-the-badge&color=7AA64E&labelColor=333A47
[Release Badge]: https://img.shields.io/github/v/release/fallsimply/bilit?style=for-the-badge&color=4E5CA6&labelColor=333A47
[Open Badge]: https://img.shields.io/static/v1?label=Open&message=Example%20Folder&style=for-the-badge&color=7AA64E&labelColor=333A47
[Docs Badge]: https://img.shields.io/static/v1?label=pkg.go.dev&message=Docs&style=for-the-badge&logo=go&color=4E98A6&labelColor=333A47

[Modern API]: https://play.golang.org
[Function API]: https://play.golang.org/p/n9MtRW_RIlg

[Latest Version]: https://github.com/fallsimply/bilit/releases/latest
[Doc Site]: https://pkg.go.dev/github.com/fallsimply/bilit
[License]: ./LICENSE
[Nujucks Site]: https://mozilla.github.io/nunjucks/templating.html