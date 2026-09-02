package excelize

import (
	"encoding/xml"
	"sync"
)

type xlsxRelationships struct {
	mu            sync.Mutex
	XMLName       xml.Name           `xml:"http://schemas.openxmlformats.org/package/2006/relationships Relationships"`
	Relationships []xlsxRelationship `xml:"Relationship"`
}

type xlsxRelationship struct {
	ID         string `xml:"Id,attr"`
	Target     string `xml:",attr"`
	Type       string `xml:",attr"`
	TargetMode string `xml:",attr,omitempty"`
}

type xlsxWorkbook struct {
	XMLName                xml.Name                 `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main workbook"`
	Conformance            string                   `xml:"conformance,attr,omitempty"`
	FileVersion            *xlsxFileVersion         `xml:"fileVersion"`
	FileSharing            *xlsxExtLst              `xml:"fileSharing"`
	WorkbookPr             *xlsxWorkbookPr          `xml:"workbookPr"`
	AlternateContent       *xlsxAlternateContent    `xml:"mc:AlternateContent"`
	DecodeAlternateContent *xlsxInnerXML            `xml:"http://schemas.openxmlformats.org/markup-compatibility/2006 AlternateContent"`
	WorkbookProtection     *xlsxWorkbookProtection  `xml:"workbookProtection"`
	BookViews              *xlsxBookViews           `xml:"bookViews"`
	Sheets                 xlsxSheets               `xml:"sheets"`
	FunctionGroups         *xlsxFunctionGroups      `xml:"functionGroups"`
	ExternalReferences     *xlsxExternalReferences  `xml:"externalReferences"`
	DefinedNames           *xlsxDefinedNames        `xml:"definedNames"`
	CalcPr                 *xlsxCalcPr              `xml:"calcPr"`
	OleSize                *xlsxExtLst              `xml:"oleSize"`
	CustomWorkbookViews    *xlsxCustomWorkbookViews `xml:"customWorkbookViews"`
	PivotCaches            *xlsxPivotCaches         `xml:"pivotCaches"`
	SmartTagPr             *xlsxExtLst              `xml:"smartTagPr"`
	SmartTagTypes          *xlsxExtLst              `xml:"smartTagTypes"`
	WebPublishing          *xlsxExtLst              `xml:"webPublishing"`
	FileRecoveryPr         *xlsxFileRecoveryPr      `xml:"fileRecoveryPr"`
	WebPublishObjects      *xlsxExtLst              `xml:"webPublishObjects"`
	ExtLst                 *xlsxExtLst              `xml:"extLst"`
}

type xlsxFileRecoveryPr struct {
	AutoRecover     bool `xml:"autoRecover,attr,omitempty"`
	CrashSave       bool `xml:"crashSave,attr,omitempty"`
	DataExtractLoad bool `xml:"dataExtractLoad,attr,omitempty"`
	RepairLoad      bool `xml:"repairLoad,attr,omitempty"`
}

type xlsxWorkbookProtection struct {
	LockRevision           bool   `xml:"lockRevision,attr,omitempty"`
	LockStructure          bool   `xml:"lockStructure,attr,omitempty"`
	LockWindows            bool   `xml:"lockWindows,attr,omitempty"`
	RevisionsAlgorithmName string `xml:"revisionsAlgorithmName,attr,omitempty"`
	RevisionsHashValue     string `xml:"revisionsHashValue,attr,omitempty"`
	RevisionsSaltValue     string `xml:"revisionsSaltValue,attr,omitempty"`
	RevisionsSpinCount     int    `xml:"revisionsSpinCount,attr,omitempty"`
	WorkbookAlgorithmName  string `xml:"workbookAlgorithmName,attr,omitempty"`
	WorkbookHashValue      string `xml:"workbookHashValue,attr,omitempty"`
	WorkbookSaltValue      string `xml:"workbookSaltValue,attr,omitempty"`
	WorkbookSpinCount      int    `xml:"workbookSpinCount,attr,omitempty"`
}

type xlsxFileVersion struct {
	AppName      string `xml:"appName,attr,omitempty"`
	CodeName     string `xml:"codeName,attr,omitempty"`
	LastEdited   string `xml:"lastEdited,attr,omitempty"`
	LowestEdited string `xml:"lowestEdited,attr,omitempty"`
	RupBuild     string `xml:"rupBuild,attr,omitempty"`
}

type xlsxWorkbookPr struct {
	Date1904                   bool   `xml:"date1904,attr,omitempty"`
	ShowObjects                string `xml:"showObjects,attr,omitempty"`
	ShowBorderUnselectedTables *bool  `xml:"showBorderUnselectedTables,attr"`
	FilterPrivacy              bool   `xml:"filterPrivacy,attr,omitempty"`
	PromptedSolutions          bool   `xml:"promptedSolutions,attr,omitempty"`
	ShowInkAnnotation          *bool  `xml:"showInkAnnotation,attr"`
	BackupFile                 bool   `xml:"backupFile,attr,omitempty"`
	SaveExternalLinkValues     *bool  `xml:"saveExternalLinkValues,attr"`
	UpdateLinks                string `xml:"updateLinks,attr,omitempty"`
	CodeName                   string `xml:"codeName,attr,omitempty"`
	HidePivotFieldList         bool   `xml:"hidePivotFieldList,attr,omitempty"`
	ShowPivotChartFilter       bool   `xml:"showPivotChartFilter,attr,omitempty"`
	AllowRefreshQuery          bool   `xml:"allowRefreshQuery,attr,omitempty"`
	PublishItems               bool   `xml:"publishItems,attr,omitempty"`
	CheckCompatibility         bool   `xml:"checkCompatibility,attr,omitempty"`
	AutoCompressPictures       *bool  `xml:"autoCompressPictures,attr"`
	RefreshAllConnections      bool   `xml:"refreshAllConnections,attr,omitempty"`
	DefaultThemeVersion        string `xml:"defaultThemeVersion,attr,omitempty"`
}

type xlsxBookViews struct {
	WorkBookView []xlsxWorkBookView `xml:"workbookView"`
}

type xlsxWorkBookView struct {
	Visibility             string  `xml:"visibility,attr,omitempty"`
	Minimized              bool    `xml:"minimized,attr,omitempty"`
	ShowHorizontalScroll   *bool   `xml:"showHorizontalScroll,attr"`
	ShowVerticalScroll     *bool   `xml:"showVerticalScroll,attr"`
	ShowSheetTabs          *bool   `xml:"showSheetTabs,attr"`
	XWindow                string  `xml:"xWindow,attr,omitempty"`
	YWindow                string  `xml:"yWindow,attr,omitempty"`
	WindowWidth            int     `xml:"windowWidth,attr,omitempty"`
	WindowHeight           int     `xml:"windowHeight,attr,omitempty"`
	TabRatio               float64 `xml:"tabRatio,attr,omitempty"`
	FirstSheet             int     `xml:"firstSheet,attr,omitempty"`
	ActiveTab              int     `xml:"activeTab,attr,omitempty"`
	AutoFilterDateGrouping *bool   `xml:"autoFilterDateGrouping,attr"`
}

type xlsxSheets struct {
	Sheet []xlsxSheet `xml:"sheet"`
}

type xlsxSheet struct {
	Name    string `xml:"name,attr,omitempty"`
	SheetID int    `xml:"sheetId,attr,omitempty"`
	ID      string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
	State   string `xml:"state,attr,omitempty"`
}

type xlsxFunctionGroup struct {
	Name string `xml:"name,attr"`
}

type xlsxFunctionGroups struct {
	BuiltInGroupCount *int                `xml:"builtInGroupCount,attr"`
	FunctionGroup     []xlsxFunctionGroup `xml:"functionGroup"`
}

type xlsxExternalReferences struct {
	ExternalReference []xlsxExternalReference `xml:"externalReference"`
}

type xlsxExternalReference struct {
	RID string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type xlsxPivotCaches struct {
	PivotCache []xlsxPivotCache `xml:"pivotCache"`
}

type xlsxPivotCache struct {
	CacheID int    `xml:"cacheId,attr"`
	RID     string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
}

type xlsxExtLst struct {
	Ext string `xml:",innerxml"`
}

type xlsxExt struct {
	XMLName xml.Name `xml:"ext"`
	URI     string   `xml:"uri,attr"`
	Content string   `xml:",innerxml"`
	xmlns   []xml.Attr
}

type xlsxAlternateContent struct {
	XMLNSMC string `xml:"xmlns:mc,attr,omitempty"`
	Content string `xml:",innerxml"`
}

type xlsxChoice struct {
	XMLName    xml.Name `xml:"mc:Choice"`
	XMLNSA14   string   `xml:"xmlns:a14,attr,omitempty"`
	XMLNSSle15 string   `xml:"xmlns:sle15,attr,omitempty"`
	Requires   string   `xml:"Requires,attr,omitempty"`
	Content    string   `xml:",innerxml"`
}

type xlsxFallback struct {
	XMLName xml.Name `xml:"mc:Fallback"`
	Content string   `xml:",innerxml"`
}

type xlsxInnerXML struct {
	Content string `xml:",innerxml"`
}

type decodeExtLst struct {
	XMLName xml.Name   `xml:"extLst"`
	Ext     []*xlsxExt `xml:"ext"`
}

type decodeExt struct {
	URI     string `xml:"uri,attr,omitempty"`
	Content string `xml:",innerxml"`
}

type xlsxDefinedNames struct {
	DefinedName []xlsxDefinedName `xml:"definedName"`
}

type xlsxDefinedName struct {
	Comment           string `xml:"comment,attr,omitempty"`
	CustomMenu        string `xml:"customMenu,attr,omitempty"`
	Description       string `xml:"description,attr,omitempty"`
	Function          bool   `xml:"function,attr,omitempty"`
	FunctionGroupID   int    `xml:"functionGroupId,attr,omitempty"`
	Help              string `xml:"help,attr,omitempty"`
	Hidden            bool   `xml:"hidden,attr,omitempty"`
	LocalSheetID      *int   `xml:"localSheetId,attr"`
	Name              string `xml:"name,attr,omitempty"`
	PublishToServer   bool   `xml:"publishToServer,attr,omitempty"`
	ShortcutKey       string `xml:"shortcutKey,attr,omitempty"`
	StatusBar         string `xml:"statusBar,attr,omitempty"`
	VbProcedure       bool   `xml:"vbProcedure,attr,omitempty"`
	WorkbookParameter bool   `xml:"workbookParameter,attr,omitempty"`
	Xlm               bool   `xml:"xml,attr,omitempty"`
	Data              string `xml:",chardata"`
}

type xlsxCalcPr struct {
	CalcCompleted         bool    `xml:"calcCompleted,attr,omitempty"`
	CalcID                int     `xml:"calcId,attr,omitempty"`
	CalcMode              string  `xml:"calcMode,attr,omitempty"`
	CalcOnSave            bool    `xml:"calcOnSave,attr,omitempty"`
	ConcurrentCalc        *bool   `xml:"concurrentCalc,attr"`
	ConcurrentManualCount int     `xml:"concurrentManualCount,attr,omitempty"`
	ForceFullCalc         bool    `xml:"forceFullCalc,attr,omitempty"`
	FullCalcOnLoad        bool    `xml:"fullCalcOnLoad,attr,omitempty"`
	FullPrecision         bool    `xml:"fullPrecision,attr,omitempty"`
	Iterate               bool    `xml:"iterate,attr,omitempty"`
	IterateCount          int     `xml:"iterateCount,attr,omitempty"`
	IterateDelta          float64 `xml:"iterateDelta,attr,omitempty"`
	RefMode               string  `xml:"refMode,attr,omitempty"`
}

type xlsxCustomWorkbookViews struct {
	CustomWorkbookView []xlsxCustomWorkbookView `xml:"customWorkbookView"`
}

type xlsxCustomWorkbookView struct {
	ActiveSheetID        *int     `xml:"activeSheetId,attr"`
	AutoUpdate           *bool    `xml:"autoUpdate,attr"`
	ChangesSavedWin      *bool    `xml:"changesSavedWin,attr"`
	GUID                 *string  `xml:"guid,attr"`
	IncludeHiddenRowCol  *bool    `xml:"includeHiddenRowCol,attr"`
	IncludePrintSettings *bool    `xml:"includePrintSettings,attr"`
	Maximized            *bool    `xml:"maximized,attr"`
	MergeInterval        int      `xml:"mergeInterval,attr"`
	Minimized            *bool    `xml:"minimized,attr"`
	Name                 *string  `xml:"name,attr"`
	OnlySync             *bool    `xml:"onlySync,attr"`
	PersonalView         *bool    `xml:"personalView,attr"`
	ShowComments         *string  `xml:"showComments,attr"`
	ShowFormulaBar       *bool    `xml:"showFormulaBar,attr"`
	ShowHorizontalScroll *bool    `xml:"showHorizontalScroll,attr"`
	ShowObjects          *string  `xml:"showObjects,attr"`
	ShowSheetTabs        *bool    `xml:"showSheetTabs,attr"`
	ShowStatusbar        *bool    `xml:"showStatusbar,attr"`
	ShowVerticalScroll   *bool    `xml:"showVerticalScroll,attr"`
	TabRatio             *float64 `xml:"tabRatio,attr"`
	WindowHeight         *int     `xml:"windowHeight,attr"`
	WindowWidth          *int     `xml:"windowWidth,attr"`
	XWindow              *int     `xml:"xWindow,attr"`
	YWindow              *int     `xml:"yWindow,attr"`
}

type DefinedName struct {
	Name     string
	Comment  string
	RefersTo string
	Scope    string
}

type CalcPropsOptions struct {
	CalcID                *uint
	CalcMode              *string
	FullCalcOnLoad        *bool
	RefMode               *string
	Iterate               *bool
	IterateCount          *uint
	IterateDelta          *float64
	FullPrecision         *bool
	CalcCompleted         *bool
	CalcOnSave            *bool
	ConcurrentCalc        *bool
	ConcurrentManualCount *uint
	ForceFullCalc         *bool
}

type WorkbookPropsOptions struct {
	Date1904      *bool
	FilterPrivacy *bool
	CodeName      *string
}

type WorkbookProtectionOptions struct {
	AlgorithmName string
	Password      string
	LockStructure bool
	LockWindows   bool
}
