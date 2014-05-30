http://howtonode.org/understanding-process-next-tick

Event loop a queue of callbacks that are processed by Node on every tick of the event loop.

```javascript
function foo() {
    console.error('foo');
}

process.nextTick(foo);
console.error('bar');
---
bar
foo
---
setTimeout(foo, 0);
console.log('bar');
---
bar
foo
```

`process.nextTick()` defers the function until a completely new stack. 
* Far more performant than `setTimeout`. https://gist.github.com/mmalecki/1257394


Enforce callback is called in the next tick.
```javascript
function asyncReal(data, callback) {
    process.nextTick(function() {
        callback(data === 'foo');       
    });
}
```

```javascript
function StreamLibrary(resourceName) {      
    var self = this;

    process.nextTick(function() {
        self.emit('start');
    });

    // read from the file, and for every chunk read, do:        
    this.emit('data', chunkRead);       
}

var stream = new StreamLibrary('fooResource');
// if StreamLibrary emitted 'start' synchronously, 'start' will never be heard.
stream.on('start', function() {
    console.log('Reading has started');
});
```
