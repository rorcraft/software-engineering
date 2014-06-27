http://thenoisychannel.com/2011/08/08/retiring-a-great-interview-problem/

__Question:__ Given an input string and a dictionary of words,
segment the input string into a space-separated
sequence of dictionary words if possible. For
example, if the input string is "applepie" and
dictionary contains a standard set of English words,
then we would return the string "apple pie" as output.

```
Q: What if the input string is already a word in the
   dictionary?
A: A single word is a special case of a space-separated
   sequence of words.

Q: Should I only consider segmentations into two words?
A: No, but start with that case if it's easier.

Q: What if the input string cannot be segmented into a
   sequence of words in the dictionary?
A: Then return null or something equivalent.

Q: What about stemming, spelling correction, etc.?
A: Just segment the exact input string into a sequence
   of exact words in the dictionary.

Q: What if there are multiple valid segmentations?
A: Just return any valid segmentation if there is one.

Q: I'm thinking of implementing the dictionary as a
   trie, suffix tree, Fibonacci heap, ...
A: You don't need to implement the dictionary. Just
   assume access to a reasonable implementation.

Q: What operations does the dictionary support?
A: Exact string lookup. That's all you need.

Q: How big is the dictionary?
A: Assume it's much bigger than the input string,
   but that it fits in memory.
```

If you could change the dictionary into a Trie, walk the string with the Trie can achieve O(n).

```javascript
var dict = { 'a': true, 'award': true, 'word': true, 'etc': true }
var str = "awordaward"
function wordBreak(str) {
  var foundAt = 0
  if (str.length === 0) { return [] }
  for (var i = 0; i <= str.length; i++) {
    if (dict[str.substring(0, i)]) { foundAt = i }
  }
  return [str.substring(0, foundAt)]
    .concat(wordBreak(str.substring(foundAt)))
}

console.log(wordBreak("awordaward"))
console.log(wordBreak("aaaaaa"))

[ 'a', 'word', 'award' ]
[ 'a', 'a', 'a', 'a', 'a', 'a' ]

O(n^2)
```
