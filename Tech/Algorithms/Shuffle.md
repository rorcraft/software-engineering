### Shuffle

> Fisher–Yates shuffle (named after Ronald Fisher and Frank Yates), also known as the Knuth shuffle (after Donald Knuth)

https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle

__O(n)__

In place.

```golang
func shuffle(slice []int) []int {
	for i := len(slice) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		slice[i], slice[r] = slice[r], slice[i]
	}
	return slice
}
```

The "inside-out" algorithm, for copying into new slice with n random elements.
http://golang.org/src/pkg/math/rand/rand.go?s=3157:3189#L94

```golang
// rand.Perm
// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *Rand) Perm(n int) []int {
	m := make([]int, n)
	for i := 0; i < n; i++ {
		m[i] = i
	}
	for i := 0; i < n; i++ {
		j := r.Intn(i + 1)
		m[i], m[j] = m[j], m[i]
	}
	return m
}
```
