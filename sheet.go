package excelize

import (
	"encoding/xml"
)

type IgnoredErrorsType byte

const (
	IgnoredErrorsEvalError = iota
	IgnoredErrorsTwoDigitTextYear
	IgnoredErrorsNumberStoredAsText
	IgnoredErrorsFormula
	IgnoredErrorsFormulaRange
	IgnoredErrorsUnlockedFormula
	IgnoredErrorsEmptyCellReference
	IgnoredErrorsListDataValidation
	IgnoredErrorsCalculatedColumn
)

func (f *File) NewSheet(sheet string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *File) contentTypesReader() (*xlsxTypes, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) contentTypesWriter() { _ = "STUB: not implemented"; return }

func (f *File) getWorksheetPath(relTarget string) (path string) {
	_ = "STUB: not implemented"
	return ""
}

func (f *File) mergeExpandedCols(ws *xlsxWorksheet) { _ = "STUB: not implemented"; return }

func (f *File) workSheetWriter() { _ = "STUB: not implemented"; return }

func trimRow(sheetData *xlsxSheetData) []xlsxRow { _ = "STUB: not implemented"; return nil }

func trimCell(row xlsxRow) xlsxRow { _ = "STUB: not implemented"; return *new(xlsxRow) }

func (f *File) setContentTypes(partName, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) setSheet(index int, name string) { _ = "STUB: not implemented"; return }

func (f *File) relsWriter() { _ = "STUB: not implemented"; return }

func replaceRelationshipsBytes(content []byte) []byte { _ = "STUB: not implemented"; return nil }

func (f *File) SetActiveSheet(index int) { _ = "STUB: not implemented"; return }

func (f *File) GetActiveSheetIndex() (index int) { _ = "STUB: not implemented"; return 0 }

func (f *File) getActiveSheetID() int { _ = "STUB: not implemented"; return 0 }

func (f *File) SetSheetName(source, target string) error { _ = "STUB: not implemented"; return nil }

func (f *File) GetSheetName(index int) (name string) { _ = "STUB: not implemented"; return "" }

func (f *File) getSheetID(sheet string) int { _ = "STUB: not implemented"; return 0 }

func (f *File) GetSheetIndex(sheet string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (f *File) GetSheetMap() map[int]string { _ = "STUB: not implemented"; return nil }

func (f *File) GetSheetList() (list []string) { _ = "STUB: not implemented"; return nil }

func (f *File) getSheetMap() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) getSheetXMLPath(sheet string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (f *File) SetSheetBackground(sheet, picture string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) SetSheetBackgroundFromBytes(sheet, extension string, picture []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) setSheetBackground(sheet, extension string, file []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) DeleteSheet(sheet string) error { _ = "STUB: not implemented"; return nil }

func (f *File) MoveSheet(source, target string) error { _ = "STUB: not implemented"; return nil }

func deleteAndAdjustDefinedNames(wb *xlsxWorkbook, deleteLocalSheetID int) {
	_ = "STUB: not implemented"
	return
}

func (f *File) deleteSheetFromWorkbookRels(rID string) string { _ = "STUB: not implemented"; return "" }

func (f *File) deleteSheetRelationships(sheet, rID string) { _ = "STUB: not implemented"; return }

func (f *File) getSheetRelationshipsTargetByID(sheet, rID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *File) CopySheet(from, to int) error { _ = "STUB: not implemented"; return nil }

func (f *File) copySheet(from, to int) error { _ = "STUB: not implemented"; return nil }

func getSheetState(visible bool, veryHidden []bool) string { _ = "STUB: not implemented"; return "" }

func (f *File) SetSheetVisible(sheet string, visible bool, veryHidden ...bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) setPanes(panes *Panes) error { _ = "STUB: not implemented"; return nil }

func (f *File) SetPanes(sheet string, panes *Panes) error { _ = "STUB: not implemented"; return nil }

func (ws *xlsxWorksheet) getPanes() Panes { _ = "STUB: not implemented"; return *new(Panes) }

func (f *File) GetPanes(sheet string) (Panes, error) {
	_ = "STUB: not implemented"
	return *new(Panes), nil
}

func (f *File) GetSheetVisible(sheet string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *File) SearchSheet(sheet, value string, reg ...bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) searchSheet(name, value string, regSearch bool) (result []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func attrValToInt(name string, attrs []xml.Attr) (val int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func attrValToFloat(name string, attrs []xml.Attr) (val float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func attrValToBool(name string, attrs []xml.Attr) (val bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (f *File) SetHeaderFooter(sheet string, opts *HeaderFooterOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetHeaderFooter(sheet string) (*HeaderFooterOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) ProtectSheet(sheet string, opts *SheetProtectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) UnprotectSheet(sheet string, password ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetSheetProtection(sheet string) (SheetProtectionOptions, error) {
	_ = "STUB: not implemented"
	return *new(SheetProtectionOptions), nil
}

func checkSheetName(name string) error { _ = "STUB: not implemented"; return nil }

func (f *File) SetPageLayout(sheet string, opts *PageLayoutOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (ws *xlsxWorksheet) newPageSetUp() { _ = "STUB: not implemented"; return }

func (ws *xlsxWorksheet) setPageSetUp(opts *PageLayoutOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetPageLayout(sheet string) (PageLayoutOptions, error) {
	_ = "STUB: not implemented"
	return *new(PageLayoutOptions), nil
}

func (f *File) SetDefinedName(definedName *DefinedName) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) DeleteDefinedName(definedName *DefinedName) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetDefinedName() []DefinedName { _ = "STUB: not implemented"; return nil }

func (f *File) GroupSheets(sheets []string) error { _ = "STUB: not implemented"; return nil }

func (f *File) UngroupSheets() error { _ = "STUB: not implemented"; return nil }

func (f *File) InsertPageBreak(sheet, cell string) error { _ = "STUB: not implemented"; return nil }

func (ws *xlsxWorksheet) insertPageBreak(cell string) error { _ = "STUB: not implemented"; return nil }

func (f *File) RemovePageBreak(sheet, cell string) error { _ = "STUB: not implemented"; return nil }

func (f *File) relsReader(path string) (*xlsxRelationships, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ws *xlsxWorksheet) prepareSheetXML(col, row int) { _ = "STUB: not implemented"; return }

func fillColumns(rowData *xlsxRow, col, row int) { _ = "STUB: not implemented"; return }

func (ws *xlsxWorksheet) makeContiguousColumns(fromRow, toRow, colCount int) {
	_ = "STUB: not implemented"
	return
}

func (f *File) SetSheetDimension(sheet, rangeRef string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetSheetDimension(sheet string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) AddIgnoredErrors(sheet, rangeRef string, ignoredErrorsType IgnoredErrorsType) error {
	_ = "STUB: not implemented"
	return nil
}
