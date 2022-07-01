# str: An experimental string package
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0) [![Go Reference](https://pkg.go.dev/badge/github.com/tarndt/str.svg)](https://pkg.go.dev/github.com/tarndt/str) [![Go Report Card](https://goreportcard.com/badge/github.com/tarndt/str)](https://goreportcard.com/report/github.com/tarndt/str)

This is an experimental string package that imagines strings in Go as a complex (non-primitive) type. The purpose of it's creation is to facilitate my own private experimentation and development of an experience report to evaluate if the value added by a package like this is significant enough to warrant the complexity of its existence. 

## Caveats

* Use at your own risk. No stability or defect guarantees are made and, unless this experiment proves fruitful
* This project may end up abandoned.

Those caveats aside, if you choose to use this, feedback and thoughts in the form of pull requests or issues are welcome.

## Using it

Declare a `Str` "string":
```go 
var s str.Str
s := str.Str("foobar")
fromInt := str.New(3)
fromStringer := str.New(myStringer)
fromAnything := str.New(someObject)
formatedStr := str.Newf("No need to call %s", "Sprintf")
```

Use any of the stdlib strings.* functions as methods:
```go
lowerAlpha := str.Str(" -abc ").TrimSpace().TrimSuffix("-")
upperAlpha := lowerAlpha.ToUpper()
```

Use it as a HTTP request body:
```go
resp, err: = http.Post("http://example.test.com", "text/plain", upperAlpha.Reader())
```

Or, get it in a buffer and append it to itself:
```go
buf := upperAlpha.Buffer()
_, err := upperAlpha.WriteTo(buf)
```

Or, start a list of `Strs` and play with it:
```go
slice := str.Slice(lowerAlpha, upperAlpha)
slice = slice.Append("somethingElse")
slice = slice.Append("x", "y").AppendAny(3.14159).Remove(0).RemoveStr("x")
slice.Sort()

if slice.Contains("x") {
	 panic("I removed x!")
}
```


## Future

* TODO: Further unit tests will be added if the experiment has initial positive returns.
* TODO: Better comments will be added if the experiment has initial positive returns.
