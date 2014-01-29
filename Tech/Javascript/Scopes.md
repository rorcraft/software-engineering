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
