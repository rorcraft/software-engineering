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
