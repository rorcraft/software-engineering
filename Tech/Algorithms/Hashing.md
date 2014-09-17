http://algs4.cs.princeton.edu/34hash/

We have three primary requirements in implementing a good hash function for a given data type:
- It should be deterministic—equal keys must produce the same hash value.
- It should be efficient to compute.
- It should uniformly distribute the keys.

For integers.

```java
k % M // M is a prime closest to desirable array size.
```

For strings.

```java
int hash = 0;
for (int i = 0; i < s.length(); i++)
  hash = (R * hash + s.charAt(i)) % M;

// R is small prime (java use 31)
```

Compound keys.

```java
// e.g. area, exch, ext - treat them like string chars.

int hash = (((area * R + exch) % M) * R + ext) % M; 
```

(Java) Converting a hashCode() to an array index.

```java
private int hash(Key key) {
   return (key.hashCode() & 0x7fffffff) % M;
}
```
