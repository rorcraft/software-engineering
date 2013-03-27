* The only way a team can write quality software is when the entire team is responsible for quality. ie. PM, Dev, Testers... everyone. Make testing a feature of the code base.

* "Engineering Productivity" has to overcome bias against testing and a company culture that favored heroic effort over engineering rigor.

* Testing must not create friction that slows down innovation and development.

* Google engineers prefer quality over features.

* Get it right from the beginning or you've created a permanent mess.

* "You build it, you break it"

* A large part of productivity is avoiding re-work because of sloppy development.

Roles:
SWE: fault tolerant designs, failure recovery, TDD, unit tests
SET: review designs, check code quality and risk. refactor code to make it more testable. Write testing frameworks and automation.
TE: user focused, product experts, quality advisers and analyzers of risk. automation scripts.
obvious bugs are signs of inadequate testing in earlier cycle. When such bugs are rare, TE turn to the primary task of ensuring that the software runs common user scenarios, meets performance expectataions, secure, i18n'd, accessible etc.
Coordination with contract testers, dogfooders, beta users etc.
Test director, managers.

SETs / TEs in Engineering Productivity

dev process centers around Code review:
Review designs, code quality and risk. Refactor code to make it more testable.
ensure the entire codebase has the look and feel of having been written by a single developer.

For test code, mindset is about breaking and writing code that will draw out cases that disrupt the user and his workflow.

common user scenarios
performance 
secure
internationalized
accessible

Testing is a feature = Testable and tested. 

Types of tests: small, medium, large

Small: mostly automated, exercise the code within signle function or module. 
Medium: "Does a set of near neighbor functions interoperate with each other the way they are supposed to?"
Large: "More results-driven, checking that the software satisfies user needs" "Does the product operate the way a user would expect and produce the desired results?"

Test sizes guarantee an upper bound on each test's execution time and resources.
< 1 min, < 5 mins, < 15 mins, < 1 hr

small: mocked
medium: db, localhost
large: everything.

70/20/10

"Quality is not important until the software is important"

Good SET: well documented, testable, working stable test automation and clear processes.

Automate as much as you can to keep manual testers always focusing on new issues.

Test code, the mindset is about breaking the software and writing code that will draw out cases that disrupt the user and his workflow.

Review design documents:

* completeness: identify parts that require special knowledge or not clear.
* correctness: grammer, spelling, punctuation mistakes.
* consistency: wording matchin diagrams, does not contradict claims from other documents.
* design: achievable? pitfalls, too complex? too simple?
* interfaces & protocals: are they standard?
* testing: testing hooks? can design be tweaked to make testing easier.

automation planning:
* Designs that seek to automate everything end-to-end all in one master test suite are generally a mistake.
* The larger an automation effort is, the harder it is to maintain and the more brittle it becomes as the systemm evovles. 

Single repository - little to relearn, available to any engineers, allow "20 percent contributors". Keeping it simple and uniform is a specific goal of Google platform. 
CI: submit queue, most large projects adopted this.
should provide the exact change that a test failed.

Test code is only useful in the context of being executable in a way that accelerates development.

Test suite:
randomize order of test execution.
test coverage
run tests in parallel on single machine - rewritten to drop dependencies on singular resources.

encourage testing:
OKRs, certification program (test coverage, test speeds, 0 flaky tests, automated smoke test).

TEs:
- weak points in software
- security, privacy, performance, reliability, usability, compatibility, globalization concerns
- primary user scenarios
- interoperate with other products
- how good are diagnostics, fall back
- usage metrics, user feedbacks

Test plans , hard to maintain
* ACC (Attribute Component Capability)
We want to test in the order of risk.
* frequency of failure and impact - relative risk
* analyze bugs using ACC

Exploratory testing with tours for crowd sourced testers, dogfooders.
Google Test Case Manager (GTGM)

Buganizer: triage, accountability, reports, search etc
Severity: how much it affects functionality.
Priority: how soon it should be fixed.

Type: Bug, Feature request, Customer issue, Internal Cleanup, Process

Teams simply stop adding features when incoming bug rate exceeds the team's ability to fix them.

Better testers ask above specifics - UX, race conditions, privacy, xss, seo, boundaries, i18n.

Playback framework (RPF) - recalculates XPaths of the elements they change.

The goal of engineer in google is to create impact.

Any product where user experience is important must undergo exploratory testing.

(Gmail TE) Put out a version to a few users and get some feedback before investing in a lot of test automation.


Roles in future:

SET == SWE, ownership of testing features should fall to new team members, particularly the more junior ones.

TE = Test design and coordination, identify resources required for testing.
