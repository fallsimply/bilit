# Bilit Design Doc
## Pull
Template
```
Hello, I'm {{name}} from {{City}}, {{from_state}}
```
to Regex using `(?m){{(?P<groupName>[^{}]*)}}`
``` regexp
Hello, I'm (?P<name>.*) from (?P<City>.*), (?P<from_state>.*)
```
Run on
```
Hello, I'm John from Dallas, TX
```
To Get
``` js
{
	name: "John",
	City: "Dallas",
	from_state: "TX"
}
```
[Template to Regex - Regex101](https://regex101.com/r/faPa7N/2)<br>
[String to Data - Regex101](https://regex101.com/r/YPDAUW/2)
## Populate
Template and
``` 
Hello, I'm {{name}} from {{City}}, {{from_state}}
```
Data substituted into
``` js
{
	name: "John",
	City: "Dallas",
	from_state: "TX"
}
```
substituted into String
```
Hello, I'm John from Dallas, TX
```