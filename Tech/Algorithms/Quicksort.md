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

```java
/**
 *  <at> author Vladimir Yaroslavskiy
 *  <at> version 2009.09.10 m765
 */
public class DualPivotQuicksort {

    public static void sort(int[] a) {
        sort(a, 0, a.length);
    }

    public static void sort(int[] a, int fromIndex, int toIndex) {
        rangeCheck(a.length, fromIndex, toIndex);
        dualPivotQuicksort(a, fromIndex, toIndex - 1, 3);
    }

    private static void rangeCheck(int length, int fromIndex, int toIndex) {
        if (fromIndex > toIndex) {
            throw new IllegalArgumentException("fromIndex(" + fromIndex + ") > toIndex(" + toIndex + ")");
        }
        if (fromIndex < 0) {
            throw new ArrayIndexOutOfBoundsException(fromIndex);
        }
        if (toIndex > length) {
            throw new ArrayIndexOutOfBoundsException(toIndex);
        }
    }

    private static void swap(int[] a, int i, int j) {
        int temp = a[i];
        a[i] = a[j];
        a[j] = temp;
    }

    private static void dualPivotQuicksort(int[] a, int left, int right, int div) {
        int len = right - left;

        if (len < 27) { // insertion sort for tiny array
            for (int i = left + 1; i <= right; i++) {
                for (int j = i; j > left && a[j] < a[j - 1]; j--) {
                    swap(a, j, j - 1);
                }
            }
            return;
        }
        int third = len / div;

        // "medians"
        int m1 = left  + third;
        int m2 = right - third;

        if (m1 <= left) {
            m1 = left + 1;
        }
        if (m2 >= right) {
            m2 = right - 1;
        }
        if (a[m1] < a[m2]) {
            swap(a, m1, left);
            swap(a, m2, right);
        }
        else {
            swap(a, m1, right);
            swap(a, m2, left);
        }
        // pivots
        int pivot1 = a[left];
        int pivot2 = a[right];

        // pointers
        int less  = left  + 1;
        int great = right - 1;

        // sorting
        for (int k = less; k <= great; k++) {
            if (a[k] < pivot1) {
                swap(a, k, less++);
            } 
            else if (a[k] > pivot2) {
                while (k < great && a[great] > pivot2) {
                    great--;
                }
                swap(a, k, great--);

                if (a[k] < pivot1) {
                    swap(a, k, less++);
                }
            }
        }
        // swaps
        int dist = great - less;

        if (dist < 13) {
           div++;
        }
        swap(a, less  - 1, left);
        swap(a, great + 1, right);

        // subarrays
        dualPivotQuicksort(a, left,   less - 2, div);
        dualPivotQuicksort(a, great + 2, right, div);

        // equal elements
        if (dist > len - 13 && pivot1 != pivot2) {
            for (int k = less; k <= great; k++) {
                if (a[k] == pivot1) {
                    swap(a, k, less++);
                }
                else if (a[k] == pivot2) {
                    swap(a, k, great--);

                    if (a[k] == pivot1) {
                        swap(a, k, less++);
                    }
                }
            }
        }
        // subarray
        if (pivot1 < pivot2) {
            dualPivotQuicksort(a, less, great, div);
        }
    }
}
```
