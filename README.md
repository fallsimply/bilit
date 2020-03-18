![Bilit Logo](./assets/Bilit%20(README).svg)

[![MIT Licensed][License Badge]](./LICENSE) [![Latest Version][Latest Badge]](https://github.com/simplycodin/bilit/releases/latest)
[![pkg.go.dev docs][Docs Badge]](https://pkg.go.dev/github.com/simplycodin/bilit)

Bilit (pronounced /ˈbilit/) is a bidirectional template library for Go.
Inspired by Javascript's Template Literals and [Nunjucks](https://mozilla.github.io/nunjucks/templating.html)

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
[![Open Example Folder][Open Badge]](./examples/modern) [![Try it on Go Playground][Play Badge]][Function API]
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

[Play Badge]: https://img.shields.io/static/v1?label=Go%20Playground&message=Try%20It&color=00addc&style=for-the-badge&logo=go
[Modern API]: https://play.golang.org
[Function API]: https://play.golang.org/p/n9MtRW_RIlg
[License Badge]: https://img.shields.io/github/license/simplycodin/bilit?style=for-the-badge
[Latest Badge]: https://img.shields.io/github/v/release/simplycodin/bilit?style=for-the-badge
[Open Badge]: https://img.shields.io/static/v1?label=Open&message=Example%20Folder&color=green&style=for-the-badge
[Docs Badge]: https://img.shields.io/static/v1?label=pkg.go.dev&message=Docs&color=00addc&style=for-the-badge&logo=go