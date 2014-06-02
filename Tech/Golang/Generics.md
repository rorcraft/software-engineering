Using `reflect`

http://play.golang.org/p/WGDOX4_csT

```go
package main

import "reflect"
import "fmt"

func Filter(xs interface{}, fn func(i interface{}) bool) interface{} {
	refValue := reflect.ValueOf(xs)
	refType := reflect.TypeOf(xs)

	if !(refValue.Kind() == reflect.Slice ||
		refValue.Kind() == reflect.Array) {

		panic("Not an iterable")
	}

	var l = refValue.Len()

	var _xs = reflect.MakeSlice(refType,0,0)

	var val reflect.Value

	for i := 0; i < l; i++ {
		val = refValue.Index(i)
		if fn(val.Interface()) {
			_xs = reflect.Append(_xs, val)
		}

	}
	return _xs.Interface()
}
```
