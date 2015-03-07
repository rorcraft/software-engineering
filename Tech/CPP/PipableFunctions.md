### Pipable functions 

Overload `operator|` 
* http://pfultz2.com/blog/2014/09/05/pipable-functions/

https://github.com/facebook/folly/tree/master/folly/gen
```cpp
    string& longestName = from(names)
                       | maxBy([](string& s) { return s.size() });
```
