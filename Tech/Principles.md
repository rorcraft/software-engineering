### Object Organization:

**SOLID principle:**

**Single responsibility principle**
an object should have only a single responsibility.
> small enough to lower coupling, but
> large enough to maximize cohesion.

**Open/closed principle**
“software entities … should be open for extension, but closed for modification”.

**Liskov substitution principle**
“objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program”. See also design by contract.

**Interface segregation principle**
“many client specific interfaces are better than one general purpose interface.”

**Dependency inversion principle**
one should “Depend upon Abstractions. Do not depend upon concretions.”
Dependency injection is one method of following this principle.

**Minimize Coupling**

**Tell, don't ask.**

This is what we are trying to prevent.
```
SortedList thingy = someObject.getEmployeeList();
thingy.addElementWithKey(foo.getKey(), foo);
```
We also have an example of Asking instead of Telling in foo.getKey()). 
Direct access of a child like this extends coupling from the caller farther than it needs to be. The caller is depending on these facts:

* someObject holds employees in a SortedList.
* SortedList’s add method is addElementWithKey()
* foo’s method to query its key is getKey()

Instead, this should be:
```
someObject.addToThingy(foo);
```

**Law of Demeter**

http://en.wikipedia.org/wiki/Law_of_Demeter

Any method of an object should only call methods belonging to:

* itself.
* any parameters that were passed in to the method.
* any objects it created.
* any composite objects.
* NOT methods belonging to objects that were returned from some other call

### Logic Organization:

DRY - Don’t repeat yourself 

### Maintenance:

Avoid Creating a **YAGNI** (You aren’t going to need it), **Premature Optimization**
Don’t even think about optimization unless your code is working, but slower than you want. Only then should you start thinking about optimizing, and then only with the aid of empirical data.

**Principle of least astonishment, KISS**

**Maximize Cohesion** - Code that has similar functionality should be found within the same component.

**Embrace Change**
