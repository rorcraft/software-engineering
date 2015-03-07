## Document delineation
1. Determine encoding. (heuristic methods, user selection, metadata)
2. e.g. Arabic (vocalization), expanding combined versions

In languages that are written right-to-left, such as Hebrew and Arabic,
it is quite common to also have left-to-right text interspersed, such as
numbers and dollar amounts.

In Unicode, order of charactersa in files matches conceptual order, and
the reversal of displayed characters is handled by rendering system.

__Choose a document unit__
e.g. PDF vs multi HTML, email threads vs per message + attachments.
Issue of index granularity.

## Normalization (classing of terms)
- canonicalizing tokens
- alternative is synonyms
* Accents and diacritics. e.g. cliché and cliche
* Capitalization/case-folding e.g. C.A.T -> CAT, cat, most commonly lower case
* English - American vs British, date formats. ~60% web pages.
* French (the) - le, la, l', les
* stop words, sort by collection frequency. (not used anymore)
  problems: "to be or not to be", "let it be"
  modern: term weighting, impact-sorted index
* Soundex algorithm - phonetic equivalents (Beijing, Peking)
* Stemming - chops off ends of word. Porter Stemmer - popular 4 phase.
* Lemmatization - proper use of vocabulary and morphological analysis
  reduce to base / dictionary form - lemma.

### Skiplist
* hash table is at best O(n)
* Length P, Sqrt P even-spaced skip pointers

## Positional Postings and phrase queries

### Biword indexes
* every pair of consecutive terms
* nouns (N), articles (X) - NX\*N e.g. cost overruns on power plant - 'overruns power'

### Positional Index
* much more common
* docId: position1, position2...
* after matching, calculate the distance between term positions for filtering.
* can be 2-4 times as large as non-positional index
* compressed positional index ~1/3 of raw text.

### Combination - Biword and Positional
* e.g. "The Who" - biword, may speed up query by 1000x without going through positional index.

## Determining the vocabulary of terms
