__From code to Heap__

* http://www.ibm.com/developerworks/library/j-codetoheap/
* http://www.youtube.com/watch?v=FLcXf9pO27w

__Native Process Memory__

http://www.ibm.com/developerworks/java/library/j-nativememory-linux/

### Anatomy of a Java object

When your Java code uses the new operator to create an instance of a Java object, much more data is allocated than you might expect. For example, it might surprise you to know that the size ratio of an int value to an Integer object — the smallest object that can hold an int value — is typically 1:4. The additional overhead is metadata that the JVM uses to describe the Java object, in this case an Integer.

The amount of object metadata varies by JVM version and vendor, but it typically consists of:

* Class : A pointer to the class information, which describes the object type. In the case of a java.lang.Integer object, for example, this is a pointer to the java.lang.Integer class.
* Flags : A collection of flags that describe the state of the object, including the hash code for the object if it has one, and the shape of the object (that is, whether or not the object is an array).
* Lock : The synchronization information for the object — that is, whether the object is currently synchronized.

###  Summary of collections attributes

<table border="0" cellpadding="0" cellspacing="0" class="ibm-data-table ibm-alternating" summary="Summary of collections attributes"><tbody><tr><th scope="col">Collection</th><th scope="col">Performance</th><th scope="col">Default capacity</th><th scope="col">Empty size</th><th scope="col">10K entry overhead</th><th scope="col">Accurately sized?</th><th scope="col">Expansion algorithm</th></tr><tr class="ibm-alt-row"><td><code>HashSet</code></td><td>O(1)</td><td>16</td><td>144</td><td>360K</td><td>No</td><td>x2</td></tr><tr><td><code>HashMap</code></td><td>O(1)</td><td>16</td><td>128</td><td>360K</td><td>No</td><td>x2</td></tr><tr class="ibm-alt-row"><td><code>Hashtable</code></td><td>O(1)</td><td>11</td><td>104</td><td>360K</td><td>No</td><td>x2+1</td></tr><tr><td><code>LinkedList</code></td><td>O(n)</td><td>1</td><td>48</td><td>240K</td><td>Yes</td><td>+1</td></tr><tr class="ibm-alt-row"><td><code>ArrayList</code></td><td>O(n)</td><td>10</td><td>88</td><td>40K</td><td>No</td><td>x1.5</td></tr><tr><td><code>StringBuffer</code></td><td>O(1)</td><td>16</td><td>72</td><td>24</td><td>No</td><td>x2</td></tr></tbody></table>
