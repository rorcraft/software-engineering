* https://github.com/cujojs/when/blob/master/docs/api.md
* http://www.html5rocks.com/en/tutorials/es6/promises/

```javascript
var when = require('when');
function inc(x) {
  console.log("inc: " + x);
  return x + 1;
}
function setval(x) {
  val = x;
  console.log("setval: "+ x);
}
function log() {
  console.log("val: " + val);
}

var d = when.defer();
var p = d.promise;
var d2 = when.defer();
var p2 = d2.promise;
var val;

p.then(inc).then(inc).then(setval).then(log);
p2.then(inc).then(inc).then(setval).then(log);

// d.resolve(1);
// d.resolve(0);// wont do anything, already resolved.
/*
  inc: 1
  inc: 2
  setval: 3
  val: 3
*/

/*********************************************/
// d.resolve(1);
// d2.resolve(0);
/*
  inc: 1 // d.then
  inc: 0 // d2.then
  inc: 2 // d.then.then
  inc: 1 // d2.then.then
  setval: 3
  setval: 2
  val: 2
  val: 2
*/
/*********************************************/
// d.resolve(1);
// var r = when.resolve(0);
// r.then(inc).then(inc).then(setval).then(log);
/* same as d.resolve(1); d2.resolve(0);
  inc: 1
  inc: 0
  inc: 2
  inc: 1
  setval: 3
  setval: 2
  val: 2
  val: 2
*/
/********************************************/
// var r = when.resolve(0);
// r.then(inc).then(inc).then(setval).then(log);
// r.then(inc).then(inc).then(setval).then(log);
/* // doubles the work
  inc: 0
  inc: 0
  inc: 1
  inc: 1
  setval: 2
  setval: 2
  val: 2
  val: 2
*/
/*********************************************/
```
