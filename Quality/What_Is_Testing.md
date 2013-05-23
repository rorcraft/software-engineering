https://www.udacity.com/wiki/cs258/unit_1

### What is testing

* Range of input, Range of output.
* finding bugs as early as possible
* more testing is not always better. In fact, the quality of testing is all about the cost/benefit tradeoff
* design software to be testable

### Bugs can be of many reasons

* bug in SUT
* bug in our acceptability check
* bug in our specification
* bug in our understanding

### Assertions in production

* A run-time check that tests a condition which, when false, reveals a bug in the software.
* A mechanism by which the program is halted (maybe after really minimal clean-up work).
* Skip assertions for mission critical programs when continuing with error is better than halting.
* http://docs.oracle.com/javase/1.4.2/docs/guide/lang/assert.html
* http://c2.com/cgi/wiki?ShipWithAssertionsOn

### Domains and Range

* interface that span trust boundaries are special must be test on the full range of representatives. e.g. GUI
* fault injection - faults that we want our code to be robust with respect to.
* consider time dependent testing

### Kinds of testing

* White box testing
* Black box testing

* Unit Testing
* Integration Testing
* Differential Testing
* System Testing
* Stress Testing
* Random Testing
