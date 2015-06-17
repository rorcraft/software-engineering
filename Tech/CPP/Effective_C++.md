# Effective C++ 
Meyers, Scott (2014-11-11). Effective Modern C++: 42 Specific Ways to Improve Your Use of C++11 and C++14

## Deducing Types

### 1. Template Type deduction
* References arguments are treated as non-references (reference-ness ignored)
* `const` or `volatile` arguments are treated as non-`const` and non-`volatile`
* array or function names decay to pointers unless they're used to initialize references.

### 2. `auto` type deduction
* auto type deduction is usually the same as template type deduction, but auto type deduction assumes that a braced initializer represents a std:: initializer_list, and template type deduction doesn’t. 
* auto in a function return type or a lambda parameter implies template type deduction, not auto type deduction.

### 3. `decltype`

### 4. typeid().name, demangle
```
typeid(T).name();
// PK6Widget
// PK = "pointer to (k)const", 6 = Name with 6 characters (Widget)

using boost:: typeindex:: type_id_with_cvr;
cout << type_id_with_cvr<T>().pretty_name();

// or folly::demangle
```
## auto

### 6. `auto` watch out for 'invisible' proxy types
```c++
std::vector<bool> features(const Widget& w);

auto highPriority = features(w)[5]; // deduced std::vector<bool>::reference
// use explicit typed initializer idiom
auto highPriority = static_cast<bool>(features(w)[5])
```
## Moving to Modern C++

### 7. Initialize with `()` and `{}`
```c++
// {} - uniform initialize
int x{0} // defaults to 0
int y = {0} // same as int y = 0
int z(0) // error if member declaration
vector<int> v{ 1, 2, 3}; // C++11 initialize vector

// Watch out for std::initializer_list overloads, overshadows where other overloads are hardly considered.
std::vector<int>(10, 20); // 10 elements
std::vector<int>{10, 20}; // 2 elements 

Widget w(); // watch out for most vexing parse, declares a function
Widget w1{}; // default ctor
Widget w2;   // default ctor
Widget w3({}); // std::initializer_list ctor with empty list

```

