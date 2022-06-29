package str

import (
	"strings"
	"testing"
)

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
		expected        = "THIS-IS-AN-EXAMPLE"
		prefix   String = "preFIX-"
		suffix   String = "-SUFFix"
		in              = "\t" + prefix + "This,Is,An,Example" + suffix + "\n"
	)

	for i := 0; i < b.N; i++ {
		out := String(in).TrimSpace().TrimPrefix(prefix).ToLower().TrimSuffix(suffix.ToLower()).ToUpper().Split(",").Join("-")
		if out != expected {
			b.Fatalf("%q became %q rather than %q as expected", in, out, expected)
		}
	}
}
