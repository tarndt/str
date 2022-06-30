package str

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

var basicTestCases = [][]string{
	nil,
	[]string{},
	[]string{"singleton"},
	[]string{"a", "tuple"},
	[]string{"is", "a", "triple"},
}

func TestSliceStringsLen(t *testing.T) {
	t.Parallel()

	for _, tcase := range basicTestCases {
		t.Run(fmt.Sprint(tcase), func(t *testing.T) {
			slice := Slice(tcase...)
			if len(tcase) != slice.Len() {
				t.Fatalf("Test input %v with length %d did not match output %v with length %d", tcase, len(tcase), slice, len(slice))
			}

			out := slice.Strings()
			if !reflect.DeepEqual(tcase, out) {
				t.Fatalf("Test input %v did not match output %v", tcase, slice)
			}
		})
	}
}

func TestSliceString(t *testing.T) {
	for _, tcase := range basicTestCases {
		t.Run(fmt.Sprint(tcase), func(t *testing.T) {
			if expected, actual := fmt.Sprint(tcase), Slice(tcase...); expected != actual.String() {
				t.Fatalf("For %v expected .String() was %q, got: %q", tcase, expected, actual)
			} else if Str(actual.String()) != actual.Str() {
				t.Fatalf("For %v .String() and .Str() differened in content", tcase)
			}
		})
	}
}

func TestSliceEqual(t *testing.T) {
	t.Parallel()

	sameSlice := Slice(basicTestCases[4]...)
	tcases := []struct {
		desc  Str
		slice [2]Strs
		equal bool
	}{
		{desc: "same slice", slice: [2]Strs{sameSlice, sameSlice}, equal: true},
		{desc: "nil and nil", slice: [2]Strs{nil, nil}, equal: true},
		{desc: "empty and nil", slice: [2]Strs{Strs{}, nil}, equal: true},
		{desc: "nil and empty", slice: [2]Strs{nil, Strs{}}, equal: true},
		{desc: "singleton and nil", slice: [2]Strs{Strs{"a"}, nil}, equal: false},
		{desc: "nil and singleton", slice: [2]Strs{nil, Strs{"a"}}, equal: false},
		{desc: "singleton and empty", slice: [2]Strs{Strs{"a"}, Strs{}}, equal: false},
		{desc: "empty and singleton", slice: [2]Strs{Strs{}, Strs{"a"}}, equal: false},
		{desc: "singleton match", slice: [2]Strs{Strs{"a"}, Strs{"a"}}, equal: true},
		{desc: "singleton mismatch", slice: [2]Strs{Strs{"a"}, Strs{"b"}}, equal: false},
		{desc: "singleton mismatch reverse", slice: [2]Strs{Strs{"b"}, Strs{"a"}}, equal: false},
		{desc: "tuple and singleton", slice: [2]Strs{Strs{"a", "b"}, Strs{"a"}}, equal: false},
		{desc: "singleton and tuple", slice: [2]Strs{Strs{"a"}, Strs{"a", "b"}}, equal: false},
		{desc: "tuple match", slice: [2]Strs{Strs{"a", "b"}, Strs{"a", "b"}}, equal: true},
		{desc: "tuple mismatch", slice: [2]Strs{Strs{"a", "b"}, Strs{"b", "a"}}, equal: false},
		{desc: "triple and tuple", slice: [2]Strs{Strs{"a", "b", "c"}, Strs{"a", "b"}}, equal: false},
		{desc: "tuple and triple", slice: [2]Strs{Strs{"a", "b"}, Strs{"a", "b", "c"}}, equal: false},
		{desc: "triple match", slice: [2]Strs{Strs{"a", "b", "c"}, Strs{"a", "b", "c"}}, equal: true},
		{desc: "triple mismatch first of first", slice: [2]Strs{Strs{"x", "b", "c"}, Strs{"a", "b", "c"}}, equal: false},
		{desc: "triple mismatch first of second", slice: [2]Strs{Strs{"a", "b", "c"}, Strs{"x", "b", "c"}}, equal: false},
		{desc: "triple mismatch second of first", slice: [2]Strs{Strs{"a", "x", "c"}, Strs{"a", "b", "c"}}, equal: false},
		{desc: "triple mismatch second of second", slice: [2]Strs{Strs{"a", "b", "c"}, Strs{"a", "x", "c"}}, equal: false},
		{desc: "triple mismatch third of first", slice: [2]Strs{Strs{"a", "b", "x"}, Strs{"a", "b", "c"}}, equal: false},
		{desc: "triple mismatch third of second", slice: [2]Strs{Strs{"a", "b", "c"}, Strs{"a", "b", "x"}}, equal: false},
		{desc: "triple mismatch all", slice: [2]Strs{Strs{"a", "b", "c"}, Strs{"1", "2", "3"}}, equal: false},
	}

	for _, tcase := range tcases {
		t.Run(tcase.desc.String(), func(t *testing.T) {
			if tcase.slice[0].Equal(tcase.slice[1]) != tcase.equal {
				t.Fatalf("%s == %s should be %t, but was not!", tcase.slice[0], tcase.slice[1], tcase.equal)
			}
		})
	}
}

func TestSliceAppend(t *testing.T) {
	t.Parallel()

	slice := Strs{}
	if !slice.Equal(nil) {
		t.Fatal("Empty Strings should be equal to nil")
	}

	slice = slice.Append("a")
	if !slice.Equal(Strs{"a"}) {
		t.Fatal("Singtleton Strings should be equal")
	}

	slice = slice.Append("b")
	if !slice.Equal(Strs{"a", "b"}) {
		t.Fatal("Tuple Strings should be equal")
	}

	if !slice.Append("c", "d").Equal(Strs{"a", "b", "c", "d"}) {
		t.Fatal("Multi-append Strings should be equal")
	}

	if !slice.Equal(Strs{"a", "b"}) {
		t.Fatal("Append should not have modified original slice")
	}
}

func TestSliceAppendAny(t *testing.T) {
	t.Parallel()

	slice, ref := Strs{}, (Strs)(nil)
	if !slice.Equal(nil) {
		t.Fatal("Empty Strings should be equal to nil")
	}

	slice, ref = slice.AppendAny("a"), Strs{"a"}
	if !slice.Equal(ref) {
		t.Fatalf("string primitive append failed: %s", slice)
	}

	slice, ref = slice.AppendAny(Str("b")), append(ref, "b")
	if !slice.Equal(ref) {
		t.Fatalf("String type append failed: %s", slice)
	}

	primStr := "c"
	slice, ref = slice.AppendAny(&primStr), append(ref, "c")
	if !slice.Equal(ref) {
		t.Fatalf("string pointer type append failed: %s", slice)
	}

	typeStr := Str("d")
	slice, ref = slice.AppendAny(&typeStr), append(ref, "d")
	if !slice.Equal(ref) {
		t.Fatalf("String pointer append failed: %s", slice)
	}

	slice, ref = slice.AppendAny(3), append(ref, "3")
	if !slice.Equal(ref) {
		t.Fatalf("int append failed: %s", slice)
	}

	slice, ref = slice.AppendAny(3.14), append(ref, "3.14")
	if !slice.Equal(ref) {
		t.Fatalf("float append failed: %s", slice)
	}

	slice, ref = slice.AppendAny(true), append(ref, "true")
	if !slice.Equal(ref) {
		t.Fatalf("bool append failed: %s", slice)
	}

	slice, ref = slice.AppendAny(point{10, 5}), append(ref, "(10,5)")
	if !slice.Equal(ref) {
		t.Fatalf("Stringer append failed: %s", slice)
	}

	slice, ref = slice.AppendAny(fmt.Errorf("error")), append(ref, "error")
	if !slice.Equal(ref) {
		t.Fatalf("Error append failed: %s", slice)
	}

	slice, ref = slice.AppendAny([]byte("bytes")), append(ref, "bytes")
	if !slice.Equal(ref) {
		t.Fatalf("Error append failed: %s", slice)
	}

	slice, ref = slice.AppendAny([]rune("runes")), append(ref, "runes")
	if !slice.Equal(ref) {
		t.Fatalf("Error append failed: %s", slice)
	}
}

type point struct {
	x, y int
}

func (pt point) String() string {
	return Newf("(%d,%d)", pt.x, pt.y).String()
}

func TestSliceJoin(t *testing.T) {
	t.Parallel()

	const sep = ","

	for _, tcase := range basicTestCases {
		t.Run(fmt.Sprint(tcase), func(t *testing.T) {
			if expected, actual := strings.Join(tcase, sep), Slice(tcase...).Join(sep); expected != actual.String() {
				t.Fatalf("For %v expected .Join() was %q, got: %q", tcase, expected, actual)
			}
		})
	}
}

func TestSliceMaps(t *testing.T) {
	t.Parallel()

	f1 := func(s string) string {
		return "*" + s + "#"
	}
	f2 := func(s Str) Str {
		return "*" + s + "#"
	}

	t.Run("map", func(t *testing.T) {
		for _, tcase := range basicTestCases {
			t.Run(fmt.Sprint(tcase), func(t *testing.T) {
				clone := make([]string, len(tcase))
				copy(clone, tcase)
				slice := Slice(clone...)

				slice.Map(f2)
				for i := range tcase {
					if expected, actual := f1(tcase[i]), slice[i]; expected != actual.String() {
						t.Fatalf("For %v expected mapped value to be %q, got: %q", tcase, expected, actual)
					}
				}
			})
		}
	})

	t.Run("mapnew", func(t *testing.T) {
		for _, tcase := range basicTestCases {
			t.Run(fmt.Sprint(tcase), func(t *testing.T) {
				clone := make([]string, len(tcase))
				copy(clone, tcase)

				slice := Slice(clone...).MapNew(f2)
				if unsafe.Pointer(&slice) == unsafe.Pointer(&clone) {
					t.Fatalf("Returned map should have been newly allocated")
				}

				for i := range tcase {
					if expected, actual := f1(tcase[i]), slice[i]; expected != actual.String() {
						t.Fatalf("For %v expected mapped value to be %q, got: %q", tcase, expected, actual)
					}
				}
			})
		}
	})
}
