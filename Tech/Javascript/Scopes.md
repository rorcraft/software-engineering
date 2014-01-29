### Forward-referenceable within the same scope.

```javascript
var assert = require('assert');
assert(typeof futureFunction === "function"); // futureFunction is defined
function futureFunction() {
  assert(typeof inner === "function"); // inner is defined
  function inner() {};
};
futureFunction();
assert(typeof inner === "undefined"); // inner is undefined
```
### Scopes are declared by functions, and not by blocks.

```javascript
if(true) {
  var x = 123;
}
assert(x === 123); //x is defined
//---------------------
function outer() {
  console.log(x);
  console.log(this.x);
}
var obj = {
  x: "x of obj",
  outer: outer
}
obj.outer();
// 123
// x of obj
outer();
// 123
// undefined
```

### Invocations

* As a function, in which the function is invoked in a straightforward manner
* As a method, which ties the invocation to an object, enabling object-oriented programming
* As a constructor, in which a new object is brought into being
* Via its apply() or call() methods, which is kind of complicated, so we’ll cover that when we get to it
