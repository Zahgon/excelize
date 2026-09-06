package excelize

func (f *File) getSheetView(sheet string, viewIndex int) (*xlsxSheetView, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (view *xlsxSheetView) setSheetView(opts *ViewOptions) { _ = "STUB: not implemented"; return }

func (f *File) SetSheetView(sheet string, viewIndex int, opts *ViewOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetSheetView(sheet string, viewIndex int) (ViewOptions, error) {
	_ = "STUB: not implemented"
	return *new(ViewOptions), nil
}
