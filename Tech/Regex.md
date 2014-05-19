### Implementations:

* https://github.com/k-takata/Onigmo

### Avoid backtracking:

* http://www.regular-expressions.info/catastrophic.html
* http://blog.codinghorror.com/regex-performance/

### Quantifiers
https://docs.oracle.com/javase/tutorial/essential/regex/quant.html

Quantifiers allow you to specify the number of occurrences to match against. 

|Greedy	| Reluctant	| Possessive | Meaning|
|-------|-----------|------------|--------|
|X?	| X??	| X?+	| X, once or not at all|
|X*	| X*?	| X*+	| X, zero or more times|
|X+	| X+?	| X++	| X, one or more times|
|X{n}	| X{n}?	| X{n}+	| X, exactly n times|
|X{n,}	| X{n,}?	| X{n,}+	| X, at least n times|
|X{n,m}	| X{n,m}?	| X{n,m}+	| X, at least n but not more than m times |

### Greedy
``` 
"".match(/a?/g)
=> [""]

"".match(/a*/g)
=> [""]

"".match(/a+/g)
=> null

//-------------------------

"a".match(/a?/g)
=> ["a"]

"a".match(/a*/g)
=> ["a"]

"a".match(/a+/g)
=> ["a"]

//--------------------------

"aaaaa".match(/a?/g)
=> ["a", "a", "a", "a", "a", ""]

"aaaaa".match(/a*/g)
=> ["aaaaa", ""]

"aaaaa".match(/a+/g)
=> ["aaaaa"]

//----------------------------

"ababaaaab".match(/a?/g)
=> ["a", "", "a", "", "a", "a", "a", "a", "", ""]

"ababaaaab".match(/a*/g)
=> ["a", "", "a", "", "aaaa", "", ""]

"ababaaaab".match(/a+/g)
=> ["a", "a", "aaaa"]

//----------------------------

"aa".match(/a{3}/g)
=> null

"aaa".match(/a{3}/g)
=> ["aaa"]

"aaaaaaaaa".match(/a{3}/g)
=> ["aaa", "aaa", "aaa"]

"aaaaaaaaa".match(/a{3,}/g) // at least 3 times
=> ["aaaaaaaaa"]

"aaaaaaaaa".match(/a{3,6}/g) // at least 3 times, but no more than 6
["aaaaaa", "aaa"]
```

### Capturing Groups and Character Classes with Quantifiers

```
"dogdogdogdogdogdog".match(/(dog){3}/g)
=> ["dogdogdog", "dogdogdog"]

"dogdogdogdogdogdog".match(/dog{3}/g)
=> null // expecting 3 g.

"abccabaaaccbbbc".match(/[abc]{3}/g) // applies to entire character class
=> ["abc", "cab", "aaa", "ccb", "bbc"]

```


### Differences Among Greedy, Reluctant, and Possessive Quantifiers

There are subtle differences among greedy, reluctant, and possessive quantifiers.

Greedy quantifiers: force the matcher to read in the entire input string prior to attempting the first match. If the first match attempt (the entire input string) fails, the matcher backs off the input string by one character and tries again, repeating the process until a match is found or there are no more characters left to back off from.

The reluctant quantifiers: start at the beginning of the input string, then reluctantly eat one character at a time looking for a match. The last thing they try is the entire input string.

Finally, the possessive quantifiers always eat the entire input string, trying once (and only once) for a match. Unlike the greedy quantifiers, __possessive quantifiers never back off__, even if doing so would allow the overall match to succeed.

```
// greedy
"xfooxxxxxxfoo".match(/.*foo/g)
=> ["xfooxxxxxxfoo"]
 
// reluctant quantifier
"xfooxxxxxxfoo".match(/.*?foo/g)
=> ["xfoo", "xxxxxxfoo"]

// possessive quantifier
"xfooxxxxxxfoo".match(/.*+foo/g)
=> invalid in javascript ?
=> null
```

