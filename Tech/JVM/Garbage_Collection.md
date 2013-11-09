### Eden (Young Gen)
 
### PermGen

### 2 survivor spaces (Copying Garbage Collector)

The role of two survivor spaces gets reversed after the operation of a minor garbage collection

The two survivor spaces. These hold objects that have survived at least one minor garbage collection but have been given another chance to become unreachable before being promoted to the old generation. Only one of them holds objects, while the other is most of the time unused.

During the operation of a minor garbage collection, objects that have been found to be garbage will be marked. Live objects in the eden that survive the collection are copied to the unused survivor space. Live objects in the survivor space that is in use, which will be given another chance to be reclaimed in the young generation, are also copied to the unused survivor space. Finally, live objects in the survivor space that is in use, that are deemed “old enough,” are promoted to the old generation.

At the end of the minor garbage collection, the two survivor spaces swap roles. The eden is entirely empty; only one survivor space is in use; and the occupancy of the old generation has grown slightly. Because live objects are copied during its operation, this type of garbage collector is called a copying garbage collector.

* http://blogs.msdn.com/b/abhinaba/archive/2009/02/02/back-to-basics-copying-garbage-collection.aspx

__primary advantage is that allocation is extremely fast__

[Concurrency & Garbage Collection - Considerations as the JVM Goes to Big Data](http://www.youtube.com/watch?v=8BwXijVmvKk)
