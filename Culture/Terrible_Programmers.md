### Good vs Bad

> The difference between a bad programmer and a good programmer is understanding. That is, bad programmers don't understand what they are doing, and good programmers do.

Code Simplicity: The Science of Software Development by Max Kanat-Alexander

> Bad programmers worry about the code. Good programmers worry about data structures and their relationships.

Linus Torvald

### Terrible Programmers

http://www.aaronstannard.com/post/2013/12/19/The-Taxonomy-of-Terrible-Programmers.aspx

The Pet Technologist

My personal favorite.
A pet technologist is born when they make the fatal mistake of falling in love with a piece of technology. It’s not a gentle, appreciative, “man this is a well-designed framework” sort of love – it’s inseparable, unrequited, Misery-obsessive love. And just like with spying, falling in love is a liability in our business.
No matter what the question is, you can trust that the pet technologist will have an answer: his or her pet technology.
“Hey, we need to implement a content management system in Rails, which database should we use?” Mongo.
“Multi-tenant blog engine?” Mongo.
“Business-critical compliance system?” Mongo.
“Inventory management system?” Mongo.
“Electronic medical records system?” Mongo.
“Distributed data warehouse?” Mongo.
And so forth.
They will invent reasons to include their pet technology in any project you work on, regardless of whether there’s a practical reason for it or not. They will vehemently, emotionally fight any decision against including their pets. Sometimes they might even resort to not telling anyone they’re using it, and will try to sneak it in at the last minute.

The Arcanist

Anyone who has worked on a legacy system of any import has dealt with an Arcanist. The Arcanist’s goal is noble: to preserve the uptime and integrity of the system, but at a terrible cost.
The Arcanist has a simple philosophy that guides his or her software development or administrative practices: if it ain’t broke, don’t fix it – to an extreme.
The day a piece of software under his or her auspices ships, it will forever stay on that development platform, with that database, with that operating system, with that deployment procedure. The Arcanist will see to it, to the best of his ability. He may not win every battle, but he will fight ferociously always.
All change is the enemy – it’s a vampire, seducing less vigilant engineers to gain entry to the system, only to destroy it from within.
The past is the future in the Arcanists’ worldview, and he’ll fight anyone tries to upgrade his circa 1981 PASCAL codebase to the bitter, tearful end.

The Futurist

The Futurist is the antithesis of the Arcanist – today is the future, and any code written with yesterday’s tools fills the Futurist with unparalleled disgust and shame. The Futurist’s goal is not noble – it’s to be seen as new and cutting edge. The Futurist’s measure of success is Hacker News karma and well-attended User Group meetups from his war stories, not effective programming.
The Futurist will breathlessly tell you with spittle-flying gasps about the latest JavaScript-does-something-it-shouldn’t experimental project he read about on Hacker News thirty minutes ago and has already forked to his Github. He’ll squeal like a teenage girl at a Justin Bieber concert every time Microsoft Research or the Server and Tools Team releases an obscure alpha of some technology he knows the name of but doesn’t understand.
The Futurist is responsible for more reverted commits than any other developer, and is often flabbergasted when his attempts to upgrade your database driver package to v1.0.13-alpha-unstable-prelease-DONOTUSE are rejected. His pleadings of “but it has Java Futures, so we get pure async!!11!1” do not stir the vigilant release manager.
The Futurist cares not for quaint, passing concerns about stability, maintainability, or teachability – it doesn’t matter to him if it’s impossible to hire Erlang developers. New is everything.
The DevOps Engineer, the QA Engineer, and the Release Manager are the natural enemies of the Futurist.

The Hoarder

The Hoarder is a cautious creature, perpetually unsure of itself. The Hoarder lives in a world of perpetual cognitive dissonance: extremely proud of his work, but so unsure of himself that he won’t let anyone see it if it can be helped.
So he hides his code. Carefully avoiding check-ins until the last possible minute, when he crams it all into one monolithic commit and hopes no one can trace the changes back to him. His greatest fear is the dreaded merge conflict, where the risk of exposure is greatest.
The Hoarder will openly tell you that his work is awesome but in confidence knows his code might suck. It probably does – but his fear of facing that possibility is what makes him who he is. The Hoarder is responsible for many last minute bugs sprinkled throughout the catacombs of the code base. His fellow engineers, tired of slowly going insane from invasion of subtle bugs, rose to fight him.
They invented tools like git blame and other weapons of accountability, and vengeance.
Ultimately, the Hoarder is damned and doomed – but there is hope for him in the short run. The day accountability comes for him at one job, his Dice resume is updated for another. The Hoarder will live to fight another day.

The Artist

A cousin of The Hoarder and The Futurist, the Artist pours his soul into every thoughtfully constructed line of code. The Artist is an emotionalist – his software is an expression of himself, a living embodiment of the genius he represents.
The Artist considers the important questions, like will my JavaScript be more syntactically beautiful without semicolons? How can I turn this perfectly readable for..each statement into a single line of LINQ? Will wrapping this block in a promise make it more… elegant?
The Artist and his code are one, which creates inevitable problems in the real world. Every JIRA issue containing documented, indisputable proof of a bug in his code is an act of artistic censorship on the behalf of users who “just don’t understand” or jealous colleagues who envy his genius. His lower lip will quavers for minutes each time a UserVoice ticket announces its presence in HipChat, titled with the name of a feature he owns – even if the ticket in question is replete with the gentle patience and understanding of an early adopter.
The Artist is not long for this world, unable to have an objective discussion about his body of work with even the most sympathetic of colleagues, he withdraws from the company of his fellow developers and metamorphoses into The Island.

The Island

The Island is the ultimate loner in the taxonomy of terrible software developers, as he desires above all things to be left in peace with his favorite text editor and devoid himself of all human contact. The ideal condition for the Island is one where communication with the outside world is kept at a minimum and strictly at his convenience. Just code, no humans.
Unfortunately, reality and ideal often being far afield, the Island has to seek the company of others in order to live. So he is forced to communicate with co-workers or clients, and it is a tremendous burden for him indeed.
So hides – he’ll miss meetings, fail to return phone calls, stay signed off of IM, keep the email client closed, and so forth. He’ll gladly spend hundreds of man-hours Gandalfing the documentation and project specs rather than simply ask someone on his team questions.
The Island is usually also a Hoarder – keeping his releases close to the vest until it’s absolutely necessary to share them. Anything to avoid people and their judgments.
And like the Hoarder, the Island is doomed. Software is a team sport and does not suffer those who do not play by its rules.

The “Agile” Guy

The “Agile” Guy is a utilitarian, who ultimately seeks to improve the efficiency and productivity of himself and his team. Unfortunately, both his understanding of “agile” philosophies and his implementation strategy are hilariously inflexible and rigid, an irony which is completely lost on him.
The “Agile” Guy’s intentions are noble: to improve the way software is made. He’ll introduce kanban boards with precisely four tiers for exactly every project and a meticulously calculated method for determining the exact number of business points and sprint points for every issue the team can encounter.
Any issue which takes longer than four hours must be broken up; any sprint longer than two weeks must be truncated. NO EXCEPTIONS.
All personnel must pair program with their designated pair programming partners at all times, no exceptions. All git commits must be done in this exact format and hours of work against a JIRA issue must be logged at regular intervals. No status updates are allowed at standups, which are strictly cut off at 10 minutes.
The “Agile” Guy forgets the purpose of the agile process to begin with – to be flexible and dynamic, and instead imposes round-hole order on square-peg problems. The JIRA board turns into a ghetto, a wasteland of broken promises and infeasible commitments. And from dusk to dawn the developers toil, pair programming all the way, while their “Agile” Guy overload looks proudly over his new empire, forged on discipline and process.
But the “Agile” Guy creates powerful enemies in his wake: the actual programmers responsible for getting things done, men and women who live in a world without luxuries of time or realizable schedules. Men and women who try to create order from chaos themselves, but are often at the mercy of failing networked file systems and poorly implemented drivers. Such intrepid souls do not suffer having the will of idiots imposed on them. Dissent and strife spread throughout the team; what follows is rebellion.
All “Agile” Guys meet their end by being torn down like the statue of a deposed dictator, complete with an Ewok luau in the Endor moonlight and the burning wreckage of their kanban Death Star smoldering in the distance.

The Human Robot

The Human Robot tries his best and his intentions are good. But he has a handicap: he interprets everything literally. The Human Robot is the world’s first true organic computer, which also means that every user-space detail of their work must be explained literally, byte-for-byte.
The Human Robot has his uses – he can find (and patch himself!) the subtle race condition created by minor JVM differences on your particular version of the Linux kernel, but ask him to implement a new feature and a monster is born. The Human Robot, unable to grasp concepts such as figurative language, imagination, abstraction, and creative interpretation, is stuck in a world where he can only process commands.
Your product lead asks the Human Robot to create a button that allows the end-user to share their documents via email with another end user; a week later the Human Robot delivers a fully functional high-throughput transactional email server embedded inside your application.
When a Human Robot is confronted about an issue with his or her work on the product, they will respond with the following sentence every time: “your requirement was not found in the specification for this project. We require additional pylons.”
This handicap rears its head most clearly in the team dynamics of a software organization. The Human Robot is the sort of person who needs a four-page-long decision tree and a finite state machine diagram to help him understand what does and does not qualify as sexual harassment in the workplace.
Human Robots often tend to be conference organizers (see PyCon 2013) and moderators on StackExchange.

The Stream of Consciousness

The Stream of Consciousness programmer is related to the Illiterate in that he too cannot read code. However, what’s fundamentally different about the Stream of Consciousness is that he can’t read his own code in addition to that of every other developer on the team.
In fact, the Stream of Consciousness programmer is best described as a “forward-only” cursor – the only way for them to solve their problems is to write new code, every time. Code reuse and refactoring are alien programming practices that the Stream of Consciousness will tell you they “understand,” but that’s only because the know the names of those concepts.
The Stream of Consciousness will happily add a third and fourth new interface to your application for writing to the filesystem, because their cursor has already moved past the first and second interfaces they wrote last week. Their code will be totally free of circular references, but that’s only because nothing ever gets used more than once.
The easiest way to determine if you’re working with a Stream of Consciousness programmer is to read their source code and look for comments along the lines of “hmm, I wonder if this works?” and “I really wish the kitchen had more non-dairy creamer.”

The Illiterate

As the name suggests, the Illiterate has a massive problem when it comes to read other people’s source code.
The Illiterate, a close cousin of the Island, understands basic programming constructs and has a full grasp on the syntax of his / her preferred programming languages, but is totally blind when it comes to code written by other developers. We call this “code-blind” in extreme cases.
The illiterate can understand the basic “hello world” example, but beyond that he or she never developed the capacity to understand another programmer’s intent or the “Find Usages” button in Visual Studio. So the Illiterate is forced to work around this alien “other programmer” code in all of his or her day-to-day assignments – often duplicating the work of other developers and contributing to code bloat.
When confronted by other developers and asked “why didn’t you use our standard interface for rendering a dialog?,” the Illiterate will simply stare at his or her shoes and mumble inaudibly.

The Agitator

The Agitator, like the Human Robot, is a social retard. But unlike the Human Robot, has no good intentions. An Agitator is not born, they are forged through years of suffering through undesirable work environments and programming assignments.
Having been through shit work for years and years, the Agitator believes he or she now knows best and is determined to run things they see fit, whether they actually have the authority to or not.
The goal of the Agitator is to establish dominance and control over their work environment through the use of force and intimidation.
The Agitator sees themselves as the confluence of Grace Hopper and Che Guevara, a brilliant technical visionary casting off the chains of oppression by an ignorant and ineffective ruling class. The Agitator is seen by his or her teammates as an idiotic wannabe-alpha bully who gets into high volume arguments over when it’s appropriate to use Pascal case versus CamelCase.
The Agitator will routinely try to talk down peers and superiors both in an attempt to assert dominance. In many cases he will win, seeding dissatisfaction and discord in the workplace in the process.
In most cases the Agitator is beyond help, and management has no choice but to enforce the “No Asshole” rule and terminate him.
