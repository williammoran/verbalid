package verbalid

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestVerbalIDFuzz(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for tn := 0; tn < 100; tn++ {
		t.Run(fmt.Sprintf("%d", tn), func(t *testing.T) {
			o := rand.Int()
			iv := IntToVerbalID(o)
			r := VerbalIDToInt(iv)
			if r != o {
				t.Logf("VID = '%s'", iv)
				t.Fatalf("%d != %d", r, o)
			}
		})
	}
}

func TestVerbalIDValues(t *testing.T) {
	cases := []struct {
		i int
		s string
	}{
		{},
		{i: 5, s: "C"},
		{i: 35, s: "4P"},
	}
	for id, tc := range cases {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			rs := IntToVerbalID(tc.i)
			if rs != tc.s {
				t.Logf("%d -> '%s' but should be '%s'", tc.i, rs, tc.s)
				t.Fail()
			}
			ri := VerbalIDToInt(tc.s)
			if ri != tc.i {
				t.Logf("'%s' -> %d but should be %d", tc.s, ri, tc.i)
				t.Fail()
			}
		})
	}
}

func TestVerbalIDCase(t *testing.T) {
	r := VerbalIDToInt("a3X")
	const expected = 1783
	if r != expected {
		t.Fatalf("%d != %d", r, expected)
	}
}

func TestVerbalIDInvalidPanics(t *testing.T) {
	defer func() {
		rec := recover()
		if rec == nil {
			t.Fatal("Invalid char did not panic")
		}
	}()
	_ = VerbalIDToInt("aBx")
}

func TestIntPow(t *testing.T) {
	cases := []struct {
		x, y, e int
	}{
		{x: 0, y: 2, e: 0},
		{x: 42, y: 0, e: 1},
		{x: 2, y: 2, e: 4},
		{x: 3, y: 3, e: 27},
		{x: 10, y: 4, e: 10000},
		{x: 5, y: 2, e: 25},
	}
	for id, tc := range cases {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			r := intPow(tc.x, tc.y)
			if r != tc.e {
				t.Fatalf("intPow(%d, %d) = %d != %d", tc.x, tc.y, r, tc.e)
			}
		})
	}
}
