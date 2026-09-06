package excelize

import "encoding/xml"

type xlsxTable struct {
	XMLName              xml.Name            `xml:"table"`
	XMLNS                string              `xml:"xmlns,attr"`
	ID                   int                 `xml:"id,attr"`
	Name                 string              `xml:"name,attr"`
	DisplayName          string              `xml:"displayName,attr,omitempty"`
	Comment              string              `xml:"comment,attr,omitempty"`
	Ref                  string              `xml:"ref,attr"`
	TableType            string              `xml:"tableType,attr,omitempty"`
	HeaderRowCount       *int                `xml:"headerRowCount,attr"`
	InsertRow            bool                `xml:"insertRow,attr,omitempty"`
	InsertRowShift       bool                `xml:"insertRowShift,attr,omitempty"`
	TotalsRowCount       int                 `xml:"totalsRowCount,attr,omitempty"`
	TotalsRowShown       *bool               `xml:"totalsRowShown,attr"`
	Published            bool                `xml:"published,attr,omitempty"`
	HeaderRowDxfID       int                 `xml:"headerRowDxfId,attr,omitempty"`
	DataDxfID            int                 `xml:"dataDxfId,attr,omitempty"`
	TotalsRowDxfID       int                 `xml:"totalsRowDxfId,attr,omitempty"`
	HeaderRowBorderDxfID int                 `xml:"headerRowBorderDxfId,attr,omitempty"`
	TableBorderDxfID     int                 `xml:"tableBorderDxfId,attr,omitempty"`
	TotalsRowBorderDxfID int                 `xml:"totalsRowBorderDxfId,attr,omitempty"`
	HeaderRowCellStyle   string              `xml:"headerRowCellStyle,attr,omitempty"`
	DataCellStyle        string              `xml:"dataCellStyle,attr,omitempty"`
	TotalsRowCellStyle   string              `xml:"totalsRowCellStyle,attr,omitempty"`
	ConnectionID         int                 `xml:"connectionId,attr,omitempty"`
	AutoFilter           *xlsxAutoFilter     `xml:"autoFilter"`
	TableColumns         *xlsxTableColumns   `xml:"tableColumns"`
	TableStyleInfo       *xlsxTableStyleInfo `xml:"tableStyleInfo"`
}

type xlsxAutoFilter struct {
	XMLName      xml.Name            `xml:"autoFilter"`
	Ref          string              `xml:"ref,attr"`
	FilterColumn []*xlsxFilterColumn `xml:"filterColumn"`
}

type xlsxFilterColumn struct {
	ColID         int                `xml:"colId,attr"`
	HiddenButton  bool               `xml:"hiddenButton,attr,omitempty"`
	ShowButton    bool               `xml:"showButton,attr,omitempty"`
	CustomFilters *xlsxCustomFilters `xml:"customFilters"`
	Filters       *xlsxFilters       `xml:"filters"`
	ColorFilter   *xlsxColorFilter   `xml:"colorFilter"`
	DynamicFilter *xlsxDynamicFilter `xml:"dynamicFilter"`
	IconFilter    *xlsxIconFilter    `xml:"iconFilter"`
	Top10         *xlsxTop10         `xml:"top10"`
}

type xlsxCustomFilters struct {
	And          bool                `xml:"and,attr,omitempty"`
	CustomFilter []*xlsxCustomFilter `xml:"customFilter"`
}

type xlsxCustomFilter struct {
	Operator string `xml:"operator,attr,omitempty"`
	Val      string `xml:"val,attr,omitempty"`
}

type xlsxFilters struct {
	Blank         bool                 `xml:"blank,attr,omitempty"`
	CalendarType  string               `xml:"calendarType,attr,omitempty"`
	Filter        []*xlsxFilter        `xml:"filter"`
	DateGroupItem []*xlsxDateGroupItem `xml:"dateGroupItem"`
}

type xlsxFilter struct {
	Val string `xml:"val,attr,omitempty"`
}

type xlsxColorFilter struct {
	CellColor bool `xml:"cellColor,attr"`
	DxfID     int  `xml:"dxfId,attr"`
}

type xlsxDynamicFilter struct {
	MaxValISO string  `xml:"maxValIso,attr,omitempty"`
	Type      string  `xml:"type,attr,omitempty"`
	Val       float64 `xml:"val,attr,omitempty"`
	ValISO    string  `xml:"valIso,attr,omitempty"`
}

type xlsxIconFilter struct {
	IconID  int    `xml:"iconId,attr"`
	IconSet string `xml:"iconSet,attr,omitempty"`
}

type xlsxTop10 struct {
	FilterVal float64 `xml:"filterVal,attr,omitempty"`
	Percent   bool    `xml:"percent,attr,omitempty"`
	Top       bool    `xml:"top,attr"`
	Val       float64 `xml:"val,attr,omitempty"`
}

type xlsxDateGroupItem struct {
	DateTimeGrouping string `xml:"dateTimeGrouping,attr,omitempty"`
	Day              int    `xml:"day,attr,omitempty"`
	Hour             int    `xml:"hour,attr,omitempty"`
	Minute           int    `xml:"minute,attr,omitempty"`
	Month            int    `xml:"month,attr,omitempty"`
	Second           int    `xml:"second,attr,omitempty"`
	Year             int    `xml:"year,attr,omitempty"`
}

type xlsxTableColumns struct {
	Count       int                `xml:"count,attr"`
	TableColumn []*xlsxTableColumn `xml:"tableColumn"`
}

type xlsxTableColumn struct {
	ID                 int    `xml:"id,attr"`
	UniqueName         string `xml:"uniqueName,attr,omitempty"`
	Name               string `xml:"name,attr"`
	TotalsRowFunction  string `xml:"totalsRowFunction,attr,omitempty"`
	TotalsRowLabel     string `xml:"totalsRowLabel,attr,omitempty"`
	QueryTableFieldID  int    `xml:"queryTableFieldId,attr,omitempty"`
	HeaderRowDxfID     int    `xml:"headerRowDxfId,attr,omitempty"`
	DataDxfID          int    `xml:"dataDxfId,attr,omitempty"`
	TotalsRowDxfID     int    `xml:"totalsRowDxfId,attr,omitempty"`
	HeaderRowCellStyle string `xml:"headerRowCellStyle,attr,omitempty"`
	DataCellStyle      string `xml:"dataCellStyle,attr,omitempty"`
	TotalsRowCellStyle string `xml:"totalsRowCellStyle,attr,omitempty"`
}

type xlsxTableStyleInfo struct {
	Name              string `xml:"name,attr,omitempty"`
	ShowFirstColumn   bool   `xml:"showFirstColumn,attr"`
	ShowLastColumn    bool   `xml:"showLastColumn,attr"`
	ShowRowStripes    bool   `xml:"showRowStripes,attr"`
	ShowColumnStripes bool   `xml:"showColumnStripes,attr"`
}

type xlsxSingleXMLCells struct {
	XMLName       xml.Name            `xml:"singleXmlCells"`
	SingleXmlCell []xlsxSingleXMLCell `xml:"singleXmlCell"`
}

type xlsxSingleXMLCell struct {
	XMLName      xml.Name      `xml:"singleXmlCell"`
	ID           int           `xml:"id,attr"`
	R            string        `xml:"r,attr"`
	ConnectionID int           `xml:"connectionId,attr"`
	XMLCellPr    xlsxXMLCellPr `xml:"xmlCellPr"`
	ExtLst       *xlsxInnerXML `xml:"extLst"`
}

type xlsxXMLCellPr struct {
	XMLName    xml.Name      `xml:"xmlCellPr"`
	ID         int           `xml:"id,attr"`
	UniqueName string        `xml:"uniqueName,attr,omitempty"`
	XMLPr      *xlsxInnerXML `xml:"xmlPr"`
	ExtLst     *xlsxInnerXML `xml:"extLst"`
}

type Table struct {
	tID               int
	rID               string
	tableXML          string
	Range             string
	Name              string
	StyleName         string
	ShowColumnStripes bool
	ShowFirstColumn   bool
	ShowHeaderRow     *bool
	ShowLastColumn    bool
	ShowRowStripes    *bool
}

type AutoFilterOptions struct {
	Column     string
	Expression string
}
