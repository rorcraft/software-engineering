[Kent Beck on TDD @ Quora](https://www.quora.com/In-Test-Driven-Development-how-do-unit-tests-help-drive-good-design/answer/Kent-Beck?srid=Vih&share=1)

TDD doesn't drive good design. TDD gives you immediate feedback about what is likely to be bad design. If a test is hard to write, if a test is non-deterministic, if a test is slow, then something is wrong with the design. When I'm not ready to make good design decisions, I still don't end up with a good design, but I certainly know I have room for improvement.

[Is TDD Dead?](https://www.youtube.com/watch?v=z9quxZsLcfo) DHH, Kent Beck, Martin Fowler.

Kent - "I don't go very far down the mocking path".

TDD != Continuous Integration != Fully Isolated Tests.

Fast feedback loop is the goal. Whichever way achieves that goal depends on the problem/trade offs.

Fast feedback loop means:
* quicker to find what is likely to be bad design.
* drives better design because of confidence in refactoring and experimentation.

If whatever you're doing is slowing down the feedback loop, thats the wrong.
