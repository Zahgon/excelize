package excelize

import "encoding/xml"

type xlsxSlicers struct {
	XMLName   xml.Name     `xml:"http://schemas.microsoft.com/office/spreadsheetml/2009/9/main slicers"`
	XMLNSXMC  string       `xml:"xmlns:mc,attr"`
	XMLNSX    string       `xml:"xmlns:x,attr"`
	XMLNSXR10 string       `xml:"xmlns:xr10,attr"`
	Slicer    []xlsxSlicer `xml:"slicer"`
}

type xlsxSlicer struct {
	Name           string `xml:"name,attr"`
	XR10UID        string `xml:"xr10:uid,attr,omitempty"`
	Cache          string `xml:"cache,attr"`
	Caption        string `xml:"caption,attr,omitempty"`
	StartItem      *int   `xml:"startItem,attr"`
	ColumnCount    *int   `xml:"columnCount,attr"`
	ShowCaption    *bool  `xml:"showCaption,attr"`
	Level          int    `xml:"level,attr,omitempty"`
	Style          string `xml:"style,attr,omitempty"`
	LockedPosition bool   `xml:"lockedPosition,attr,omitempty"`
	RowHeight      int    `xml:"rowHeight,attr"`
}

type xlsxSlicerCacheDefinition struct {
	XMLName     xml.Name                    `xml:"http://schemas.microsoft.com/office/spreadsheetml/2009/9/main slicerCacheDefinition"`
	XMLNSXMC    string                      `xml:"xmlns:mc,attr"`
	XMLNSX      string                      `xml:"xmlns:x,attr"`
	XMLNSX15    string                      `xml:"xmlns:x15,attr,omitempty"`
	XMLNSXR10   string                      `xml:"xmlns:xr10,attr"`
	Name        string                      `xml:"name,attr"`
	XR10UID     string                      `xml:"xr10:uid,attr,omitempty"`
	SourceName  string                      `xml:"sourceName,attr"`
	PivotTables *xlsxSlicerCachePivotTables `xml:"pivotTables"`
	Data        *xlsxSlicerCacheData        `xml:"data"`
	ExtLst      *xlsxExtLst                 `xml:"extLst"`
}

type xlsxSlicerCachePivotTables struct {
	PivotTable []xlsxSlicerCachePivotTable `xml:"pivotTable"`
}

type xlsxSlicerCachePivotTable struct {
	TabID int    `xml:"tabId,attr"`
	Name  string `xml:"name,attr"`
}

type xlsxSlicerCacheData struct {
	OLAP    *xlsxInnerXML           `xml:"olap"`
	Tabular *xlsxTabularSlicerCache `xml:"tabular"`
}

type xlsxTabularSlicerCache struct {
	PivotCacheID   int                          `xml:"pivotCacheId,attr"`
	SortOrder      string                       `xml:"sortOrder,attr,omitempty"`
	CustomListSort *bool                        `xml:"customListSort,attr"`
	ShowMissing    *bool                        `xml:"showMissing,attr"`
	CrossFilter    string                       `xml:"crossFilter,attr,omitempty"`
	Items          *xlsxTabularSlicerCacheItems `xml:"items"`
	ExtLst         *xlsxExtLst                  `xml:"extLst"`
}

type xlsxTabularSlicerCacheItems struct {
	Count int                          `xml:"count,attr,omitempty"`
	I     []xlsxTabularSlicerCacheItem `xml:"i"`
}

type xlsxTabularSlicerCacheItem struct {
	X  int  `xml:"x,attr"`
	S  bool `xml:"s,attr,omitempty"`
	ND bool `xml:"nd,attr,omitempty"`
}

type xlsxTableSlicerCache struct {
	XMLName        xml.Name    `xml:"x15:tableSlicerCache"`
	TableID        int         `xml:"tableId,attr"`
	Column         int         `xml:"column,attr"`
	SortOrder      string      `xml:"sortOrder,attr,omitempty"`
	CustomListSort *bool       `xml:"customListSort,attr"`
	CrossFilter    string      `xml:"crossFilter,attr,omitempty"`
	ExtLst         *xlsxExtLst `xml:"extLst"`
}

type xlsxX14SlicerList struct {
	XMLName xml.Name         `xml:"x14:slicerList"`
	Slicer  []*xlsxX14Slicer `xml:"x14:slicer"`
}

type xlsxX14Slicer struct {
	XMLName xml.Name `xml:"x14:slicer"`
	RID     string   `xml:"r:id,attr"`
}

type xlsxX14SlicerCaches struct {
	XMLName xml.Name `xml:"x14:slicerCaches"`
	XMLNS   string   `xml:"xmlns:x14,attr"`
	Content string   `xml:",innerxml"`
}

type xlsxX14SlicerCache struct {
	XMLName xml.Name `xml:"x14:slicerCache"`
	RID     string   `xml:"r:id,attr"`
}

type xlsxX15SlicerCaches struct {
	XMLName xml.Name `xml:"x15:slicerCaches"`
	XMLNS   string   `xml:"xmlns:x14,attr"`
	Content string   `xml:",innerxml"`
}

type decodeTableSlicerCache struct {
	XMLName   xml.Name `xml:"tableSlicerCache"`
	TableID   int      `xml:"tableId,attr"`
	Column    int      `xml:"column,attr"`
	SortOrder string   `xml:"sortOrder,attr"`
}

type decodeSlicerList struct {
	XMLName xml.Name        `xml:"slicerList"`
	Slicer  []*decodeSlicer `xml:"slicer"`
}

type decodeSlicer struct {
	RID string `xml:"id,attr"`
}

type decodeSlicerCaches struct {
	XMLName xml.Name `xml:"slicerCaches"`
	Content string   `xml:",innerxml"`
}

type xlsxTimelines struct {
	XMLName   xml.Name       `xml:"http://schemas.microsoft.com/office/spreadsheetml/2010/11/main timelines"`
	XMLNSXMC  string         `xml:"xmlns:mc,attr"`
	XMLNSX    string         `xml:"xmlns:x,attr"`
	XMLNSXR10 string         `xml:"xmlns:xr10,attr"`
	Timeline  []xlsxTimeline `xml:"timeline"`
}

type xlsxTimeline struct {
	Name                    string `xml:"name,attr"`
	XR10UID                 string `xml:"xr10:uid,attr,omitempty"`
	Cache                   string `xml:"cache,attr"`
	Caption                 string `xml:"caption,attr,omitempty"`
	ShowHeader              *bool  `xml:"showHeader,attr"`
	ShowSelectionLabel      *bool  `xml:"showSelectionLabel,attr"`
	ShowTimeLevel           *bool  `xml:"showTimeLevel,attr"`
	ShowHorizontalScrollbar *bool  `xml:"showHorizontalScrollbar,attr"`
	Level                   int    `xml:"level,attr"`
	SelectionLevel          int    `xml:"selectionLevel,attr"`
	ScrollPosition          string `xml:"scrollPosition,attr,omitempty"`
	Style                   string `xml:"style,attr,omitempty"`
}
