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
### 5. Prefer auto to explicit type declarations.
* must be initialized, immune to type mismatches that can lead to portability or efficiency problems, can ease the process of refactoring.
* less typing than variables with explicitly specified types.
* Can handle undeclarable types.

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
### 8. Prefer `nullptr` to `0` and `NULL`
* Avoid overloading on integral and pointer types.
```
if (result == 0) {} // unclear
if (result == nullptr) {} // better

e.g.
void f(int);
void f(bool);
void f(void*);
f(0); // calls f(int), not f(void*)

//nullptr is circular definition std::nullptr_t
```
### 9. `alias` instead of `typedef`
* typedef don't support templatization
```
template<typename T>
using MyAllocList = std::list<T, MyAlloc<T>>;

// instead of
template<typename T>
struct MyAllocList {
  typedef std::list<T, MyAlloc<T>> type;
}
```
### 10. Scoped `enums` instead of unscoped `enums`
```
enum Color { black, white, red } // unscoped
enum class Color { black, white, red } // scoped

Color c = Color::white;

enum class Status; // scoped enum always fwd decl
enum Status: std::uint32_t; // unscoped enum need to fwd decl with type
// unscoped enum benefit
enum UserInfoFields { uiName, uiEmail, uiReputation };
UserInfo uInfo; // std::tuple<string, string, size_t>;
std::get<uiEmail>(uInfo); // instead of std::get<1>(uInfo);
// if scoped to enum class UserInfoFields
std::get<static_cast<std::size_t>(UserInfoFields::uiEmail)>(uInfo);
```
### 11. Prefer deleted functions to private undefined ones.
* any function may be deleted including non-member functions and template instantiations.
```
bool isLucky(int number);
bool isLucky(char) = delete; // reject chars
bool isLucky(double) = delete; // reject doubles and floats

// c++98
class Widget {
 private:
  isLucky(char x); // not defined
};
```
### 12. `override` virtual functions
* member function reference qualifiers make it possible to treat lvalue and rvalue objects differently.
```
class Derived: public Base {
 public:
  virtual void f1() override; // virtual is optional
  void override(); // legal
  // reference qualifiers
  Data& data() & { return values; } // lvalue instance return lvalue
  Data&& data() & { return std::move(values); } // rvalue instance return rvalue
}
auto val1 = w1.data();
auto val2 = makeWidget().data(); // rvalue
```

