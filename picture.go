package excelize

import (
	"image"
)

type PictureInsertType byte

const (
	PictureInsertTypePlaceOverCells PictureInsertType = iota
	PictureInsertTypePlaceInCell
	PictureInsertTypeIMAGE
	PictureInsertTypeDISPIMG
)

func (opts *GraphicOptions) parseGraphicOptions(defaults *GraphicOptions) (*GraphicOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parsePictureOptions(pic *Picture) (*GraphicOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) AddPicture(sheet, cell, name string, opts *GraphicOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) AddPictureFromBytes(sheet, cell string, pic *Picture) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addSheetLegacyDrawing(sheet string, rID int) { _ = "STUB: not implemented"; return }

func (f *File) addSheetLegacyDrawingHF(sheet string, rID int) { _ = "STUB: not implemented"; return }

func (f *File) addSheetDrawing(sheet string, rID int) { _ = "STUB: not implemented"; return }

func (f *File) addSheetPicture(sheet string, rID int) error { _ = "STUB: not implemented"; return nil }

func (f *File) countDrawings() int { _ = "STUB: not implemented"; return 0 }

func (f *File) addDrawingPicture(sheet, drawingXML, cell, ext string, rID, hyperlinkRID int, img image.Config, opts *GraphicOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) countMedia() int { _ = "STUB: not implemented"; return 0 }

func (f *File) addMedia(file []byte, ext string) string { _ = "STUB: not implemented"; return "" }

func (f *File) GetPictures(sheet, cell string) ([]Picture, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) GetPictureCells(sheet string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) DeletePicture(sheet, cell string) error { _ = "STUB: not implemented"; return nil }

func (f *File) getPicture(row, col int, drawingXML, drawingRelationships string) (pics []Picture, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) extractPictureFromAnchor(drawingRelationships string, a *xdrCellAnchor, r *xlsxRelationship) *Picture {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) calculatePictureScale(pic *Picture, cx, cy int) { _ = "STUB: not implemented"; return }

func (f *File) extractPictureFromDecodeAnchor(drawingRelationships string, a *decodeCellAnchor, r *xlsxRelationship) *Picture {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) extractCellAnchor(anchor *xdrCellAnchor, drawingRelationships string,
	cond func(from *xlsxFrom) bool,
	cb func(anchor *xdrCellAnchor, rels *xlsxRelationship, drawingRelationships string),
	cond2 func(from *decodeFrom) bool,
	cb2 func(anchor *decodeCellAnchor, rels *xlsxRelationship, drawingRelationships string),
) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractDecodeCellAnchor(anchor *xdrCellAnchor, drawingRelationships string,
	cond func(from *decodeFrom) bool, cb func(anchor *decodeCellAnchor, rels *xlsxRelationship, drawingRelationships string),
) {
	_ = "STUB: not implemented"
	return
}

func (f *File) getDrawingRelationships(rels, rID string) *xlsxRelationship {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawingsWriter() { _ = "STUB: not implemented"; return }

func (f *File) drawingResize(sheet, cell string, width, height float64, opts *GraphicOptions) (w, h, c, r int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0, nil
}

func (f *File) getPictureCells(drawingXML, drawingRelationships string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) cellImagesReader() (*decodeCellImages, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getImageCells(sheet string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getRichDataRichValueRel(val string) (*xlsxRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getRichDataWebImagesRel(val string) (*xlsxRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getImageCellRel(c *xlsxC, pic *Picture) (*xlsxRelationship, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *xlsxRichValueStructure) getRichDataValueIdx(n string) int {
	_ = "STUB: not implemented"
	return 0
}

func (f *File) getCellImages(sheet, cell string) ([]Picture, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getDispImages(sheet, cell string) ([]Picture, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
