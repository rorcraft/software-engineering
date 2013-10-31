__From code to Heap__

http://www.ibm.com/developerworks/library/j-codetoheap/
http://www.youtube.com/watch?v=FLcXf9pO27w

__Native Process Memory__

http://www.ibm.com/developerworks/java/library/j-nativememory-linux/

### Anatomy of a Java object

When your Java code uses the new operator to create an instance of a Java object, much more data is allocated than you might expect. For example, it might surprise you to know that the size ratio of an int value to an Integer object — the smallest object that can hold an int value — is typically 1:4. The additional overhead is metadata that the JVM uses to describe the Java object, in this case an Integer.

The amount of object metadata varies by JVM version and vendor, but it typically consists of:

* Class : A pointer to the class information, which describes the object type. In the case of a java.lang.Integer object, for example, this is a pointer to the java.lang.Integer class.
* Flags : A collection of flags that describe the state of the object, including the hash code for the object if it has one, and the shape of the object (that is, whether or not the object is an array).
* Lock : The synchronization information for the object — that is, whether the object is currently synchronized.

###  Summary of collections attributes

__Collection	Performance	Default capacity	Empty size	10K entry overhead	Accurately sized?	Expansion algorithm__
HashSet	O(1)	16  144	360K	No	x2
HashMap	O(1)	16	128	360K	No	x2
Hashtable	O(1)	11	104	360K	No	x2+1
LinkedList	O(n)	1	48	240K	Yes	+1
ArrayList	O(n)	10	88	40K	No	x1.5
StringBuffer	O(1)	16	72	24	No	x2
