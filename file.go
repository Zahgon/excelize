package excelize

import (
	"bytes"
	"io"
)

func NewFile(opts ...Options) *File { _ = "STUB: not implemented"; return nil }

func (f *File) Save(opts ...Options) error { _ = "STUB: not implemented"; return nil }

func (f *File) SaveAs(name string, opts ...Options) error { _ = "STUB: not implemented"; return nil }

func (f *File) Close() error { _ = "STUB: not implemented"; return nil }

func (f *File) Write(w io.Writer, opts ...Options) error { _ = "STUB: not implemented"; return nil }

func (f *File) WriteTo(w io.Writer, opts ...Options) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) WriteToBuffer() (*bytes.Buffer, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) writeToZip(zw ZipWriter) error { _ = "STUB: not implemented"; return nil }

func (f *File) writeZip64LFH(buf *bytes.Buffer) error { _ = "STUB: not implemented"; return nil }
