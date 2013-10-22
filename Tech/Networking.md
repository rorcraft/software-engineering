http://chimera.labs.oreilly.com/books/1230000000545/

### Latency

* Propagation delay: distance, equatorial light in fiber 200ms (RTT)
* Transmission delay: data length, bandwidth
* Processing delay: checking packet header, error, destination
* Queuing delay: waiting in queue (Bufferbloat)

### TCP/UDP  -> IP -> Hardware

IP (OSI3)
* unreliable, connectionless
* header - source and destination for routing
* supplies a checksum that includes its own header

UDP (OSI4)
* unreliable and connectionless
* checksum of content and port numbers

TCP
* reliable and connection oriented
* IPv4 32 bit integer
* Class A  - first byte network, last 3 bytes device. 128 Class A.
* Class B - first 2 bytes network, last 2 bytes device. 2^16 devices.
* Class C - first 3 bytes network, last byte device. 2^8 devices.
* IPv6 128 bit address. hexadecimal notation.

SCTP ?

Ports
* Telnet 23
* DNS 53
* FTP 21, 20
* HTTP 80, 8000, 8080, 8088
* XWindow 6000-6007
/etc/services

IPv6 with port '[::1]:23' , ':80'

__SPDY__

http://www.ietf.org/proceedings/83/slides/slides-83-httpbis-3

__HTTP2__

http://www.igvita.com/slides/2012/http2-spdy-devconf.pdf
