### goimjson
a Go Package for maintaining immutable json object

The immutable property means that every change to the object will result in a version number.
As long as you can hold a version number, you can always retrieve the object back even though the object is mutated by someone else

### Importing
```
    import github.com/hsinhoyeh/goimjson
```
### Usage (as a library)
```
package main

import "github.com/hsinhoyeh/goimjson"

func main() {

    start := []byte("{}") // {}
    imj, _ := goimjson.NewWithBody(start)
    ver1 := imj.Set("field1", "value1") // +{"field1": "value1"}
    ver2 := imj.Set("field1", "value2") // +{"field1": "value2"}
    imj.Encode() // should be {"field1":"value2"}
    imj.Get(ver1, "field1") // will get "value1"
    imj.Get(ver2, "field2") // will get "value2"
}
```
see testcase for more.


### Usage (as a service)
```
cd github.com/hsinhoyeh/goimjson/http/server
go run server.go

// then use post to write data
curl -X POST -H "Content-Type: application/json" http://localhost:9000 -d '{"bar":"bar","foo":"foo"}'

// the response will be a json with version field
{
  "ver": "4592134183702187642"
}

// then use get to retrieve data
curl -X GET http://localhost:9000/4592134183702187642
{
  "bar": "bar",
  "foo": "foo"
}

```