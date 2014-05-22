```javascript

var factorial = function ycombinator(fn) {
  // return fn(fn(fn(...)))
  return function (gen) {
    return gen(gen)
  }(function (gen) {
    return fn(function (x) {
      return gen(gen)(x)
    })
  })
}(function (fac) {
  return function (n) {
    return n <= 2 ? n : n * fac(n - 1);
  }
})

factorial(5)
> 120
```
