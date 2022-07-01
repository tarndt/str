package str

import (
	"sort"
	"strings"
	"unsafe"
)

//Strs is a slice of Str representing a list of Strs.
type Strs []Str

//Slice  constructs a new list of Str.
func Slice(strs ...string) Strs {
	return Strs(*((*[]Str)(unsafe.Pointer(&strs))))
}

//Clone creates a copy of this slice
func (s Strs) Clone() Strs {
	clone := make(Strs, len(s))
	copy(clone, s)
	return clone
}

//Strings converts this slice of Strs to a slice of strings in an optimized manner.
func (s Strs) Strings() []string {
	return *((*[]string)(unsafe.Pointer(&s)))
}

//Str returns a single Str representing the content of this slice
// formatted the same as the default string format of a slice.
func (s Strs) Str() Str {
	return "[" + s.Join(" ") + "]"
}

//String returns a single string representing the content of this slice
// formatted the same as the default string format of a slice.
func (s Strs) String() string {
	return string(s.Str())
}

//Len returns the cardinality of the slice. Equivalent to to len(slice).
func (s Strs) Len() int {
	return len(s)
}

//Equal tests if this slice of Strs is equal in content to another.
func (s Strs) Equal(other Strs) bool {
	strLen := len(s)
	switch {
	case strLen != len(other):
		return false
	case strLen < 1:
		return true
	case &s[0] == &other[0]:
		return true
	}

	for i, v := range s {
		if v != other[i] {
			return false
		}
	}
	return true
}

//Append is similar to the built-in append function but infers the first argument
// to be itself.
func (s Strs) Append(others ...Str) Strs {
	return append(s, others...)
}

//AppendAny is like Append but uses reflection to convert the arguments to Strs/
func (s Strs) AppendAny(others ...any) Strs {
	slice := s
	for _, val := range others {
		slice = append(slice, New(val))
	}
	return slice
}

//Remove removes the ith element from the slice and returns it.
func (s Strs) Remove(i int) Strs {
	lastIdx := len(s) - 1
	if i < 0 || i > lastIdx {
		return s
	}

	copy(s[i:], s[i+1:])
	return s[:lastIdx]
}

//RemoveStr removes the first instance the the Str found from the slice and returns it.
func (s Strs) RemoveStr(str Str) Strs {
	return s.Remove(s.Index(str))
}

//Contains returns if the slice contains the Str
func (s Strs) Contains(str Str) bool {
	return s.Index(str) >= 0
}

//Index returns the index of the first occurrence of the Str
func (s Strs) Index(str Str) int {
	for i, elem := range s {
		if elem == str {
			return i
		}
	}
	return -1
}

//Join concatenates the elements of its first argument to create a single string
// delimited by sep. See strings.Join.
func (s Strs) Join(sep Str) Str {
	return Str(strings.Join(s.Strings(), string(sep)))
}

//Map applies the provided transformation function to sever element of the slice
// in place.
func (s Strs) Map(f func(s Str) Str) {
	for i, str := range s {
		s[i] = f(str)
	}
}

//MappedNew is like Map but rather than modifying this slice a new one is created
// from the mapping.
func (s Strs) MappedNew(f func(s Str) Str) Strs {
	mapped := make(Strs, 0, len(s))
	for _, str := range s {
		mapped = append(mapped, f(str))
	}
	return mapped
}

//Sort sorts the Strs in this slice in ascending order
func (s Strs) Sort() {
	sort.Strings(s.Strings())
}

//SortedNew is like Sort but rather than modifying this slice a new one sorted
// ascending is created
func (s Strs) SortedNew() Strs {
	clone := s.Clone()
	clone.Sort()
	return clone
}
