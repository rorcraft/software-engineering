Object - { Behavior { State } } 
* Abstraction - concept of describing something in simpler terms.
* Encapsulation - hide the values or state of a structured data object inside a class
* Dependencies.
* Decoupling and Coherence.

http://www.amazon.com/dp/0321721330

* Depend on behaviour, not data
* Single Responsibility, isolate extra responsibility, and dependencies.
* Control the dependency direction. Depend on things that change less often than you do.
* Dependency injection.
* Tell, don't ask.
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

__Anemic Domain Model (antipattern)__
http://martinfowler.com/bliki/AnemicDomainModel.html

"The basic symptom of an Anemic Domain Model is that at first blush it looks like the real thing. There are objects, many named after the nouns in the domain space, and these objects are connected with the rich relationships and structure that true domain models have. The catch comes when you look at the behavior, and you realize that there is hardly any behavior on these objects, making them little more than bags of getters and setters. Indeed often these models come with design rules that say that you are not to put any domain logic in the the domain objects. Instead there are a set of service objects which capture all the domain logic. These services live on top of the domain model and use the domain model for data."

> Application Layer [his name for Service Layer]: Defines the jobs the software is supposed to do and directs the expressive domain objects to work out problems. The tasks this layer is responsible for are meaningful to the business or necessary for interaction with the application layers of other systems. This layer is kept thin. It does not contain business rules or knowledge, but only coordinates tasks and delegates work to collaborations of domain objects in the next layer down. It does not have state reflecting the business situation, but it can have state that reflects the progress of a task for the user or the program.

> Domain Layer (or Model Layer): Responsible for representing concepts of the business, information about the business situation, and business rules. State that reflects the business situation is controlled and used here, even though the technical details of storing it are delegated to the infrastructure. This layer is the heart of business software.

The key point here is that the Service Layer is thin - all the key logic lies in the domain layer. He reiterates this point in his service pattern:

> Now, the more common mistake is to give up too easily on fitting the behavior into an appropriate object, gradually slipping toward procedural programming.
