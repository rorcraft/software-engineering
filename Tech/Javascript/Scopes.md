### Forward-referenceable within the same scope.

```javascript
var assert = require('assert');
assert(typeof futureFunction === "function", "futureFunction() is undefined");
function futureFunction() {
  assert(typeof inner === "function", "inner() is undefined");
  function inner() {};
};
futureFunction();
```
