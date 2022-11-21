package main

func quickSort(a []int) []int {
	length := len(a)
	l, r := 0, length-1

	if length < 2 { // if lenght is 1 or 0 no need to check for sort
		return a
	}

	p := length / 2 // lets take middle as pivot

	/*
		shift pivot to right side to compare all elements with pivot
		now assume pivot is a[r]
	*/
	a[p], a[r] = a[r], a[p]

	for i, _ := range a {
		if a[i] < a[r] {
			a[l], a[i] = a[i], a[l] // swap to left if it smaller than pivot
			l++                     // increment l
		}
	}
	/*
	  bring pivot to position where left stop
	  now all left are small and all right are big to pivot
	*/
	a[r], a[l] = a[l], a[r]
	quickSort(a[l:])   // recursive for left which are small to pivot
	quickSort(a[l+1:]) // recursive for  right which are big to pivot
	return a
}
