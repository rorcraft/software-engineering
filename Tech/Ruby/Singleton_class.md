http://ruby-hacking-guide.github.io/class.html

* all Data are objects
* classes are data

```
Super Class (Parent) -> Super Class's Singleton Class
                        |- singleton methods (Super Class methods inheritable)
  |- instance methods inheritable
  ^
  |
Class -> Singleton Class
           |- singleton methods (Class methods overriding super class's class methods (singleton methods))
  | 
  |- instance methods overriding super
  ^
  |
Singleton Class (not inheritable, only for the instance)
  |- singleton methods
  ^
  |
Instance
```

