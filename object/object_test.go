package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have diffrent hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have diffrent hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with diffrent content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	v1 := &Integer{Value: 123}
	v2 := &Integer{Value: 123}
	diff1 := &Integer{Value: 234}
	diff2 := &Integer{Value: 234}

	if v1.HashKey() != v2.HashKey() {
		t.Errorf("integers with same content have diffrent hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integers with same content have diffrent hash keys")
	}

	if v1.HashKey() == diff1.HashKey() {
		t.Errorf("integers with diffrent content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	t1 := &Boolean{Value: true}
	t2 := &Boolean{Value: true}
	f1 := &Boolean{Value: false}
	f2 := &Boolean{Value: false}

	if t1.HashKey() != t2.HashKey() {
		t.Errorf("booleans with same content have diffrent hash keys")
	}

	if f1.HashKey() != f2.HashKey() {
		t.Errorf("booleans with same content have diffrent hash keys")
	}

	if t1.HashKey() == f1.HashKey() {
		t.Errorf("booleans with diffrent content have same hash keys")
	}
}
