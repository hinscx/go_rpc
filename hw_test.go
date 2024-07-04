package main

import (
	"math/rand"
	"testing"

	"github.com/hinscx/goecho/echov2/newecho"
)


func TestEcho(t *testing.T) {

	want := "testin"
	got := newecho.Echo(want)
	rand.Int()
	if want != got {
		t.Errorf("got %s want %s given, %s", got, want, want)
	}
}