### Hash functions

* Murmur
* CityHash
* SipHash

### Dynamic Sizing


__Distributed Hash Table__

http://engineering.bittorrent.com/2013/01/22/bittorrent-tech-talks-dht/

__Redis__

http://stackoverflow.com/questions/10004565/redis-10x-more-memory-usage-than-data

```
# Hashes are encoded in a special way (much more memory efficient) when they
# have at max a given number of elements, and the biggest element does not
# exceed a given threshold. You can configure this limits with the following
# configuration directives.
hash-max-zipmap-entries 512
hash-max-zipmap-value 64
```
