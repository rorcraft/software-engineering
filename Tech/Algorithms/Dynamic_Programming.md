### Dynamic Programmig

5 Steps:
1. Subproblems
2. Guess
3. Recurrence

### Fibonacci

```
def fib:
  if n < 2: return 1
  else: return fib(n-1) + fib(n-2)
  
```

recursive + memoization

bottom-up
(space = only last 2)

### Shortest path

http://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-006-introduction-to-algorithms-spring-2008/lecture-notes/lec20.pdf

### Text Justification

1. Subproblems: suffixes words[i:]
2. Guess: where to start 2nd line, no. of choices <= n-i = O(n)
3. Recurrence: DP(i) = min { DP(j) + badness(i,j) for j in range(i+1, n+1) }
   time/subproblem = O(n)
4. Topological Order: i = n, n-1, n-2 ... 0, Total time: O(n^2)
5. Original problem: DP(0)

Parent Pointers: remember which guess was best.
parent[i] = argmin(...) = j value

### Perfect Information Black Jack

* Deck = C0, C1 .. Cn-1
* 1 player vs dealer

1. Subproblems: 

3. Recurrence: BJ(i) = max { outcome in { -1, 0, 1 } + BJ(j) for hits in range(0,n) if valid play } 


### Subproblems for strings

1. prefixes
2. suffixes
3. substring (n^2)

Parenthesization:
Optimal evaluation of associative expression. (A0. * A1) (.) . ) An-1
e.g. Matrix multiplication.

