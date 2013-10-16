# TOREAD
* http://www.ucolick.org/~sla/leapsecs/onlinebib.html
* http://en.wikipedia.org/wiki/Logical_clock

http://aphyr.com/posts/299-the-trouble-with-timestamps

POSIX time itself is not monotonic.
POSIX days are defined as 86400 seconds in length. However, real days aren’t exactly 86400 seconds. The powers-that-be occasionally schedule leap seconds to correct for the drift. On those occasions, the system clock will either skip a second, or double-count a second–e.g., __counting 59:60.7, 59:60.8, 59:60.9, 59:60.0, 59:60.1__, and then repeating the previous second’s worth of timestamps before continuing on
