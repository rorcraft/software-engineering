https://github.com/loop-recur/drbooleanchallenges

```javascript
Boolean(new Boolean(false))
=> true

Boolean(!!null)
=> false

if Array(2,3,4).toString() == '2,3,4'

Boolean(Array(2) == ',')
=> true

Boolean(eval(eval(eval)) === eval)
=> true

Boolean({a:'a'}.toString === '[object Object]')
=> true

Boolean({a:'a',a:'b',a:'c'}['a'] === 'c')
=> true

Boolean(([][0]+'')['0'] === 'u')
=> true

Boolean((![][0]+'')['0'] === 't')
=> true
```
