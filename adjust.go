package excelize

import (
	"github.com/xuri/efp"
)

type adjustDirection bool

const (
	columns adjustDirection = false
	rows    adjustDirection = true
)

var adjustHelperFunc = [9]func(*File, *xlsxWorksheet, string, adjustDirection, int, int, int) error{
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustConditionalFormats(ws, sheet, dir, num, offset, sheetID)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustDataValidations(ws, sheet, dir, num, offset, sheetID)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustDefinedNames(sheet, dir, num, offset)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustDrawings(ws, sheet, dir, num, offset)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustMergeCells(ws, sheet, dir, num, offset, sheetID)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustAutoFilter(ws, sheet, dir, num, offset, sheetID)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustCalcChain(ws, sheet, dir, num, offset, sheetID)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustTable(ws, sheet, dir, num, offset, sheetID)
	},
	func(f *File, ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
		return f.adjustVolatileDeps(ws, sheet, dir, num, offset, sheetID)
	},
}

func (f *File) adjustHelper(sheet string, dir adjustDirection, num, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustCols(ws *xlsxWorksheet, col, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustColDimensions(sheet string, ws *xlsxWorksheet, col, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustRowDimensions(sheet string, ws *xlsxWorksheet, row, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *xlsxRow) adjustSingleRowDimensions(offset int) { _ = "STUB: not implemented"; return }

func (f *File) adjustSingleRowFormulas(sheet, sheetN string, r *xlsxRow, num, offset int, si bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustCellRef(cellRef string, dir adjustDirection, num, offset int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) adjustFormula(sheet, sheetN string, cell *xlsxC, dir adjustDirection, num, offset int, si bool) error {
	_ = "STUB: not implemented"
	return nil
}

func escapeSheetName(name string) string { _ = "STUB: not implemented"; return "" }

func adjustFormulaColumnName(name, operand string, abs, keepRelative bool, dir adjustDirection, num, offset int) (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func adjustFormulaRowNumber(name, operand string, abs, keepRelative bool, dir adjustDirection, num, offset int) (string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func adjustFormulaOperandRef(row, col, operand string, abs, keepRelative bool, dir adjustDirection, num int, offset int) (string, string, string, bool, error) {
	_ = "STUB: not implemented"
	return "", "", "", false, nil
}

func (f *File) adjustFormulaOperand(sheet, sheetN string, keepRelative bool, token efp.Token, dir adjustDirection, num int, offset int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isExternalSheetReference(ref string) bool { _ = "STUB: not implemented"; return false }

func (f *File) adjustFormulaRef(sheet, sheetN, formula string, keepRelative bool, dir adjustDirection, num, offset int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func transformParenthesesToken(token efp.Token) string { _ = "STUB: not implemented"; return "" }

func adjustRangeSheetName(rng, source, target string) string { _ = "STUB: not implemented"; return "" }

type arrayFormulaOperandToken struct {
	operandTokenIndex, topLeftCol, topLeftRow, bottomRightCol, bottomRightRow int
	sheetName, sourceCellRef, targetCellRef                                   string
}

func (af *arrayFormulaOperandToken) setCoordinates() error { _ = "STUB: not implemented"; return nil }

func transformArrayFormula(tokens []efp.Token, afs []arrayFormulaOperandToken) string {
	_ = "STUB: not implemented"
	return ""
}

func getArrayFormulaTokens(sheet, formula string, definedNames []DefinedName) ([]efp.Token, []arrayFormulaOperandToken, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (f *File) adjustHyperlinks(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset int) {
	_ = "STUB: not implemented"
	return
}

func (f *File) adjustTable(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustAutoFilter(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustAutoFilterHelper(dir adjustDirection, coordinates []int, num, offset int) []int {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustMergeCells(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustMergeCellsHelper(p1, p2, num, offset int) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (f *File) deleteMergeCell(ws *xlsxWorksheet, idx int) { _ = "STUB: not implemented"; return }

func adjustCellName(cell string, dir adjustDirection, c, r, offset int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) adjustCalcChain(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (vt *xlsxVolTypes) adjustVolatileDepsTopic(cell string, dir adjustDirection, indexes []int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) adjustVolatileDeps(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustConditionalFormats(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustDataValidations(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset, sheetID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (from *xlsxFrom) adjustDrawings(dir adjustDirection, num, offset int, editAs string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (to *xlsxTo) adjustDrawings(dir adjustDirection, num, offset int, ok bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *xdrCellAnchor) adjustDrawings(dir adjustDirection, num, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *xlsxCellAnchorPos) adjustDrawings(dir adjustDirection, num, offset int, editAs string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustDrawings(ws *xlsxWorksheet, sheet string, dir adjustDirection, num, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) adjustDefinedNames(sheet string, dir adjustDirection, num, offset int) error {
	_ = "STUB: not implemented"
	return nil
}
