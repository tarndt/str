package str

import (
	"bytes"
	"strings"
	"testing"
)

func TestStrString(t *testing.T) {
	t.Parallel()

	for _, expected := range []string{"", "x", "12", "123"} {
		if actual := Str(expected).String(); actual != expected {
			t.Fatalf("%q.String() was %q rather than %q", expected, actual, expected)
		}
	}
}

func TestStrBytes(t *testing.T) {
	t.Parallel()

	for _, str := range []string{"", "x", "12", "123"} {
		if expected, actual := []byte(str), Str(str).Bytes(); !bytes.Equal(expected, actual) {
			t.Fatalf("%q.Bytes() was %q rather than %q", expected, actual, expected)
		}
	}
}

func TestStrBuffer(t *testing.T) {
	t.Parallel()

	for _, str := range []string{"", "x", "12", "123"} {
		if expected, actual := bytes.NewBufferString(str), Str(str).Buffer(); !bytes.Equal(expected.Bytes(), actual.Bytes()) {
			t.Fatalf("%q.Buffer() was %q rather than %q", str, actual.Bytes(), expected.Bytes())
		}
	}
}

func TestStrReader(t *testing.T) {
	t.Parallel()

	for _, str := range []string{"", "x", "12", "123"} {
		strLen := int64(len(str))
		expectedRdr, actualRdr := strings.NewReader(str), Str(str).ReaderCloser()
		var expected, actual bytes.Buffer
		if n, err := expected.ReadFrom(expectedRdr); err != nil {
			t.Fatalf("Reading from strings reader failed: %s", err)
		} else if n != strLen {
			t.Fatalf("Read wrong number of bytes (%d vs expected %d) from strings reader", n, len(str))
		}

		if n, err := actual.ReadFrom(actualRdr); err != nil {
			t.Fatalf("Reading from String reader failed: %s", err)
		} else if n != strLen {
			t.Fatalf("Read wrong number of bytes (%d vs expected %d) from String reader", n, len(str))
		}

		if !bytes.Equal(expected.Bytes(), actual.Bytes()) {
			t.Fatalf("%q.Buffer() was %q rather than %q", str, actual.Bytes(), expected.Bytes())
		}
		if err := actualRdr.Close(); err != nil {
			t.Fatalf("%q.ReaderCloser().Close() should never fail; got: %s", str, err)
		}
	}
}

func TestStrWriteTo(t *testing.T) {
	t.Parallel()

	const (
		prefix   = "The "
		in       = Str("mississippi is a river.")
		expected = prefix + in
	)

	buf := bytes.NewBufferString(prefix)
	if n, err := in.WriteTo(buf); err != nil {
		t.Fatalf("%q.WriteTo(...) failed: %s", in, err)
	} else if n != int64(in.Len()) {
		t.Fatalf("%q.WriteTo(...) wrote %d rather than %d bytes", in, n, in.Len())
	}
}
