## Scoring, term weighting and vector space model

### Parametric and zone indexes

parametric index
* index of fields, e.g. author, creation dates

zone
* free text, e.g. abstract

weighted zone scoring
* aka ranked Boolean retrieval.

learning weights
* use training examples, optimize for minimizing errors; machine-learned relevance.

optimal weight g

### Term frequency and weighting

inverse document frequency
```
idf = log( N / df )
```

tf-idf weighting = tf * idf
1. highest when t occurs many times within small number of documents
2. lower when the term occurs fewer times in a document, or occurs in many
   documents.
3. lowest when the term occurs in virually all documents.

overlap score measure = sum( tf-idf of all terms in query ) 

vector space model: representation of a set of documents as vectors in a common vector space.

cosine similarity, computes the similarity of two documents d1, d2
```
sim(d1,d2) = V(d1).V(d2) / |V(d1)|.|V(d2)|
= v(d1).v(d2) // unit vectors
```
This measure is the cosine of angle between two vectors
Also used for 'more like this' feature - computing the dot products and max of v(d).[v(d1)..v(dn)]

Term-document matrix - M x N matrix
M terms (dimentions)
N columns

### Queries as vectors

* by viewing query as a 'bag of words', like a short document.
* use cosine similarity between query vector and document vector

see algorithm...

### Variant tf-idf functions

Sublinear tf scaling

Maximum tf normalization

SMART notation for specific combination of weights

Term frequency: (n)atural, (l)ogarithm, (a)ugmented, (b)oolean, (L)og ave.
Document frequency: (n)o, (t)idf, (p)rob idf
Normalization: (n)one, (c)osine, (u) pivoted unique, (b)yte size

Pivoted normalized document length
* longer documents contain more terms, have higher tf values.
* longer documents contain more distinct terms.
1. verbose documents, repeat the same content
2. multiple different topics, a term matching small segments of doc but not all (lower value)
* compute probability of relevance as a function of doc length, avg over all queries.

* pivot - crossing pt between prob of relevance and cosine normalization.
* pivot length - document length at pivot
* rotate the cosine normalization curve counter clockwise, so it more closely matches 
  prob of relevance.
