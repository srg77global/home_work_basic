package main

import (
	"errors"
)

func BinarySearch(args []int, val int) (int, error) {
	low := 0
	high := len(args) - 1

	if args[low] == val {
		return low, nil
	}
	if args[high] == val {
		return high, nil
	}

	for low < high-1 {
		mid := (low + high) / 2
		if args[mid] == val {
			return mid, nil
		}
		if args[mid] < val {
			low = mid
		}
		if args[mid] > val {
			high = mid
		}
	}
	return -1, errors.New("val is out of range")
}
