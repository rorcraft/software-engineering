/*
 * Given an integer, print the next smallest and next largest number that have the
 * same number of 1 bits in their binary representation.
 */

 // next largest = (least 1 digit > 0) >> 1
 // next smallest = (least 1 digit << 1) + carry

package main

import (
	"fmt"
	"strconv"
)

func least1(num int64) uint {
	i := uint(0)
	for num > 0 {
		if num & 1 == 1 {
			return i
		}
		num >>= 1
		i++
	}
	return i
}
func next0(num int64, i uint) uint {
	j := i
	num >>= j
	for num > 0 {
		if num & 1 == 0 {
			return j
		}
		num >>= 1
		j++
	}
	return j
}

func swap(num int64, i, j uint) int64 {
	m := mask(num)
	numi := ((num >> i) & 1) << j
	// fmt.Printf("mask %b\n", m)
	numj := int64((1 << i)) ^ m
	// fmt.Printf("numi %b\n", numi)
	// fmt.Printf("numj %b\n", numj)
	num ^= numi
	num &= numj
	return num
}

func mask(num int64) int64 {
	var mask int64
	mask = 1 << 1
	for n := num; n > 0; n >>= 1 {
		mask <<= 1
	}
	mask -= 1
	return mask
}
func nextSmallest(num int64) int64 {
	i := least1(num)
	return swap(num, i, next0(num, i))
}

func least10(num int64) uint {
	i := uint(1)
	num >>= 1
	for num > 0 {
		if (num & 1) == 1 {
			return i
		}
		num >>= 1
		i++
	}
	return i
}

func next10(num int64, i uint) uint {
	j := i
	for j > 0 {
		if ((num >> j) & 1 == 0) {
			return j
		}
		j--
	}
	return j
}

func swapDown(num int64, i, j uint) int64 {
	m := mask(num)
	numi := int64(1 << j)
	// fmt.Printf("mask %b\n", m)
	numj := int64((1 << i)) ^ m
	fmt.Printf("numi %b\n", numi)
	fmt.Printf("numj %b\n", numj)
	num |= numi
	num &= numj
	return num
}

func nextLargest(num int64) int64 {
	i := least10(num)
	j := next10(num, i)
	return swapDown(num, i, j)
}


func main() {
	num, _ := strconv.ParseInt("1111", 2, 64)
	fmt.Printf("num: %b\n", num)
	small := nextSmallest(num)
	fmt.Printf("next smallest %b\n", small)
	fmt.Printf("next largest %b\n", nextLargest(small))
	num, _ = strconv.ParseInt("100010", 2, 64)
	fmt.Printf("num: %b\n", num)
	small = nextSmallest(num)
	fmt.Printf("next smallest %b\n", small)
	fmt.Printf("next largest %b\n", nextLargest(small))
}
