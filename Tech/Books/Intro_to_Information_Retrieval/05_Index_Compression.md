## Index Compression

* compression ratios of 1:4 are easy to achieve.
* increased use of caching
* for improved cache utilization and faster disk-to-memory transfer, decompression speeds must be high
* posting = docID in a postings list.

### Statistical properties of terms in IR

* preprocessing affects the size of the dictionary and number of nonpositional postings greatly.
* stemming and case folding reduce number of distinct terms by 17%
* rule of 30 - 30 most common words account for 30% of tokens in written text.
* Oxford English Dictionary has 600k words
* collections includes names, locations, products

Heaps' law: Estimating the number of terms.
Term size as function of collection size:
```
M = kT^b
30 <= k <= 100
b ~= 0.5
```
* the dictionary size continues to increase with more documents in collection rather than 
  reaching a maximum
* the size of dictionary is quite large for large collections.

Zipf's law: Modeling the distribution of terms
* if t1 is the most common term, t2 is the next most common, and so on, then collection frequency
  cfi of ith most common term is proportional to 1/i:
* if t1 occurs cf1 times, then t2 is 1/2 * cf1 times.

### Dictionary compression

Dictionary as a string
* fixed-width entries, sorted lexicographically.
  20 bytes per term, 4 bytes per doc freq, 4 bytes per pointer to postings list.
* one long string, saves 60% compared to fixed-width. Use pointers to mark the end 
  of preceding term. freq, posting ptr, term ptr -> position of term in the one long string.
* blocked storage. Keep a term pointer only for first term of each block. 
  Use 1 byte to store length of string at beginning of term.
  Eliminates the k-1 term pointers.
* front coding (common prefixes) - e.g.:
  8automata8automate9automatic10automation (k=4 block)
  `8automat*a1<>e2<>ic3<>ion` (front coding)

### Postings file compression

* gaps between postings are short
* rare terms have large gaps
* use variable encoding

bytewise compression
* integral number of bytes to encode a gap
* first bit of byte is continuation bit, last 7 bits are payload.

bitwise compression
* simplet bit-level code is unary code. string of n 1s followed by a 0.
* gamma codes implement variabl-lenght encoding by splitting the representation of a gap G
  into a pair of length and offset.
* without knowing distribution of gaps, we can apply gammer codes and be certain that they
  ar within a factor of ~ 2 of optimal code for distributions of large entropy.
  i.e. universal code.
* prefix free
* parameter free
* expensive to decode than for variable byte codes.
