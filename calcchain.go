package excelize

func (f *File) calcChainReader() (*xlsxCalcChain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) calcChainWriter() { _ = "STUB: not implemented"; return }

func (f *File) deleteCalcChain(index int, cell string) error { _ = "STUB: not implemented"; return nil }

type xlsxCalcChainCollection []xlsxCalcChainC

func (c xlsxCalcChainCollection) Filter(fn func(v xlsxCalcChainC) bool) []xlsxCalcChainC {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) volatileDepsReader() (*xlsxVolTypes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) volatileDepsWriter() { _ = "STUB: not implemented"; return }

func (vt *xlsxVolTypes) deleteVolTopicRef(i1, i2, i3, i4 int) { _ = "STUB: not implemented"; return }
