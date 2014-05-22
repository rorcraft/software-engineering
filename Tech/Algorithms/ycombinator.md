```javascript

var factorial = function ycombinator(fn) {
  // need to simlate this fn(fn(...n times))
  return function (gen) {
    // pass fac partial back into itself
    // gen() return fn
    return gen(gen)
  }(function (gen) {
    // create the fac partial
    return fn(function (x) {
      // pass the fac partial back into itself
      // fac partial call be called infinitely because
      // n - 1 will become <= 2
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
