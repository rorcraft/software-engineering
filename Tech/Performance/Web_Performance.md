### Web Applications Performance

* 0 latency = do nothing. (do as little as possible)
* Measure RUM, 95%, 75%, compare between locations, mobile vs desktop.

```
 ___________
|  Browser  |
| --------- |
|  Network  |
| --------- |
|  App      |
| --------- |
|  Storage  |
|___________|

```

__What is fast enough?__

<table border="1" class="delay" width="100%">
<tbody><tr>
<td class="header">Delay</td>
<td class="header">User Reaction</td>
</tr>
<tr>
<td class="green">0 - 100ms</td>
<td>Instant</td>
</tr>
<tr>
<td>100 - 300ms</td>
<td>Small perceptible delay</td>
</tr>
<tr>
<td>300 - 1000ms</td>
<td>Machine is working</td>
</tr>
<tr>
<td class="red">1s+</td>
<td>Mental context switch</td>
</tr>
<tr>
<td class="red">10s+</td>
<td>I'll come back later...</td>
</tr>
</tbody></table>

__Browser__

  * HTML - DOM (Bytes -> Characters -> Tokens -> Nodes -> Dom)
  * Sync Javascript will block render of page. (defer, async)
  * CSS tree
  * Executing javascript - V8
  * render (paint and composite) into rasterized images - blink (multithreaded)
  * Rails4 - Turbolinks - ajax replace content to avoid reparsing css, js.
  * Chrome does a lot of prediction.
  * HTTP cache headers (we're not doing them)
  * LocalStorage
 
__Network__

* Light speed (locality)
* DNS
* TCP 
  - Ack connection
  - CWND
  - MTU
* Bufferbloat
* Mobile 3G, HSPA, LTE, Wifi 802.11n/ac
* DosArrest - ip filtering.
* SSL
* HTTP (watch out for redirect)
* websocket less overhead than ajax.
* App data access sequence.
* SOA/RPC protocol
  * HTTP
  * Protobuf / MessagePack
  * Typhoeus - libcurl - reuse TCP connection.
* chunked encoding (facebook BigPipe, PageSpeed)
* SPDY
  * Over SSL
  * binary protocol
  * multiplexed (warm)
* HTTP 2.0
* QUIC (over UDP)
  * fast open (remove connect RTT)
  * snapstart 
* chrome://net-internals/ 
* chrome://flags/

__Application__

* Garbage collection - parallel, stop the world, generational. (95%tiles)
* Cache fill race condition (95%tile)
* Generating new objects - hash.merge(k: v)
* less mem = less data to copy to CPU registers.
* modular code for JIT optimization.
* balance between too much methods (extra call pointers)
* meta programming, usually bad 
* dynamic programming / memoization.
* concurrency - contention (ConcurrentHashMap, message passing)
* code path, data access sequence
* non blocking vs threaded (watch out for synchronised cpu intensive code)
* CPU - IPC, instructions per cycle
* bit hacking.

__Data__

* MySQL 
  * connection pool
  * sql cache (by hash)
  * parse (prepared statements?)
  * optimizer
  * index
  * fetch data (unnecessary data)
  * protocol (binary?)
  * stored procedures
  * contention / locking
  * storage engine - innodb
  * fetch from disk (RAID, disk buffer, file system (journalling mode), SSD)
  * true sequential read (patch from Facebook)
* Memcache
  * distributed, hashed.
  * multi get
* Redis
  * connection pool
  * pipeline
  * Lua scripts
  
__Data structures__

* Mysql Index - B+tree
* Redis Hash (limiting size can optimize memory)
* Redis ordered sets - skip list
* trie - autocomplete, Rails routes
* Ruby Hash - doubling linked list, thats why they are ordered but also why slower than Python Dict.

__All together__

  * Locality
  * Warm - pipe, cache
  * Compact data - protocol, storage, retrieval
  * Respond early to start download
  * fetch data from memory, data structure efficiency
  * Limit necessary JS/CSS for initial rendering.
 
__References__

* [Blink’s Rendering Pipeline](https://docs.google.com/a/change.org/document/d/1wYNK2q_8vQuhVSWyUHZMVPGELzI0CYJ07gTPWP1V1us/pub)
* http://www.igvita.com/2013/01/15/faster-websites-crash-course-on-web-performance/
* http://www.igvita.com/posa/high-performance-networking-in-google-chrome/
* http://www.chromium.org/spdy/
* Bufferbloat - http://queue.acm.org/detail.cfm?id=2209336
* http://chimera.labs.oreilly.com/books/1230000000545
