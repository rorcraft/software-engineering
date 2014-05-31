http://www.sorting-algorithms.com/static/QuicksortIsOptimal.pdf

### Quicksort

http://permalink.gmane.org/gmane.comp.java.openjdk.core-libs.devel/2628

* Tony Hoare developed Quicksort in the early 1960s.
* Bentley & McIlroy - 3 way optimized
* Sedgewick - 3 way optimized
* Vladimir - 3 dual pivot

### Classic

```
[ <= p | >= p ] ]
```

http://algs4.cs.princeton.edu/23quicksort/Quick.java.html
```go

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func InsertionSort(data Interface) {
	for i := 1; i < data.Len(); i++ {
		for j := i; j >= 1 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}

func QuickSort(data Interface) {
	quicksort(data, 0, data.Len()-1)
}

func quicksort(data Interface, low, high int) {
	if high <= low { return }
	if high-low < 7 {
		InsertionSort(data)
	}
	p := qsPartition(data, low, high)
	quicksort(data, low, p - 1)
	quicksort(data, p + 1, high)
}

func qsPartition(data Interface, low, high int) int {
	N := high - low + 1
	p := median3(data, low, low + N/2, high)
	data.Swap(low, p) // keep pivot at low
	i, j := low+1, high
	for {
		// find from low not less than pivot value.
		for data.Less(i, low) {
			if i == high { break }
			i++
		}
		// find from high not larger than pivot value.
		for data.Less(low, j) {
			if j == low { break }
			j--
		}
		if i >= j { break }
		data.Swap(i, j)
	}

	data.Swap(low, j) // swap stored pivot to j
	return j
}

func median3(data Interface, i, j, k int) int {
	if data.Less(i, j) {
		if data.Less(j, k) {
			return j
		} else {
			if data.Less(i, k) {
				return k
			} else {
				return i
			}
		}
	} else {
		if data.Less(k, j) {
			return j
		} else {
			if data.Less(k, i) {
				return k
			} else {
				return i
			}
		}
	}
}
```

### 3 way

* http://algs4.cs.princeton.edu/23quicksort/Quick3way.java.html
* http://algs4.cs.princeton.edu/23quicksort/QuickX.java.html

```
[ < p | = p | > p ]  or  [ = p | < p | > p | = p ] ] ]
```

### 3 way dual pivot

```
[ < P1 | P1 <= & <= P2 } > P2 ] ]
```
