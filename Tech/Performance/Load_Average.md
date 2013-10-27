### Load Averages

`top - 21:35:08 up 8 days, 16:22,  2 users,  load average: 1.11, 1.14, 1.16`
1-minute, 5-minute and 15-minute average.

also same as - `cat /proc/loadavg`

http://www.linuxjournal.com/article/9001?page=0,1
> In the Linux kernel, each dispatchable process is given a fixed amount of time on the CPU per dispatch. By default, this amount is 10 milliseconds, or 1/100th of a second. For that short time span, the process is assigned a physical CPU on which to run its instructions and allowed to take over that processor. More often than not, the process will give up control before the 10ms are up through socket calls, I/O calls or calls back to the kernel. (On an Intel 2.6GHz processor, 10ms is enough time for approximately 50-million instructions to occur. That's more than enough processing time for most application cycles.) If the process uses its fully allotted CPU time of 10ms, an interrupt is raised by the hardware, and the kernel regains control from the process. The kernel then promptly penalizes the process for being such a hog. As you can see, that time slicing is an important design concept for making your system seem to run smoothly on the outside. It also is the vehicle that produces the load-average values.
> When the kernel takes back control, it increment its jiffies counter. When the quantum timer pops, timer.c is entered at a function in the kernel called timer.c:do_timer(). Here, all interrupts are disabled so the code is not working with moving targets. The jiffies counter is incremented by 1, and the load-average calculation is checked to see if it should be computed. In actuality, the load-average computation is not truly calculated on each quantum tick, but driven by a variable value that is based on the HZ frequency setting and tested on each quantum tick. (HZ is not to be confused with the processor's MHz rating. This variable sets the pulse rate of particular Linux kernel activity and 1HZ equals one quantum or 10ms by default.) Although the HZ value can be configured in some versions of the kernel, it is normally set to 100. The calculation code uses the HZ value to determine the calculation frequency. Specifically, the timer.c:calc_load() function will run the averaging algorithm every 5 * HZ, or roughly every five seconds. Following is that function in its entirety:

```c
   unsigned long avenrun[3];

   static inline void calc_load(unsigned long ticks)
   {
      unsigned long active_tasks; /* fixed-point */
      static int count = LOAD_FREQ;

      count -= ticks;
      if (count < 0) {
         count += LOAD_FREQ;
         active_tasks = count_active_tasks();
         CALC_LOAD(avenrun[0], EXP_1, active_tasks);
         CALC_LOAD(avenrun[1], EXP_5, active_tasks);
         CALC_LOAD(avenrun[2], EXP_15, active_tasks);
      }
   }
```

sched.h, a header used by much of the kernel code. In there, the CALC_LOAD macro and its associated values are available:

```c
   extern unsigned long avenrun[];	/* Load averages */

   #define FSHIFT   11		/* nr of bits of precision */
   #define FIXED_1  (1<<FSHIFT)	/* 1.0 as fixed-point */
   #define LOAD_FREQ (5*HZ)	/* 5 sec intervals */
   #define EXP_1  1884		/* 1/exp(5sec/1min) as fixed-point */
   #define EXP_5  2014		/* 1/exp(5sec/5min) */
   #define EXP_15 2037		/* 1/exp(5sec/15min) */

   #define CALC_LOAD(load,exp,n) \
      load *= exp; \
      load += n*(FIXED_1-exp); \
      load >>= FSHIFT;
```

> The macro takes in three parameters: the load-average bucket (one of the three elements in avenrun[]), a constant exponent and the number of running/uninterruptible processes currently on the run queue. The possible exponent constants are listed above: EXP_1 for the 1-minute average, EXP_5 for the 5-minute average and EXP_15 for the 15-minute average. The important point to notice is that the value decreases with age.

> When x=1, then y=1884; when x=5, then y=2014; and when x=15, then y=2037. The purpose of the magical numbers is that it allows the CALC_LOAD macro to use precision fixed-point representation of fractions. The magic numbers are then nothing more than multipliers used against the running load average to make it a moving average. 

> As the load average goes above the number of physical CPUs, the more the CPU is being used and the more demand there is for it. And, as it recedes, the less of a demand there is. With this understanding, the load average can be used with the CPU percentage to obtain a more accurate view of CPU activity.
