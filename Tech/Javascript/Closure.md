### Closure

An enclosing scope available to the function that contains the functions and variables that are in scope at the time the function is declared.  
It survives beyond the life of the original scope at the point the function is declared.

```javascript
var outerValue = "ninja";
var later;

function outer() {
  var innerValue = "samurai";
  
  function inner() {
    console.log(outerValue);
    console.log(innerValue);
  }
  later = inner;
}
outer();
later();
// ninja
// samurai
```
