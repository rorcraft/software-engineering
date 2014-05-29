## Algebraic Javascript Spec

https://github.com/fantasyland/fantasy-land

https://github.com/fantasyland/fantasy-land/raw/master/figures/dependencies.png

### Semigroup


* Associativity `a.concat(b).concat(c)` is equivalent to `a.concat(b.concat(c))`

### Functor

* things that can be mapped over.

__Identity__: Mapping an identity function over a Functor is the same as applying identity to the Functor itself
e.g.
```javascript
var identity = function(x) { return x; }
fmap (identify, functor) === identity(functor)
```
__Composition: If you compose the functions f and g and map the resulting function over the Functor, that is the same as first mapping g over the Functor and then mapping f over the resulting Functor.
```
fmap( compose(f, g), functor ) === fmap( f, fmap( g, functor )
```

### Applicative Functor
