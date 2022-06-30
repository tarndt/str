package str

import (
	"bytes"
	"fmt"
	"testing"
)

func TestStrNewf(t *testing.T) {
	if expected, actual := Str(fmt.Sprintf("a %s c", "b")), Newf("a %s c", "b"); actual != expected {
		t.Fatalf("Newf resulted in %q rather than %q", actual, expected)
	}
}

func TestStrReadFrom(t *testing.T) {
	t.Parallel()

	const (
		in       = "mississippi is a river."
		expected = Str(in)
	)

	buf := bytes.NewBufferString(in)
	actual, err := ReadFrom(buf)
	if err != nil {
		t.Fatalf("ReadFrom(%q) failed: %s", in, err)
	}

	if actual != expected {
		t.Fatalf("ReadFrom(%q) resulted in %q rather than %q", in, actual, expected)
	}

	t.Run("limit", func(t *testing.T) {
		in := Str("0").Repeat(11)
		buf := bytes.NewBufferString(in.String())
		actual, err := ReadFromLimit(buf, int64(in.Len()-1))
		if err != nil {
			t.Fatalf("ReadFrom(%q) failed: %s", in, err)
		}

		if expected := in[0 : in.Len()-1]; actual != expected {
			t.Fatalf("ReadFrom(%q) resulted in %q rather than %q", in, actual, expected)
		}
	})
}
