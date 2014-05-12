http://www.ebooks-it.net/ebook/javascript-the-good-parts

__Semicolon insertion__

returned before the object.
```
return 
{
  status: true
};
```

__Unicode__

JavaScript’s characters are 16 bits

__typeof__

`typeof null === 'object'`

__parseInt__

```
parseInt("16") === parseInt("16 tons") === 16
```

If the first character of the string is 0, then the string is evaluated in base 8 instead of base 10.
```
parseInt("08") === 0
parseInt("08", 10) === 8
```

__floating point__

```
0.1 + 0.2 !== 0.3
```

__NaN is Number__

```
typeof NaN === 'number'
NaN !== NaN

function isNumber(value) { 
  return typeof value === 'number' && isFinite(value);
}
```

__Falsy Values__

```
0
NaN
''
false
null
undefined
```

__==__

```
0 == '' // true
0 == '0' // true
false == 'false' // false
false == '0' // true
false == undefined // false
null == undefined // true
```
