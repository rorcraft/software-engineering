### How to break software
http://www.math.uaa.alaska.edu/~afkjm/cs470/handouts/breaking.pdf

* Testing often not as glamorous as developing, but requires 
lots of creativity
* Testing is not something one can master, always 
something new to learn

__UI Attackes__

* Apply inputs that force error messages to occur
* Apply inputs that force the software to (re) 
establish default values
* Explore allowable character sets and data 
types
* Overflow input buffers
* Find inputs that interact and test combinations of
their values
* Repeat the same input or series of inputs 
numerous times
* Force different outputs to be generated for 
each input
* Force invalid outputs to be generated
* Force properties of an output to change
* Force the screen to refresh
* Apply inputs using a variety of initial 
conditions
* Force a data structure to store too many or too few 
values
* Investigate ways to modify internal data 
constraints
* Experiment with invalid operand and operator 
combinations
* Force a function to call itself recursively
* Find features that share data or interact poorly

__System Interface Attacks__

* Fill the file system to its capacity
* Force the media to be busy or unavailable
* Damage the media
* Assign an invalid file name
* Vary file access permissions
* Vary or corrupt file contents

__System Attacks__

* Exhaust the amount of physical memory
* Inject Network Faults

