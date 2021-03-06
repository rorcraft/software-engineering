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

### 20. use `std::weak_ptr` for `std::shared_ptr` like pointers that can dangle.
* dangle (`expired()`) - when object it is supposed to point to no longer exists.
* same size as shared_ptr
* uses the secondary reference count in the control block.
* e.g. A1 -> B <- A2 (use shared_ptr), B -> A1 (use weak_ptr), prevent shared_ptr cycle which will not destroy the objects.

### 21. Prefer `make_shared` and `make_unique` instead of `new`.
* make_unique is not in C++11, but folly has one, and it is in C++14.
```
template <typename T, typename... Ts>
std::unique_ptr<T> make_unique(Ts&&... params) {
  return std::unique_ptr<T>(new T(std::forward<Ts>(params)...);
}
```
* Watch out for out of order execution, exception at computePriority may leak the new Widget
```
void processWidget(std::shared_ptr<Widget>(new Widget), computePriority());
void processWidget(std::make_shared<Widget>(), computePriority());
```
* Limitation, none of the `make` functions permit custom deleters but both constructors do.
* Within the `make` functions, the perfect forward use parenthese `()`, not braces `{}`.
```
auto initList = { 10, 20 }; // work around
auto spv = std::make_shared<std::vector<int>>(initList); 
```
* `make_shared` is more performant vs `new` because control block is being placed in the same chunk of memory as the managed object. 
* If the object type is quite large and the time between destruction of last shared_ptr and the last weak_ptr is significant, a lag can occur between when an object is destroyed and when the memory it occupied is freed.
```
// try to write exception safe but efficient initialization
std::shared_ptr<Widget> spw(new Widget, cusDel); // lvalue
processWidget(std::move(spw), computePriority()); // move otherwise a copy will be slow.
```
### 22. When using Pimpl Idiom, define special member functions in the implementation file.
* "pointer to implementation" idiom. Replace data members of a class with a pointer to an implementation class.
* Speeds up compilation by reducing dependencies between class clients and implementations.
* _incomplete type_:
  1. declaration of a data member that's a pointer to an incomplete type.
  2. dynamic allocation and deallocation of the object that holds the data members that used to be in original class.
```
Widget::Widget() 
: pImpl(std::make_unique<Impl>()) {}

Widget::~Widget() = default; // dtor definition, in order to be a complete type.
```
* default deleter employ C++11’ s static_assert to ensure that the raw pointer doesn’t point to an incomplete type.

## Rvalue References, Move Semantics and Perfect Forwarding.
```
// !important to remember
void f(Widget&& w); // w is lvalue even thought its type is rvalue reference
```

### 23. `std::move` and `std::forward`
* perform casts, `std::move` unconditionally casts its argument to an rvalue. `std::forward` performs this cast only if a particular condition is fulfilled.
```
template<typename T>
decltype(auto) move(T&& param) {
  using ReturnType = remove_reference_t<T>&&;
  return static_cast<ReturnType>(param);
}
```
* move request on `const` objects are silently transformed into copy operations.
* forward, cast lvalue param to rvalue to be passed down to other functions expecting rvalue, only if param is bound to an rvalue.

### 24. Distinguish univeral references from rvalue references.
* `T&&` 1. rvalue reference. 2. universal references. If you see `T&&` without type deduction, you're seeing rvalue reference.
```
void f(Widget&& param); // rvalue
Widget&& var1 = Widget(); // rvalue
template <typename T>
void f(T&& param); // universal
void f(const T&& param); // rvalue const
```
* to be universal, type deduction is neccesary. 

### 25. Use `std::move` on rvalue ref, `std::forward` on universal ref.
```
// instead of overloading (doesn't scale if multiple params, unperformant)
void setName(const std::string& newName) { name = newName; }
void setName(std::string&& newName) { name = std::move(newName); } 
// use universal ref and forward
void setName(T&& newName) { name = std::forward(newName); }
```
```
// return by value
Matrix operator+(Matrix&& lhs, Matrix& rhs) {
  lhs += rhs;
  return std::move(lhs); // move into return value location.
  return lhs; // copy 
}
```
```
// watch out, don't return rvalue of local var
Widget makeWidget() {
  Widget w;
  return std::move(w); // compiler already has "Return Value Optimization"
}
```
Condition for RVO.
* the type of local object is the same as that returned by the function
* the local object is what's being returned.

### 26. Avoid overloading on universal references.
* universal references overload is greedy, don't mix with other types.
* Perfect-forwarding constructors are especially problematic, because they’re typically better matches than copy constructors for non-const lvalues, and they can hijack derived class calls to base class copy and move constructors.
* in situations where a template instantiation and a non-template function (i.e., a “normal” function) are equally good matches for a function call, the normal function is preferred.

### 27. Alternatives to overloading on universal references.
* __Abandon overloading__: use different function names.
* __Pass by const T&__: more code
* __Pass by value__: cleaner, could be performant if you'll need to copy them anyway.
* __Tag dispatch__: Add hints for overload resolution.
```
template<typename T>
void logAndAdd(T&& name) {
  // T would be deduced to int& if name is lvalue int.
  logAndAddImpl(std::forward<T>(name), std::is_integral(typename std::remove_reference<T>::type>());
}
void logAndAddImpl(T&& name, std::false_type) {
  names.emplace(std::forward<T>name)
}
void logAndAddImpl(int idx, std::true_type) {
  logAndAdd(nameFromIdx(idx));
}
```
* __Constraining templates that takes universal references__: `std::enable_if` - force compilers to behave as if particular template didn't exist. e.g. enable Person perfect-forwarding constructor only if the type being passed isn't Person. 
```
class Person {
 public:
   template<typename T, typename = typename std::enable_if<condition>::type>
   explicit Person(T&& n);
// strip any references, consts, volatiles from T before checking type
// std::decay<T>::type
// condition:
// !std::is_same<Person, typename std::decay<T>::type>::value
// watch out is_same will run into problems when derived class pass in derived objects into base constructors.
SpecialPerson(cost SpecialPerson& rhs) : Person(rhs) {...}
// use is_base_of.
   template <typename T,
             typename = typename std::enable_if<
                          !std::is_base_of<Person, typename std::decay<T>::type>::value
                          &&
                          !std::is_integral<std::move_reference<T>>::value
                        >::type
            >
   explicit Person(T&& n) : name(std::foward<T>(n)) {}
   explicit Person(int idx) : name(nameFromIdx(idx)) {}
   // error messages can be crypic, use is_constructible
   static_assert(std::is_constructible<std::string, T>::value,
     "Parameter n cannot be used to construct std::string"
   );
```
* SFINAE - ?

### 28. reference collapsing.
* template instantiation, `auto` type generation, `typedef` and alias declarations and `deltype`.
* when compilers generate reference to a reference in a reference collapsing context, result becomes a single reference. 

### 29. assume move operations are not present, not cheap and not used.

### 30. Familiarize yourself with perfect forwarding failure cases.
* Perfect forwarding means we don’t just forward objects, we also forward their salient characteristics: their types, whether they’re lvalues or rvalues, and whether they’re const or volatile.
```
f({1,2,3});
fwd({1,2,3}); // doesn't compile, forbidden to deduce std::initializer_list. "non-deduced context"
auto il = {1,2,3};
fwd(il); // works
```
* Compiler unable to deduce type.
* Compiler deduced the wrong type.
* `0` and `NULL` type deduction can be awry.
* There's no need to define integral `static const` and `constexpr` data members in classes. Compilers perform const propagation. Passing integral static const and constexpr data members by reference "generally" requires that they be defined.
```
constexpr std::size_t Widget::MinVals; // defined in .cpp file
```
* Overloaded function names and template names. Compile figures out which function to overload by comparing declaration.
```
using ProcessFuncType = int (*)(int);
ProcessFuncType processValPtr = processVal; // processVal is a function
fwd(processValPtr);
```
* bit-field, a non-const reference shall not bound to a bit-field. Simply pass a copy.

## Lambda Expressions.
* lambda expression
* closure object - instantiated lambda expression.
* closure class - class from which the closure is instantiated. lambda generates a unique closure class.

### 31. Avoid default capture modes.
* default by-reference capture and lead to dangling references.
* default by-value may also not be self-contained.
* If the lifetime of the closure created by lambda exceeds the lifetime of local variable, the reference will dangle.
```
auto divisor = computeDivisor(calc1, calc2); 
filters.emplace_back( 
  [&](int value) { return value % divisor == 0; } // divisor could dangle
);
static auto divisor;
[=](int value) { return value % divisor == 0; } // does not copy stsatic var divisor
```
* Every non-static member functino has an implicit `this` pointer
* Default capture `[=]` captures the `this` pointer and not `divisor`.

### 32. Use init capture for capturing move only objects into closure.
* C++14 has init capture (generalized lambda capture). It takes an expression.
```
[divisor = divisor](int value) { return value % divisor == 0; };

auto pw = std::make_unique<Widget>();
[pw = std::move(pw)](int value) { return pw->isValidated(); };
// or directly
[pw = std::make_unique<Widget>()](int value) { return pw->isValidated(); };
```
* C++11 can use class with operator() or `std::bind`
```
auto func =
  std::bind(
    [](const std:vector<double>& data) {},
    std::move(data)
  )
)
```
### 33. Use `decltype` on `auto&&` parameters to `std::forward` them
```
auto func = [](auto&& x) { return normalize(std::forward<decltype(x)>(x)); };
```

### 34. Prefer lambda to `std::bind`
* lambdas are more readable
```
auto setSoundB = std::bind(
  setAlarm, // need to cast to function pointer type if setAlarm is overloaded
  std::bind(std::plus<>(), // to delay evaluation
            std::bind(steady_clock::now),
            1h),
  _1,
  30);
```
* `std::bind` always copies its arguments, but callers can achieve the effect of having an argument stored by reference by applying `std::ref` to it. 
```
auto compressRateB = std:: bind( compress, std:: ref( w), _1);
```
* compressRateB acts as if it holds a reference to w, rather than a copy.

## Currency API

### 35. Prefer task based programming to thread-based.
* `std::thread t(doAsyncWork);` － thread based.
* `auto fut = std::async(doAsyncWork);` - task based.
* std::thread cannot return values from called function, if it throws, the program terminates.
* thread based programming calls for manual management of thread exhaustion, oversubscription, load balancing and adaptation to new platforms.

### 36. Specify `std::launch::async` launch policy if asynchroncity is essential
* `std::launch::async` - f must be run asynchronously, ie on a different thread.
* `std::launch::deferred` - f is only run when `get` or `wait` is called on the future.
* default policy without specification can be either or, ie unpredictable to be async or sync, therefore cannot expect to use `thread_local` variables. WATCH OUT for using wait_for on the future from default std::async.
Okay to use `std::async` when:
* task need not run concurrently with the thread calling `get` or `wait`
* doesn't matter which thread's thread_local variables.
* acceptable task may never execute
* wait_for, wait_until checks for deferred status.

### 37. Make `std::thread` unjoinable on all paths.
* state is either joinable (thread is blocked or waiting) or unjoinable (default constructed, moved, joined or detached).
* use RAII (Resource Acquisition Is Initialization) object to explicitly join or detach on destruct.
WATCH OUT:
* `join`-on-destruction can lead to difficult to debug performance anomalies.
* `detach`-on-destruction can lead to undefined behaviour.
* declare `std::thread` last in list of data members.

### 38. Beware of varying thread handle destructor behavior.
* destructor for a future sometimes behaves as if it did an implicit `join` or `detach` or neither. It never causes program to terminate. 
* result of std::promise is stored in shared state, heap based object.
* The destructor of the last future referring to a shared state for a non-deferred task launched via std::async __blocks__ until the task completes.
* The destructor for all other futures destroys the future object. (detached, and never run)
* use `std::packaged_task` to store future result into shared state.

### 39. Consider `void` futures for one-shot event communication.
```
std::condition_variable cv;
std::mutex m;
cv.notify_one(); // or notify_all
/*******************/
std::unique_lock<std::mutex> lk(m);
cv.wait(lk);
```
* If the detecting task notifies before reacting task waits, the reacting task will hang.
* The wait statement fails to account for spurious wakeups.
```
std::atmoic<bool> flag(false);
// ...
while(!flag); // polling still runs, not true blocking.
```
```
bool flag(false);
{ 
  std::lock_guard<std::mutext> g(m);
  flag = true;
}
cv.notify_one();
// ...
{
  std::unique_lock<std::mutext> lk(m);
  cv.wait(lk, [] { return flag; })
}
```
```
std::promise<void> p;
p.set_value(); // notifier
p.get_future().wait(); // reactor
```
* requires no mutex, works regardless of whether the detecting task sets it std::promise before the reacting task waits, and is immune to spurious wakeups. 

### 40. Use `std::atomic` for concurrency, `volatile` (don't optimize) for special memory.

## Tweaks

### 41. Consider pass by value for copyable parameters that are cheap to move and always copied.

### 42: Consider emplacement instead of insertion.
