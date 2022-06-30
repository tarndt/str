package str

import (
	"bytes"
	"fmt"
	"io"
)

//String is a string with all the strings.* functions as methods
type Str string

//New returns a new String from am arbitrary Go value. Use a simple cast,
// Str(s string), for variables known to be strings to avoid reflection overhead.
func New(val any) Str {
	switch val := val.(type) {
	case Str:
		return val
	case string:
		return Str(val)
	case []byte:
		return Str(val)
	case []rune:
		return Str(val)
	case fmt.Stringer:
		return Str(val.String())
	case error:
		return Str(val.Error())
	case *string:
		return Str(*val)
	default:
		return Str(fmt.Sprint(val))
	}
}

//Newf formats according to a format specifier and returns the resulting Str.
// See the similar fmt.Sprintf.
func Newf(format string, args ...any) Str {
	return Str(fmt.Sprintf(format, args...))
}

//ReadFrom consumes the provided reader and returns a Str with its content. To
//limit read bytes wrap the reader or use ReadFromLimit.
func ReadFrom(rdr io.Reader) (Str, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(rdr)
	if err != nil {
		return "", err
	}

	return Str(buf.String()), nil
}

//ReadFromLimit is like ReadFrom but limits the maximum number of bytes read from
// the provided reader.
func ReadFromLimit(rdr io.Reader, n int64) (Str, error) {
	return ReadFrom(&io.LimitedReader{R: rdr, N: n})
}
