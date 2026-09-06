package excelize

import "encoding/xml"

type xlsxPivotCacheDefinition struct {
	XMLName               xml.Name               `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main pivotCacheDefinition"`
	RID                   string                 `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
	Invalid               bool                   `xml:"invalid,attr,omitempty"`
	SaveData              bool                   `xml:"saveData,attr"`
	RefreshOnLoad         bool                   `xml:"refreshOnLoad,attr,omitempty"`
	OptimizeMemory        bool                   `xml:"optimizeMemory,attr,omitempty"`
	EnableRefresh         bool                   `xml:"enableRefresh,attr,omitempty"`
	RefreshedBy           string                 `xml:"refreshedBy,attr,omitempty"`
	RefreshedDate         float64                `xml:"refreshedDate,attr,omitempty"`
	RefreshedDateIso      float64                `xml:"refreshedDateIso,attr,omitempty"`
	BackgroundQuery       bool                   `xml:"backgroundQuery,attr"`
	MissingItemsLimit     int                    `xml:"missingItemsLimit,attr,omitempty"`
	CreatedVersion        int                    `xml:"createdVersion,attr,omitempty"`
	RefreshedVersion      int                    `xml:"refreshedVersion,attr,omitempty"`
	MinRefreshableVersion int                    `xml:"minRefreshableVersion,attr,omitempty"`
	RecordCount           int                    `xml:"recordCount,attr,omitempty"`
	UpgradeOnRefresh      bool                   `xml:"upgradeOnRefresh,attr,omitempty"`
	TupleCacheAttr        bool                   `xml:"tupleCache,attr,omitempty"`
	SupportSubquery       bool                   `xml:"supportSubquery,attr,omitempty"`
	SupportAdvancedDrill  bool                   `xml:"supportAdvancedDrill,attr,omitempty"`
	CacheSource           *xlsxCacheSource       `xml:"cacheSource"`
	CacheFields           *xlsxCacheFields       `xml:"cacheFields"`
	CacheHierarchies      *xlsxCacheHierarchies  `xml:"cacheHierarchies"`
	Kpis                  *xlsxKpis              `xml:"kpis"`
	TupleCache            *xlsxTupleCache        `xml:"tupleCache"`
	CalculatedItems       *xlsxCalculatedItems   `xml:"calculatedItems"`
	CalculatedMembers     *xlsxCalculatedMembers `xml:"calculatedMembers"`
	Dimensions            *xlsxDimensions        `xml:"dimensions"`
	MeasureGroups         *xlsxMeasureGroups     `xml:"measureGroups"`
	Maps                  *xlsxMaps              `xml:"maps"`
	ExtLst                *xlsxExtLst            `xml:"extLst"`
}

type xlsxCacheSource struct {
	Type            string               `xml:"type,attr"`
	ConnectionID    int                  `xml:"connectionId,attr,omitempty"`
	WorksheetSource *xlsxWorksheetSource `xml:"worksheetSource"`
	Consolidation   *xlsxConsolidation   `xml:"consolidation"`
	ExtLst          *xlsxExtLst          `xml:"extLst"`
}

type xlsxWorksheetSource struct {
	RID   string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr,omitempty"`
	Ref   string `xml:"ref,attr,omitempty"`
	Name  string `xml:"name,attr,omitempty"`
	Sheet string `xml:"sheet,attr,omitempty"`
}

type xlsxConsolidation struct{}

type xlsxCacheFields struct {
	Count      int               `xml:"count,attr"`
	CacheField []*xlsxCacheField `xml:"cacheField"`
}

type xlsxCacheField struct {
	Name                string           `xml:"name,attr"`
	Caption             string           `xml:"caption,attr,omitempty"`
	PropertyName        string           `xml:"propertyName,attr,omitempty"`
	ServerField         bool             `xml:"serverField,attr,omitempty"`
	UniqueList          bool             `xml:"uniqueList,attr,omitempty"`
	NumFmtID            int              `xml:"numFmtId,attr"`
	Formula             string           `xml:"formula,attr,omitempty"`
	SQLType             int              `xml:"sqlType,attr,omitempty"`
	Hierarchy           int              `xml:"hierarchy,attr,omitempty"`
	Level               int              `xml:"level,attr,omitempty"`
	DatabaseField       bool             `xml:"databaseField,attr,omitempty"`
	MappingCount        int              `xml:"mappingCount,attr,omitempty"`
	MemberPropertyField bool             `xml:"memberPropertyField,attr,omitempty"`
	SharedItems         *xlsxSharedItems `xml:"sharedItems"`
	FieldGroup          *xlsxFieldGroup  `xml:"fieldGroup"`
	MpMap               *xlsxX           `xml:"mpMap"`
	ExtLst              *xlsxExtLst      `xml:"extLst"`
}

type xlsxSharedItems struct {
	ContainsSemiMixedTypes *bool            `xml:"containsSemiMixedTypes,attr"`
	ContainsNonDate        *bool            `xml:"containsNonDate,attr"`
	ContainsDate           bool             `xml:"containsDate,attr,omitempty"`
	ContainsString         *bool            `xml:"containsString,attr"`
	ContainsBlank          bool             `xml:"containsBlank,attr,omitempty"`
	ContainsMixedTypes     bool             `xml:"containsMixedTypes,attr,omitempty"`
	ContainsNumber         bool             `xml:"containsNumber,attr,omitempty"`
	ContainsInteger        bool             `xml:"containsInteger,attr,omitempty"`
	MinValue               float64          `xml:"minValue,attr,omitempty"`
	MaxValue               float64          `xml:"maxValue,attr,omitempty"`
	MinDate                string           `xml:"minDate,attr,omitempty"`
	MaxDate                string           `xml:"maxDate,attr,omitempty"`
	Count                  int              `xml:"count,attr"`
	LongText               bool             `xml:"longText,attr,omitempty"`
	Items                  []xlsxSharedItem `xml:",any"`
}

type xlsxSharedItem struct {
	XMLName xml.Name
	V       string      `xml:"v,attr,omitempty"`
	U       bool        `xml:"u,attr,omitempty"`
	F       bool        `xml:"f,attr,omitempty"`
	C       string      `xml:"c,attr,omitempty"`
	Cp      int         `xml:"cp,attr,omitempty"`
	In      int         `xml:"in,attr,omitempty"`
	Bc      string      `xml:"bc,attr,omitempty"`
	Fc      string      `xml:"fc,attr,omitempty"`
	I       bool        `xml:"i,attr,omitempty"`
	Un      bool        `xml:"un,attr,omitempty"`
	St      bool        `xml:"st,attr,omitempty"`
	B       bool        `xml:"b,attr,omitempty"`
	Tpls    *xlsxTuples `xml:"tpls"`
	X       *attrValInt `xml:"x"`
}

type xlsxTuples struct{}

type xlsxFieldGroup struct{}

type xlsxCacheHierarchies struct{}

type xlsxKpis struct{}

type xlsxTupleCache struct{}

type xlsxCalculatedItems struct{}

type xlsxCalculatedMembers struct{}

type xlsxDimensions struct{}

type xlsxMeasureGroups struct{}

type xlsxMaps struct{}

type xlsxX14PivotCacheDefinition struct {
	XMLName      xml.Name `xml:"x14:pivotCacheDefinition"`
	PivotCacheID int      `xml:"pivotCacheId,attr"`
}

type decodeX14PivotCacheDefinition struct {
	XMLName      xml.Name `xml:"pivotCacheDefinition"`
	PivotCacheID int      `xml:"pivotCacheId,attr"`
}
