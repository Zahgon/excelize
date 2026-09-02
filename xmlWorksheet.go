package excelize

import (
	"encoding/xml"
	"sync"
)

type xlsxWorksheet struct {
	mu                     sync.Mutex
	formulaSI              sync.Map
	XMLName                xml.Name                     `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main worksheet"`
	SheetPr                *xlsxSheetPr                 `xml:"sheetPr"`
	Dimension              *xlsxDimension               `xml:"dimension"`
	SheetViews             *xlsxSheetViews              `xml:"sheetViews"`
	SheetFormatPr          *xlsxSheetFormatPr           `xml:"sheetFormatPr"`
	Cols                   *xlsxCols                    `xml:"cols"`
	SheetData              xlsxSheetData                `xml:"sheetData"`
	SheetCalcPr            *xlsxInnerXML                `xml:"sheetCalcPr"`
	SheetProtection        *xlsxSheetProtection         `xml:"sheetProtection"`
	ProtectedRanges        *xlsxInnerXML                `xml:"protectedRanges"`
	Scenarios              *xlsxInnerXML                `xml:"scenarios"`
	AutoFilter             *xlsxAutoFilter              `xml:"autoFilter"`
	SortState              *xlsxSortState               `xml:"sortState"`
	DataConsolidate        *xlsxInnerXML                `xml:"dataConsolidate"`
	CustomSheetViews       *xlsxCustomSheetViews        `xml:"customSheetViews"`
	MergeCells             *xlsxMergeCells              `xml:"mergeCells"`
	PhoneticPr             *xlsxPhoneticPr              `xml:"phoneticPr"`
	ConditionalFormatting  []*xlsxConditionalFormatting `xml:"conditionalFormatting"`
	DataValidations        *xlsxDataValidations         `xml:"dataValidations"`
	Hyperlinks             *xlsxHyperlinks              `xml:"hyperlinks"`
	PrintOptions           *xlsxPrintOptions            `xml:"printOptions"`
	PageMargins            *xlsxPageMargins             `xml:"pageMargins"`
	PageSetUp              *xlsxPageSetUp               `xml:"pageSetup"`
	HeaderFooter           *xlsxHeaderFooter            `xml:"headerFooter"`
	RowBreaks              *xlsxRowBreaks               `xml:"rowBreaks"`
	ColBreaks              *xlsxColBreaks               `xml:"colBreaks"`
	CustomProperties       *xlsxInnerXML                `xml:"customProperties"`
	CellWatches            *xlsxInnerXML                `xml:"cellWatches"`
	IgnoredErrors          *xlsxIgnoredErrors           `xml:"ignoredErrors"`
	SmartTags              *xlsxInnerXML                `xml:"smartTags"`
	Drawing                *xlsxDrawing                 `xml:"drawing"`
	LegacyDrawing          *xlsxLegacyDrawing           `xml:"legacyDrawing"`
	LegacyDrawingHF        *xlsxLegacyDrawingHF         `xml:"legacyDrawingHF"`
	DrawingHF              *xlsxDrawingHF               `xml:"drawingHF"`
	Picture                *xlsxPicture                 `xml:"picture"`
	OleObjects             *xlsxInnerXML                `xml:"oleObjects"`
	Controls               *xlsxInnerXML                `xml:"controls"`
	WebPublishItems        *xlsxInnerXML                `xml:"webPublishItems"`
	AlternateContent       *xlsxAlternateContent        `xml:"mc:AlternateContent"`
	TableParts             *xlsxTableParts              `xml:"tableParts"`
	ExtLst                 *xlsxExtLst                  `xml:"extLst"`
	DecodeAlternateContent *xlsxInnerXML                `xml:"http://schemas.openxmlformats.org/markup-compatibility/2006 AlternateContent"`
}

type xlsxDrawing struct {
	XMLName xml.Name `xml:"drawing"`
	RID     string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type xlsxHeaderFooter struct {
	XMLName          xml.Name `xml:"headerFooter"`
	DifferentOddEven bool     `xml:"differentOddEven,attr,omitempty"`
	DifferentFirst   bool     `xml:"differentFirst,attr,omitempty"`
	ScaleWithDoc     *bool    `xml:"scaleWithDoc,attr"`
	AlignWithMargins *bool    `xml:"alignWithMargins,attr"`
	OddHeader        string   `xml:"oddHeader,omitempty"`
	OddFooter        string   `xml:"oddFooter,omitempty"`
	EvenHeader       string   `xml:"evenHeader,omitempty"`
	EvenFooter       string   `xml:"evenFooter,omitempty"`
	FirstHeader      string   `xml:"firstHeader,omitempty"`
	FirstFooter      string   `xml:"firstFooter,omitempty"`
}

type xlsxDrawingHF struct {
	Content string `xml:",innerxml"`
}

type xlsxPageSetUp struct {
	XMLName            xml.Name `xml:"pageSetup"`
	BlackAndWhite      bool     `xml:"blackAndWhite,attr,omitempty"`
	CellComments       string   `xml:"cellComments,attr,omitempty"`
	Copies             int      `xml:"copies,attr,omitempty"`
	Draft              bool     `xml:"draft,attr,omitempty"`
	Errors             string   `xml:"errors,attr,omitempty"`
	FirstPageNumber    string   `xml:"firstPageNumber,attr,omitempty"`
	FitToHeight        *int     `xml:"fitToHeight,attr"`
	FitToWidth         *int     `xml:"fitToWidth,attr"`
	HorizontalDPI      string   `xml:"horizontalDpi,attr,omitempty"`
	RID                string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
	Orientation        string   `xml:"orientation,attr,omitempty"`
	PageOrder          string   `xml:"pageOrder,attr,omitempty"`
	PaperHeight        string   `xml:"paperHeight,attr,omitempty"`
	PaperSize          *int     `xml:"paperSize,attr"`
	PaperWidth         string   `xml:"paperWidth,attr,omitempty"`
	Scale              int      `xml:"scale,attr,omitempty"`
	UseFirstPageNumber bool     `xml:"useFirstPageNumber,attr,omitempty"`
	UsePrinterDefaults bool     `xml:"usePrinterDefaults,attr,omitempty"`
	VerticalDPI        string   `xml:"verticalDpi,attr,omitempty"`
}

type xlsxPrintOptions struct {
	XMLName            xml.Name `xml:"printOptions"`
	GridLines          bool     `xml:"gridLines,attr,omitempty"`
	GridLinesSet       bool     `xml:"gridLinesSet,attr,omitempty"`
	Headings           bool     `xml:"headings,attr,omitempty"`
	HorizontalCentered bool     `xml:"horizontalCentered,attr,omitempty"`
	VerticalCentered   bool     `xml:"verticalCentered,attr,omitempty"`
}

type xlsxPageMargins struct {
	XMLName xml.Name `xml:"pageMargins"`
	Left    float64  `xml:"left,attr"`
	Right   float64  `xml:"right,attr"`
	Top     float64  `xml:"top,attr"`
	Bottom  float64  `xml:"bottom,attr"`
	Header  float64  `xml:"header,attr"`
	Footer  float64  `xml:"footer,attr"`
}

type xlsxSheetFormatPr struct {
	XMLName          xml.Name `xml:"sheetFormatPr"`
	BaseColWidth     uint8    `xml:"baseColWidth,attr,omitempty"`
	DefaultColWidth  float64  `xml:"defaultColWidth,attr,omitempty"`
	DefaultRowHeight float64  `xml:"defaultRowHeight,attr"`
	CustomHeight     bool     `xml:"customHeight,attr,omitempty"`
	ZeroHeight       bool     `xml:"zeroHeight,attr,omitempty"`
	ThickTop         bool     `xml:"thickTop,attr,omitempty"`
	ThickBottom      bool     `xml:"thickBottom,attr,omitempty"`
	OutlineLevelRow  uint8    `xml:"outlineLevelRow,attr,omitempty"`
	OutlineLevelCol  uint8    `xml:"outlineLevelCol,attr,omitempty"`
}

type xlsxSheetViews struct {
	XMLName   xml.Name        `xml:"sheetViews"`
	SheetView []xlsxSheetView `xml:"sheetView"`
}

type xlsxSheetView struct {
	WindowProtection         bool             `xml:"windowProtection,attr,omitempty"`
	ShowFormulas             bool             `xml:"showFormulas,attr,omitempty"`
	ShowGridLines            *bool            `xml:"showGridLines,attr"`
	ShowRowColHeaders        *bool            `xml:"showRowColHeaders,attr"`
	ShowZeros                *bool            `xml:"showZeros,attr,omitempty"`
	RightToLeft              bool             `xml:"rightToLeft,attr,omitempty"`
	TabSelected              bool             `xml:"tabSelected,attr,omitempty"`
	ShowRuler                *bool            `xml:"showRuler,attr,omitempty"`
	ShowWhiteSpace           *bool            `xml:"showWhiteSpace,attr"`
	ShowOutlineSymbols       bool             `xml:"showOutlineSymbols,attr,omitempty"`
	DefaultGridColor         *bool            `xml:"defaultGridColor,attr"`
	View                     string           `xml:"view,attr,omitempty"`
	TopLeftCell              string           `xml:"topLeftCell,attr,omitempty"`
	ColorID                  int              `xml:"colorId,attr,omitempty"`
	ZoomScale                float64          `xml:"zoomScale,attr,omitempty"`
	ZoomScaleNormal          float64          `xml:"zoomScaleNormal,attr,omitempty"`
	ZoomScalePageLayoutView  float64          `xml:"zoomScalePageLayoutView,attr,omitempty"`
	ZoomScaleSheetLayoutView float64          `xml:"zoomScaleSheetLayoutView,attr,omitempty"`
	WorkbookViewID           int              `xml:"workbookViewId,attr"`
	Pane                     *xlsxPane        `xml:"pane,omitempty"`
	Selection                []*xlsxSelection `xml:"selection"`
}

type xlsxSelection struct {
	ActiveCell   string `xml:"activeCell,attr,omitempty"`
	ActiveCellID *int   `xml:"activeCellId,attr"`
	Pane         string `xml:"pane,attr,omitempty"`
	SQRef        string `xml:"sqref,attr,omitempty"`
}

type xlsxPane struct {
	ActivePane  string  `xml:"activePane,attr,omitempty"`
	State       string  `xml:"state,attr,omitempty"`
	TopLeftCell string  `xml:"topLeftCell,attr,omitempty"`
	XSplit      float64 `xml:"xSplit,attr,omitempty"`
	YSplit      float64 `xml:"ySplit,attr,omitempty"`
}

type xlsxSheetPr struct {
	XMLName                           xml.Name         `xml:"sheetPr"`
	SyncHorizontal                    bool             `xml:"syncHorizontal,attr,omitempty"`
	SyncVertical                      bool             `xml:"syncVertical,attr,omitempty"`
	SyncRef                           string           `xml:"syncRef,attr,omitempty"`
	TransitionEvaluation              bool             `xml:"transitionEvaluation,attr,omitempty"`
	TransitionEntry                   bool             `xml:"transitionEntry,attr,omitempty"`
	Published                         *bool            `xml:"published,attr"`
	CodeName                          string           `xml:"codeName,attr,omitempty"`
	FilterMode                        bool             `xml:"filterMode,attr,omitempty"`
	EnableFormatConditionsCalculation *bool            `xml:"enableFormatConditionsCalculation,attr"`
	TabColor                          *xlsxColor       `xml:"tabColor"`
	OutlinePr                         *xlsxOutlinePr   `xml:"outlinePr"`
	PageSetUpPr                       *xlsxPageSetUpPr `xml:"pageSetUpPr"`
}

type xlsxOutlinePr struct {
	ApplyStyles        *bool `xml:"applyStyles,attr"`
	SummaryBelow       *bool `xml:"summaryBelow,attr"`
	SummaryRight       *bool `xml:"summaryRight,attr"`
	ShowOutlineSymbols *bool `xml:"showOutlineSymbols,attr"`
}

type xlsxPageSetUpPr struct {
	AutoPageBreaks bool `xml:"autoPageBreaks,attr,omitempty"`
	FitToPage      bool `xml:"fitToPage,attr,omitempty"`
}

type xlsxCols struct {
	XMLName xml.Name  `xml:"cols"`
	Col     []xlsxCol `xml:"col"`
}

type xlsxCol struct {
	BestFit      bool     `xml:"bestFit,attr,omitempty"`
	Collapsed    bool     `xml:"collapsed,attr,omitempty"`
	CustomWidth  bool     `xml:"customWidth,attr,omitempty"`
	Hidden       bool     `xml:"hidden,attr,omitempty"`
	Max          int      `xml:"max,attr"`
	Min          int      `xml:"min,attr"`
	OutlineLevel uint8    `xml:"outlineLevel,attr,omitempty"`
	Phonetic     bool     `xml:"phonetic,attr,omitempty"`
	Style        int      `xml:"style,attr,omitempty"`
	Width        *float64 `xml:"width,attr"`
}

type xlsxDimension struct {
	XMLName xml.Name `xml:"dimension"`
	Ref     string   `xml:"ref,attr"`
}

type xlsxSheetData struct {
	XMLName xml.Name  `xml:"sheetData"`
	Row     []xlsxRow `xml:"row"`
}

type xlsxRow struct {
	C            []xlsxC  `xml:"c"`
	R            int      `xml:"r,attr,omitempty"`
	Spans        string   `xml:"spans,attr,omitempty"`
	S            int      `xml:"s,attr,omitempty"`
	CustomFormat bool     `xml:"customFormat,attr,omitempty"`
	Ht           *float64 `xml:"ht,attr"`
	Hidden       bool     `xml:"hidden,attr,omitempty"`
	CustomHeight bool     `xml:"customHeight,attr,omitempty"`
	OutlineLevel uint8    `xml:"outlineLevel,attr,omitempty"`
	Collapsed    bool     `xml:"collapsed,attr,omitempty"`
	ThickTop     bool     `xml:"thickTop,attr,omitempty"`
	ThickBot     bool     `xml:"thickBot,attr,omitempty"`
	Ph           bool     `xml:"ph,attr,omitempty"`
}

type xlsxSortState struct {
	ColumnSort    bool   `xml:"columnSort,attr,omitempty"`
	CaseSensitive bool   `xml:"caseSensitive,attr,omitempty"`
	SortMethod    string `xml:"sortMethod,attr,omitempty"`
	Ref           string `xml:"ref,attr"`
	Content       string `xml:",innerxml"`
}

type xlsxCustomSheetViews struct {
	XMLName         xml.Name               `xml:"customSheetViews"`
	CustomSheetView []*xlsxCustomSheetView `xml:"customSheetView"`
}

type xlsxBrk struct {
	ID  int  `xml:"id,attr,omitempty"`
	Min int  `xml:"min,attr,omitempty"`
	Max int  `xml:"max,attr,omitempty"`
	Man bool `xml:"man,attr,omitempty"`
	Pt  bool `xml:"pt,attr,omitempty"`
}

type xlsxRowBreaks struct {
	XMLName xml.Name `xml:"rowBreaks"`
	xlsxBreaks
}

type xlsxColBreaks struct {
	XMLName xml.Name `xml:"colBreaks"`
	xlsxBreaks
}

type xlsxBreaks struct {
	Brk              []*xlsxBrk `xml:"brk"`
	Count            int        `xml:"count,attr,omitempty"`
	ManualBreakCount int        `xml:"manualBreakCount,attr,omitempty"`
}

type xlsxCustomSheetView struct {
	Pane           *xlsxPane         `xml:"pane"`
	Selection      *xlsxSelection    `xml:"selection"`
	RowBreaks      *xlsxBreaks       `xml:"rowBreaks"`
	ColBreaks      *xlsxBreaks       `xml:"colBreaks"`
	PageMargins    *xlsxPageMargins  `xml:"pageMargins"`
	PrintOptions   *xlsxPrintOptions `xml:"printOptions"`
	PageSetup      *xlsxPageSetUp    `xml:"pageSetup"`
	HeaderFooter   *xlsxHeaderFooter `xml:"headerFooter"`
	AutoFilter     *xlsxAutoFilter   `xml:"autoFilter"`
	ExtLst         *xlsxExtLst       `xml:"extLst"`
	GUID           string            `xml:"guid,attr"`
	Scale          int               `xml:"scale,attr,omitempty"`
	ColorID        int               `xml:"colorId,attr,omitempty"`
	ShowPageBreaks bool              `xml:"showPageBreaks,attr,omitempty"`
	ShowFormulas   bool              `xml:"showFormulas,attr,omitempty"`
	ShowGridLines  bool              `xml:"showGridLines,attr,omitempty"`
	ShowRowCol     bool              `xml:"showRowCol,attr,omitempty"`
	OutlineSymbols bool              `xml:"outlineSymbols,attr,omitempty"`
	ZeroValues     bool              `xml:"zeroValues,attr,omitempty"`
	FitToPage      bool              `xml:"fitToPage,attr,omitempty"`
	PrintArea      bool              `xml:"printArea,attr,omitempty"`
	Filter         bool              `xml:"filter,attr,omitempty"`
	ShowAutoFilter bool              `xml:"showAutoFilter,attr,omitempty"`
	HiddenRows     bool              `xml:"hiddenRows,attr,omitempty"`
	HiddenColumns  bool              `xml:"hiddenColumns,attr,omitempty"`
	State          string            `xml:"state,attr,omitempty"`
	FilterUnique   bool              `xml:"filterUnique,attr,omitempty"`
	View           string            `xml:"view,attr,omitempty"`
	ShowRuler      bool              `xml:"showRuler,attr,omitempty"`
	TopLeftCell    string            `xml:"topLeftCell,attr,omitempty"`
}

type xlsxMergeCell struct {
	Ref  string `xml:"ref,attr,omitempty"`
	rect []int
}

type xlsxMergeCells struct {
	XMLName xml.Name         `xml:"mergeCells"`
	Count   int              `xml:"count,attr,omitempty"`
	Cells   []*xlsxMergeCell `xml:"mergeCell,omitempty"`
}

type xlsxDataValidations struct {
	XMLName        xml.Name              `xml:"dataValidations"`
	Count          int                   `xml:"count,attr,omitempty"`
	DisablePrompts bool                  `xml:"disablePrompts,attr,omitempty"`
	XWindow        int                   `xml:"xWindow,attr,omitempty"`
	YWindow        int                   `xml:"yWindow,attr,omitempty"`
	DataValidation []*xlsxDataValidation `xml:"dataValidation"`
}

type xlsxDataValidation struct {
	AllowBlank       bool          `xml:"allowBlank,attr"`
	Error            *string       `xml:"error,attr"`
	ErrorStyle       *string       `xml:"errorStyle,attr"`
	ErrorTitle       *string       `xml:"errorTitle,attr"`
	Operator         string        `xml:"operator,attr,omitempty"`
	Prompt           *string       `xml:"prompt,attr"`
	PromptTitle      *string       `xml:"promptTitle,attr"`
	ShowDropDown     bool          `xml:"showDropDown,attr,omitempty"`
	ShowErrorMessage bool          `xml:"showErrorMessage,attr,omitempty"`
	ShowInputMessage bool          `xml:"showInputMessage,attr,omitempty"`
	Sqref            string        `xml:"sqref,attr"`
	XMSqref          string        `xml:"sqref,omitempty"`
	Type             string        `xml:"type,attr,omitempty"`
	Formula1         *xlsxInnerXML `xml:"formula1"`
	Formula2         *xlsxInnerXML `xml:"formula2"`
}

type xlsxX14DataValidation struct {
	XMLName          xml.Name      `xml:"x14:dataValidation"`
	AllowBlank       bool          `xml:"allowBlank,attr"`
	Error            *string       `xml:"error,attr"`
	ErrorStyle       *string       `xml:"errorStyle,attr"`
	ErrorTitle       *string       `xml:"errorTitle,attr"`
	Operator         string        `xml:"operator,attr,omitempty"`
	Prompt           *string       `xml:"prompt,attr"`
	PromptTitle      *string       `xml:"promptTitle,attr"`
	ShowDropDown     bool          `xml:"showDropDown,attr,omitempty"`
	ShowErrorMessage bool          `xml:"showErrorMessage,attr,omitempty"`
	ShowInputMessage bool          `xml:"showInputMessage,attr,omitempty"`
	Sqref            string        `xml:"sqref,attr"`
	Type             string        `xml:"type,attr,omitempty"`
	Formula1         *xlsxInnerXML `xml:"x14:formula1"`
	Formula2         *xlsxInnerXML `xml:"x14:formula2"`
	XMSqref          string        `xml:"xm:sqref,omitempty"`
}

type xlsxX14DataValidations struct {
	XMLName        xml.Name `xml:"x14:dataValidations"`
	XMLNSXM        string   `xml:"xmlns:xm,attr,omitempty"`
	Count          int      `xml:"count,attr,omitempty"`
	DisablePrompts bool     `xml:"disablePrompts,attr,omitempty"`
	XWindow        int      `xml:"xWindow,attr,omitempty"`
	YWindow        int      `xml:"yWindow,attr,omitempty"`
	DataValidation []*xlsxX14DataValidation
}

type xlsxC struct {
	XMLName  xml.Name `xml:"c"`
	XMLSpace xml.Attr `xml:"space,attr,omitempty"`
	R        string   `xml:"r,attr,omitempty"`
	S        int      `xml:"s,attr,omitempty"`
	T        string   `xml:"t,attr,omitempty"`
	Cm       *uint    `xml:"cm,attr"`
	Vm       *uint    `xml:"vm,attr"`
	Ph       *bool    `xml:"ph,attr"`
	F        *xlsxF   `xml:"f"`
	V        string   `xml:"v,omitempty"`
	IS       *xlsxSI  `xml:"is"`
	f        string
}

type xlsxF struct {
	Content string `xml:",chardata"`
	T       string `xml:"t,attr,omitempty"`
	Aca     bool   `xml:"aca,attr,omitempty"`
	Ref     string `xml:"ref,attr,omitempty"`
	Dt2D    bool   `xml:"dt2D,attr,omitempty"`
	Dtr     bool   `xml:"dtr,attr,omitempty"`
	Del1    bool   `xml:"del1,attr,omitempty"`
	Del2    bool   `xml:"del2,attr,omitempty"`
	R1      string `xml:"r1,attr,omitempty"`
	R2      string `xml:"r2,attr,omitempty"`
	Ca      bool   `xml:"ca,attr,omitempty"`
	Si      *int   `xml:"si,attr"`
	Bx      bool   `xml:"bx,attr,omitempty"`
}

type xlsxSheetProtection struct {
	XMLName             xml.Name `xml:"sheetProtection"`
	AlgorithmName       string   `xml:"algorithmName,attr,omitempty"`
	Password            string   `xml:"password,attr,omitempty"`
	HashValue           string   `xml:"hashValue,attr,omitempty"`
	SaltValue           string   `xml:"saltValue,attr,omitempty"`
	SpinCount           int      `xml:"spinCount,attr,omitempty"`
	Sheet               bool     `xml:"sheet,attr"`
	Objects             bool     `xml:"objects,attr"`
	Scenarios           bool     `xml:"scenarios,attr"`
	FormatCells         bool     `xml:"formatCells,attr"`
	FormatColumns       bool     `xml:"formatColumns,attr"`
	FormatRows          bool     `xml:"formatRows,attr"`
	InsertColumns       bool     `xml:"insertColumns,attr"`
	InsertRows          bool     `xml:"insertRows,attr"`
	InsertHyperlinks    bool     `xml:"insertHyperlinks,attr"`
	DeleteColumns       bool     `xml:"deleteColumns,attr"`
	DeleteRows          bool     `xml:"deleteRows,attr"`
	SelectLockedCells   bool     `xml:"selectLockedCells,attr"`
	Sort                bool     `xml:"sort,attr"`
	AutoFilter          bool     `xml:"autoFilter,attr"`
	PivotTables         bool     `xml:"pivotTables,attr"`
	SelectUnlockedCells bool     `xml:"selectUnlockedCells,attr"`
}

type xlsxPhoneticPr struct {
	XMLName   xml.Name `xml:"phoneticPr"`
	Alignment string   `xml:"alignment,attr,omitempty"`
	FontID    *int     `xml:"fontId,attr"`
	Type      string   `xml:"type,attr,omitempty"`
}

type xlsxConditionalFormatting struct {
	XMLName xml.Name      `xml:"conditionalFormatting"`
	Pivot   bool          `xml:"pivot,attr,omitempty"`
	SQRef   string        `xml:"sqref,attr,omitempty"`
	CfRule  []*xlsxCfRule `xml:"cfRule"`
}

type xlsxCfRule struct {
	Type         string          `xml:"type,attr,omitempty"`
	DxfID        *int            `xml:"dxfId,attr"`
	Priority     int             `xml:"priority,attr,omitempty"`
	StopIfTrue   bool            `xml:"stopIfTrue,attr,omitempty"`
	AboveAverage *bool           `xml:"aboveAverage,attr"`
	Percent      bool            `xml:"percent,attr,omitempty"`
	Bottom       bool            `xml:"bottom,attr,omitempty"`
	Operator     string          `xml:"operator,attr,omitempty"`
	Text         string          `xml:"text,attr,omitempty"`
	TimePeriod   string          `xml:"timePeriod,attr,omitempty"`
	Rank         int             `xml:"rank,attr,omitempty"`
	StdDev       int             `xml:"stdDev,attr,omitempty"`
	EqualAverage bool            `xml:"equalAverage,attr,omitempty"`
	Formula      []string        `xml:"formula,omitempty"`
	ColorScale   *xlsxColorScale `xml:"colorScale"`
	DataBar      *xlsxDataBar    `xml:"dataBar"`
	IconSet      *xlsxIconSet    `xml:"iconSet"`
	ExtLst       *xlsxExtLst     `xml:"extLst"`
}

type xlsxColorScale struct {
	Cfvo  []*xlsxCfvo  `xml:"cfvo"`
	Color []*xlsxColor `xml:"color"`
}

type xlsxDataBar struct {
	MaxLength int          `xml:"maxLength,attr,omitempty"`
	MinLength int          `xml:"minLength,attr,omitempty"`
	ShowValue *bool        `xml:"showValue,attr"`
	Cfvo      []*xlsxCfvo  `xml:"cfvo"`
	Color     []*xlsxColor `xml:"color"`
}

type xlsxIconSet struct {
	Cfvo      []*xlsxCfvo `xml:"cfvo"`
	IconSet   string      `xml:"iconSet,attr,omitempty"`
	ShowValue *bool       `xml:"showValue,attr"`
	Percent   bool        `xml:"percent,attr,omitempty"`
	Reverse   bool        `xml:"reverse,attr,omitempty"`
}

type xlsxCfvo struct {
	Gte    bool        `xml:"gte,attr,omitempty"`
	Type   string      `xml:"type,attr,omitempty"`
	Val    string      `xml:"val,attr,omitempty"`
	ExtLst *xlsxExtLst `xml:"extLst"`
}

type xlsxHyperlinks struct {
	XMLName   xml.Name        `xml:"hyperlinks"`
	Hyperlink []xlsxHyperlink `xml:"hyperlink"`
}

type xlsxHyperlink struct {
	Ref      string `xml:"ref,attr"`
	Location string `xml:"location,attr,omitempty"`
	Display  string `xml:"display,attr,omitempty"`
	Tooltip  string `xml:"tooltip,attr,omitempty"`
	RID      string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type xlsxTableParts struct {
	XMLName    xml.Name         `xml:"tableParts"`
	Count      int              `xml:"count,attr,omitempty"`
	TableParts []*xlsxTablePart `xml:"tablePart"`
}

type xlsxTablePart struct {
	RID string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type xlsxPicture struct {
	XMLName xml.Name `xml:"picture"`
	RID     string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type xlsxIgnoredError struct {
	XMLName            xml.Name `xml:"ignoredError"`
	Sqref              string   `xml:"sqref,attr"`
	EvalError          bool     `xml:"evalError,attr,omitempty"`
	TwoDigitTextYear   bool     `xml:"twoDigitTextYear,attr,omitempty"`
	NumberStoredAsText bool     `xml:"numberStoredAsText,attr,omitempty"`
	Formula            bool     `xml:"formula,attr,omitempty"`
	FormulaRange       bool     `xml:"formulaRange,attr,omitempty"`
	UnlockedFormula    bool     `xml:"unlockedFormula,attr,omitempty"`
	EmptyCellReference bool     `xml:"emptyCellReference,attr,omitempty"`
	ListDataValidation bool     `xml:"listDataValidation,attr,omitempty"`
	CalculatedColumn   bool     `xml:"calculatedColumn,attr,omitempty"`
}

type xlsxIgnoredErrors struct {
	XMLName      xml.Name           `xml:"ignoredErrors"`
	IgnoredError []xlsxIgnoredError `xml:"ignoredError"`
	ExtLst       *xlsxExtLst        `xml:"extLst"`
}

type xlsxLegacyDrawing struct {
	XMLName xml.Name `xml:"legacyDrawing"`
	RID     string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type xlsxLegacyDrawingHF struct {
	XMLName xml.Name `xml:"legacyDrawingHF"`
	RID     string   `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type decodeX14SparklineGroups struct {
	XMLName xml.Name `xml:"sparklineGroups"`
	XMLNSXM string   `xml:"xmlns:xm,attr"`
	Content string   `xml:",innerxml"`
}

type decodeX14ConditionalFormattingExt struct {
	XMLName xml.Name `xml:"ext"`
	ID      string   `xml:"id"`
}

type decodeX14ConditionalFormattings struct {
	XMLName xml.Name `xml:"conditionalFormattings"`
	XMLNSXM string   `xml:"xmlns:xm,attr"`
	Content string   `xml:",innerxml"`
}

type decodeX14ConditionalFormattingRules struct {
	XMLName xml.Name                         `xml:"conditionalFormattings"`
	XMLNSXM string                           `xml:"xmlns:xm,attr"`
	CondFmt []decodeX14ConditionalFormatting `xml:"conditionalFormatting"`
}

type decodeX14ConditionalFormatting struct {
	XMLName xml.Name           `xml:"conditionalFormatting"`
	Pivot   bool               `xml:"pivot,attr,omitempty"`
	CfRule  []*decodeX14CfRule `xml:"cfRule"`
	Sqref   string             `xml:"sqref,omitempty"`
	ExtLst  *xlsxExtLst        `xml:"x14:extLst"`
}

type decodeX14CfRule struct {
	XMLName       xml.Name          `xml:"cfRule"`
	Type          string            `xml:"type,attr,omitempty"`
	Priority      int               `xml:"priority,attr,omitempty"`
	StopIfTrue    bool              `xml:"stopIfTrue,attr,omitempty"`
	AboveAverage  *bool             `xml:"aboveAverage,attr"`
	Percent       bool              `xml:"percent,attr,omitempty"`
	Bottom        bool              `xml:"bottom,attr,omitempty"`
	Operator      string            `xml:"operator,attr,omitempty"`
	Text          string            `xml:"text,attr,omitempty"`
	TimePeriod    string            `xml:"timePeriod,attr,omitempty"`
	Rank          int               `xml:"rank,attr,omitempty"`
	StdDev        int               `xml:"stdDev,attr,omitempty"`
	EqualAverage  bool              `xml:"equalAverage,attr,omitempty"`
	ActivePresent bool              `xml:"activePresent,attr,omitempty"`
	ID            string            `xml:"id,attr,omitempty"`
	F             []string          `xml:"f"`
	ColorScale    *xlsxInnerXML     `xml:"colorScale"`
	DataBar       *decodeX14DataBar `xml:"dataBar"`
	IconSet       *decodeX14IconSet `xml:"iconSet"`
	Dxf           *xlsxInnerXML     `xml:"dxf"`
	ExtLst        *xlsxExtLst       `xml:"extLst"`
}

type decodeX14DataBar struct {
	XMLName           xml.Name    `xml:"dataBar"`
	MaxLength         int         `xml:"maxLength,attr"`
	MinLength         int         `xml:"minLength,attr"`
	Border            bool        `xml:"border,attr,omitempty"`
	Gradient          *bool       `xml:"gradient,attr"`
	ShowValue         bool        `xml:"showValue,attr,omitempty"`
	Direction         string      `xml:"direction,attr,omitempty"`
	Cfvo              []*xlsxCfvo `xml:"cfvo"`
	BorderColor       *xlsxColor  `xml:"borderColor"`
	NegativeFillColor *xlsxColor  `xml:"negativeFillColor"`
	AxisColor         *xlsxColor  `xml:"axisColor"`
}

type decodeX14IconSet struct {
	XMLName   xml.Name         `xml:"iconSet"`
	IconSet   string           `xml:"iconSet,attr,omitempty"`
	ShowValue *bool            `xml:"showValue,attr"`
	Percent   *bool            `xml:"percent,attr"`
	Reverse   bool             `xml:"reverse,attr,omitempty"`
	Custom    bool             `xml:"custom,attr,omitempty"`
	Cfvo      []*decodeX14Cfvo `xml:"cfvo"`
	CfIcon    []*xlsxInnerXML  `xml:"cfIcon"`
}

type decodeX14Cfvo struct {
	XMLName xml.Name    `xml:"cfvo"`
	Type    string      `xml:"type,attr"`
	Gte     *bool       `xml:"gte,attr"`
	F       string      `xml:"f"`
	ExtLst  *xlsxExtLst `xml:"extLst"`
}

type xlsxX14ConditionalFormattings struct {
	XMLName xml.Name `xml:"x14:conditionalFormattings"`
	Content string   `xml:",innerxml"`
}

type xlsxX14ConditionalFormatting struct {
	XMLName xml.Name         `xml:"x14:conditionalFormatting"`
	XMLNSXM string           `xml:"xmlns:xm,attr"`
	Pivot   bool             `xml:"pivot,attr,omitempty"`
	CfRule  []*xlsxX14CfRule `xml:"x14:cfRule"`
	Sqref   string           `xml:"xm:sqref,omitempty"`
	ExtLst  *xlsxExtLst      `xml:"x14:extLst"`
}

type xlsxX14CfRule struct {
	Type          string         `xml:"type,attr,omitempty"`
	Priority      int            `xml:"priority,attr,omitempty"`
	StopIfTrue    bool           `xml:"stopIfTrue,attr,omitempty"`
	AboveAverage  *bool          `xml:"aboveAverage,attr"`
	Percent       bool           `xml:"percent,attr,omitempty"`
	Bottom        bool           `xml:"bottom,attr,omitempty"`
	Operator      string         `xml:"operator,attr,omitempty"`
	Text          string         `xml:"text,attr,omitempty"`
	TimePeriod    string         `xml:"timePeriod,attr,omitempty"`
	Rank          int            `xml:"rank,attr,omitempty"`
	StdDev        int            `xml:"stdDev,attr,omitempty"`
	EqualAverage  bool           `xml:"equalAverage,attr,omitempty"`
	ActivePresent bool           `xml:"activePresent,attr,omitempty"`
	ID            string         `xml:"id,attr,omitempty"`
	F             []string       `xml:"xm:f"`
	ColorScale    *xlsxInnerXML  `xml:"x14:colorScale"`
	DataBar       *xlsx14DataBar `xml:"x14:dataBar"`
	IconSet       *xlsx14IconSet `xml:"x14:iconSet"`
	Dxf           *xlsxInnerXML  `xml:"x14:dxf"`
	ExtLst        *xlsxExtLst    `xml:"x14:extLst"`
}

type xlsx14DataBar struct {
	MaxLength         int         `xml:"maxLength,attr"`
	MinLength         int         `xml:"minLength,attr"`
	Border            bool        `xml:"border,attr"`
	Gradient          bool        `xml:"gradient,attr"`
	ShowValue         bool        `xml:"showValue,attr,omitempty"`
	Direction         string      `xml:"direction,attr,omitempty"`
	Cfvo              []*xlsxCfvo `xml:"x14:cfvo"`
	BorderColor       *xlsxColor  `xml:"x14:borderColor"`
	NegativeFillColor *xlsxColor  `xml:"x14:negativeFillColor"`
	AxisColor         *xlsxColor  `xml:"x14:axisColor"`
}

type xlsx14IconSet struct {
	IconSet   string          `xml:"iconSet,attr,omitempty"`
	ShowValue *bool           `xml:"showValue,attr"`
	Percent   *bool           `xml:"percent,attr"`
	Reverse   bool            `xml:"reverse,attr,omitempty"`
	Custom    bool            `xml:"custom,attr,omitempty"`
	Cfvo      []*xlsx14Cfvo   `xml:"x14:cfvo"`
	CfIcon    []*xlsxInnerXML `xml:"x14:cfIcon"`
}

type xlsx14Cfvo struct {
	Type   string      `xml:"type,attr"`
	Gte    *bool       `xml:"gte,attr"`
	F      string      `xml:"xm:f"`
	ExtLst *xlsxExtLst `xml:"x14:extLst"`
}

type xlsxX14SparklineGroups struct {
	XMLName         xml.Name                 `xml:"x14:sparklineGroups"`
	XMLNSXM         string                   `xml:"xmlns:xm,attr"`
	SparklineGroups []*xlsxX14SparklineGroup `xml:"x14:sparklineGroup"`
	Content         string                   `xml:",innerxml"`
}

type xlsxX14SparklineGroup struct {
	XMLName             xml.Name          `xml:"x14:sparklineGroup"`
	ManualMax           int               `xml:"manualMax,attr,omitempty"`
	ManualMin           int               `xml:"manualMin,attr,omitempty"`
	LineWeight          float64           `xml:"lineWeight,attr,omitempty"`
	Type                string            `xml:"type,attr,omitempty"`
	DateAxis            bool              `xml:"dateAxis,attr,omitempty"`
	DisplayEmptyCellsAs string            `xml:"displayEmptyCellsAs,attr,omitempty"`
	Markers             bool              `xml:"markers,attr,omitempty"`
	High                bool              `xml:"high,attr,omitempty"`
	Low                 bool              `xml:"low,attr,omitempty"`
	First               bool              `xml:"first,attr,omitempty"`
	Last                bool              `xml:"last,attr,omitempty"`
	Negative            bool              `xml:"negative,attr,omitempty"`
	DisplayXAxis        bool              `xml:"displayXAxis,attr,omitempty"`
	DisplayHidden       bool              `xml:"displayHidden,attr,omitempty"`
	MinAxisType         string            `xml:"minAxisType,attr,omitempty"`
	MaxAxisType         string            `xml:"maxAxisType,attr,omitempty"`
	RightToLeft         bool              `xml:"rightToLeft,attr,omitempty"`
	ColorSeries         *xlsxColor        `xml:"x14:colorSeries"`
	ColorNegative       *xlsxColor        `xml:"x14:colorNegative"`
	ColorAxis           *xlsxColor        `xml:"x14:colorAxis"`
	ColorMarkers        *xlsxColor        `xml:"x14:colorMarkers"`
	ColorFirst          *xlsxColor        `xml:"x14:colorFirst"`
	ColorLast           *xlsxColor        `xml:"x14:colorLast"`
	ColorHigh           *xlsxColor        `xml:"x14:colorHigh"`
	ColorLow            *xlsxColor        `xml:"x14:colorLow"`
	Sparklines          xlsxX14Sparklines `xml:"x14:sparklines"`
}

type xlsxX14Sparklines struct {
	Sparkline []*xlsxX14Sparkline `xml:"x14:sparkline"`
}

type xlsxX14Sparkline struct {
	F     string `xml:"xm:f"`
	Sqref string `xml:"xm:sqref"`
}

type DataValidation struct {
	AllowBlank       bool
	Error            *string
	ErrorStyle       *string
	ErrorTitle       *string
	Operator         string
	Prompt           *string
	PromptTitle      *string
	ShowDropDown     bool
	ShowErrorMessage bool
	ShowInputMessage bool
	Sqref            string
	Type             string
	Formula1         string
	Formula2         string
}

type SparklineOptions struct {
	Location      []string
	Range         []string
	Max           int
	CustMax       int
	Min           int
	CustMin       int
	Type          string
	Weight        float64
	DateAxis      bool
	Markers       bool
	High          bool
	Low           bool
	First         bool
	Last          bool
	Negative      bool
	Axis          bool
	Hidden        bool
	Reverse       bool
	Style         int
	SeriesColor   string
	NegativeColor string
	MarkersColor  string
	FirstColor    string
	LastColor     string
	HightColor    string
	LowColor      string
	EmptyCells    string
}

type Selection struct {
	SQRef      string
	ActiveCell string
	Pane       string
}

type Panes struct {
	Freeze      bool
	Split       bool
	XSplit      int
	YSplit      int
	TopLeftCell string
	ActivePane  string
	Selection   []Selection
}

type ConditionalFormatOptions struct {
	Type           string
	AboveAverage   bool
	Percent        bool
	Format         *int
	Criteria       string
	Value          string
	MinType        string
	MidType        string
	MaxType        string
	MinValue       string
	MidValue       string
	MaxValue       string
	MinColor       string
	MidColor       string
	MaxColor       string
	BarColor       string
	BarBorderColor string
	BarDirection   string
	BarOnly        bool
	BarSolid       bool
	IconStyle      string
	ReverseIcons   bool
	IconsOnly      bool
	StopIfTrue     bool
}

type SheetProtectionOptions struct {
	AlgorithmName       string
	AutoFilter          bool
	DeleteColumns       bool
	DeleteRows          bool
	EditObjects         bool
	EditScenarios       bool
	FormatCells         bool
	FormatColumns       bool
	FormatRows          bool
	InsertColumns       bool
	InsertHyperlinks    bool
	InsertRows          bool
	Password            string
	PivotTables         bool
	SelectLockedCells   bool
	SelectUnlockedCells bool
	Sort                bool
}

type HeaderFooterOptions struct {
	AlignWithMargins *bool
	DifferentFirst   bool
	DifferentOddEven bool
	ScaleWithDoc     *bool
	OddHeader        string
	OddFooter        string
	EvenHeader       string
	EvenFooter       string
	FirstHeader      string
	FirstFooter      string
}

type PageLayoutMarginsOptions struct {
	Bottom       *float64
	Footer       *float64
	Header       *float64
	Left         *float64
	Right        *float64
	Top          *float64
	Horizontally *bool
	Vertically   *bool
}

type PageLayoutOptions struct {
	Size *int

	Orientation *string

	FirstPageNumber *uint

	AdjustTo *uint

	FitToHeight *int

	FitToWidth *int

	BlackAndWhite *bool

	PageOrder *string
}

type ViewOptions struct {
	DefaultGridColor *bool

	RightToLeft *bool

	ShowFormulas *bool

	ShowGridLines *bool

	ShowRowColHeaders *bool

	ShowRuler *bool

	ShowZeros *bool

	TopLeftCell *string

	View *string

	ZoomScale *float64
}

type SheetPropsOptions struct {
	CodeName *string

	EnableFormatConditionsCalculation *bool

	Published *bool

	AutoPageBreaks *bool

	FitToPage *bool

	TabColorIndexed *int

	TabColorRGB *string

	TabColorTheme *int

	TabColorTint *float64

	OutlineSummaryBelow *bool

	OutlineSummaryRight *bool

	BaseColWidth *uint8

	DefaultColWidth *float64

	DefaultRowHeight *float64

	CustomHeight *bool

	ZeroHeight *bool

	ThickTop *bool

	ThickBottom *bool
}
