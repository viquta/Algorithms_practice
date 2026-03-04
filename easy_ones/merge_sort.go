/*
//Merge Sort:
A := Array []int{p...r}
func merge_sort()
divide

	compute mid = (p + r) / 2 --> get A[p...mid] and A[mid+1...r]

conquer

	  sort A[p...mid] and A[mid+1...r] recursively by calling
		merge_sort() on each half

combine

	merge the sorted halves to produce a single sorted array A[p...r]

pseudocode (p. 39):
MERGE-SORT(A, p, r)
1. if p > r                        // base case: if the array has one or zero elements, it's already sorted
2.     return
2.     mid = ⌊(p + r) / 2⌋
3.     MERGE-SORT(A, p, mid)       // ← RECURSION (left half)
4.     MERGE-SORT(A, mid+1, r)     // ← RECURSION (right half)
5.     MERGE(A, p, mid, r)         // combine step (your pseudocode)

pseudocode (p. 36):
MERGE-SORT(A, p, mid, r)
1. nL = mid - p + 1
2. nR = r - mid
3. create arrays L[1...nL + 1] and R[1...nR + 1]
4. for i = 1 to nL - 1 //copy A[p...mid] to L[0...nL - 1]
5.     L[i] = A[p + i]
6. for j = 0 to nR - 1 //copy A[mid + 1...r] to R[0...nR - 1]
7.     R[j] = A[mid + 1 + j]
8. i = 0
9. j = 0
10. k = p //k is the index for the merged array A[p...r]
11. as long as each of the arrays L and R has an unprocessed element
we copy the smallest of the two unprocessed elements to A[p...r]
12. while i < nL and j < nR
13.     if L[i] <= R[j]
14.         A[k] = L[i]
15.         i = i + 1
16.     else
17.         A[k] = R[j]
18.         j = j + 1
19.     k = k + 1
20. copy the remaining elements of L, if there are any
21. while i < nL
22.     A[k] = L[i]
23.     i = i + 1
24.     k = k + 1
25. copy the remaining elements of R, if there are any
26. while j < nR
27.     A[k] = R[j]
28.     j = j + 1
29.     k = k + 1

//Here is an implementation of Merge Sort in Go:
*/
package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr // exiting the recursion
	}
	mid := len(arr) / 2           //divide
	left := mergeSort(arr[:mid])  // conquer
	right := mergeSort(arr[mid:]) // conquer
	return merge(left, right)     //combine
}

func merge(left, right []int) []int {
	// Merge the two sorted halves
	//invariant: at each iteration, result contains all el smaller than
	// left and right in a sorted order
	//O(n)
	result := make([]int, 0, len(left)+len(right)) //the make() is just performance optimization
	//could also use:
	//result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i]) //result is A[k]
			i++
		} else {
			result = append(result, right[j]) //result is A[k]
			j++
		}
	}
	//one of the slices may still have remaining elements
	//append the remaining elements to the result
	result = append(result, left[i:]...) //"take all remaining elements from left starting at index i and append each one to result."
	result = append(result, right[j:]...)
	return result
}

func main() {
	arr := []int{5, 3, 2, 8, 1, 4, 7, 6, 9, 11, 10, 12, 18, 77, 176, 19, 23, 45, 67, 89, 34, 56, 78, 90, 21, 43, 65, 87, 32, 54, 76, 98}
	fmt.Println("Original array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
