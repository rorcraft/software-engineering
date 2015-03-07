## Computing scores in a complete search system.

* relative scores rather than absolute.
* Fast-cosine score: compute cosine similarity from each document unit vector
  non-zero components of the query vector are set to 1.

e.g. d1, d2
```
V(q).v(d1) > V(q).v(d2) <=> v(q).v(d1) > v(q).v(d2)
```
* use a heap to keep top K documents in order.

### Inexact top K, heuristics:
1. Find a set A, not necessarily contain the K top docs, but likely to have many docs close 
   to top K.
2. return K top of set A.
suited for free text queries, but not for Boolean / phrase queries.

Index elimination
1. only terms with idf > threshold.
2. only consider docs containing many / all of the terms.

Champion lists
* precompute for each term the set (r) of highest weights docs.
* take the union of champion lists for each term

Static quality scores and ordering (query independent)


