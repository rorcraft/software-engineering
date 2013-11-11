### Eden (Young Gen)
 
### PermGen

### 2 survivor spaces (Copying Garbage Collector)

The role of two survivor spaces gets reversed after the operation of a minor garbage collection

The two survivor spaces. These hold objects that have survived at least one minor garbage collection but have been given another chance to become unreachable before being promoted to the old generation. Only one of them holds objects, while the other is most of the time unused.

During the operation of a minor garbage collection, objects that have been found to be garbage will be marked. Live objects in the eden that survive the collection are copied to the unused survivor space. Live objects in the survivor space that is in use, which will be given another chance to be reclaimed in the young generation, are also copied to the unused survivor space. Finally, live objects in the survivor space that is in use, that are deemed “old enough,” are promoted to the old generation.

At the end of the minor garbage collection, the two survivor spaces swap roles. The eden is entirely empty; only one survivor space is in use; and the occupancy of the old generation has grown slightly. Because live objects are copied during its operation, this type of garbage collector is called a copying garbage collector.

* http://blogs.msdn.com/b/abhinaba/archive/2009/02/02/back-to-basics-copying-garbage-collection.aspx

__primary advantage is that allocation is extremely fast__

### Strategies

* http://www.javaperformancetuning.com/news/qotm026.shtml
* http://www.oracle.com/technetwork/java/javase/gc-tuning-6-140523.html

__Stop the world (STW)__

- Serial (single threaded)
- Parallel (multi threaded, higher throughput) (not exactly mean concurrent!)

__Concurrent__

- = single thread running along side with other threads.

__Flags for young generation:__

`-XX:+UseSerialGC` (Copying Collector) (STW)

`-XX:+UseParallelGC` (Parallel Scavenge Collector) (STW) 
* tuned for gigabytes heaps, (over 10GB). 
* It has an optional adaptive tuning policy which will automatically resize heap spaces.
* Does not work with CMS.  

`-XX:+UseParNewGC` (Parallel Copying Collector) (STW)  
* Works with CMS
 
__Flags for old generation__

`-XX:+UseParallelOldGC` (STW) _by default_

`-XX:+UseConcMarkSweepGC` (Concurrent Mark and Sweep - CMS) - 6 phases

1. the initial-mark phase (stop-the-world, snapshot the old generation so that we can run most of the rest of the collection concurrent to the application threads);
2. the mark phase (concurrent, mark the live objects traversing the object graph from the roots);
3. the pre-cleaning phase (concurrent);
4. the re-mark phase (stop-the-world, another snapshot to capture any changes to live objects since the collection started);
5. the sweep phase (concurrent, recycles memory by clearing unreferenced objects);
6. the reset phase (concurrent).
If "the rate of creation" of objects is too high, and the concurrent collector is not able to keep up with the concurrent collection, it falls back to the traditional mark-sweep collector.

* `-XX:+CMSIncrementalMode` 

`-Xincgc` (Incremental Collector)
* Uses a "train" algorithm to collect small portions of the old generation at a time. STW pause is minimized at the cost of total garbage collection taking longer.

`-XX:+UseG1GC` Garbage first (G1)

__Compaction__


__Pauseless GC__

http://www.artima.com/lejava/articles/azul_pauseless_gc.html


JVM internals
- http://www.slideshare.net/waterfox1/an-introduction-to-jvm-internals-and-garbage-collection-in-java
- http://www.artima.com/insidejvm/ed2/jvm.html (old)
- http://www.cubrid.org/blog/dev-platform/understanding-jvm-internals/

Collector
- http://sureshsvn.com/jvm.html
- http://www.cubrid.org/blog/dev-platform/understanding-java-garbage-collection/
- http://www.oracle.com/technetwork/java/javase/memorymanagement-whitepaper-150215.pdf
- http://www.infoq.com/articles/G1-One-Garbage-Collector-To-Rule-Them-All

Talks:
- [Concurrency & Garbage Collection - Considerations as the JVM Goes to Big Data](http://www.youtube.com/watch?v=8BwXijVmvKk)
