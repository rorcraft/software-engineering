__From code to Heap__

* http://www.ibm.com/developerworks/library/j-codetoheap/
* http://www.youtube.com/watch?v=FLcXf9pO27w

__Native Process Memory__

* http://www.ibm.com/developerworks/java/library/j-nativememory-linux/
* http://www.oracle.com/technetwork/java/javase/memleaks-137499.html#gbyvk
* http://www.oracle.com/technetwork/java/javase/clopts-139448.html#gbmtq

- Generated (JIT:ed) code (both the input (the bytecode) and the output (the executable code))
- Loaded libraries (including jar and class files)
- Control structures for the java heap
- Thread Stacks
- User native memory (malloc:ed in JNI)

Reserving native memory is not the same as allocating it. When native memory is reserved, it is not backed with physical memory or other storage. Although reserving chunks of the address space will not exhaust physical resources, it does prevent that memory from being used for other purposes. 

NIO - Direct ByteBuffers can be passed directly to native OS library functions for performing I/O — making them significantly faster in some scenarios because they can avoid copying data between Java heap and native heap.  
The application still uses an object on the Java heap to orchestrate I/O operations, but the buffer that holds the data is held in native memory — the Java heap object only contains a reference to the native heap buffer. A non-direct ByteBuffer holds its data in a byte[] array on the Java heap

Some implementations allow you to specify the stack size for Java threads. Values between 256KB and 756KB are typical.

### Anatomy of a Java object

When your Java code uses the new operator to create an instance of a Java object, much more data is allocated than you might expect. For example, it might surprise you to know that the size ratio of an int value to an Integer object — the smallest object that can hold an int value — is typically 1:4. The additional overhead is metadata that the JVM uses to describe the Java object, in this case an Integer.

The amount of object metadata varies by JVM version and vendor, but it typically consists of:

* Class : A pointer to the class information, which describes the object type. In the case of a java.lang.Integer object, for example, this is a pointer to the java.lang.Integer class.
* Flags : A collection of flags that describe the state of the object, including the hash code for the object if it has one, and the shape of the object (that is, whether or not the object is an array).
* Lock : The synchronization information for the object — that is, whether the object is currently synchronized.

###  Summary of collections attributes

<table border="0" cellpadding="0" cellspacing="0" class="ibm-data-table ibm-alternating" summary="Summary of collections attributes"><tbody><tr><th scope="col">Collection</th><th scope="col">Performance</th><th scope="col">Default capacity</th><th scope="col">Empty size</th><th scope="col">10K entry overhead</th><th scope="col">Accurately sized?</th><th scope="col">Expansion algorithm</th></tr><tr class="ibm-alt-row"><td><code>HashSet</code></td><td>O(1)</td><td>16</td><td>144</td><td>360K</td><td>No</td><td>x2</td></tr><tr><td><code>HashMap</code></td><td>O(1)</td><td>16</td><td>128</td><td>360K</td><td>No</td><td>x2</td></tr><tr class="ibm-alt-row"><td><code>Hashtable</code></td><td>O(1)</td><td>11</td><td>104</td><td>360K</td><td>No</td><td>x2+1</td></tr><tr><td><code>LinkedList</code></td><td>O(n)</td><td>1</td><td>48</td><td>240K</td><td>Yes</td><td>+1</td></tr><tr class="ibm-alt-row"><td><code>ArrayList</code></td><td>O(n)</td><td>10</td><td>88</td><td>40K</td><td>No</td><td>x1.5</td></tr><tr><td><code>StringBuffer</code></td><td>O(1)</td><td>16</td><td>72</td><td>24</td><td>No</td><td>x2</td></tr></tbody></table>

### CompressedOops

Oop - "ordinary object pointer"

https://wikis.oracle.com/display/HotSpotInternals/CompressedOops

The term "oop" is traditional to certain VMs that derive from Smalltalk and Self, including the following:

> On an LP64 system, though, the heap for any given run may have to be around 1.5 times as large as for the corresponding ILP32 system (assuming the run fits both modes). This is due to the expanded size of managed pointers.

(Additionally, on x86 chips, the ILP32 mode provides half the usable registers that the LP64 mode does. SPARC is not affected this way; RISC chips start out with lots of registers and just widen them for LP64 mode.)

Compressed oops represent managed pointers (in many but not all places in the JVM) as 32-bit values which must be scaled by a factor of 8 and added to a 64-bit base address to find the object they refer to. This allows applications to address up to four billion objects (not bytes), or a heap size of up to about 32Gb. At the same time, data structure compactness is competitive with ILP32 mode.

If UseCompressedOops is true, the following oops in the heap will be compressed:

* the klass field of every object
* every oop instance field
* every element of an oop array (objArray)

The Hotspot VM's data structures to manage Java classes are not compressed. These are generally found in the section of the Java heap known as the Permanent Generation (PermGen).

__e.g.__

Self (an prototype-based relative of Smalltalk) https://github.com/russellallen/self/blob/master/vm/src/any/objects/oop.hh
Strongtalk (a Smalltalk implementation) http://code.google.com/p/strongtalk/wiki/VMTypesForSmalltalkObjects
Hotspot  http://hg.openjdk.java.net/hsx/hotspot-main/hotspot/file/0/src/share/vm/oops/oop.hpp
V8 http://code.google.com/p/v8/source/browse/trunk/src/objects.h (mentions "smi" but not "oop")
