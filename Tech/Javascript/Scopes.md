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
var assert = require('assert');
if(true) {
  var x = 123;
}
assert(x === 123); //x is defined
function something() {
  var y = 456;
  assert(y === 456);
}
something();
assert(typeof y === 'undefined');
```

### Invocations

```javascript
var assert = require('assert');
/**
 * `this` and `arguments` are implicitly passed parameters to the function.
 * `this` is different based on the way the function is invoked.
 */
var x = 123;
function outer() {
  console.log(x);
  console.log(this.x);
}
var obj = {
  x: "x of obj",
  outer: outer
}
/**
 * Invocate as function, `this` is the local context
 */
outer();
// 123
// // undefined

/**
 * Invocate as method on object, `this` is the object
 */
obj.outer();
// 123
// x of obj

/**
 * Invocate as constructor by calling `new` <function>()
 * - `this` is a new empty object, returned if no explicit return defined.
 */
function ObjClass() {
  this.x = "instance of ObjClass";
  this.outer = outer;
}
var objInstance = new ObjClass();
objInstance.outer();
// 123
// instance of ObjClass

/**
 * call - takes (context, arg1, arg2, arg3)
 * apply - takes (context, args[])
 */
outer.call(obj, null);
// 123
// x of obj
outer.call(this, null);
// 123
// undefined
```

### Caveats of using `var`

```javascript
var assert = require('assert');

var object = {
  x: 123,
  puts: function() {
    console.log(this.x);
  }
}
object.puts();
// 123

function wrappedBindedPuts() {
  var puts = object.puts.bind(this);
  puts();
}
wrappedBindedPuts.call(object);
// 123

function thisPuts() {
  this.puts();
}
thisPuts.call(object);
// 123

function wrappedPuts() {
  var puts = object.puts; // var puts is in a different scope not attached to `this`
  assert(this.x === 123);
  puts();
}
wrappedPuts.call(object);
// undefined
```
