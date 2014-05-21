* http://www.cs.tau.ac.il/~shanir/concurrent-data-structures.pdf
* http://www.ibm.com/developerworks/aix/library/au-multithreaded_structures1/index.html
* http://www.liblfds.org/ written in C
* http://sysmagazine.com/posts/196834/ libcds
* http://kukuruku.co/hub/cpp/lock-free-data-structures-introduction
* http://kukuruku.co/hub/cpp/lock-free-data-structures-basics-atomicity-and-atomic-primitives

### Issues

_Amdahl’s law_ http://en.wikipedia.org/wiki/Amdahl's_law

> find the maximum expected improvement to an overall system when only part of the system is improved

```
# b - fraction of the program that is subject to a sequential bottleneck
# If program takes 1 time unit to execute on single processor
# P - number of cores

concurrent = (1 - b)/P time units
```

_Memory contention_

one a cache-coherent multiprocessor, exclusive ownership of the cache line containing the lock must be repeatedly transferred from one processor to another.

_blocking_ - OS can preempt a thread while it holds the lock.

### Blocking Techniques

_fine-grained_ locking scheme - multiple locks to protect different parts of the data structure.

_combining tree_ http://www.cs.nyu.edu/~lerner/spring11/proj_counting.pdf
- one leaf per thread
- root of tree = counter value
- threads climb the free from leaves and attempt to combine with other concurrent operations.
- a winner thread adds its sum to the counter in a single addition, thereby effecting the increments of all combined operations.
- descends the tree distributing a return value to each waiting loser thread.
- A loser operation waits by repeatedly reading a memory location in a tree node - "spinning"
- If all threads manage to repeatedly combine, then a tree of width P  
will allow P threads to return P values after every O(log P) operations required to ascend  
and descend the tree, oﬀering a speedup of O(P=log P) 

### Non Blocking Techniques

* _wait-free_ operation is guaranteed to complete after a finite number of its own steps, regardless of the timing behavior of other operations.
* _lock-free_ operation gaurantees that after a finite number of its own steps, some operation completes.
* _obstruction-free_ operation is guaranteed to complete within a finite number of its own steps after it stops encounrtering interference from other operations.

`load` and `store` instructions are not enough to complete _wait-free_ or _lock-free_.

`compare-and-swap` (CAS) and `load-linked/store-conditional` are atomic operations in modern multiprocessors.
* The CAS operation atomically loads from a memory location, compares the value read to an expected value, and stores a new value to the location if the comparison succeeds.
* The CAS instruction fails to increment the counter only if it changes between the load and the CAS. In this case, the operation can retry, as the failing CAS had no effect.
* However, it is not wait-free because a single fetch-and-inc operation can repeatedly fail its CAS.

