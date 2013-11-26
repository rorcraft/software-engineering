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

### Including Module

similar to inheritance, using 'Include Class', at the time of include.

```
Super
 ^
 |
Include Class Module1 -> Module1
 ^
 |
Class
```


```
module M
 def test; puts 'm'; end
end
module M2
 def test; puts 'm2'; end
end
class C
  include M   # M2 is not yet included in M
end           # therefore M2 is not in C's superclasses

>C.new.test
'm'

module M
  include M2  # as there M2 is included in M,
end

>C.new.test
'm'

class C
  include M
end

>C.new.test
'm2'

```

