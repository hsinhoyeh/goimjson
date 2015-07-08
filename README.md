### goimjson
a Go Package for maintaining immutable json object

### Importing
```
    import github.com/hsinhoyeh/goimjson
```
### Usage
```
package main

import "github.com/hsinhoyeh/goimjson"

func main() {

    start := []byte("{}") // {}
    imj, _ := goimjson.NewWithBody(start)
    ver1 := imj.Set("field1", "value1") // +{"field1": "value1"}
    ver2 := imj.Set("field2", "value2") // +{                  "field2": "value2"}
    imj.Encode() // should be {"field1":"value1","field2":"value2"}
    imj.Get(ver1, "field1") // will get "value1"
    imj.Get(ver1, "field2") // will get an empty json
}
```
see testcase for more.