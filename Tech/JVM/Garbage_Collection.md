__Intro__

- http://www.slideshare.net/waterfox1/an-introduction-to-jvm-internals-and-garbage-collection-in-java
- http://sureshsvn.com/jvm.html

![](http://sureshsvn.com/jvm.png)

> weak generational hypothesis - most objects die young

__Eden (Young Gen)__

__Old Gen__
 
__PermGen__

__2 survivor spaces (Copying Garbage Collector)__

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

* `-XX:+CMSIncrementalMode` - http://www.oracle.com/technetwork/java/javase/gc-tuning-6-140523.html 

`-Xincgc` (Incremental Collector)
* Uses a "train" algorithm to collect small portions of the old generation at a time. STW pause is minimized at the cost of total garbage collection taking longer.

`-XX:+UseG1GC` Garbage first (G1)
* An incremental parallel compacting GC that provides more predictable pause times 
* http://docs.oracle.com/javase/7/docs/technotes/guides/vm/G1.html
* http://www.drdobbs.com/jvm/g1-javas-garbage-first-garbage-collector/219401061?pgno=2
* http://www.infoq.com/articles/G1-One-Garbage-Collector-To-Rule-Them-All
* In the strictest sense, the heap doesn't contain generational areas, although a subset of the regions can be treated as such.
* The heap is broken down into equally sized regions.
* Regions are further broken down into 512 byte sections called cards.
* Subsets of these cards are tracked, and referred to as Remembered Sets (RS)
* G1 performs a concurrent global marking phase to determine the liveness of objects throughout the heap. After the mark phase completes, G1 knows which regions are mostly empty. It collects in these regions first, which usually yields a large amount of free space. This is why this method of garbage collection is called Garbage-First. As the name suggests, G1 concentrates its collection and compaction activity on the areas of the heap that are likely to be full of reclaimable objects, that is, garbage. G1 uses a pause prediction model to meet a user-defined pause time target and selects the number of regions to collect based on the specified pause time target.
* The regions identified by G1 as ripe for reclamation are garbage collected using evacuation. G1 copies objects from one or more regions of the heap to a single region on the heap, and in the process both compacts and frees up memory. This evacuation is performed in parallel on multi-processors, to decrease pause times and increase throughput. Thus, with each garbage collection, G1 continuously works to reduce fragmentation, working within the user defined pause times
* RS Maintenance
  * Cards are placed in a region's RS via a write barrier, which is an efficient block of code that all mutator threads must execute when modifying an object reference. To be precise, for a particular region (i.e., region a), only cards that contain pointers from other regions to an object in region a are recorded in region a's RS
* Concurrent Marking
  * Marking Stage. The heap regions are traversed and live objects are marked.
  * Re-marking Stage. When the heap reaches a certain percentage filled, as indicated by the number of allocations since the snapshot in the Marking Stage, the heap is re-marked.
  * Cleanup Stage. When the Re-mark Stage completes, counts of live objects are maintained.
* Evacuation and Collection
  * all mutator threads are paused, and live objects are moved from their respective regions and compacted into other regions.
* Once G1 GC successfully completes the concurrent marking cycle, it has the information that it needs to start the old generation collection.
* A collection that facilitates the compaction and evacuation of old generation is appropriately called a 'mixed' collection. A mixed collection can (and usually does) happen over multiple mixed garbage collection cycles. 

> The first focus of G1 is to provide a solution for users running applications that require large heaps with limited GC latency. This means heap sizes of around 6GB or larger, and stable and predictable pause time below 0.5 seconds.

* http://www.oracle.com/technetwork/articles/java/g1gc-1984535.html (Tuning options)
* http://www.oracle.com/technetwork/java/javase/tech/vmoptions-jsp-140102.html#G1Options

> Young Generation Size: Avoid explicitly setting young generation size with the -Xmn option or any or other related option such as -XX:NewRatio. Fixing the size of the young generation overrides the target pause-time goal.

__Compaction__

http://java.dzone.com/articles/ghost-java-virtual-machine

* Since compaction involves moving objects and updating referring objects, it can only be done during a major garbage collection while the entire JVM is paused.
* Another detrimental effect of dark matter re-allocation is locality of related objects. When a thread creates objects simultaneously and links them together, these related objects can end up being physically scattered across the heap in different memory pages.
* ConcurrentMarkAndSweep keeps fragmentation under control via free lists
* Serial Collector, Parallel Collector and the Parallel Compacting Collector performs compaction during every major collection.
* Compaction squeezes all live objects towards the beginning part of the heap.
* Use a dense prefix that acts like a watermark denoting the optimal point at which compaction should start.

__Pauseless GC__

* http://www.artima.com/lejava/articles/azul_pauseless_gc.html
* https://www.usenix.org/legacy/events/vee05/full_papers/p46-click.pdf
* http://www.javaworld.com/javaworld/jw-11-2012/121107-jvm-performance-optimization-low-latency-garbage-collection.html

> They establish a "read barrier" that allows the GC to intercept dereferencing, and this way they can lazily update the references that are actually used.


JVM internals
- http://www.artima.com/insidejvm/ed2/jvm.html (old)
- http://www.cubrid.org/blog/dev-platform/understanding-jvm-internals/

Garbage Collection 
- http://www.cubrid.org/blog/dev-platform/understanding-java-garbage-collection/
- http://www.oracle.com/technetwork/java/javase/memorymanagement-whitepaper-150215.pdf

Talks:
- [Concurrency & Garbage Collection - Considerations as the JVM Goes to Big Data](http://www.youtube.com/watch?v=8BwXijVmvKk)
