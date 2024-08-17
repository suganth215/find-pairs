package main

import (
	"slices"
)

//removing duplicates by using Sort and Compact functions from slices package
func (request Request) calculate() (result Response) {
	slices.Sort(request.Numbers)
	res := slices.Compact(request.Numbers)
	result = findSumPair(res, request.Target)
	return result
}

//function to find the pairs
func findSumPair(arr []int, sum int) (result Response) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == sum {
				var unit []int
				//appending one solution to a list of integers
				unit = append(unit, i, j)

				//appending to the total set of result
				result.Solutions = append(result.Solutions, unit)
			}
		}
	}
	return result
}
