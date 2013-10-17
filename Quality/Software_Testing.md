https://www.udacity.com/wiki/cs258/

## Domains, Ranges, Oracles, and Kinds of Testing

* Always require a source of test inputs which result in test outputs. 
* The quality of testing is all about the cost/benefit tradeoff
* Quality cannot be tested into software

__Assertion (defensive programming)__

Code that we write fail proactively when something's wrong, and this can lead to more effective testing.

Trade off between mission critical or fail early and recover.

NASA rocket - that is to say for the launch phase, for the cruise phase, for orbiting Mars-- they had plenty of assertions enabled in the actual mission software-- that is to say, in the software running on the spacecraft. On the other hand, for a short period of time during which the spacecraft was actually landing they disabled all of the assertions on the idea that they got only one shot to land on the planet and that if anything went wrong with the software during the landing process it was better to just try to keep going, because an assertion violation leading to a system reset could cause the lander to become unstable in such a way that it wouldn't have been recoverable.

__Specification__ Often when we're testing software, we're not so much just looking for bugs in the software, but we're helping to refine the specification for the software under test.

Domain (inputs) => Range (outputs)

Trust Boundaries - The principle is that if you have an interface or an API that represents a trust boundary, that is a boundary between part of the system that we trust and users on the other side who we can't trust to stay within the domain of the API that we're implementing, then we have to test that API with all possible representative values, not just ones that the developers happen to think are in the domain

Fault injection - faults that we want our code to be robust with respect to.

Take timing of input into account. Usually software under test that interacts directly with hardware devices.

White box testing - using detailed knowledge about the internals of the system in order to construct better test cases.

Black box testig - rather testing the system based on what we know about how it's supposed to respond to our test cases

## Coverage Testing

Partitioning the input domain
* Partition into number of different classes. For purposes of finding defects in the system under test, any point within that subdomain is as good as any other point within that subdomain.
(this sort of scheme hasn’t worked out for large systems under test)

__Prefer Coverage over Partitioning__
Test coverage is an automatic way of partitioning the input domain.

* Line coverage
* Statement coverage
* Branch coverage
* Loop coverage
* Path coverage
* Boundary value coverage

Used in the right way, coverage can be a relatively low cost way to improve the testing that we do for a piece of software. Used incorrectly, it can waste our time, and perhaps worst, lead to a false sense of security.

## Random Testing

* should be a separate test suite.
* should generate mostly valid input

## Fuzzing

