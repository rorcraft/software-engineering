### Memory management

```
int* invalid_pointer() {
  int x;
  return &x; // when function completes, x is deallocated.
}
```

```
int* pointer_to_static() {
  static int x; // one instance of persisted x
  return &x;
}
```

```
int* p = new int(42); // allocates and let p handles the object.
delete p; // dealloc once it's done.

int* pointer_to_dynamic() {
  return new int(10);
}
// caller has responsibility of freeing the object.
int* p = pointer_to_dynamic(); // any copying going on?
```

### Unique pointer 

### Shared pointer
* always use a named smart pointer variable to hold the result of new
* (old) http://www.boost.org/doc/libs/1_43_0/libs/smart_ptr/shared_ptr.htm#BestPractices
* example: http://www.boost.org/doc/libs/1_43_0/libs/smart_ptr/example/shared_ptr_example.cpp
* http://www.boost.org/doc/libs/1_57_0/libs/smart_ptr/sp_techniques.html
* http://www.umich.edu/~eecs381/handouts/C++11_smart_ptrs.pdf

