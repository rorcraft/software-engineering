## Algebraic Javascript Spec

https://github.com/fantasyland/fantasy-land

https://github.com/fantasyland/fantasy-land/raw/master/figures/dependencies.png

### Semigroup

must provide a `concat` method. 

* `a.concat(b).concat(c)` is equivalent to `a.concat(b.concat(c))` (associativity)

1. `b` must be a value of the same Semigroup.
1. `concat` must return a value of the same Semigroup.

### Monoid 

must provide an `empty` method on itself or its `constructor` object.  
must also implement Semigroup spec.

```
m.empty()
m.constructor.empty()
```

1. `m.concat(m.empty())` is equivalent to `m` (right identity)
2. `m.empty().concat(m)` is equivalent to `m` (left identity)

### Functor

must provide a `map` method. 

`u.map(f)`

1. `u.map(function(a) { return a; }))` is equivalent to `u` (identity)
2. `u.map(function(x) { return f(g(x)); })` is equivalent to `u.map(g).map(f)` (composition)

### Apply

must provide `ap` method.  
must also implement Functor spec

`a.ap(b)` - must apply the function in Apply `a` to value in Apply `b`

* `a.map(function(f) { return function(g) { return function(x) { return f(g(x)) }}}).ap(u).ap(v)` is equivalent to `a.ap(u.ap(v))` (composition)

### Applicative

must provide an `of` method on itself or its `constructor` object.  
must also implement Apply spec

`a.of(b)`
`a.constructor.of(b)`

1. `a.of(function (a) { return a }).ap(v) is equivalent to `v` (identity)
1. `a.of(f).ap(a.of(x))` is equivalent to `a.of(f(x))` (homomorphism)
1. `u.ap(a.of(y))` is equivalent to `a.of(function (f) { return f(y); }).ap(u)` (interchange)

### Chain

must provide a `chain` method  
must also implement Apply spec

`m.chain(f)`

1. Apply's `ap`; derivable as `m.chain(function (f) { return m.map(f) })`
2. `m.chain(f).chain(g)` is equivalent to `m.chain(function (x) { return f(x).chain(g) })` (associative)

### Monad

must also implement Applicative and Chain spec.

1. Apply's `ap`; derivable as `function (m) { return this.chain(function (f) { return m.map(f) }) }`
2. Functor's map; derivable as `function (f) { var m = this; return m.chain(function(a) { return m.of(f(a)); }) }`
3. `m.of(a).chain(f)` is equivalent to `f(a)` (left identity)
4. `m.chain(m.of) is equivalent to `m` (right identity)
