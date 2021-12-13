package goscriptable

import "testing"

func Test_SplitStringBy(t *testing.T) {
	arr := SplitStringBy("a, b, c", ",")
	if len(arr) != 3 {
		t.Fatal("Len is not 3: ", len(arr))
	}
	if arr[0] != "a" || arr[1] != "b" || arr[2] != "c" {
		t.Fatalf("Data wrong (Should be 'a', 'b', 'c'): '%s', '%s', '%s'", arr[0], arr[1], arr[2])
	}
	arr = SplitStringBy("a", ",")
	if len(arr) != 1 {
		t.Fatal("Len is not 1: ", len(arr))
	}
	if arr[0] != "a" {
		t.Fatal("Value wrong 'a': ", arr[0])
	}

	arr = SplitStringBy(" a    b  c", " ")
	if len(arr) != 3 {
		t.Fatal("Len is not 3: ", len(arr))
	}
	if arr[0] != "a" || arr[1] != "b" || arr[2] != "c" {
		t.Fatalf("Value wrong 'a', 'b', 'c': '%s', '%s', '%s'", arr[0], arr[1], arr[2])
	}
}
