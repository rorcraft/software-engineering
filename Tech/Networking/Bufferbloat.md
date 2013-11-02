__Bufferbloat in Your Local Router__

Bufferbloat is a term that was coined and popularized by Jim Gettys in 2010, and is a great example of queuing delay affecting the overall performance of the network.

The underlying problem is that many routers are now shipping with large incoming buffers under the assumption that dropping packets should be avoided at all costs. However, this breaks TCP’s congestion avoidance mechanisms (which we will cover in the next chapter), and introduces high and variable latency delays into the network.

The good news is that the new __CoDel active queue management algorithm__ has been proposed to address this problem, and is now implemented within the Linux 3.5+ kernels. To learn more, refer to "Controlling Queue Delay" in ACM Queue.

http://queue.acm.org/detail.cfm?id=2209336
