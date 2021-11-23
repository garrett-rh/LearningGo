package v3_main

import "testing"
import "reflect"

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) { //reflect.DeepEqual is not type safe, will check any type even if they arent the same
		t.Errorf("got %v want %v", got, want)
	}
}
