package str

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"unsafe"
)

func (s Str) Len() int {
	return len(s)
}

func (s Str) String() string {
	return string(s)
}

func (s Str) Bytes() []byte {
	return []byte(s)
}

func (s Str) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(s.Bytes())
}

func (s Str) Reader() io.Reader {
	return strings.NewReader(string(s))
}

func (s Str) ReaderCloser() io.ReadCloser {
	return io.NopCloser(s.Reader())
}

var _ io.WriterTo = Str("")

func (s Str) WriteTo(wtr io.Writer) (int64, error) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var slice []byte
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHdr.Data, sliceHdr.Cap, sliceHdr.Len = uintptr(strHdr.Data), strHdr.Len, strHdr.Len

	n, err := wtr.Write(slice)
	return int64(n), err
}
