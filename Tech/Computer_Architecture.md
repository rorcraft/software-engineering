http://www.yosefk.com/blog/how-fpgas-work-and-why-youll-buy-one.html


__What Every Programmer Should Know About Memory__
* http://www.akkadia.org/drepper/cpumemory.pdf

__L1 cache__ is very small and very tightly bound to the actual processing units of the CPU, it can typically fulfil data requests within 3 CPU clock ticks. L1 cache tends to be around 4-32KB depending on CPU architecture and is split between instruction and data caches.

__L2 cache__ is generally larger but a bit slower and is generally tied to a CPU core. Recent processors tend to have 512KB of cache per core and this cache has no distinction between instruction and data caches, it is a unified cache. I believe the response time for in-cache data is typically under 20 CPU "ticks"

__L3 cache__ tends to be shared by all the cores present on the CPU and is much larger and slower again, but it is still a lot faster than going to main memory. L3 cache tends to be of the order of 4-8MB these days.

__Static RAM__, a form of flip-flop holds each bit of memory. A flip-flop for a memory cell takes 4 or 6 transistors along with some wiring, but never has to be refreshed. This makes static RAM significantly faster than dynamic RAM. However, because it has more parts, a static memory cell takes a lot more space on a chip than a dynamic memory cell. Therefore you get less memory per chip, and that makes static RAM a lot more expensive.

__Dynamic RAM__, each memory cell holds one bit of information and is made up of two parts: a transistor and a capacitor. The transistor acts as a switch that lets the control circuitry on the memory chip read the capacitor or change its state.

This refresh operation is where dynamic RAM gets its name. Dynamic RAM has to be dynamically refreshed all of the time or it forgets what it is holding. The downside of all of this refreshing is that it takes time and slows down the memory.
