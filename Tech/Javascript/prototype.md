### Prototype

```javascript
function Ninja() {
  this.swung = false
}
Ninja.prototype.swung = true;

var ninja1 = new Ninja(); // you can still exptend the prototype after instantiation

Ninja.prototype.swingSword = function(){
    return this.swung;
};

ninja1.constructor
> Ninja
ninja1.prototype
> undefined
ninja1.swung 
> false // constructor instantiated attributes are looked up first
ninja1.swingSword()
> false
ninja1 instanceof Ninja
> true

var ninja2 = new ninja1.constructor();
ninja2 instanceof Ninja
> true
```

### Prototype Chain & Inheritance

ninja1 -> constructor.prototype -> (prototype.constructor).prototype ... etc

```javascript
function Male() {
  this.gender = 'Male';
}
function Female() {
  this.gender = 'Female';
}
function Ninja() {}

ninja = new Ninja();
Ninja.prototype = new Male(); // prototype.constructor = Male
ninja.gender
> undefined // cannot change super once instantiated.
ninja1 = new Ninja();
ninja1.gender
> 'Male'
Ninja.prototype = new Female();
ninja2 = new Ninja(0;
ninja2.gender
> 'Female'

'gender' in ninja // inherited property
> true
ninja.hasOwnProperty('gender') 
> false
ninja.gender = 'Secret'
ninja.hasOwnProperty('gender')
> true
```

`__proto__`

