### Probability Ranking Principle
* order results by estimated probability of relevance with respect to information need.

Bayes Theorem lets us derive a posterior probability P(A|B) after
having seen evidence B, based on the likelihood of B.

Odds(A) = P(A) / P(-A) = P(A) / (1 - P(A))
- a kind of multiplier for how probabilities change.

1/0 loss case
1 for relevant, 0 for not - R(d,q) relevance for document and query.
* lose a point for either returning a nonrelevant doc or failing to return a relevant doc.

Bayes Optimal Decision Rule - return docs more likely relevant than nonrelevant.
d is relevant iff P(R=1|d,q) > P(R=0|d,q)
Proof - Ripley (1996) (but never the case in practice)

### Binary Independence Model
'binary' = Boolean.
'independence' - terms are modeled as occuring in documents independently, no associations.
same as the multivariate Bernoulli Naive Bayes model.

Prabilistic relevance feedback
1. Guess pt and ut. Assume pt is constant over all xt in the query. = 1/2.
2. Best guess relevant documents R based on pt, ut.
3. Learn from user relevance judgements.
4. Reestimate pt and ut.

Requires major assumptions.
* A Boolean representation of documents/queries/relevance
* term independence
* terms no in the query don't affect the outcome
* document relevance values are independent

Tree-structured dependencies between terms
* Tree Augmented Naive Bayes model by Friedman 1996

### Okapi BM25 - non-binary model
BM25 weighting scheme, aka Okapi weighting, building a probabilistic model sensitive
to these quantities while not introducing too many additional parameters into the model.

### Bayesian Network
Directed graphs to show probabilistic dependencies between variables.
