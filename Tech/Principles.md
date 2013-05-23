### Object Organization:

**SOLID principle:**

**Single responsibility principle**
an object should have only a single responsibility.

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

**Law of Demeter**
Code components should only communicate with their direct relations (e.g. classes that they inherit from, objects that they contain, objects passed by argument, etc.) http://en.wikipedia.org/wiki/Law_of_Demeter

### Logic Organization:

DRY - Don’t repeat yourself 

### Maintenance:

Avoid Creating a **YAGNI** (You aren’t going to need it), **Premature Optimization**
Don’t even think about optimization unless your code is working, but slower than you want. Only then should you start thinking about optimizing, and then only with the aid of empirical data.

**Principle of least astonishment, KISS**

**Maximize Cohesion** - Code that has similar functionality should be found within the same component.

**Embrace Change**
