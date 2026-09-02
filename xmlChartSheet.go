package excelize

import "encoding/xml"

type xlsxChartsheet struct {
	XMLName          xml.Name                   `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main chartsheet"`
	SheetPr          *xlsxChartsheetPr          `xml:"sheetPr"`
	SheetViews       *xlsxChartsheetViews       `xml:"sheetViews"`
	SheetProtection  *xlsxChartsheetProtection  `xml:"sheetProtection"`
	CustomSheetViews *xlsxCustomChartsheetViews `xml:"customSheetViews"`
	PageMargins      *xlsxPageMargins           `xml:"pageMargins"`
	PageSetup        *xlsxPageSetUp             `xml:"pageSetup"`
	HeaderFooter     *xlsxHeaderFooter          `xml:"headerFooter"`
	Drawing          *xlsxDrawing               `xml:"drawing"`
	DrawingHF        *xlsxDrawingHF             `xml:"drawingHF"`
	Picture          *xlsxPicture               `xml:"picture"`
	WebPublishItems  *xlsxInnerXML              `xml:"webPublishItems"`
	ExtLst           *xlsxExtLst                `xml:"extLst"`
}

type xlsxChartsheetPr struct {
	XMLName       xml.Name   `xml:"sheetPr"`
	PublishedAttr bool       `xml:"published,attr,omitempty"`
	CodeNameAttr  string     `xml:"codeName,attr,omitempty"`
	TabColor      *xlsxColor `xml:"tabColor"`
}

type xlsxChartsheetViews struct {
	XMLName   xml.Name              `xml:"sheetViews"`
	SheetView []*xlsxChartsheetView `xml:"sheetView"`
	ExtLst    []*xlsxExtLst         `xml:"extLst"`
}

type xlsxChartsheetView struct {
	XMLName            xml.Name      `xml:"sheetView"`
	TabSelectedAttr    bool          `xml:"tabSelected,attr,omitempty"`
	ZoomScaleAttr      uint32        `xml:"zoomScale,attr,omitempty"`
	WorkbookViewIDAttr uint32        `xml:"workbookViewId,attr"`
	ZoomToFitAttr      bool          `xml:"zoomToFit,attr,omitempty"`
	ExtLst             []*xlsxExtLst `xml:"extLst"`
}

type xlsxChartsheetProtection struct {
	XMLName           xml.Name `xml:"sheetProtection"`
	AlgorithmNameAttr string   `xml:"algorithmName,attr,omitempty"`
	HashValueAttr     []byte   `xml:"hashValue,attr,omitempty"`
	SaltValueAttr     []byte   `xml:"saltValue,attr,omitempty"`
	SpinCountAttr     uint32   `xml:"spinCount,attr,omitempty"`
	ContentAttr       bool     `xml:"content,attr,omitempty"`
	ObjectsAttr       bool     `xml:"objects,attr,omitempty"`
}

type xlsxCustomChartsheetViews struct {
	XMLName         xml.Name                    `xml:"customSheetViews"`
	CustomSheetView []*xlsxCustomChartsheetView `xml:"customSheetView"`
}

type xlsxCustomChartsheetView struct {
	XMLName       xml.Name            `xml:"customSheetView"`
	GUIDAttr      string              `xml:"guid,attr"`
	ScaleAttr     uint32              `xml:"scale,attr,omitempty"`
	StateAttr     string              `xml:"state,attr,omitempty"`
	ZoomToFitAttr bool                `xml:"zoomToFit,attr,omitempty"`
	PageMargins   []*xlsxPageMargins  `xml:"pageMargins"`
	PageSetup     []*xlsxPageSetUp    `xml:"pageSetup"`
	HeaderFooter  []*xlsxHeaderFooter `xml:"headerFooter"`
}
