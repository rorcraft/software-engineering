* http://blog.libtorrent.org/2012/12/principles-of-high-performance-programs/
* http://blog.libtorrent.org/2012/10/asynchronous-disk-io/
* http://www.slideshare.net/bittorrentinc/performance-software


----------

http://www.crowl.org/lawrence/programming/Bentley82.html

A Summary of Jon Louis Bentley's
``Writing Efficient Programs''

Lawrence A. Crowl 29 August 1990

This summary of Jon Louis Bentley's book ``Writing Efficient Programs'' consists of selected text from that book with editorial changes, primarily removing references to specific examples.

© 1982 by Prentice-Hall, Inc., Englewood Cliffs, N.J. 07632 (ISBN 0-13-970251-2)
© 1990 by Lawrence A. Crowl, University of Rochester, Computer Science Department, Rochester, N.Y. 14627

## Methodology

The following steps are essential parts of a methodology of building efficient computing systems.
The most important properties of a large system are a __clean design and implementation, useful documentation, and a maintainable modularity.__

1. If the overall system performance is not satisfactory, then the programmer should monitor the program to identify where the scarce resources are being consumed.
This usually reveals that most of the time is used by a few percent of the code.
Proper data structure selection and algorithm design are often the key to large reductions in the running time of the expensive parts of the program.
1. If the performance of the critical parts is still unsatisfactory, then use the techniques below to recode them. __The original code should usually be left in the program as documentation.__
1. If additional speed is still needed, then the programmer should work at lower design levels, including hand-written assembly code, operating system modifications, microcode, and special-purpose hardware design.

### Applying the Rules

Once the hot spots of a system have been isolated, there are four critical steps in applying the following rules.
1. __Identify the code to be changed.__ We should identify the code to be changed by monitoring the program, and then concentrate on the parts that are taking the majority of the time.
1. __Choose a rule and apply it.__ Once we know what code to change, we see if any of our rules can be used to change it. The rules are presented in groups as they relate to different parts of programs; we should therefore identify whether the expense of our program is going to data structures, loops, logic, procedures or expressions, and then search in the appropriate list for a candidate rule. When we apply a rule, we should make sure that the application preserves the correctness of the program; this is usually done by applying the spirit, if not the actual formulas, of program verification.
1. __Measure the effect of the modification.__ Removing a common subexpression in a traveling salesman tour program is typical of many changes we make: it appears that it would increase the program's speed by a factor of two but in fact it gives less than a three precent improvement. Even if we believe that we understand the effect of a transformation by reasoning alone, it is absolutely necessary to support that analysis with observation; we often find that we are quite mistaken.
1. __Document the resulting program.__ The final modified program should include a description of the clean code and of the modification that was incorporated for efficiency. That description can range from a brief comment to including a copy of the original code enclosed within comment characters and a thorough discussion of the rule used to modify it (with an appropriate reference to this book, of course).

### Fundamental Rules

* __Code Simplification:__ Most fast programs are simple. Therefore, keep code simple to make it faster.
* __Problem Simplification:__ To increase efficiency of a program, simplify the problem it solves.
* __Relentless Suspicion:__ Question the necessity of each instruction in a time-critical piece of code and each field in a space-critical data structure.
* __Early Binding:__ Move work forward in time. Specifically, do work now just once in hope of avoiding doing it many times later.

### Space for Time Rules

Introducing redundant information can decrease run time at the cost of increasing space used. (The complements of these rules may apply.)
* __Data Structure Augmentation:__ The time required for common operations on data can often be reduced by augmenting the structure with extra information or by changing the information within the structure so that it can be accessed more easily. Examples are reference counters and hints.
* __Store Precomputed Results:__ The cost of recomputing an expensive function can be reduced by computing the function only once and storing the results. Subsequent requests for the function are then handled by table lookup rather than by computing the function.
* __Caching:__ Data that is accessed most often should be the cheapest to access. (Caching can ``backfire'' and increase the run time of a program if locality is not present in the underlying data.)
* __Lazy Evaluation:__ The strategy of never evaluating an item until it is needed avoids evaluations of unnecessary items.

### Time for Space Rules

Recomputing information can decrease a program's space requirements, at the cost of increased execution time. (The complements of these rules may apply.)
* __Packing:__ Dense storage representations can decrease storage costs by increasing the time required to store and retrieve data. (Decreasing data space may increase code space.)
* __Overlaying:__ Storing data items that are never simultaneously active in the same memory space reduces data space. Code overlaying reduces code space by using the same storage for routines that are never simultaneously needed. Many operating systems provide this service automatically in their virtual memory systems.
* __Interpreters:__ The space required to represent a program can often be decreased by the use of interpreters in which common sequences of operations are represented compactly. (Finite state machines are simple, compact interpreters.)

### Loop Rules

The efficiency rules used most frequently deal with loops for the simple reason that the hot spots in most programs involve loops.
* __Code Motion Out of Loops:__ Instead of performing a certain computation in each iteration of a loop, it is better to perform it only once, outside the loop. (Code cannot be moved out of loops if it has side effects that are desired on every iteration.)
* __Combining Tests:__ An efficient inner loop should contain as few tests as possible, and preferably only one. The programmer should therefore try to simulate some of the exit conditions of the loop by other exit conditions. Sentinels are a common application of this rule: we place a sentinel at the boundary of a data structure to reduce the cost of testing whether our search has exhausted the structure.
* __Loop Unrolling:__ A large cost of some short loops is in modifying the loop indices. That cost can often be reduced by unrolling the loop.
* __Transfer-Driven Loop Unrolling:__ If a large cost of an inner loop is devoted to trivial assignments, then those assignments can often be removed by repeating the code and changing the use of variables. Specifically, to remove the assignment I:=J, the subsequent code must treat J as though it were I.
* __Unconditional Branch Removal:__ A fast loop should contain no unconditional branches. An unconditional branch at the end of a loop can be removed by ``rotating'' the loop to have a conditional branch at the bottom.
* __Loop Fusion:__ If two nearby loops operate on the same set of elements, then combine their operational parts and use only one set of loop control operations.

### Logic Rules

These rules deal with efficiency problems that arise when evaluating the program state by making various kinds of tests. They sacrifice clarity and robustness for speed of execution.
* __Exploit Algebraic Identities:__ If the evaluation of a logical expression is costly, replace it by an algebraically equivalent expression that is cheaper to evaluate.
* __Short-Circuiting Monotone Functions:__ If we wish to test whether some monotone nondecreasing function of several variables is over a certain threshold, then we need not evaluate any of the variables once the threshold has been reached. Short-circuit evaluation of boolean expressions is an example of this rule. (A more complex application of this rule exits from a loop as soon as the purpose of the loop has been accomplished.)
* __Reordering Tests:__ Logical tests should be arranged such that inexpensive and often successful tests precede expensive and rarely successful tests.
* __Precompute Logical Functions:__ A logical function over a small finite domain can be replaced by a lookup in a table that represents the domain.
* __Boolean Variable Elimination:__ We can remove boolean variables from a program by replacing the assignment to a boolean variable V by an if-then-else statement in which one branch represents the case that V is true and the other represents the case that V is false. (This generalizes to case statements and other logical control structures.) This rule usually decreases time slightly (say, less than 25 percent), but greatly increases code space. More complex applications of this rule remove boolean variables from data structures by keeping separate structures for the true and false records.

### Procedure Rules

Procedures form good abstraction boundaries, but a given boundary may have an unacceptable execution cost.
* __Collapsing Procedure Hierarchies:__ The run times of the elements of a set of procedures that (nonrecursively) call themselves can often be reduced by rewriting procedures in line and binding the passed variables.
* __Exploit Common Cases:__ Procedures should be organized to handle all cases correctly and common cases efficiently.
* __Coroutines:__ A multiple-pass algorithm can often be turned into a single-pass algorithm by use of coroutines.
* __Transformations on Recursive Procedures:__ The run time of recursive procedures can often be reduced by applying the following transformations:
* __Code the recursion explicitly by use of a program stack.__ 
  * If the final action of a procedure P is to call itself recursively, replace that call by a goto to its first statement; this is usually known as removing tail recursion. That goto can often be transformed into a loop.
  * If a procedure contains only one recursive call on itself, then it is not necessary to store the return address on the stack.
  * It is often more efficient to solve small subproblems by use of an auxiliary procedure, rather than by recurring down to problems of size zero or one.

### Parallelism:

A program should be structured to exploit as much of the parallelism as possible in the underlying hardware.

### Expression Rules

Be careful that applying these rules does not work against your compiler.

* __Compile-Time Initialization:__ As many variables as possible should be initialized before program execution.
* __Exploit Algebraic Identities:__ If the evaluation of an expression is costly, replace it by an algebraically equivalent expression that is cheaper to evaluate.
  An algebraic identity yields a fast range test that compiler writers can use on two's complement architectures.
* We can often multiply or divide by powers of two by shifting left or right.
* Strength reduction on a loop that iterates through the elements of an array replaces a multiplication by an addition. This technique generalizes to a large class of incremental algorithms.
* __Common Subexpression Elimination:__ If the same expression is evaluated twice with none of its variables altered between evaluations, then the second evaluation can be avoided by storing the result of the first and using that in place of the second. (We cannot eliminate the common evaluation of an expression with important side-effects.)
* __Pairing Computation:__ If two similar expressions are frequently evaluated together, then we should make a new procedure that evaluates them as a pair. Examples include sine/cosine and minimum/maximum.
* __Exploit Word Parallelism:__ Use the full word width of the underlying computer architecture to evaluate expensive expressions. The bitwise representation of sets exploits this rule.

### System Dependent Optimization

Some optimizations require an understanding of their effect on the underlying system.

* __The High-Level Language:__ What are the relative costs of the various constructs in the language? What are the cheapest possible ways to accomplish common operations (i.e., input and output, or iterating through the elements of a vector)?
* __The Compiler:__ What optimizations does the compiler perform, and what optimization should the programmer perform? Which language constructs encourage the compiler to perform optimizations, and which constructs block optimizations? Current optimizing compilers can move code out of loops, unroll loops, unroll transfer-driven loops, remove unconditional branches, exploit algebraic identities, short-circuit simple monotone functions, collapse procedure hierarchies, eliminate tail recursion, and eliminate common subexpressions.
* __The Machine Architecture:__ Which efficient operations in the underlying hardware are easy to access, and how can the compiler give access to them? Is there an instruction buffering scheme that makes certain instruction organizations more efficient than others? What paging or caching schemes does the machine support, and how can the program's locality structure be organized to exploit the schemes? Issues include the following.
* __Instruction Costs:__ Different instructions may have different relative costs for different implementations of an architecture.
* __Register Allocation:__ Storing a variable in a register increases a program's efficiency in two ways. The running time is usually decreased because registers are usually realized by a faster technology than main memory and are architecturally (and usually physically) closer to the arithmetic and logical unit. A variable in a register also requires fewer bits to address; this means that the instruction is faster to access and that the overall program will require less space.
* __Instruction Caching:__ Many high-speed computers have a cache for the recently accessed instructions; if a tight loop fits in that cache then it will run more quickly than if the instructions must be fetched repeatedly from main memory. Unrolling a loop such that it exceeds the size of the cache may reduce performance.
* __Storage Layout:__ The layout of a program's storage in memory can have a significant effect on the execution speed of the program. Placing records of a length that is a power of two in arrays enables the multiplication of the array index by a left shift (which is often cheaper than multiplication by an arbitrary integer). However, banked memories are often implemented in powers of two, which interferes badly with this approach.
* __Virtual Memory Caching:__ Knowledge of the interaction of the paging structure and the storage layout can have a profound effect on efficiency. Traverse row major matrices by row rather than column.

Lawrence A. Crowl, crowl@cs.orst.edu, 14 September 1995
