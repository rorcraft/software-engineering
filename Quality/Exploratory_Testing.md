### Exploratory Testing

A systematic approach for __discovering risks__ using __rigorous analysis__ techniques coupled with __testing heuristics__

* Discover things you can vary
* State Models
* Data Models
* Deployment Diagrams

### Test Heuristics Cheat Sheet
http://testobsessed.com/wp-content/uploads/2011/04/testheuristicscheatsheetv1.pdf

### Explore with a focus using Tours
http://msdn.microsoft.com/en-us/library/jj620911.aspx

__The Guidebook Tour__

* Guidebooks for tourists identify the best hotels, the best bargains, and the top attractions, without going into too much detail or overwhelming a tourist with too many options. The analogous artifact for exploratory testing is the user manual, whether it is printed or implemented as online help (in which case, I often call this the F1 tour to denote the shortcut to most help systems). For this tour, we will follow the user manual’s advice just like the wary traveler, by never deviating from its lead.

__The Money Tour__

* Every location that covets tourists must have some good reasons for them to come. For Las Vegas, it’s the casinos and the strip, and for Egypt it’s the pyramids. For exploratory testers finding the money features leads directly to the sales force. Sales folk spend a great deal of time giving demos of applications and are a fantastic source of information for the Money tour. To execute the tour, simply run through the demos yourself and look for problems. As the product code is modified for bug fixes and new features, it may be that the demo breaks and you’ve not only found a great bug, but you’ve saved your sales force from some pretty serious embarrassment.

__The Landmark Tour__

* As a boy growing up in the fields, meadows, and woods of Kentucky, I learned to use a compass by watching my older brother. The process was simple. Use the compass to locate a landmark (a tree, rock, cliff face, and so forth) in the direction you want to go, make your way to that landmark, and then locate the next landmark, and so on and so forth. As long as the landmarks were all in the same direction, you could get yourself through a patch of dense Kentucky woods.
* The Landmark tour for exploratory testers is similar in that we will choose landmarks and perform the same landmark hopping through the software that we would through a forest. Choose a set of landmarks, decide on an ordering for them, and then explore the application going from landmark to landmark until you’ve visited all of them in your list. Keep track of which landmarks you’ve used and create a landmark coverage map to track your progress.

__The Intellectual Tour__

* I was once on a walking tour of London in which the guide was a gentleman in his fifties who claimed at the outset to have lived in London all his life. A fellow tourist happened to be a scholar who was knowledgeable in English history and was constantly asking hard questions of the guide. He didn’t mean to be a jerk, but he was curious, and that combined with his knowledge ended up being a dangerous combination … at least to the guide. When applied to exploratory testing, this tour takes on the approach of asking the software hard questions. How do we make the software work as hard as possible? Which features will stretch it to its limits? What inputs and data will cause it to perform the most processing? Which inputs might fool its error-checking routines? Which inputs and internal data will stress its capability to produce any specific output?

__The FedEx Tour__

* FedEx is an icon in the package-delivery world. They pick up packages, move them around their various distribution centers, and send them to their final destination. For this tour, instead of packages moving around the planet through the FedEx system, think of data moving through the software. During this tour, a tester must concentrate on this data. Try to identify inputs that are stored and “follow” them around the software. For example, when an address is entered into a shopping site, where does it gets displayed? What features consume it? If it is used as a billing address, make sure you exercise that feature. If it is used as a shipping address, make sure you use that feature. If can be updated, update it. Does it ever get printed or purged or processed? Try to find every feature that touches the data so that, just as FedEx handles their packages, you are involved in every stage of the data’s life cycle.

__The Garbage Collector’s Tour__

* Those who collect curbside garbage often know neighborhoods better than even residents and police because they go street by street, house by house, and become familiar with every bump in the road. However, because they are in a hurry, they don’t stay in one place very long. For software, this is like a methodical spot check. We can decide to spot check the interface where we go screen by screen, dialog by dialog (favoring, like the garbage collector, the shortest route), and not stopping to test in detail, but checking the obvious things (perhaps like the Supermodel tour). We could also use this tour to go feature by feature, module by module, or any other landmark that makes sense for our specific application.

__The Bad-Neighborhood Tour__

* Every city worth visiting has bad neighborhoods and areas that a tourist is well advised to avoid. Software also has bad neighborhoods—those sections of the code populated by bugs. Clearly, we do not know in advance which features are likely to represent bad neighborhoods. But as bugs are found and reported, we can connect certain features with bug counts and can track where bugs are occurring on our product. Because bugs tend to congregate, revisiting buggy sections of the product is a tour worth taking. Indeed, once a buggy section of code is identified, it is recommended to take a Garbage Collector’s tour through nearby features to verify that the fixes didn’t introduce any new bugs.

__The Museum Tour__

* Museums that display antiquities are a favorite of tourists. Antiquities within a code base deserve the same kind of attention from testers. In this case, software’s antiquities are legacy code. Older code files that undergo revision or that are put into a new environment tend to be failure prone. With the original developers long gone and documentation often poor, legacy code is hard to modify, hard to review, and evades the unit testing net of developers (who usually write such tests only for new code). During this tour, testers should identify older code and executable artifacts and ensure they receive a fair share of testing attention.

__The Back Alley Tour__

* In many peoples’ eye, a good tour is one in which you visit popular places. The opposite of these tours would be one in which you visited places no one else was likely to go. In exploratory testing terms, these are the least likely features to be used and the ones that are the least attractive to users. If your organization tracks feature usage, this tour will direct you to test the ones at the bottom of the list. If your organization tracks code coverage, this tour implores you to find ways to test the code yet to be covered.

__The All-Nighter Tour__

* Also known as the Clubbing tour, this one is for those folks who stay out late and hit the nightspots. The key here is all night. Exploratory testers on the All-Nighter tour will keep their application running without closing it. They will open files and not close them. Often, they don’t even bother saving them so as to avoid any potential resetting effect that might occur at save time. They connect to remote resources and never disconnect. And while all these resources are in constant use, they may even run tests using other tours to keep the software working and moving data around. If they do this long enough, they may find bugs that other testers will not find because the software is denied that clean reset that occurs when it is restarted.

__The Supermodel Tour__

* For this tour, I want you to think superficially. Whatever you do, don’t go beyond skin deep. This tour is not about function or substance; it’s about looks and first impressions. During the Supermodel tour, the focus is not on functionality or real interaction. It’s only on the interface. Take the tour and watch the interface elements. Do they look good? Do they render properly, and is the performance good? As you make changes, does the GUI refresh properly? Does it do so correctly or are there unsightly artifacts left on the screen? If the software is using color in a way to convey some meaning, is this done consistently? Are the GUI panels internally consistent with buttons and controls where you would expect them to be? Does the interface violate any conventions or standards?

__The Couch Potato Tour__

* There’s always one person on a group tour who just doesn’t participate. He stands in the back with his arms folded. He’s bored, unenergetic, and makes one wonder exactly why he bothered paying for the tour in the first place. A Coach Potato tour means doing as little actual work as possible. This means accepting all default values (values prepopulated by the application), leaving input fields blank, filling in as little form data as possible, never clicking on an advertisement, paging through screens without clicking any buttons or entering any data, and so forth. If there is any choice to go one way in the application or another, the coach potato always takes the path of least resistance.

__The Obsessive-Compulsive Tour__

* OCD testers will enter the same input over and over. They will perform the same action over and over. They will repeat, redo, copy, paste, borrow, and then do all that some more. Mostly, the name of the game is repetition. Order an item on a shopping site and then order it again to check if a multiple purchase discount applies. Enter some data on a screen, then return immediately to enter it again. These are actions developers often don’t program error cases for. They can wreak significant havoc.
Developers are often thinking about a user doing things in a specific order and using the software with purpose. But users make mistakes and have to backtrack, and they often don’t understand what specific path the developer had in mind for them, and they take their own. This can cause a usage scheme carefully laid by developers to fall by the wayside quickly.
