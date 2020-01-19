package sn

import (
	"reflect"
	"testing"
)

func TestJump(t *testing.T) {
	if reflect.ValueOf(Jump(2)).Pointer() != reflect.ValueOf(SN02).Pointer() {
		t.Fatal("jump 2 incorrect")
	}
	if reflect.ValueOf(Jump(11)).Pointer() != reflect.ValueOf(SN11).Pointer() {
		t.Fatal("jump 11 incorrect")
	}
	if Jump(-1) != nil {
		t.Fatal("jump -1 incorrect")
	}
	if Jump(1) != nil {
		t.Fatal("jump 1 incorrect")
	}
	if Jump(12) != nil {
		t.Fatal("jump 12 incorrect")
	}
}
