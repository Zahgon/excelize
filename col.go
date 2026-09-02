package excelize

import (
	"encoding/xml"
)

type Cols struct {
	err                                    error
	curCol, totalCols, totalRows, stashCol int
	rawCellValue                           bool
	sheet                                  string
	f                                      *File
	sheetXML                               []byte
	sst                                    *xlsxSST
}

func (f *File) GetCols(sheet string, opts ...Options) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cols *Cols) Next() bool { _ = "STUB: not implemented"; return false }

func (cols *Cols) Error() error { _ = "STUB: not implemented"; return nil }

func (cols *Cols) Rows(opts ...Options) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type columnXMLIterator struct {
	err                  error
	cols                 Cols
	cellCol, curRow, row int
}

func columnXMLHandler(colIterator *columnXMLIterator, xmlElement *xml.StartElement) {
	_ = "STUB: not implemented"
	return
}

func (cols *Cols) rowXMLHandler(rowIterator *rowXMLIterator, xmlElement *xml.StartElement, decoder *xml.Decoder) {
	_ = "STUB: not implemented"
	return
}

func (f *File) Cols(sheet string) (*Cols, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) GetColVisible(sheet, col string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *File) SetColVisible(sheet, columns string, visible bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) setColVisible(minVal, maxVal int, visible bool) {
	_ = "STUB: not implemented"
	return
}

func (f *File) GetColOutlineLevel(sheet, col string) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) parseColRange(columns string) (minVal, maxVal int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (f *File) SetColOutlineLevel(sheet, col string, level uint8) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) setColOutlineLevel(colNum int, level uint8) {
	_ = "STUB: not implemented"
	return
}

func (f *File) SetColStyle(sheet, columns string, styleID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) setColStyle(minVal, maxVal, styleID int) {
	_ = "STUB: not implemented"
	return
}

func (f *File) SetColWidth(sheet, startCol, endCol string, width float64) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) setColWidth(minVal, maxVal int, width float64) {
	_ = "STUB: not implemented"
	return
}

func flatCols(col xlsxCol, cols []xlsxCol, replacer func(fc, c xlsxCol) xlsxCol) []xlsxCol {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) positionObjectPixels(sheet string, col, row, width, height int, opts *GraphicOptions) (int, int, int, int, int, int, int, int) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, 0, 0, 0, 0
}

func (f *File) getColWidth(sheet string, col int) int { _ = "STUB: not implemented"; return 0 }

func (f *File) GetColStyle(sheet, col string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) GetColWidth(sheet, col string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) InsertCols(sheet, col string, n int) error { _ = "STUB: not implemented"; return nil }

func (f *File) RemoveCol(sheet, col string) error { _ = "STUB: not implemented"; return nil }

func convertColWidthToPixels(width float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fnt *Font) calcTextWidth(text string) float64 { _ = "STUB: not implemented"; return 0 }

func (fnt *Font) calcRichTextWidth(runs []RichTextRun) float64 { _ = "STUB: not implemented"; return 0 }

func (f *File) autoFitColWidth(sheet string, col, rows int, defaultFnt *Font) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) AutoFitColWidth(sheet, columns string) error { _ = "STUB: not implemented"; return nil }
