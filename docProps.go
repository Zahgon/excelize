package excelize

func (f *File) SetAppProps(appProperties *AppProperties) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetAppProps() (ret *AppProperties, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) SetDocProps(docProperties *DocProperties) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetDocProps() (ret *DocProperties, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) SetCustomProps(prop CustomProperty) error { _ = "STUB: not implemented"; return nil }

func (prop *xlsxProperty) setCustomProps(value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetCustomProps() ([]CustomProperty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *decodeProperty) getCustomProps() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
