### Hash functions

* Numeric Recipes - http://www.nr.com/aboutNR3book.html
* http://www.azillionmonkeys.com/qed/hash.html
* http://burtleburtle.net/bob/hash/spooky.html
* Murmur - https://en.wikipedia.org/wiki/MurmurHash
* CityHash - (Google) https://en.wikipedia.org/wiki/CityHash,  *has not been tested much on big-endian platforms.
* SipHash - (2012) https://en.wikipedia.org/wiki/SipHash,  64-bit message, 128-bit secret key.

### Collision resolution

http://eternallyconfuzzled.com/tuts/datastructures/jsw_tut_hashtable.aspx

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

The step size is almost always 1 with linear probing, but it is acceptable to use other step sizes as long as the step size is relatively prime to the table size so that every index is eventually visited.  
Keys tend to cluster. That is, several parts of the table may become full quickly while others remain completely empty.
__Quadratic probing__ Simply start with a step size of 1 and quadratically increase the step size until the desired index is found. Alleviates primary clustering because probes are no longer close together according to some small constant.

__Double hashing__  The first hash function is used as usual, and the second hash function is used to create a step size. Because the key itself determines the step size, primary clustering is avoided.

__Cuckoo Hash__ Using two hash function.
* http://da-data.blogspot.com/2013/03/optimistic-cuckoo-hashing-for.html

### Dynamic Sizing

* http://openmymind.net/Back-To-Basics-Hashtables/
* http://openmymind.net/Back-To-Basics-Hasthables-Part-2/ 

__Redis__ https://github.com/antirez/redis/blob/unstable/src/dict.c

Redis incrementally does a rehash. How? First it starts the same way by doubling the number of buckets. Instead of moving all elements over, it simply marks the hashtable as being in a "rehashing mode". As long as the the hashtable is in "rehashing mode" two sets of buckets exist, the old and the twice-larger new one.

This means that hashtables which see a lot of activity get quickly rehashed, while hashtables which see little activity don't take up as much precious cycles. Eventually, when all the buckets are moved over, the old set of buckets can be deleted and the hashtable is no longer in "rehashing mode".

Every method needs to be aware of the rehashing state of the table.

__Golang Hashmap__ http://golang.org/src/pkg/runtime/hashmap.c
```
// This file contains the implementation of Go's map type.
//
// The map is just a hash table.  The data is arranged
// into an array of buckets.  Each bucket contains up to
// 8 key/value pairs.  The low-order bits of the hash are
// used to select a bucket.  Each bucket contains a few
// high-order bits of each hash to distinguish the entries
// within a single bucket.
//
// If more than 8 keys hash to a bucket, we chain on
// extra buckets.
//
// When the hashtable grows, we allocate a new array
// of buckets twice as big.  Buckets are incrementally
// copied from the old bucket array to the new bucket array.
//
// Map iterators walk through the array of buckets and
// return the keys in walk order (bucket #, then overflow
// chain order, then bucket index).  To maintain iteration
// semantics, we never move keys within their bucket (if
// we did, keys might be returned 0 or 2 times).  When
// growing the table, iterators remain iterating through the
// old table and must check the new table if the bucket
// they are iterating through has been moved ("evacuated")
// to the new table.
```
