http://www.amazon.com/dp/0321721330

* Depend on behaviour, not data
* Single Responsibility, isolate extra responsibility, and dependencies.
* Control the dependency direction. Depend on things that change less often than you do.
* Dependency injection.
* Ask, don't tell.
* Hook methods allow subclass to contribute specializations.
* honor the contract - don't change the subclass interfaces - Liskov Substitution Principle.
* template method pattern
* shallow hierachies
* use modules for role extension.
* Duck Typing: behaves-like-a relationship - by behaviour instead of by class type.
* inheritance: is-a relationship
* composition: has-a relationship

__Design Goal: Cheaply Reduce Coupling__
http://www.threeriversinstitute.org/blog/?p=104

Coupling measures the spread of a change across elements. Cohesion measures the cost of a change within an element. An element is cohesive to the degree that the entire element changes when the system needs to change.

If one element can be insulated from a likely change in another at reasonable cost, then it’s worth doing sooner rather than later. Breaking other forms of coupling will be more expensive and might be better defered until just before a triggering change.

