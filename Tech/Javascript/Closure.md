### Closure

An enclosing scope available to the function that contains the functions and variables that are in scope at the time the function is declared.  
It survives beyond the life of the original scope at the point the function is declared.

```javascript
var outerValue = "ninja";
var later;

function outer() {
  var innerValue = "samurai";
  
  function inner() {
    // outerValue and innerValue are 'hoisted' in this function
    // 
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

### Immediate function 

```javascript
var keepcount = (function() {
  var private = 0;
  
  var increment = function() {
    return ++private;
  }
  
  return {
    increment: increment
  };
})()

keepcount.increment();
> 1
keepcount.increment();
> 2

// cannot directly access private.
```

jQuery plugins can also use this.

```javascript
var plugin = (function($) {
  // $ doesn't need to interfere with global scope
  return $('.plugin').html();
})(jQuery)
```

Javascript is statically scoped.

```
myVar=0;

function runMe(){
    var myVar = 10;
    callMe();
}

function callMe(){
   addMe = myVar+10;
}

runMe();
console.log(addMe);
> 10 // 0 + 10
console.log(myVar);
> 0 // unchanged
```

