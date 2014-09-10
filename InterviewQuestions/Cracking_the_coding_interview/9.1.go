// You are given two sorted arrays, A and B, and A has a large enough buffer at the end to hold B.
// Write a method to merge B into A in sorted order.

// http://hawstein.com/posts/9.1.html
// e.g.
// a[1,3,4,8]
// b[2,5,6,7]
//
// a[3] > b[3]
// a[7] = a[3]
// a[2] < b[3]
// a[6] = b[3]
// a[2] < b[2]
// a[5] = b[2]
// a[2] < b[1]
// a[4] = b[1]
// a[2] > b[0]
// a[3] = a[2]
// a[1] > b[0]
// a[2] = a[1]
// a[0] < b[0]
// a[1] = b[0]
// a[0] = a[0]
