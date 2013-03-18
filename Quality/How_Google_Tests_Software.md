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
TE: product experts, quality advisers and analyzers of risk.

Code review:
Review designs, code quality and risk. Refactor code to make it more testable.

For test code, mindset is about breaking and writing code that will draw out cases that disrupt the user and his workflow.

common user scenarios
performance 
secure
internationalized
accessible


Types of tests: small, medium, large

Small: mostly automated, exercise the code within signle function or module. 
Medium: "Does a set of near neighbor functions interoperate with each other the way they are supposed to?"
Large: "More results-driven, checking that the software satisfies user needs" "Does the product operate the way a user would expect and produce the desired results?"

Single repository - little to relearn, available to any engineers, allow "20 percent contributors". Keeping it simple and uniform is a specific goal of Google platform. 

"Quality is not important until the software is important"

Good SET: well documented, testable, working stable test automation and clear processes.

