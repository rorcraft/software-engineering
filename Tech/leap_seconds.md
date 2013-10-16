TOREAD

* http://www.ucolick.org/~sla/leapsecs/onlinebib.html
* http://en.wikipedia.org/wiki/Logical_clock

http://aphyr.com/posts/299-the-trouble-with-timestamps

POSIX time itself is not monotonic.
POSIX days are defined as 86400 seconds in length. However, real days aren’t exactly 86400 seconds. The powers-that-be occasionally schedule leap seconds to correct for the drift. On those occasions, the system clock will either skip a second, or double-count a second–e.g., __counting 59:60.7, 59:60.8, 59:60.9, 59:60.0, 59:60.1__, and then repeating the previous second’s worth of timestamps before continuing on

### Google's Leap Smear:

http://googleblog.blogspot.com/2011/09/time-technology-and-leaping-seconds.html

Usually when a leap second is almost due, the NTP protocol says a server must indicate this to its clients by setting the “Leap Indicator” (LI) field in its response. This indicates that the last minute of that day will have 61 seconds, or 59 seconds. (Leap seconds can, in theory, be used to shorten a day too, although that hasn’t happened to date.) Rather than doing this, we applied a patch to the NTP server software on our internal Stratum 2 NTP servers to not set LI, and tell a small “lie” about the time, modulating this “lie” over a time window w before midnight:
```lie(t) = (1.0 - cos(pi * t / w)) / 2.0```
What this did was make sure that the “lie” we were telling our servers about the time wouldn’t trigger any undesirable behavior in the NTP clients, such as causing them to suspect the time servers to be wrong and applying local corrections themselves.
