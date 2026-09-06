package excelize

func (mc *xlsxMergeCell) Rect() ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) MergeCell(sheet, topLeftCell, bottomRightCell string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) UnmergeCell(sheet, topLeftCell, bottomRightCell string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetMergeCells(sheet string, withoutValues ...bool) ([]MergeCell, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ws *xlsxWorksheet) mergeOverlapCells() error { _ = "STUB: not implemented"; return nil }

type MergeCell []string

func (m *MergeCell) GetCellValue() string { _ = "STUB: not implemented"; return "" }

func (m *MergeCell) GetStartAxis() string { _ = "STUB: not implemented"; return "" }

func (m *MergeCell) GetEndAxis() string { _ = "STUB: not implemented"; return "" }
