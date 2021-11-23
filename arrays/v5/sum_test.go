package v5_main

import "testing"
import "reflect"

func TestSumAllTails(t *testing.T) {

	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
