> Go does not support function overloading and does not support user defined operators.

https://stackoverflow.com/questions/2032149/optional-parameters

Best you can do is...
```golang
type Params struct {
  a, b, c int
}

func doIt(p Params) int {
  return p.a + p.b + p.c 
}

// you can call it without specifying all parameters
doIt(Params{a: 1, c: 9})
```
or
```
func doIt(params ...int) {} //variable params
```

> In Go, it is not possible to create methods on Interfaces.
http://golang.org/doc/effective_go.html#allocation_new

```
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

> Type system

https://functionwhatwhat.com/go%E2%80%99s-type-system-is-an-embarrassment/
