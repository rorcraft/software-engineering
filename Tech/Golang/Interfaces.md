http://research.swtch.com/interfaces

### Duck typing

Interfaces aren't restricted to static checking, though. You can check dynamically whether a particular interface value has an additional method. For example:

```golang
type Stringer interface {
    String() string
}

func ToString(any interface{}) string {
    if v, ok := any.(Stringer); ok {
        return v.String()
    }
    switch v := any.(type) {
    case int:
        return strconv.Itoa(v)
    case float:
        return strconv.Ftoa(v, 'g', -1)
    }
    return "???"
}

// A value of type Binary can be passed to ToString,
// which will format it using the String method, even though the program never says
// that Binary intends to implement Stringer.
// There's no need: the runtime can see that Binary has a String method,
// so it implements Stringer, even if the author of Binary has never heard of Stringer.
type Binary uint64

func (i Binary) String() string {
    return strconv.Uitob64(i.Get(), 2)
}

func (i Binary) Get() uint64 {
    return uint64(i)
}
```
> Go has method tables but computes them at run time

http://golang.org/src/pkg/runtime/iface.goc

* Interface values are represented as a two-word pair giving a pointer to information about the type stored in the interface and a pointer to the associated data.

__Method Lookup Performance__

Smalltalk and the many dynamic systems that have followed it perform a method lookup every time a method gets called. For speed, many implementations use a simple one-entry cache at each call site, often in the instruction stream itself. In a multithreaded program, these caches must be managed carefully, since multiple threads could be at the same call site simultaneously. Even once the races have been avoided, the caches would end up being a source of memory contention.

Because Go has the hint of static typing to go along with the dynamic method lookups, it can move the lookups back from the call sites to the point when the value is stored in the interface. For example, consider this code snippet:

```
1   var any interface{}  // initialized elsewhere
2   s := any.(Stringer)  // dynamic conversion
3   for i := 0; i < 100; i++ {
4       fmt.Println(s.String())
5   }
```

In Go, the itable gets computed (or found in a cache) during the assignment on line 2; the dispatch for the s.String() call executed on line 4 is a couple of memory fetches and a single indirect call instruction.

In contrast, the implementation of this program in a dynamic language like Smalltalk (or JavaScript, or Python, or ...) would do the method lookup at line 4, which in a loop repeats needless work. The cache mentioned earlier makes this less expensive than it might be, but it's still more expensive than a single indirect call instruction.

### Examples

http://www.golang-book.com/9/index.htm

```golang
type Rectangle struct {
  x1, y1, x2, y2 float64
}
func (r *Rectangle) area() float64 {
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return l * w
}
type Circle struct {
    x, y, r float64
}
func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}
// the interface
type Shape interface {
  area() float64
}
type MultiShape struct {
    shapes []Shape
}
func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area
}
```
