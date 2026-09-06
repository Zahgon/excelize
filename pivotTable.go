package excelize

import (
	"encoding/xml"
)

type PivotTableOptions struct {
	items               map[string][]*xlsxItem
	sharedItems         map[string]xlsxSharedItems
	pivotTableXML       string
	pivotCacheXML       string
	pivotSheetName      string
	pivotDataRange      string
	namedDataRange      bool
	DataRange           string
	PivotTableRange     string
	Name                string
	Rows                []PivotTableField
	Columns             []PivotTableField
	Data                []PivotTableField
	Filter              []PivotTableField
	RowGrandTotals      bool
	ColGrandTotals      bool
	ShowDrill           bool
	UseAutoFormatting   bool
	PageOverThenDown    bool
	MergeItem           bool
	ClassicLayout       bool
	CompactData         bool
	ShowError           bool
	ShowRowHeaders      bool
	ShowColHeaders      bool
	ShowRowStripes      bool
	ShowColStripes      bool
	ShowLastColumn      bool
	FieldPrintTitles    bool
	ItemPrintTitles     bool
	PivotTableStyleName string
}

type PivotTableShowValuesAsType byte

const (
	PivotTableShowValuesAsNoCalculation PivotTableShowValuesAsType = iota
	PivotTableShowValuesAsPercentOfGrandTotal
	PivotTableShowValuesAsPercentOfColumnTotal
	PivotTableShowValuesAsPercentOfRowTotal
	PivotTableShowValuesAsPercentOf
	PivotTableShowValuesAsPercentOfParentRowTotal
	PivotTableShowValuesAsPercentOfParentColumnTotal
	PivotTableShowValuesAsPercentOfParentTotal
	PivotTableShowValuesAsDifferenceFrom
	PivotTableShowValuesAsPercentDifferenceFrom
	PivotTableShowValuesAsRunningTotalIn
	PivotTableShowValuesAsPercentRunningTotalIn
	PivotTableShowValuesAsRankSmallestToLargest
	PivotTableShowValuesAsRankLargestToSmallest
	PivotTableShowValuesAsIndex
)

type PivotTableShowValuesAs struct {
	Type      PivotTableShowValuesAsType
	BaseField string
	BaseItem  string
}

type PivotTableField struct {
	Compact         bool
	Data            string
	Name            string
	Outline         bool
	ShowAll         bool
	InsertBlankRow  bool
	Subtotal        string
	DefaultSubtotal bool
	NumFmt          int
	SelectedItems   []string
	ShowValuesAs    PivotTableShowValuesAs
}

var (
	pivotTableShowValuesAsMap = map[PivotTableShowValuesAsType]string{
		PivotTableShowValuesAsPercentOfGrandTotal:        "percentOfTotal",
		PivotTableShowValuesAsPercentOfColumnTotal:       "percentOfCol",
		PivotTableShowValuesAsPercentOfRowTotal:          "percentOfRow",
		PivotTableShowValuesAsPercentOf:                  "percent",
		PivotTableShowValuesAsPercentOfParentRowTotal:    "percentOfParentRow",
		PivotTableShowValuesAsPercentOfParentColumnTotal: "percentOfParentCol",
		PivotTableShowValuesAsPercentOfParentTotal:       "percentOfParent",
		PivotTableShowValuesAsDifferenceFrom:             "difference",
		PivotTableShowValuesAsPercentDifferenceFrom:      "percentDiff",
		PivotTableShowValuesAsRunningTotalIn:             "runTotal",
		PivotTableShowValuesAsPercentRunningTotalIn:      "percentOfRunningTotal",
		PivotTableShowValuesAsRankSmallestToLargest:      "rankAscending",
		PivotTableShowValuesAsRankLargestToSmallest:      "rankDescending",
		PivotTableShowValuesAsIndex:                      "index",
	}
	pivotTableShowValuesAsX14Map = map[PivotTableShowValuesAsType]bool{
		PivotTableShowValuesAsPercentOfParentRowTotal:    true,
		PivotTableShowValuesAsPercentOfParentColumnTotal: true,
		PivotTableShowValuesAsPercentOfParentTotal:       true,
		PivotTableShowValuesAsPercentRunningTotalIn:      true,
		PivotTableShowValuesAsRankSmallestToLargest:      true,
		PivotTableShowValuesAsRankLargestToSmallest:      true,
		PivotTableShowValuesAsIndex:                      true,
	}
	pivotTableShowValuesAsBaseFieldRequiredMap = map[PivotTableShowValuesAsType]bool{
		PivotTableShowValuesAsPercentOf:             true,
		PivotTableShowValuesAsPercentOfParentTotal:  true,
		PivotTableShowValuesAsDifferenceFrom:        true,
		PivotTableShowValuesAsPercentDifferenceFrom: true,
		PivotTableShowValuesAsRunningTotalIn:        true,
		PivotTableShowValuesAsPercentRunningTotalIn: true,
		PivotTableShowValuesAsRankSmallestToLargest: true,
		PivotTableShowValuesAsRankLargestToSmallest: true,
	}
	pivotTableShowValuesAsBaseItemRequiredMap = map[PivotTableShowValuesAsType]bool{
		PivotTableShowValuesAsPercentOf:             true,
		PivotTableShowValuesAsDifferenceFrom:        true,
		PivotTableShowValuesAsPercentDifferenceFrom: true,
	}
)

func (f *File) AddPivotTable(opts *PivotTableOptions) error { _ = "STUB: not implemented"; return nil }

func (f *File) parseFormatPivotTableSet(opts *PivotTableOptions) (*xlsxWorksheet, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (f *File) adjustRange(rangeStr string) (string, []int, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (f *File) getTableFieldsOrder(opts *PivotTableOptions) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (si xlsxSharedItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (si *xlsxSharedItems) addMissingItem() { _ = "STUB: not implemented"; return }

func (si *xlsxSharedItems) addNumberItem(val string) { _ = "STUB: not implemented"; return }

func (si *xlsxSharedItems) addBooleanItem(val string) { _ = "STUB: not implemented"; return }

func (si *xlsxSharedItems) addErrorItem(val string) { _ = "STUB: not implemented"; return }

func (si *xlsxSharedItems) addStringItem(val string) { _ = "STUB: not implemented"; return }

func (si *xlsxSharedItems) checkSelectedItems(field string, selectedItems []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addSharedItems(sheet string, col, fromRow, toRow int) (xlsxSharedItems, error) {
	_ = "STUB: not implemented"
	return *new(xlsxSharedItems), nil
}

func (f *File) buildPivotSharedItems(opts *PivotTableOptions, idx int, coordinates []int, field PivotTableField) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addPivotSharedItems(opts *PivotTableOptions, coordinates []int, fieldsType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addPivotCache(opts *PivotTableOptions) error { _ = "STUB: not implemented"; return nil }

func (f *File) addPivotTable(cacheID, pivotTableID int, opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addPivotRowFields(pt *xlsxPivotTableDefinition, opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addPivotPageFields(pt *xlsxPivotTableDefinition, opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addPivotDataFields(pt *xlsxPivotTableDefinition, opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (df *xlsxDataField) setPivotTableShowValuesAs(idx int, order []string, opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (df *xlsxDataField) setPivotTableShowValuesAsBaseItem(baseField, baseItem string, sharedItems *xlsxSharedItems) error {
	_ = "STUB: not implemented"
	return nil
}

func inPivotTableField(a []PivotTableField, x string) int { _ = "STUB: not implemented"; return 0 }

func (f *File) addPivotColFields(pt *xlsxPivotTableDefinition, opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (fld *xlsxPivotField) setClassicLayout(classicLayout bool) { _ = "STUB: not implemented"; return }

func (f *File) addPivotFields(pt *xlsxPivotTableDefinition, opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) countPivotTables() int { _ = "STUB: not implemented"; return 0 }

func (f *File) countPivotCache() int { _ = "STUB: not implemented"; return 0 }

func (f *File) getPivotFieldsIndex(fields []PivotTableField, opts *PivotTableOptions) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getPivotTableFieldsSubtotal(fields []PivotTableField) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getPivotTableFieldsName(fields []PivotTableField) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getPivotTableFieldName(name string, fields []PivotTableField) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *File) getPivotTableFieldsNumFmtID(fields []PivotTableField) []int {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getPivotTableFieldOptions(name string, fields []PivotTableField) (options PivotTableField, ok bool) {
	_ = "STUB: not implemented"
	return *new(PivotTableField), false
}

func (f *File) addWorkbookPivotCache(RID int) int { _ = "STUB: not implemented"; return 0 }

func (f *File) GetPivotTables(sheet string) ([]PivotTableOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getPivotTableDataRange(opts *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getPivotTable(sheet, pivotTableXML, pivotCacheRels string) (PivotTableOptions, error) {
	_ = "STUB: not implemented"
	return *new(PivotTableOptions), nil
}

func (pc *xlsxPivotCacheDefinition) getPivotCacheFieldsName() []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) pivotTableReader(path string) (*xlsxPivotTableDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) pivotCacheReader(path string) (*xlsxPivotCacheDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) extractPivotTableFields(pt *xlsxPivotTableDefinition, pc *xlsxPivotCacheDefinition, opts *PivotTableOptions) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractPivotTableShowValuesAs(pc *xlsxPivotCacheDefinition, df *xlsxDataField, dataField *PivotTableField) {
	_ = "STUB: not implemented"
	return
}

func (pc *xlsxPivotCacheDefinition) extractPivotTableField(data string, fld *xlsxPivotField) PivotTableField {
	_ = "STUB: not implemented"
	return *new(PivotTableField)
}

func (f *File) genPivotCacheDefinitionID() int { _ = "STUB: not implemented"; return 0 }

func (f *File) deleteWorkbookPivotCache(opt PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) DeletePivotTable(sheet, name string) error { _ = "STUB: not implemented"; return nil }

func (f *File) getPivotTables() (map[string][]PivotTableOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
