Rvalue - http://thbecker.net/articles/rvalue_references/section_01.html

https://isocpp.org/blog/2012/11/universal-references-in-c11-scott-meyers

C++11 performs “reference collapsing” when references to references arise in contexts such as template instantiation.

Because there are two kinds of references (lvalue references and rvalue references), there are four possible reference-reference combinations: 
* lvalue reference to lvalue reference
* lvalue reference to rvalue reference
* rvalue reference to lvalue reference
* rvalue reference to rvalue reference.  

There are only two reference-collapsing rules:  
* An rvalue reference to an rvalue reference becomes (“collapses into”) an rvalue reference.  
* All other references to references (i.e., all combinations involving an lvalue reference) collapse into an lvalue reference.

In a type declaration, “&&” indicates either an rvalue reference or a universal reference – a reference that may resolve to either an lvalue reference or an rvalue reference. Universal references always have the form T&& for some deduced type T.

Reference collapsing is the mechanism that leads to universal references (which are really just rvalue references in situations where reference-collapsing takes place) sometimes resolving to lvalue references and sometimes to rvalue references. It occurs in specified contexts where references to references may arise during compilation. Those contexts are template type deduction, auto type deduction, typedef formation and use, and decltype expressions.
