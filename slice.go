package str

import (
	"strings"
	"unsafe"
)

type Strings []String

func Slice(strs ...string) Strings {
	return Strings(*((*[]String)(unsafe.Pointer(&strs))))
}

func (s Strings) Strings() []string {
	return *((*[]string)(unsafe.Pointer(&s)))
}

func (s Strings) Len() int {
	return len(s)
}

func (s Strings) Append(others ...String) Strings {
	return append(s, others...)
}

func (s Strings) AppendVal(others ...any) Strings {
	slice := s
	for _, val := range others {
		slice = append(slice, New(val))
	}
	return slice
}

func (s Strings) Join(sep String) String {
	return String(strings.Join(s.Strings(), string(sep)))
}

func (s Strings) Map(f func(s String) String) {
	for i, str := range s {
		s[i] = f(str)
	}
}

func (s Strings) MapNew(f func(s String) String) Strings {
	mapped := make(Strings, 0, len(s))
	for _, str := range s {
		mapped = append(mapped, f(str))
	}
	return mapped
}
