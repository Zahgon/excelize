package excelize

func (f *File) SetWorkbookProps(opts *WorkbookPropsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetWorkbookProps() (WorkbookPropsOptions, error) {
	_ = "STUB: not implemented"
	return *new(WorkbookPropsOptions), nil
}

func (f *File) SetCalcProps(opts *CalcPropsOptions) error { _ = "STUB: not implemented"; return nil }

func (f *File) GetCalcProps() (CalcPropsOptions, error) {
	_ = "STUB: not implemented"
	return *new(CalcPropsOptions), nil
}

func (f *File) ProtectWorkbook(opts *WorkbookProtectionOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) UnprotectWorkbook(password ...string) error { _ = "STUB: not implemented"; return nil }

func (f *File) setWorkbook(name string, sheetID, rid int) { _ = "STUB: not implemented"; return }

func (f *File) getWorkbookPath() (path string) { _ = "STUB: not implemented"; return "" }

func (f *File) getWorkbookRelsPath() (path string) { _ = "STUB: not implemented"; return "" }

func (f *File) deleteWorkbookRels(relType, relTarget string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) workbookReader() (*xlsxWorkbook, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) workBookWriter() { _ = "STUB: not implemented"; return }

func (f *File) setContentTypePartRelsExtensions() error { _ = "STUB: not implemented"; return nil }

func (f *File) setContentTypePartImageExtensions() error { _ = "STUB: not implemented"; return nil }

func (f *File) setContentTypePartVMLExtensions() error { _ = "STUB: not implemented"; return nil }

func (f *File) addContentTypePart(index int, contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) removeContentTypesPart(contentType, partName string) error {
	_ = "STUB: not implemented"
	return nil
}
