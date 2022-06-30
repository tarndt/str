package str

import (
	"io"
	"strings"
	"testing"
)

//The tests below seek to confirm abstraction overhead is lower, or complex
// implementations justify their existance with gained performance.

func BenchmarkStringRefComplex(b *testing.B) {
	const (
		expected = "THIS-IS-AN-EXAMPLE"
		prefix   = "preFIX-"
		suffix   = "-SUFFix"
		in       = "\t" + prefix + "This,Is,An,Example" + suffix + "\n"
	)

	for i := 0; i < b.N; i++ {
		out := strings.Join(strings.Split(strings.ToUpper(strings.TrimSuffix(strings.ToLower(strings.ToLower(strings.TrimPrefix(strings.TrimSpace(in), prefix))), strings.ToLower(suffix))), ","), "-")
		if out != expected {
			b.Fatalf("%q became %q rather than %q as expected", in, out, expected)
		}
	}
}

func BenchmarkStrRealComplex(b *testing.B) {
	const (
		expected     = "THIS-IS-AN-EXAMPLE"
		prefix   Str = "preFIX-"
		suffix   Str = "-SUFFix"
		in           = "\t" + prefix + "This,Is,An,Example" + suffix + "\n"
	)

	for i := 0; i < b.N; i++ {
		out := Str(in).TrimSpace().TrimPrefix(prefix).ToLower().TrimSuffix(suffix.ToLower()).ToUpper().Split(",").Join("-")
		if out != expected {
			b.Fatalf("%q became %q rather than %q as expected", in, out, expected)
		}
	}
}

func BenchmarkStringWriteRef(b *testing.B) {
	const str = "mississippi is a river."

	for i := 0; i < b.N; i++ {
		io.Discard.Write([]byte(str))
	}
}

func BenchmarkStringWriteTo(b *testing.B) {
	const str = Str("mississippi is a river.")

	for i := 0; i < b.N; i++ {
		str.WriteTo(io.Discard)
	}
}

var refSlice = []string{"The", "fox", "runs", "very", "fast", "in", "the", "moonlight", "."}

func BenchmarkSliceNewRef(b *testing.B) {
	//slice := append(make(Strs, len(refSlice)), refSlice...) //Doesn't work!
	//slice := make(Strs, len(refSlice))
	//copy(slice, refSlice)  //Does work either...

	for i := 0; i < b.N; i++ {
		slice := make(Strs, 0, len(refSlice))
		for _, str := range refSlice {
			slice = append(slice, Str(str))
		}
		if len(slice) != len(refSlice) {
			b.Fatalf("Bug: Length mismatch")
		}
	}
}

func BenchmarkSliceNewOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := Slice(refSlice...)
		if len(slice) != len(refSlice) { //Ensure compiler didn't eliminate any work
			b.Fatalf("Bug: Length mismatch")
		}
	}
}
