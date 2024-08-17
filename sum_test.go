package main

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {

	var nums []int
	nums = append(nums, 1, 2, 3, 4, 5)
	request := Request{Numbers: nums, Target: 6}
	got := request.calculate()
	var unit []int
	var res [][]int

	unit = append(unit, 0, 4)
	res = append(res, unit)

	unit = nil
	unit = append(unit, 1, 3)
	res = append(res, unit)

	want := Response{Solutions: res}

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("response did not match")
	}
}
