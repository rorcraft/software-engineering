## Index Construction

### Hardware
* ram - nano seconds
* disk - micro seconds
* OS read/write entire blocks - 32, 64KB
* total time of reading + decompressing is usually less than reading uncompressed data.

### Blocked sort-based indexing - BSBI

Basic steps
1. enumrate collection creating term-docId pairs.
2. sort pairs with term as primary key
3. merge term's doc-id as posting lists, compute stats like term freq, doc freq.

termIDs - instead of string, use number.

external sorting algorithm - uses disk

BSBI - read / sort by blocks, then merge intermediate results into final index.

### Single-pass in-memory indexing - SPIMI

* terms instead of termID
* write each block's dictionary to disk
* faster - no sorting required
* saves memory - keep track of the term a posting list belongs to

### Distributed Index

* map splits of input data to key-value pairs. (parsers), writes to segment files
  list(termId, docId)
* reduce - all values for given key to be stored close together.
  termId, list(docId), -> (postings_list1, postings_list2)

### Dynamic indexing

* auxiliary index to track new changes
* searches run across both indexes and results merged.
* deletions stored in an invalidation bit vector.
* merge auxiliary index when it becomes too large
* logarithmic merging.
* too complicated to maintain, most often rebuild from scratch periodically.


