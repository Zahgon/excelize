package excelize

import (
	"encoding/xml"
	"time"
)

type CellType byte

const (
	CellTypeUnset CellType = iota
	CellTypeBool
	CellTypeDate
	CellTypeError
	CellTypeFormula
	CellTypeInlineString
	CellTypeNumber
	CellTypeSharedString
)

const (
	STCellFormulaTypeArray = "array"

	STCellFormulaTypeDataTable = "dataTable"

	STCellFormulaTypeNormal = "normal"

	STCellFormulaTypeShared = "shared"
)

var cellTypes = map[string]CellType{
	"b":         CellTypeBool,
	"d":         CellTypeDate,
	"n":         CellTypeNumber,
	"e":         CellTypeError,
	"s":         CellTypeSharedString,
	"str":       CellTypeFormula,
	"inlineStr": CellTypeInlineString,
}

func (f *File) GetCellValue(sheet, cell string, opts ...Options) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) GetCellType(sheet, cell string) (CellType, error) {
	_ = "STUB: not implemented"
	return *new(CellType), nil
}

func (f *File) SetCellValue(sheet, cell string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (x xlsxSI) String() string { _ = "STUB: not implemented"; return "" }

func (c *xlsxC) hasValue() bool { _ = "STUB: not implemented"; return false }

func (f *File) removeFormula(c *xlsxC, ws *xlsxWorksheet, sheet string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) setCellIntFunc(sheet, cell string, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) setCellTimeFunc(sheet, cell string, value time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *xlsxC) setCellTime(value time.Time, date1904 bool) (isNum bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func setCellDuration(value time.Duration) (t string, v string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (f *File) SetCellInt(sheet, cell string, value int64) error {
	_ = "STUB: not implemented"
	return nil
}

func setCellInt(value int64) (t string, v string) { _ = "STUB: not implemented"; return "", "" }

func (f *File) SetCellUint(sheet, cell string, value uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func setCellUint(value uint64) (t string, v string) { _ = "STUB: not implemented"; return "", "" }

func (f *File) SetCellBool(sheet, cell string, value bool) error {
	_ = "STUB: not implemented"
	return nil
}

func setCellBool(value bool) (t string, v string) { _ = "STUB: not implemented"; return "", "" }

func (f *File) SetCellFloat(sheet, cell string, value float64, precision, bitSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *xlsxC) setCellFloat(value float64, precision, bitSize int) {
	_ = "STUB: not implemented"
	return
}

func (f *File) SetCellStr(sheet, cell, value string) error { _ = "STUB: not implemented"; return nil }

func (f *File) setCellString(value string) (t, v string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (f *File) sharedStringsLoader() (err error) { _ = "STUB: not implemented"; return nil }

func (f *File) setSharedString(val string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func trimCellValue(value string, escape bool) (v string, ns xml.Attr) {
	_ = "STUB: not implemented"
	return "", *new(xml.Attr)
}

func (c *xlsxC) setCellValue(val string) { _ = "STUB: not implemented"; return }

func (c *xlsxC) setInlineStr(val string) { _ = "STUB: not implemented"; return }

func (c *xlsxC) setStr(val string) { _ = "STUB: not implemented"; return }

func (c *xlsxC) getCellBool(f *File, raw bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *xlsxC) setCellDefault(value string) { _ = "STUB: not implemented"; return }

func (c *xlsxC) getCellDate(f *File, raw bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *xlsxC) getValueFrom(f *File, d *xlsxSST, raw bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) SetCellDefault(sheet, cell, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetCellFormula(sheet, cell string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) getCellFormula(sheet, cell string, transformed bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type FormulaOpts struct {
	Type *string
	Ref  *string
}

func (f *File) SetCellFormula(sheet, cell, formula string, opts ...FormulaOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func sharedFormulaRefToCoordinates(opts ...FormulaOpts) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ws *xlsxWorksheet) setArrayFormula(sheet string, formula *xlsxF, definedNames []DefinedName) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) setArrayFormulaCells() error { _ = "STUB: not implemented"; return nil }

func (ws *xlsxWorksheet) setSharedFormula(coordinates []int) { _ = "STUB: not implemented"; return }

func (ws *xlsxWorksheet) countSharedFormula() (count int) { _ = "STUB: not implemented"; return 0 }

func (ws *xlsxWorksheet) deleteSharedFormula(c *xlsxC) { _ = "STUB: not implemented"; return }

func (f *File) GetCellHyperLink(sheet, cell string) (bool, string, error) {
	_ = "STUB: not implemented"
	return false, "", nil
}

func (f *File) GetHyperLinkCells(sheet, linkType string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type HyperlinkOpts struct {
	Display *string
	Tooltip *string
}

func (f *File) removeHyperLink(ws *xlsxWorksheet, sheet, cell string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) SetCellHyperLink(sheet, cell, link, linkType string, opts ...HyperlinkOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func getCellRichText(si *xlsxSI) (runs []RichTextRun) { _ = "STUB: not implemented"; return nil }

func (f *File) GetCellRichText(sheet, cell string) (runs []RichTextRun, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fnt *Font) newRpr() *xlsxRPr { _ = "STUB: not implemented"; return nil }

func (rPr *xlsxRPr) getFont() *Font { _ = "STUB: not implemented"; return nil }

func setRichText(runs []RichTextRun) ([]xlsxR, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) SetCellRichText(sheet, cell string, runs []RichTextRun) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) SetSheetRow(sheet, cell string, slice interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) SetSheetCol(sheet, cell string, slice interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) setSheetCells(sheet, cell string, slice interface{}, dir adjustDirection) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) prepareCell(cell string) (*xlsxC, int, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, 0, nil
}

func (f *File) getCellStringFunc(sheet, cell string, fn func(x *xlsxWorksheet, c *xlsxC) (string, bool, error)) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) formattedValue(c *xlsxC, raw bool, cellType CellType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (ss *xlsxStyleSheet) getCustomNumFmtCode(numFmtID int) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ws *xlsxWorksheet) prepareCellStyle(col, row, style int) int {
	_ = "STUB: not implemented"
	return 0
}

func (ws *xlsxWorksheet) mergeCellsParser(cell string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) checkCellInRangeRef(cell, rangeRef string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func cellInRange(cell, ref []int) bool { _ = "STUB: not implemented"; return false }

func isOverlap(rect1, rect2 []int) bool { _ = "STUB: not implemented"; return false }

func (c *xlsxC) convertSharedFormula(cell string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getSharedFormula(ws *xlsxWorksheet, si int, cell string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func shiftCell(val string, dCol, dRow int) string { _ = "STUB: not implemented"; return "" }
