package excelize

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"strings"
	"time"
)

type StreamWriter struct {
	file            *File
	Sheet           string
	SheetID         int
	sheetWritten    bool
	worksheet       *xlsxWorksheet
	rawData         bufferedWriter
	rows            int
	mergeCellsCount int
	mergeCells      strings.Builder
	tableParts      string
}

func (f *File) NewStreamWriter(sheet string) (*StreamWriter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sw *StreamWriter) AddTable(table *Table) error { _ = "STUB: not implemented"; return nil }

func (sw *StreamWriter) getRowValues(hRow, hCol, vCol int) (res []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getRowElement(token xml.Token, hRow int) (startElement xml.StartElement, ok bool) {
	_ = "STUB: not implemented"
	return *new(xml.StartElement), false
}

type Cell struct {
	StyleID int
	Formula string
	Value   interface{}
}

type RowOpts struct {
	Height       float64
	Hidden       bool
	StyleID      int
	OutlineLevel int
}

func (r *RowOpts) marshalAttrs() (strings.Builder, error) {
	_ = "STUB: not implemented"
	return *new(strings.Builder), nil
}

func parseRowOpts(opts ...RowOpts) *RowOpts { _ = "STUB: not implemented"; return nil }

func (sw *StreamWriter) SetRow(cell string, values []interface{}, opts ...RowOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *StreamWriter) SetColVisible(minVal, maxVal int, visible bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *StreamWriter) SetColOutlineLevel(col int, level uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *StreamWriter) SetColStyle(minVal, maxVal, styleID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *StreamWriter) SetColWidth(minVal, maxVal int, width float64) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *StreamWriter) InsertPageBreak(cell string) error { _ = "STUB: not implemented"; return nil }

func (sw *StreamWriter) SetPanes(panes *Panes) error { _ = "STUB: not implemented"; return nil }

func (sw *StreamWriter) MergeCell(topLeftCell, bottomRightCell string) error {
	_ = "STUB: not implemented"
	return nil
}

func setCellFormula(c *xlsxC, formula string) { _ = "STUB: not implemented"; return }

func (sw *StreamWriter) setCellTime(c *xlsxC, val time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *StreamWriter) setCellDuration(c *xlsxC, val time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (sw *StreamWriter) setCellValFunc(c *xlsxC, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func setCellIntFunc(c *xlsxC, val interface{}) { _ = "STUB: not implemented"; return }

func writeCell(buf *bufferedWriter, c xlsxC) { _ = "STUB: not implemented"; return }

func (sw *StreamWriter) writeSheetData() { _ = "STUB: not implemented"; return }

func (sw *StreamWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func bulkAppendFields(w io.Writer, ws *xlsxWorksheet, from, to int) {
	_ = "STUB: not implemented"
	return
}

type bufferedWriter struct {
	tmpDir string
	tmp    *os.File
	buf    bytes.Buffer
}

func (bw *bufferedWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bw *bufferedWriter) WriteString(p string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bw *bufferedWriter) Reader() (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (bw *bufferedWriter) Sync() (err error) { _ = "STUB: not implemented"; return nil }

func (bw *bufferedWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (bw *bufferedWriter) Close() error { _ = "STUB: not implemented"; return nil }
