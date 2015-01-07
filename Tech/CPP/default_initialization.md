http://stackoverflow.com/questions/2417065/does-the-default-constructor-initialize-built-in-types

### Default initialization

Implicitly defined (by the compiler) default constructor of a class does not initialize members of built-in types

* POD - no user declared constructor
```
class C {
  int x;
};
```

```
C c; // Compiler-provided default constructor is used
// Here `c.x` contains garbage
```

* non-POD - with user declared constructor
```
class D {
 public:
  D() {
    x = 0;
  }
  int x;
}
```
```
D d;
d.x == 0
```

### Value initialization

* no user declared constructor
```
class C {
  int x;
}
```

```
C* c = new C();
c->x == 0; // default initialized x (int) to 0
```

```
C c{}; // c++11
c.x == 0;
```

http://stackoverflow.com/questions/1051379/is-there-a-difference-in-c-between-copy-initialization-and-direct-initializati

### Direct initialization
```
C c(x);
```
Direct initialization behaves like a function call to an overloaded function: The functions, in 
this case, are the constructors of C (including explicit ones), and the argument is x. Overload 
resolution will find the best matching constructor, and when needed will do any implicit conversion
required.

### Copy initialization
```
C c1 = x;
C c2 = C(); // copy initialize: value-initialized C() then copies that value to c;
```
Copy initialization constructs an implicit conversion sequence: It tries to convert x to an object
of type C. (It then may copy over that object into the to-initialized object, so a copy constructor
is needed too - but this is not important below)

### Note
```
C c(C()); // function declaration, takes a function pointer to a function returning C
```
