### TCP Fast Open

https://lwn.net/Articles/508865/

__Eliminating a round trip__

* does not detect duplicate SYN segments.
* must be idempotent—they must tolerate the possibility of receiving duplicate initial SYN segments containing the same data and produce the same result regardless of whether one or multiple such SYN segments arrive.
* to prevent malicious DDOS attacks, TFO employes a cookie, that is generated once by the server TCP and returned to the client TCP for later reuse.
* The cookie is constructed by encrypting the client IP address in a fashion that is reproducible (by the server TCP) but is difficult for an attacker to guess. Request, generation, and exchange of the TFO cookie happens entirely transparently to the application layer.
* For subsequent conversations with the server, the client can short circuit the three-way handshake

1. The client TCP sends a SYN that contains both the TFO cookie (specified as a TCP option) and data from the client application.
2. The server TCP validates the TFO cookie by duplicating the encryption process based on the source IP address of the new SYN. If the cookie proves to be valid, then the server TCP can be confident that this SYN comes from the address it claims to come from. This means that the server TCP can immediately pass the application data to the server application.
3. From here on, the TCP conversation proceeds as normal: the server TCP sends a SYN-ACK segment to the client, which the client TCP then acknowledges, thus completing the three-way handshake. The server TCP can also send response data segments to the client TCP before it receives the client's ACK.

If the TFO cookie proves not to be valid, then the server TCP discards the data and sends a segment to the client TCP that acknowledges just the SYN. At this point, the TCP conversation falls back to the normal three-way handshake. If the client TCP is authentic (not malicious), then it will (transparently to the application) retransmit the data that it sent in the SYN segment.

To prevent resource-exhaustion attacks, the server imposes a limit on the number of pending TFO connections that have not yet completed the three-way handshake. When this limit is exceeded, the server ignores TFO cookies and falls back to the normal three-way handshake

> The client-side support has been merged for Linux 3.6. However, the server-side TFO support has not so far been merged, and from conversations with the developers it appears that this support won't be added in the current merge window. Thus, an operational TFO implementation is likely to become available only in Linux 3.7.

