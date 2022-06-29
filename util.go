package str

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"unsafe"
)

func (s String) String() string {
	return string(s)
}

func (s String) Bytes() []byte {
	return []byte(s)
}

func (s String) Buffer() *bytes.Buffer {
	return bytes.NewBuffer(s.Bytes())
}

func (s String) Reader() io.Reader {
	return strings.NewReader(string(s))
}

func (s String) ReaderCloser() io.ReadCloser {
	return io.NopCloser(s.Reader())
}

var _ io.WriterTo = String("")

func (s String) WriteTo(wtr io.Writer) (int64, error) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))

	var slice []byte
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHdr.Data, sliceHdr.Cap, sliceHdr.Len = uintptr(strHdr.Data), strHdr.Len, strHdr.Len

	n, err := wtr.Write(slice)
	return int64(n), err
}
