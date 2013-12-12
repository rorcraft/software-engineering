__Intro__

- http://www.slideshare.net/waterfox1/an-introduction-to-jvm-internals-and-garbage-collection-in-java
- http://sureshsvn.com/jvm.html
- http://www.jclarity.com/2013/12/11/poorly-chosen-java-hotspot-garbage-collection-flags-and-how-to-fix-them/

![](http://sureshsvn.com/jvm.png)

__Generational__ 

> weak generational hypothesis - most objects die young

__Young Gen__

Most of the newly created objects are located here. When objects disappear from this area, we say a "minor GC" has occurred. 

__Old Gen (Tenured)__
 
The objects that did not become unreachable and survived from the young generation are copied here. When objects disappear from the old generation, we say a "major GC" (or a "full GC") has occurred. 
 
__PermGen__

"method area," and it stores classes or interned character strings. So, this area is definitely not for objects that survived from the old generation to stay permanently. 

__Card table__

Whenever an object in the old generation references an object in the young generation, it is recorded in this table.
"card table" in the old generation, which is a 512 byte chunk.
When a GC is executed for the young generation, only this card table is searched to determine whether or not it is subject for GC, instead of checking the reference of all the objects in the old generation. This card table is managed with write barrier.

__Young Gen is Copy Collector__

__1 Eden space, 2 survivor spaces__

_primary advantage is that allocation is extremely fast_

The order of execution process of each space is as below:
* The majority of newly created objects are located in the Eden space.
* After one GC in the Eden space, the surviving objects are moved to one of the Survivor spaces. 
* After a GC in the Eden space, the objects are piled up into the Survivor space, where other surviving objects already exist.
* Once a Survivor space is full, surviving objects are moved to the other Survivor space. Then, the Survivor space that is full will be changed to a state where there is no data at all.
* The objects that survived these steps that have been repeated a number of times are moved to the old generation.

in HotSpot VM, two techniques are used for faster memory allocations. One is called "bump-the-pointer," and the other is called "TLABs (Thread-Local Allocation Buffers)."

### Strategies

* http://www.javaperformancetuning.com/news/qotm026.shtml
* http://www.oracle.com/technetwork/java/javase/gc-tuning-6-140523.html

__Stop the world (STW)__

- Serial (single threaded)
- Parallel (multi threaded, higher throughput) (not exactly mean concurrent!)

__Concurrent__

- = single thread running along side with other threads. (!= parallel)

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

__No compaction__, uses a free-list, if compaction is needed - fall back to STW Compaction.

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
- http://www.cubrid.org/blog/dev-platform/how-to-tune-java-garbage-collection/

Talks:
- [Understanding Java Garbage Collection and what you can do about it](http://www.youtube.com/watch?v=we_enrM7TSY)
- [Concurrency & Garbage Collection - Considerations as the JVM Goes to Big Data](http://www.youtube.com/watch?v=8BwXijVmvKk)

Tools:
- http://www.azulsystems.com/products - Zing
- http://www.azulsystems.com/jHiccup
- https://github.com/bitcharmer/heaptrasher

JRuby:

`-Xmx500m -Xss1024k -Djruby.memory.max=500m -Djruby.stack.max=1024k –server -J-XX:+UseParallelGC –fast`
` -verbose:gc`
```
ruby -J-Xms500m -J-Xss1024k --server -J-verbose:gc -J-XX:+UseConcMarkSweepGC -J-XX:+UseParNewGC --fast -S bin/puma -e load_test -p 8080 -t 16:16
ruby -J-Xms500m -J-Xss1024k --server -J-verbose:gc -J-XX:+UseG1GC -J-XX:+UnlockExperimentalVMOptions -J-XX:MaxGCPauseMillis=20 --fast -S bin/puma -e load_test -p 8080 -t 16:16
```
