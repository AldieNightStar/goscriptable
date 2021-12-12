package goscriptable

import "testing"

func TestFindInside1Positive(t *testing.T) {
	txt := "This is some data [[abc: Hello]]"
	d := FindInside(txt, "abc")
	if d != "Hello" {
		t.Fatal("Data isn't Hello: ", d)
	}
}

func TestFindInside2Positive(t *testing.T) {
	txt := "This is some data [[abc: 1]2]3]4]]"
	d := FindInside(txt, "abc")
	if d != "1]2]3]4" {
		t.Fatal("Data isn't 1]2]3]4: ", d)
	}
}

func TestFindInside3Positive(t *testing.T) {
	txt := "This is some data [[abc: \\]\\]\\]]]"
	d := FindInside(txt, "abc")
	if d != "]]]" {
		t.Fatal("Data isn't ]]]: ", d)
	}
}

func TestFindInside4Positive(t *testing.T) {
	txt := "This is some data [[abc: \\\\] as sprite]]"
	d := FindInside(txt, "abc")
	if d != "\\] as sprite" {
		t.Fatal("Data isn't \\] as sprite: ", d)
	}
}

func TestFindInside1Negative(t *testing.T) {
	txt := "This is some data [[abc: ]]"
	d := FindInside(txt, "abc")
	if d != "" {
		t.Fatal("Not empty data")
	}
}

func TestFindInside2Negative(t *testing.T) {
	txt := "This is some data [[abc:]]"
	d := FindInside(txt, "abc")
	if d != "" {
		t.Fatal("Not empty data")
	}
}

func TestFindInside3Negative(t *testing.T) {
	txt := "This is some data [[ab: 111]]"
	d := FindInside(txt, "abc")
	if d != "" {
		t.Fatal("Not empty data")
	}
}

func TestFindInside4Negative(t *testing.T) {
	txt := "This is some data [[abc: 5"
	d := FindInside(txt, "abc")
	if d != "" {
		t.Fatal("Not empty data")
	}
}
