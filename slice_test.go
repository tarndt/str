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

func TestStrsStringsLen(t *testing.T) {
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

func TestStrsClone(t *testing.T) {
	orig := Strs{"a", "b", "c"}
	clone := orig.Clone()

	if unsafe.Pointer(&orig) == unsafe.Pointer(&clone) {
		t.Fatalf("Returned slice should have been newly allocated")
	}
}

func TestStrsString(t *testing.T) {
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

func TestStrsEqual(t *testing.T) {
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

func TestStrsAppend(t *testing.T) {
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

func TestStrsAppendAny(t *testing.T) {
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

func TestStrsIndexContainsRemoveStr(t *testing.T) {
	t.Parallel()

	tcases := []struct {
		desc, target Str
		slice        Strs
		expected     int
	}{
		{desc: "nil", target: "a", slice: nil, expected: -1},
		{desc: "empty", target: "a", slice: Strs{}, expected: -1},
		{desc: "singleton-match", target: "a", slice: Strs{"a"}, expected: 0},
		{desc: "singleton-nomatch", target: "a", slice: Strs{"b"}, expected: -1},
		{desc: "tuple-match-first", target: "a", slice: Strs{"a", "b"}, expected: 0},
		{desc: "tuple-match-second", target: "a", slice: Strs{"b", "a"}, expected: 1},
		{desc: "tuple-match-none", target: "a", slice: Strs{"x", "y"}, expected: -1},
		{desc: "triple-match-first", target: "a", slice: Strs{"a", "b", "c"}, expected: 0},
		{desc: "triple-match-second", target: "b", slice: Strs{"a", "b", "c"}, expected: 1},
		{desc: "triple-match-third", target: "c", slice: Strs{"a", "b", "c"}, expected: 2},
		{desc: "triple-match-none", target: "x", slice: Strs{"a", "b", "c"}, expected: -1},
	}

	for _, tcase := range tcases {
		t.Run(tcase.desc.String(), func(t *testing.T) {
			t.Run("contains", func(t *testing.T) {
				switch tcase.expected {
				case -1:
					if tcase.slice.Contains(tcase.target) {
						t.Fatalf("%s should not have contained %q but Contains returned true", tcase.slice, tcase.target)
					}
				default:
					if !tcase.slice.Contains(tcase.target) {
						t.Fatalf("%s should have contained %q but Contains returned false", tcase.slice, tcase.target)
					}
				}
			})

			t.Run("index", func(t *testing.T) {
				idx := tcase.slice.Index(tcase.target)
				if idx != tcase.expected {
					t.Fatalf("%s should have was expected to have %q at index %d but Index returned %d", tcase.slice, tcase.target, tcase.expected, idx)
				}
			})

			t.Run("removestr", func(t *testing.T) {
				origLen := tcase.slice.Len()
				removed := tcase.slice.RemoveStr(tcase.target)
				switch tcase.expected {
				case -1:
					if removed.Len() != origLen {
						t.Fatalf("After remove %s should have been the same size", removed)
					}
				default:
					if removed.Len() != origLen-1 {
						t.Fatalf("After remove %s should have contained %d rather than %d elements", removed, origLen-1, origLen)
					}
				}

				if rmIdx := removed.Index(tcase.target); rmIdx != -1 {
					t.Fatalf("After remove %s should not have contained %q but Index found it at %d", tcase.slice, tcase.target, rmIdx)
				}
			})

		})
	}
}

func TestStrsJoin(t *testing.T) {
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

func TestStrsMaps(t *testing.T) {
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

				slice := Slice(clone...).MappedNew(f2)
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

func TestStrsSorts(t *testing.T) {
	t.Parallel()

	oooSlice := Strs{"a", "d", "c", "b", "f"}

	properySorted := func(t *testing.T, oooSlice, sortedSlice Strs) {
		elems := make(map[Str]bool)
		for _, str := range oooSlice {
			elems[str] = false
		}

		var last Str
		for _, str := range sortedSlice {
			if last != "" && str < last {
				t.Fatalf("Found out of order element %q that was > preceding element %q", str, last)
			}

			if _, present := elems[str]; !present {
				t.Fatalf("Found element %q that did not exist in %s", str, oooSlice)
			}
			elems[str] = true
		}

		for elem, found := range elems {
			if !found {
				t.Fatalf("Element %q was missing in sorted slice %s but was in original %s", elem, sortedSlice, oooSlice)
			}
		}
	}

	t.Run("sort", func(t *testing.T) {
		sortedSlice := oooSlice.Clone()
		sortedSlice.Sort()
		properySorted(t, oooSlice, sortedSlice)
	})

	t.Run("sortedNew", func(t *testing.T) {
		sortedSlice := oooSlice.SortedNew()
		properySorted(t, oooSlice, sortedSlice)
		if unsafe.Pointer(&oooSlice) == unsafe.Pointer(&sortedSlice) {
			t.Fatalf("Returned slice should have been newly allocated")
		}
	})
}
