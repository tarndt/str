package str

import (
	"bytes"
	"fmt"
	"io"
)

//String is a string with all the strings.* functions as methods
type String string

//New returns a new String from a Go value. Use String(string) for strings to
// avoid reflection overhead.
func New(val any) String {
	switch val := val.(type) {
	case String:
		return val
	case string:
		return String(val)
	case []byte:
		return String(val)
	case []rune:
		return String(val)
	case fmt.Stringer:
		return String(val.String())
	case error:
		return String(val.Error())
	case *String:
		return *val
	case *string:
		return String(*val)
	default:
		return String(fmt.Sprint(val))
	}
}

func Newf(format string, args ...any) String {
	return String(fmt.Sprintf(format, args...))
}

func ReadFrom(rdr io.Reader) (String, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(rdr)
	if err != nil {
		return "", err
	}

	return String(buf.String()), nil
}

func ReadFromLimit(rdr io.Reader, n int64) (String, error) {
	return ReadFrom(&io.LimitedReader{R: rdr, N: n})
}
