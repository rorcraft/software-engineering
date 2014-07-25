http://blog.golang.org/slices

Slice is like a struct, pointing to underlying array.
```
type sliceHeader struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}

slice := sliceHeader{
    Length:        0,
    Capacity:      10,
    ZerothElement: &iBuffer[0],
}
```

Passing slice to function, it makes a copy of the slice headers. So to modify a slice, pass a pointer to sliceHeader to the function.

```
func PtrSubtractOneFromLength(slicePtr *[]byte) {
    slice := *slicePtr
    *slicePtr = slice[0 : len(slice)-1]
}

func main() {
    fmt.Println("Before: len(slice) =", len(slice))
    PtrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}
```

Increase capacity of a slice.
```
// if capacity allows
slice = slice[:len(slice)+1]
// otherwise
panic: runtime error: slice bounds out of range

// use copy
newSlice := make([]int, len(slice), 2*cap(slice))
copy(newSlice, slice)
```

Handy syntax to slice an array
```
array[:]
```

__Append__: an example implementation
```go
func Extend(slice []int, element int) []int {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]int, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}
// we can append a whole second slice by "exploding" the slice into arguments using the ... notation at the call site:
slice1 = Append(slice1, slice2...)
```
_built-in_ function does exactly what our Append example does, with equivalent efficiency, but it works for any slice type.

__Nil__

an empty slice can grow (assuming it has non-zero capacity), but a nil slice has no array to put values in and can never grow to hold even one element.

That said, a nil slice is functionally equivalent to a zero-length slice, even though it points to nothing. It has length zero and can be appended to, with allocation.

__Strings__  are actually very simple: they are just read-only slices of bytes with a bit of extra syntactic support from the language.

We can also take a normal slice of bytes and create a string from it with the simple conversion:

```
str := string(slice)
```
and go in the reverse direction as well:

```
slice := []byte(usr)
```
An important consequence of this slice-like design for strings is that creating a substring is very efficient.
