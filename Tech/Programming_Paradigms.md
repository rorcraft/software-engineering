### Concepts -> Paradigms -> Languages 

* The concepts are the basic primitive elements used to construct the paradigms.
* A paradigm almost always has to be Turing complete to be practical.

## Paradigm properties

__Observable nondeterminism__
When execution of a program is not completely determined by its specification. It is observable if a user can see different results from same internal configuration. e.g. concurrent programming race condition can be nondeterministic and observable.

__Named state__
State store a sequence of values in time. 3 axes of expressiveness (8 combinations):
* named/unamed
* deterministic/nondeterministic
* sequential/concurrent

```
               Declarative paradigms (relational and functional)
                     [ unamed, deterministic, sequential ]
                             /                                        \
  [ named, deterministic, sequential]    [ unamed, deterministic, concurrent ] 
  Imperative programming                          Deterministic concurrency
                   |                                                                  |
[ named, nondeterministic, sequential]   [ unamed, nondeterministic, concurrent]
Gaurded command programming                 concurrent logic programming
                  \                                                                /
                     [ named, nondeterministic, concurrent ]
              Message-passing and shared state concurrency
```

If the need for pervasive modifications manifests itself, we can take this as a sign that there is a new concept waiting to be discovered. (programs become complicated for reasons that have no direct relationship to the specific problem)

Many languages support two paradigms, typically one for programming in the small, another for programming in the large (abstraction and modularity).

a closure combines a procedure with its external references (the references it uses at its definition)


