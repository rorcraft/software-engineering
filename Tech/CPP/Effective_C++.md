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
### 13. Prefer `const_iterators` to `iterators`
```
auto it = std::find(values.cbegin(), values.cend(), 1983);
values.insert(it, 1998);

auto it = std::find(cbegin(container), cend(container), targetVal); // C++14 only

// template code, prefer non-member version of begin, end, etc over member function
```
### 14. declare `noexcept` if won't emit exceptions
* `std:: vector:: push_back` takes advantage of this “move if you can, but copy if you must” strategy,
* more optimizable
* destructors are `noexcept` by default. 

### 15. Use `constexpr` whenever possible.
* constexpr functions can be used in contexts that demand compile-time constants. If the values of the arguments you pass to a constexpr function in such a context are known during compilation, the result will be computed during compilation. If any of the arguments’ values is not known during compilation, your code will be rejected.
* When a constexpr function is called with one or more values that are not known during compilation, it acts like a normal function, computing its result at runtime. This means you don’t need two functions to perform the same operation, one for compile-time constants and one for all other values. The constexpr function does it all.
```
constexpr int pow(int base, int exp) noexcept {
  // C++11 no more than a return statement, but you can use ternary cond and recursion.
  return (exp == 0 ? 1 : base * pow(base, exp - 1));
  // C++14 much looser.
}
std::array<int, pow(3, numConds)> results; // compile time as 3^numConds elements
auto baseToExp = pow(base, exp); // runtime
```
* User defined type can be constexpr as well. ctor and getters.

### 16. Make `const` member functions thread safe.
* For a single variable or memory location requiring synchronization, use of a `std::atomic` is adequate, but once you get to two or more variables or memory locations that require manipulation as a unit, you should reach for a `mutex`.

### 17. Understand special member function generation.
* C++98: ctor, dtor, copy ctor, copy assignment operator. Generated only if needed.
* Generated special member functions are implicitly public and inline and nonvirtual.
* C++11: 2 more - move ctor and move assignment
* The two copy operations are independent: declaring one doesn't prevent compilers from generating the other.
* The two move operations are not independent. If you declare either, that prevents compilers from generating the other.
* move operations won't be generated for any class that explicitly declares a copy operation.
* declaring a move operation (ctor or assignment) in a class causes compilers to disable the copy operations. (using `delete`)
* Rule of three: copy ctor, copy assignment, dtor - should declare all three.
`move` operations are genearated for classes (when needed) only if these are true:
* No copy operations are declared in the class.
* No move operations are declared in the class.
* No destructor is declared in teh class.
C++11 deprecates the auto generation of copy operations for classes declasing copy operations or a destructor.
```
class Widget {
 public:
   ~Widget(); // watch out for declaring this later, move operations will not be auto generated.
   Widget(const Widget&) = default; 
   Widget& operator=(const Widget&) = default;
};
// Polymorphic base class
class Base {
 public:
  virtual ~Base() = default; // if not virtual delete or typeid on derived class may yield undefined behavior.
  Base(Base&&) = default;
  Base& operator=(Base&&) = default;
  Base(const Base&) = default;
  Base& operator=(const Base&) = default;
};
```
* Default constructor: generated only if no user-declared constructors.
* Destructor: `noexcept` by default, virtual only if a base class dtor is virtual.
* Copy Ctor: generated only if class lacks user-declared copy ctor. deleted if class declares move operation. will not auto generate if user declared copy assignment operator or dtor.
* Copy assignment operator: Generated if class lacks user-declared copy assignment operator. Deleted if class declares copy assignment operator. Will not generate if user declare ctor or dtor.
* Move ctor, move assignment operator: Each performs memberwise moving of non-static data members. Generated only if the class contains no user-declared copy operations, move operations, or destructor. 

* Applying `final` to virtual function prevents it from being overridden in derived class. Also application to class, prohibiting inheritance.
* `std::vector::push_back` call `std::move_if_noexcept`, a variation of `std::move` that conditionally casts to an rvalue.

## Smart Pointers
### 18. `std::unique_ptr` for exclusive-ownership
* by default, same as as raw pointers.
* copying isn't allowed.
* move-only type.
* on destruction, a non-null unique_ptr destroys its resource, by default calls `delete` to the raw pointer inside.
* use case: factory function return type.
* watch out: Function object deleters with extensive state can yield std:: unique_ptr objects of significant size.
* easily and efficiently converts to a `std::shared_ptr sp = myFactoryFunc();`
* factory function should return `unique_ptr` and let client decide to convert to `shared_ptr`.

### 19. `std::shared_ptr` for shared-ownership.
* all shared_ptrs collaborate to ensure its destruction at the point where it's no longer needed. (like gc)
* reference counted. ctor increments the count, dtor decrements, copy assignment operators do both.
* if count = 0 after decrement, it'll destroy it.
* __twice__ the size of raw pointers.
* memory for reference count must be dynamically allocated
* increment, decrement of the reference must be atomic.
* move constructing shared_ptr, no reference manipulation is required.
* the type of deleter is part of the type of unique_ptr, whereas it is not part of shared_ptr
* specifying a custom deleter doesn't change the size of a std::shared_ptr object.
* reference count is part of data structure - control block. (a copy of custom deleter, a copy of custom allocator, secondary reference count - weak count)
* `std::make_shared` always create a control block.
* A control block is created when a `std::shared_ptr` is constructed from unique-ownership pointer (std::unique_str, std::auto_ptr). 
* Watch out: ctor called with raw pointer, it creates a control block.
* `std::enable_shared_from_this<T>` create shared_ptr from `*this`. _The Curiously Recurring Template Pattern (CRTP)_. Defines a member function that creates a `std::shared_ptr` to the current object, `shared_from_this`.
```
class Widget: public std::enable_shared_from_this<Widget> {
 public:
  void process();
  template<typename... Ts>
  static std::shared_ptr<Widget> create(Ts&&... params);
};
void Widget::process() {
  processedWidgets.emplace_back(shared_from_this()); // not (*this)
}
```
* There must be an existing `shared_ptr` that points to current boject, otherwise its undefined. Often declare ctors private and have clients create objects by factory functions.
* Don't use share_ptr<T> arrays. 
