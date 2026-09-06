package excelize

func parseShapeOptions(opts *Shape) (*Shape, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) AddShape(sheet string, opts *Shape) error { _ = "STUB: not implemented"; return nil }

func (f *File) cellAnchorShape(sheet, drawingXML, cell string, width, height uint, format GraphicOptions) (*xlsxWsDr, *xdrCellAnchor, int, error) {
	_ = "STUB: not implemented"
	return nil, nil, 0, nil
}

func (f *File) addDrawingShape(sheet, drawingXML, cell string, opts *Shape) error {
	_ = "STUB: not implemented"
	return nil
}

func setShapeRef(color string, i int) *aRef { _ = "STUB: not implemented"; return nil }
