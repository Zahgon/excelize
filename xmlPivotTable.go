package excelize

import "encoding/xml"

type xlsxPivotTableDefinition struct {
	XMLName                 xml.Name                 `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main pivotTableDefinition"`
	Name                    string                   `xml:"name,attr"`
	CacheID                 int                      `xml:"cacheId,attr"`
	ApplyNumberFormats      bool                     `xml:"applyNumberFormats,attr,omitempty"`
	ApplyBorderFormats      bool                     `xml:"applyBorderFormats,attr,omitempty"`
	ApplyFontFormats        bool                     `xml:"applyFontFormats,attr,omitempty"`
	ApplyPatternFormats     bool                     `xml:"applyPatternFormats,attr,omitempty"`
	ApplyAlignmentFormats   bool                     `xml:"applyAlignmentFormats,attr,omitempty"`
	ApplyWidthHeightFormats bool                     `xml:"applyWidthHeightFormats,attr,omitempty"`
	DataOnRows              bool                     `xml:"dataOnRows,attr,omitempty"`
	DataPosition            int                      `xml:"dataPosition,attr,omitempty"`
	DataCaption             string                   `xml:"dataCaption,attr"`
	GrandTotalCaption       string                   `xml:"grandTotalCaption,attr,omitempty"`
	ErrorCaption            string                   `xml:"errorCaption,attr,omitempty"`
	ShowError               *bool                    `xml:"showError,attr"`
	MissingCaption          string                   `xml:"missingCaption,attr,omitempty"`
	ShowMissing             bool                     `xml:"showMissing,attr,omitempty"`
	PageStyle               string                   `xml:"pageStyle,attr,omitempty"`
	PivotTableStyle         string                   `xml:"pivotTableStyle,attr,omitempty"`
	VacatedStyle            string                   `xml:"vacatedStyle,attr,omitempty"`
	Tag                     string                   `xml:"tag,attr,omitempty"`
	UpdatedVersion          int                      `xml:"updatedVersion,attr,omitempty"`
	MinRefreshableVersion   int                      `xml:"minRefreshableVersion,attr,omitempty"`
	AsteriskTotals          bool                     `xml:"asteriskTotals,attr,omitempty"`
	ShowItems               bool                     `xml:"showItems,attr,omitempty"`
	EditData                bool                     `xml:"editData,attr,omitempty"`
	DisableFieldList        bool                     `xml:"disableFieldList,attr,omitempty"`
	ShowCalcMbrs            bool                     `xml:"showCalcMbrs,attr,omitempty"`
	VisualTotals            bool                     `xml:"visualTotals,attr,omitempty"`
	ShowMultipleLabel       bool                     `xml:"showMultipleLabel,attr,omitempty"`
	ShowDataDropDown        bool                     `xml:"showDataDropDown,attr,omitempty"`
	ShowDrill               *bool                    `xml:"showDrill,attr"`
	PrintDrill              bool                     `xml:"printDrill,attr,omitempty"`
	ShowMemberPropertyTips  bool                     `xml:"showMemberPropertyTips,attr,omitempty"`
	ShowDataTips            bool                     `xml:"showDataTips,attr,omitempty"`
	EnableWizard            bool                     `xml:"enableWizard,attr,omitempty"`
	EnableDrill             bool                     `xml:"enableDrill,attr,omitempty"`
	EnableFieldProperties   bool                     `xml:"enableFieldProperties,attr,omitempty"`
	PreserveFormatting      bool                     `xml:"preserveFormatting,attr,omitempty"`
	UseAutoFormatting       *bool                    `xml:"useAutoFormatting,attr"`
	PageWrap                int                      `xml:"pageWrap,attr,omitempty"`
	PageOverThenDown        *bool                    `xml:"pageOverThenDown,attr"`
	SubtotalHiddenItems     bool                     `xml:"subtotalHiddenItems,attr,omitempty"`
	RowGrandTotals          *bool                    `xml:"rowGrandTotals,attr"`
	ColGrandTotals          *bool                    `xml:"colGrandTotals,attr"`
	FieldPrintTitles        bool                     `xml:"fieldPrintTitles,attr,omitempty"`
	ItemPrintTitles         bool                     `xml:"itemPrintTitles,attr,omitempty"`
	MergeItem               *bool                    `xml:"mergeItem,attr"`
	ShowDropZones           bool                     `xml:"showDropZones,attr,omitempty"`
	CreatedVersion          int                      `xml:"createdVersion,attr,omitempty"`
	Indent                  int                      `xml:"indent,attr,omitempty"`
	ShowEmptyRow            bool                     `xml:"showEmptyRow,attr,omitempty"`
	ShowEmptyCol            bool                     `xml:"showEmptyCol,attr,omitempty"`
	ShowHeaders             bool                     `xml:"showHeaders,attr,omitempty"`
	Compact                 *bool                    `xml:"compact,attr"`
	Outline                 *bool                    `xml:"outline,attr"`
	OutlineData             bool                     `xml:"outlineData,attr,omitempty"`
	CompactData             *bool                    `xml:"compactData,attr"`
	Published               bool                     `xml:"published,attr,omitempty"`
	GridDropZones           bool                     `xml:"gridDropZones,attr,omitempty"`
	Immersive               bool                     `xml:"immersive,attr,omitempty"`
	MultipleFieldFilters    bool                     `xml:"multipleFieldFilters,attr,omitempty"`
	ChartFormat             int                      `xml:"chartFormat,attr,omitempty"`
	RowHeaderCaption        string                   `xml:"rowHeaderCaption,attr,omitempty"`
	ColHeaderCaption        string                   `xml:"colHeaderCaption,attr,omitempty"`
	FieldListSortAscending  bool                     `xml:"fieldListSortAscending,attr,omitempty"`
	MdxSubqueries           bool                     `xml:"mdxSubqueries,attr,omitempty"`
	CustomListSort          bool                     `xml:"customListSort,attr,omitempty"`
	Location                *xlsxLocation            `xml:"location"`
	PivotFields             *xlsxPivotFields         `xml:"pivotFields"`
	RowFields               *xlsxRowFields           `xml:"rowFields"`
	RowItems                *xlsxRowItems            `xml:"rowItems"`
	ColFields               *xlsxColFields           `xml:"colFields"`
	ColItems                *xlsxColItems            `xml:"colItems"`
	PageFields              *xlsxPageFields          `xml:"pageFields"`
	DataFields              *xlsxDataFields          `xml:"dataFields"`
	ConditionalFormats      *xlsxConditionalFormats  `xml:"conditionalFormats"`
	PivotTableStyleInfo     *xlsxPivotTableStyleInfo `xml:"pivotTableStyleInfo"`
}

type xlsxLocation struct {
	Ref            string `xml:"ref,attr"`
	FirstHeaderRow int    `xml:"firstHeaderRow,attr"`
	FirstDataRow   int    `xml:"firstDataRow,attr"`
	FirstDataCol   int    `xml:"firstDataCol,attr"`
	RowPageCount   int    `xml:"rowPageCount,attr,omitempty"`
	ColPageCount   int    `xml:"colPageCount,attr,omitempty"`
}

type xlsxPivotFields struct {
	Count      int               `xml:"count,attr"`
	PivotField []*xlsxPivotField `xml:"pivotField"`
}

type xlsxPivotField struct {
	Name                         string             `xml:"name,attr,omitempty"`
	Axis                         string             `xml:"axis,attr,omitempty"`
	DataField                    bool               `xml:"dataField,attr,omitempty"`
	SubtotalCaption              string             `xml:"subtotalCaption,attr,omitempty"`
	ShowDropDowns                bool               `xml:"showDropDowns,attr,omitempty"`
	HiddenLevel                  bool               `xml:"hiddenLevel,attr,omitempty"`
	UniqueMemberProperty         string             `xml:"uniqueMemberProperty,attr,omitempty"`
	Compact                      *bool              `xml:"compact,attr"`
	AllDrilled                   bool               `xml:"allDrilled,attr,omitempty"`
	NumFmtID                     string             `xml:"numFmtId,attr,omitempty"`
	Outline                      *bool              `xml:"outline,attr"`
	SubtotalTop                  bool               `xml:"subtotalTop,attr,omitempty"`
	DragToRow                    bool               `xml:"dragToRow,attr,omitempty"`
	DragToCol                    bool               `xml:"dragToCol,attr,omitempty"`
	MultipleItemSelectionAllowed bool               `xml:"multipleItemSelectionAllowed,attr,omitempty"`
	DragToPage                   bool               `xml:"dragToPage,attr,omitempty"`
	DragToData                   bool               `xml:"dragToData,attr,omitempty"`
	DragOff                      bool               `xml:"dragOff,attr,omitempty"`
	ShowAll                      bool               `xml:"showAll,attr"`
	InsertBlankRow               bool               `xml:"insertBlankRow,attr,omitempty"`
	ServerField                  bool               `xml:"serverField,attr,omitempty"`
	InsertPageBreak              bool               `xml:"insertPageBreak,attr,omitempty"`
	AutoShow                     bool               `xml:"autoShow,attr,omitempty"`
	TopAutoShow                  bool               `xml:"topAutoShow,attr,omitempty"`
	HideNewItems                 bool               `xml:"hideNewItems,attr,omitempty"`
	MeasureFilter                bool               `xml:"measureFilter,attr,omitempty"`
	IncludeNewItemsInFilter      bool               `xml:"includeNewItemsInFilter,attr,omitempty"`
	ItemPageCount                int                `xml:"itemPageCount,attr,omitempty"`
	SortType                     string             `xml:"sortType,attr,omitempty"`
	DataSourceSort               bool               `xml:"dataSourceSort,attr,omitempty"`
	NonAutoSortDefault           bool               `xml:"nonAutoSortDefault,attr,omitempty"`
	RankBy                       int                `xml:"rankBy,attr,omitempty"`
	DefaultSubtotal              *bool              `xml:"defaultSubtotal,attr"`
	SumSubtotal                  bool               `xml:"sumSubtotal,attr,omitempty"`
	CountASubtotal               bool               `xml:"countASubtotal,attr,omitempty"`
	AvgSubtotal                  bool               `xml:"avgSubtotal,attr,omitempty"`
	MaxSubtotal                  bool               `xml:"maxSubtotal,attr,omitempty"`
	MinSubtotal                  bool               `xml:"minSubtotal,attr,omitempty"`
	ProductSubtotal              bool               `xml:"productSubtotal,attr,omitempty"`
	CountSubtotal                bool               `xml:"countSubtotal,attr,omitempty"`
	StdDevSubtotal               bool               `xml:"stdDevSubtotal,attr,omitempty"`
	StdDevPSubtotal              bool               `xml:"stdDevPSubtotal,attr,omitempty"`
	VarSubtotal                  bool               `xml:"varSubtotal,attr,omitempty"`
	VarPSubtotal                 bool               `xml:"varPSubtotal,attr,omitempty"`
	ShowPropCell                 bool               `xml:"showPropCell,attr,omitempty"`
	ShowPropTip                  bool               `xml:"showPropTip,attr,omitempty"`
	ShowPropAsCaption            bool               `xml:"showPropAsCaption,attr,omitempty"`
	DefaultAttributeDrillState   bool               `xml:"defaultAttributeDrillState,attr,omitempty"`
	Items                        *xlsxItems         `xml:"items"`
	AutoSortScope                *xlsxAutoSortScope `xml:"autoSortScope"`
	ExtLst                       *xlsxExtLst        `xml:"extLst"`
}

type xlsxItems struct {
	Count int         `xml:"count,attr"`
	Item  []*xlsxItem `xml:"item"`
}

type xlsxItem struct {
	N  string `xml:"n,attr,omitempty"`
	T  string `xml:"t,attr,omitempty"`
	H  bool   `xml:"h,attr,omitempty"`
	S  bool   `xml:"s,attr,omitempty"`
	SD bool   `xml:"sd,attr,omitempty"`
	F  bool   `xml:"f,attr,omitempty"`
	M  bool   `xml:"m,attr,omitempty"`
	C  bool   `xml:"c,attr,omitempty"`
	X  *int   `xml:"x,attr,omitempty"`
	D  bool   `xml:"d,attr,omitempty"`
	E  bool   `xml:"e,attr,omitempty"`
}

type xlsxAutoSortScope struct{}

type xlsxRowFields struct {
	Count int          `xml:"count,attr"`
	Field []*xlsxField `xml:"field"`
}

type xlsxField struct {
	X int `xml:"x,attr"`
}

type xlsxRowItems struct {
	Count int      `xml:"count,attr"`
	I     []*xlsxI `xml:"i"`
}

type xlsxI struct {
	X []*xlsxX `xml:"x"`
}

type xlsxX struct{}

type xlsxColFields struct {
	Count int          `xml:"count,attr"`
	Field []*xlsxField `xml:"field"`
}

type xlsxColItems struct {
	Count int      `xml:"count,attr"`
	I     []*xlsxI `xml:"i"`
}

type xlsxPageFields struct {
	Count     int              `xml:"count,attr"`
	PageField []*xlsxPageField `xml:"pageField"`
}

type xlsxPageField struct {
	Fld    int         `xml:"fld,attr"`
	Item   int         `xml:"item,attr,omitempty"`
	Hier   int         `xml:"hier,attr,omitempty"`
	Name   string      `xml:"name,attr,omitempty"`
	Cap    string      `xml:"cap,attr,omitempty"`
	ExtLst *xlsxExtLst `xml:"extLst"`
}

type xlsxDataFields struct {
	Count     int              `xml:"count,attr"`
	DataField []*xlsxDataField `xml:"dataField"`
}

type xlsxDataField struct {
	Name       string      `xml:"name,attr,omitempty"`
	Fld        int         `xml:"fld,attr"`
	Subtotal   string      `xml:"subtotal,attr,omitempty"`
	ShowDataAs string      `xml:"showDataAs,attr,omitempty"`
	BaseField  *int        `xml:"baseField,attr"`
	BaseItem   *int        `xml:"baseItem,attr"`
	NumFmtID   int         `xml:"numFmtId,attr,omitempty"`
	ExtLst     *xlsxExtLst `xml:"extLst"`
}

type xlsxX14DataField struct {
	XMLName     xml.Name `xml:"x14:dataField"`
	PivotShowAs string   `xml:"pivotShowAs,attr,omitempty"`
	SourceField int      `xml:"sourceField,attr,omitempty"`
	UniqueName  string   `xml:"uniqueName,attr,omitempty"`
}

type xlsxConditionalFormats struct{}

type xlsxPivotTableStyleInfo struct {
	Name           string `xml:"name,attr"`
	ShowRowHeaders bool   `xml:"showRowHeaders,attr"`
	ShowColHeaders bool   `xml:"showColHeaders,attr"`
	ShowRowStripes bool   `xml:"showRowStripes,attr,omitempty"`
	ShowColStripes bool   `xml:"showColStripes,attr,omitempty"`
	ShowLastColumn bool   `xml:"showLastColumn,attr,omitempty"`
}

type decodeX14DataField struct {
	XMLName     xml.Name `xml:"dataField"`
	PivotShowAs string   `xml:"pivotShowAs,attr,omitempty"`
	SourceField int      `xml:"sourceField,attr,omitempty"`
	UniqueName  string   `xml:"uniqueName,attr,omitempty"`
}
