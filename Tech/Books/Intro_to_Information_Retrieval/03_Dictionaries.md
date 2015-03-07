## Dictionaries and tolerant retrieval

### Data structures
Hash table
* cannot search by prefix
* hash function design need to cater for growing needs

Search tree
* B-tree

### Wildcard query
* single B-tree, search by prefix is easy.

Middle wildcard `e.g. S*dney `
* method: B-tree and reverse B-tree, take intersection of results.

Permuterm index
* special symbol $ for string termination
* rotate `S*dney$` to `dney$S*`
* store all rotations of terms in index and search through that.

k-gram index
* split string into 3-grams - "castle" - $ca, cas, ast, stl, tle, le$
* each 3-gram -> list of terms that contains the 3-gram
* post-filtering to make sure terms are matching query

### Spelling correction

naive
* e.g. "carot" look for "carot", "carrot" and "tarot"
* + but only when "carot" is not in dictionary
* + but only when "carot" returned fewer than preset number of doc
* + present a spelling suggestion to end user

isolated term correction
* e.g. "flew form Heathrow" would fail to detect 'from'
* because not context aware

Edit distance
* minimum number of edit operations required to transform str1 and str2
* e.g. insert, delete, replace char
* also - Levenshtein distance
* ? - implement a dynamic algo to calc edit distance
* + heuristics to speed up lookup
  * same first letter
  * use a version of permuterm index...

k-gram overlap
* linear scan intersection
* Jaccard coefficient - measuring the overlap
* if coefficient exceeds a preset threshold, we add t to the output; (close enough)
  if not, we move on to the next term. (not close at all)
* to compute Jaccard coefficient we only need the length of t.
  q = bord, t = boardroom (2 bi-grams matches - bo, rd), if postings stored precomputed number
  of bigrams in 'boardroom' (8), we can compute Jaccard coefficient: 2/(8+3-2), denominator is 
  sum of bigrams in bord and boardroom, less posting lists.

context-sensitive correction
* naive to enumerate all the combinations of alternative spelling, 
  e.g. "flew from heathrow", "fled form heathrow" etc
* Use heuristics to trim the space, keep most frequent combinations
* only expand the list of top biwords (flew from) to corrections of Heathrow.
* or use biword statistics

Phonetic correction
* generate a phonetic hash, so similar sounding terms hash to same value.
* soundex algorithms.
  1. turn every term to be indexed into 4-character reduced form. Build an inverted index
     from these to original terms - soundex index.
  2. do the same for query terms.
e.g. Hermann - H655. matches Herman

