JVM GC:
http://www.oracle.com/technetwork/java/javase/gc-tuning-6-140523.html

Monitor GC:
* https://blog.engineyard.com/2010/monitoring-the-jvm-heap-with-jruby
* http://www.cubrid.org/blog/dev-platform/how-to-monitor-java-garbage-collection/

-Xmx500m -Xss1024k -Djruby.memory.max=500m -Djruby.stack.max=1024k
–server -J-XX:+UseParallelGC –fast

PyPy
http://morepypy.blogspot.com/2013/10/incremental-garbage-collector-in-pypy.html

Rubinius
http://rubini.us/2013/06/22/concurrent-garbage-collection/
