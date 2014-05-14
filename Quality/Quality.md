### Speed AND Quality, not conflicting.
http://www.slideshare.net/lazygolfer/testing-does-not-equal-quality

* Quality is the foundation of Value
* Quality is the foundation of Speed
* Accelerate the speed of ensuring high quality
  * Quality up front in planning
  * Speed up Continuous Integration, Automated Testing.
  * Fail fast with Test eng, dogfood.

### Quality != Test



### Tested = Validated + Explored

* Explore early, explore often


### Quality is not a tradeoff

https://medium.com/the-year-of-the-looking-glass/bcddf7c85553

> __"Quality is a bar, not a tradeoff."__

> Even if I told them it didn’t matter if the registration page was any good because we have proof—real, honest-to-god proof—that 99.99% of users suffered from an affliction known as “registration-blindness” and it didn’t matter what the page looked like, just so long as we had one, so please just give the engineer a mock to build so we can ship this thing and for heaven’s sake the quality doesn’t matter we just need to get this out ASAPASAPASAP.

> It just doesn’t work like that.

> Because to create high-quality work, there has to be a minimum acceptable bar.

> Building something that demonstrates craft at the highest level cannot be reasoned into. It happens because of love, and because there was an environment that nurtured that love.

http://dandemeyere.com/blog/5-most-inspiring-steve-jobs-stories

### Qualities of Quality

http://labs.spotify.com/2014/04/11/qualities-of-quality/ 

> Developer-facing quality is a completely different thing from end-user facing quality, and is usually more important.

![](http://spotifylabscom.files.wordpress.com/2014/04/kindsofquality.png?w=584)

http://martinfowler.com/bliki/DesignStaminaHypothesis.html

![](http://martinfowler.com/bliki/images/designStaminaGraph.gif)

Technical Debt https://www.youtube.com/watch?v=pqeJFYwnkjE

You can only profitably trade off implementation quality for speed in very short-lived projects. Probably, if you’re expecting a system to have a lifespan of more than a couple of weeks, it’s a good idea to pay attention to implementation quality right from the start.

http://tech.shopzilla.com/2009/11/the-cost-of-a-feature/
* Finished features lead to development cost even if you don’t touch them. 
* Not just in operating or maintaining that feature, but in development of other features.
  * Slower decision-making/added communication due to the larger amount of things to know and think about, both when defining features and implementing them. In the same vein, increased likelihood of mistakes in feature specifications or implementations due to the complexity of feature interaction.
  * Slower implementation of new features due to the additional complexity of the system. It is harder to figure out the right way to fit things together if there are many of them.
  * Regression testing costs that increase with every added feature.

Bugs reduce productivity in many ways: http://pettermahlen.com/2011/04/08/if-its-broken-fix-it/
The short of it is that unfixed bugs in your code lead to additional meetings, bug management overhead, duplicate reports of the same bug and context switching. So a lot of the time, the best thing you can do for your own productivity is to just fix pretty much everything that’s ever reported.

One common misconception about quality and its impact on delivery speed is that things like pluggability/extensibility/configurability of some technical solution are quality. Those are things often labelled as over-engineering. To me, over-engineering is engineers adding waste by inventing features that aren’t actually really needed.
