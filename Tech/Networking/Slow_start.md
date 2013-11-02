### Slow start

http://www.cdnplanet.com/blog/tune-tcp-initcwnd-for-optimum-performance/

* https://developers.google.com/speed/protocols/tcpm-IW10
* http://www.ietf.org/proceedings/80/slides/iccrg-1.pdf

![](http://st.cdnplanet.com/static/uploads/images/TCP1.png)

Step 1: Client sends SYN to server - "How are you? My receive window is 65,535 bytes."
Step 2: Server sends SYN, ACK - "Great! How are you? My receive window is 4,236 bytes"
Step 3: Client sends ACK, SEQ - "Great as well... Please send me the webpage http://www.example.com/"
Step 4: Server sends 3 data packets. Roughly 4 - 4.3 kb (3*MSS1) of data
Step 5: Client acknowledges the segment (sends ACK)
Step 6: Server sends the remaining bytes to the client

Google proposed changing CWND from between 2 and 4 segments, as specified in RFC 3390, to 10 segments (roughly 15KB)

```
      The table below compares the number of round trips between IW=3
       and IW=10 for different transfer sizes, assuming infinite
       bandwidth, no packet loss, and the standard delayed acks with
       large delay-ack timer.

             ---------------------------------------
            | total segments |   IW=3   |   IW=10   |
             ---------------------------------------
            |         3      |     1    |      1    |
            |        10      |     3    |      1    |
            |        12      |     3    |      2    |
            |        21      |     4    |      2    |
            |        25      |     5    |      2    |
            |        32      |     5    |      3    |
            |        46      |     6    |      3    |
            |        51      |     6    |      4    |
            |        79      |     7    |      4    |
            |       121      |     8    |      5    |
            |       128      |     9    |      5    |
             ---------------------------------------
```

2 parameters:

* client (receivers) advertised window size (RWIN) 
* initcwnd setting on the server

Some Operating Systems dynamically calculate RWIN based on external factors. The value here is based on SYN packets sent to CDN Planet. The Win flag can also be increased by the client before the transfer actually starts.

__Result__

![](http://st.cdnplanet.com/static/uploads/images/TCP2.png)
