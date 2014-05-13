### Hash functions

* Numeric Recipes http://www.nr.com/aboutNR3book.html
* http://www.azillionmonkeys.com/qed/hash.html
* http://burtleburtle.net/bob/hash/spooky.html
* Murmur
* CityHash
* SipHash

### Collision resolution

__Separate Chaining__

* advantages:
  * The table has no hard size limit.
  * Performance degrades gracefully with more collisions.
  * The table easily handles duplicate keys.
  * Deletion is simple and permanent.
* disadvantages:
  * Rebuilding the table if it is resized is slightly harder.
  * Potentially more memory is used because of the links.
  * Potentially slower because links need to be dereferenced.

__Linear Probing__


### Dynamic Sizing

* http://openmymind.net/Back-To-Basics-Hashtables/
* http://openmymind.net/Back-To-Basics-Hasthables-Part-2/
* https://github.com/antirez/redis/blob/unstable/src/dict.c

__Redis__ 

Redis incrementally does a rehash. How? First it starts the same way by doubling the number of buckets. Instead of moving all elements over, it simply marks the hashtable as being in a "rehashing mode". As long as the the hashtable is in "rehashing mode" two sets of buckets exist, the old and the twice-larger new one.

This means that hashtables which see a lot of activity get quickly rehashed, while hashtables which see little activity don't take up as much precious cycles. Eventually, when all the buckets are moved over, the old set of buckets can be deleted and the hashtable is no longer in "rehashing mode".

Every method needs to be aware of the rehashing state of the table.

### Cuckoo Hash Table

* http://da-data.blogspot.com/2013/03/optimistic-cuckoo-hashing-for.html
