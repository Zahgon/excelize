package excelize

import (
	"encoding/xml"
	"os"
)

var duplicateHelperFunc = [3]func(*File, *xlsxWorksheet, string, int, int) error{
	func(f *File, ws *xlsxWorksheet, sheet string, row, row2 int) error {
		return f.duplicateConditionalFormat(ws, sheet, row, row2)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, row, row2 int) error {
		return f.duplicateDataValidations(ws, sheet, row, row2)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, row, row2 int) error {
		return f.duplicateMergeCells(ws, sheet, row, row2)
	},
}

func (f *File) GetRows(sheet string, opts ...Options) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Rows struct {
	err                     error
	curRow, seekRow         int
	needClose, rawCellValue bool
	sheet                   string
	f                       *File
	tempFile                *os.File
	sst                     *xlsxSST
	decoder                 *xml.Decoder
	token                   xml.Token
	curRowOpts, seekRowOpts RowOpts
}

func (rows *Rows) Next() bool { _ = "STUB: not implemented"; return false }

func (rows *Rows) GetRowOpts() RowOpts { _ = "STUB: not implemented"; return *new(RowOpts) }

func (rows *Rows) Error() error { _ = "STUB: not implemented"; return nil }

func (rows *Rows) Close() error { _ = "STUB: not implemented"; return nil }

func (rows *Rows) Columns(opts ...Options) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractRowOpts(attrs []xml.Attr) RowOpts { _ = "STUB: not implemented"; return *new(RowOpts) }

func appendSpace(l int, s []string) []string { _ = "STUB: not implemented"; return nil }

type rowXMLIterator struct {
	err              error
	inElement        string
	cellCol, cellRow int
	cells            []string
}

func (rows *Rows) rowXMLHandler(rowIterator *rowXMLIterator, xmlElement *xml.StartElement, raw bool) {
	_ = "STUB: not implemented"
	return
}

func (cell *xlsxC) cellXMLAttrHandler(start *xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (cell *xlsxC) cellXMLHandler(decoder *xml.Decoder, start *xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) Rows(sheet string) (*Rows, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) getFromStringItem(index int) string { _ = "STUB: not implemented"; return "" }

func (f *File) xmlDecoder(name string) (bool, *xml.Decoder, *os.File, error) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil
}

func (f *File) SetRowHeight(sheet string, row int, height float64) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getRowHeight(sheet string, row int) int { _ = "STUB: not implemented"; return 0 }

func (f *File) GetRowHeight(sheet string, row int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) sharedStringsReader() (*xlsxSST, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) SetRowVisible(sheet string, row int, visible bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetRowVisible(sheet string, row int) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *File) SetRowOutlineLevel(sheet string, row int, level uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetRowOutlineLevel(sheet string, row int) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) RemoveRow(sheet string, row int) error { _ = "STUB: not implemented"; return nil }

func (f *File) InsertRows(sheet string, row, n int) error { _ = "STUB: not implemented"; return nil }

func (f *File) DuplicateRow(sheet string, row int) error { _ = "STUB: not implemented"; return nil }

func (f *File) DuplicateRowTo(sheet string, row, row2 int) error {
	_ = "STUB: not implemented"
	return nil
}

func duplicateSQRefHelper(row, row2 int, ref string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) duplicateConditionalFormat(ws *xlsxWorksheet, sheet string, row, row2 int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) duplicateDataValidations(ws *xlsxWorksheet, sheet string, row, row2 int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) duplicateMergeCells(ws *xlsxWorksheet, sheet string, row, row2 int) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) checkRow() error { _ = "STUB: not implemented"; return nil }

func (r *xlsxRow) hasAttr() bool { _ = "STUB: not implemented"; return false }

func (f *File) SetRowStyle(sheet string, start, end, styleID int) error {
	_ = "STUB: not implemented"
	return nil
}

func convertRowHeightToPixels(height float64) float64 { _ = "STUB: not implemented"; return 0 }
