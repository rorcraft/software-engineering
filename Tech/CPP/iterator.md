### Pre vs Post Increment `++it` vs `it++`

http://antonym.org/2008/05/stl-iterators-and-performance.html

```c
// prefix
T& operator++();
```

```c
// postfix
const T operator++();
```

> prefix operator returns a reference to the object itself (permitting chaining of method calls), while the postfix operator returns a const object, semantically defined as the previous value. (This is for consistency with the behaviour of ordinal types.)

> The result is that the pre-increment operator modifies the object in-place, whereas the post-increment operator will result in temporaries being created, invoking the constructor and destructor. So something as simple as ++it versus it++ turns out to have some significant side-effects when applied to an object with overloaded operators. 
