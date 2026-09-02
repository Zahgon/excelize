package excelize

func (f *File) SetPageMargins(sheet string, opts *PageLayoutMarginsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetPageMargins(sheet string) (PageLayoutMarginsOptions, error) {
	_ = "STUB: not implemented"
	return *new(PageLayoutMarginsOptions), nil
}

func (ws *xlsxWorksheet) prepareSheetPr() { _ = "STUB: not implemented"; return }

func (ws *xlsxWorksheet) setSheetOutlineProps(opts *SheetPropsOptions) {
	_ = "STUB: not implemented"
	return
}

func (ws *xlsxWorksheet) setSheetProps(opts *SheetPropsOptions) { _ = "STUB: not implemented"; return }

func (f *File) SetSheetProps(sheet string, opts *SheetPropsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetSheetProps(sheet string) (SheetPropsOptions, error) {
	_ = "STUB: not implemented"
	return *new(SheetPropsOptions), nil
}
