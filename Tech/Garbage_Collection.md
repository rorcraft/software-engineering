__To Know A Garbage Collector__
http://michaelrbernste.in/2013/06/10/to-know-a-garbage-collector-goruco-2013.html

* Heap
* Mutator - allocates memory from heap
* Collector - reclaims memory return to heap
* Roots - references directly accessible to mutator
* Barriers - mutates object on a heap

* Mark & Sweep - collect cyclic references but can have long pauses
* Reference Counting - incremental but cannot easily collect cyclic refs.
* Unified Theory - hybrid.

__Rare are GC Talks__
http://furious-waterfall-55.heroku.com/ruby-guide/internals/gc.html

__Boehm GC__:
http://en.wikipedia.org/wiki/Boehm_garbage_collector

__JVM GC__:
http://www.oracle.com/technetwork/java/javase/gc-tuning-6-140523.html
http://blog.mgm-tp.com/2013/03/garbage-collection-tuning/

Monitor GC:
* https://blog.engineyard.com/2010/monitoring-the-jvm-heap-with-jruby
* http://www.cubrid.org/blog/dev-platform/how-to-monitor-java-garbage-collection/

-Xmx500m -Xss1024k -Djruby.memory.max=500m -Djruby.stack.max=1024k
–server -J-XX:+UseParallelGC –fast

__PyPy__:

http://morepypy.blogspot.com/2013/10/incremental-garbage-collector-in-pypy.html
http://doc.pypy.org/en/latest/garbage_collection.html#minimark-gc

PyPy essentially has only the cycle finder - it does not bother with reference counting, instead it walks alive objects every now and then

Incremental GC spreads the walking of objects and cleaning them across the execution time in smaller intervals, hence smaller pauses.

__Rubinius__:
http://rubini.us/2013/06/22/concurrent-garbage-collection/

__Course__:
http://michaelrbernste.in/2013/05/20/adventures-in-GC-pedagogy.html
http://faculty.cs.byu.edu/~jay/courses/2012/fall/330/course/gc.html
