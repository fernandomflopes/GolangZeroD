package main

import "testing"

func Test_add(t *testing.T) {
	l := initList()
	l.add(1)
	ret, _ := l.get(0)
	expected := 1
	if ret.value != expected {
		t.Fail()
	}
}

func Test_addTwoValues(t *testing.T) {
	l := initList()
	l.add(1)
	l.add(2)
	ret, _ := l.get(1)
	expected := 2

	if ret.value != expected {
		t.Fail()
	}
}

