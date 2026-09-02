package excelize

import (
	"encoding/xml"
)

type SlicerOptions struct {
	slicerXML       string
	slicerCacheXML  string
	slicerCacheName string
	slicerSheetName string
	slicerSheetRID  string
	drawingXML      string
	Name            string
	Cell            string
	TableSheet      string
	TableName       string
	Caption         string
	Macro           string
	Width           uint
	Height          uint
	DisplayHeader   *bool
	ItemDesc        bool
	Format          GraphicOptions
	SelectedItems   []string
}

func (f *File) AddSlicer(sheet string, opts *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func parseSlicerOptions(opts *SlicerOptions) (*SlicerOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) countSlicers() int { _ = "STUB: not implemented"; return 0 }

func (f *File) countSlicerCache() int { _ = "STUB: not implemented"; return 0 }

func (f *File) getSlicerSource(opts *SlicerOptions) (*Table, *PivotTableOptions, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

func (f *File) addSheetSlicer(sheet, extURI string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) addSheetTableSlicer(ws *xlsxWorksheet, rID int, extURI string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addSlicer(slicerID int, slicer xlsxSlicer) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) genSlicerName(name string) string { _ = "STUB: not implemented"; return "" }

func (f *File) genSlicerCacheName(name string) string { _ = "STUB: not implemented"; return "" }

func (f *File) setSlicerCache(colIdx int, opts *SlicerOptions, table *Table, pivotTable *PivotTableOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) slicerReader(slicerXML string) (*xlsxSlicers, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) slicerCacheReader(slicerCacheXML string) (*xlsxSlicerCacheDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) timelineReader(timelineXML string) (*xlsxTimelines, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) buildSlicerItems(pivotTable *PivotTableOptions, opts *SlicerOptions) ([]xlsxTabularSlicerCacheItem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) addSlicerCache(slicerCacheName string, colIdx int, opts *SlicerOptions, table *Table, pivotTable *PivotTableOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addPivotCacheSlicer(opts *PivotTableOptions) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) addDrawingSlicer(sheet, slicerName string, ns xml.Attr, opts *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addWorkbookSlicerCache(slicerCacheID int, URI string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetSlicers(sheet string) ([]SlicerOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getSlicerCache(slicerCacheName string, opt *SlicerOptions) *xlsxSlicerCacheDefinition {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getSlicers(sheet, rID, drawingXML string) ([]SlicerOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) extractTableSlicer(slicerCache *xlsxSlicerCacheDefinition, opt *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) extractPivotTableSlicer(slicerCache *xlsxSlicerCacheDefinition, opt *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) extractSlicerSelectedItems(pivotCacheXML string, slicerCache *xlsxSlicerCacheDefinition, opt *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) extractSlicerCellAnchor(drawingXML string, opt *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) extractSlicerFromAnchor(anchor *xdrCellAnchor, opt *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) extractSlicerFromDecodeAnchor(anchor *xdrCellAnchor, opt *SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getAllSlicers() (map[string][]SlicerOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) DeleteSlicer(name string) error { _ = "STUB: not implemented"; return nil }

func (f *File) deleteSlicer(opts SlicerOptions) error { _ = "STUB: not implemented"; return nil }

func (f *File) deleteSlicerCache(sles map[string][]SlicerOptions, opts SlicerOptions) error {
	_ = "STUB: not implemented"
	return nil
}
