package excelize

import (
	"encoding/xml"
	"sync"
)

type xlsxStyleSheet struct {
	mu           sync.Mutex
	XMLName      xml.Name          `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main styleSheet"`
	NumFmts      *xlsxNumFmts      `xml:"numFmts"`
	Fonts        *xlsxFonts        `xml:"fonts"`
	Fills        *xlsxFills        `xml:"fills"`
	Borders      *xlsxBorders      `xml:"borders"`
	CellStyleXfs *xlsxCellStyleXfs `xml:"cellStyleXfs"`
	CellXfs      *xlsxCellXfs      `xml:"cellXfs"`
	CellStyles   *xlsxCellStyles   `xml:"cellStyles"`
	Dxfs         *xlsxDxfs         `xml:"dxfs"`
	TableStyles  *xlsxTableStyles  `xml:"tableStyles"`
	Colors       *xlsxStyleColors  `xml:"colors"`
	ExtLst       *xlsxExtLst       `xml:"extLst"`
}

type xlsxAlignment struct {
	Horizontal      string `xml:"horizontal,attr,omitempty"`
	Indent          int    `xml:"indent,attr,omitempty"`
	JustifyLastLine bool   `xml:"justifyLastLine,attr,omitempty"`
	ReadingOrder    uint64 `xml:"readingOrder,attr,omitempty"`
	RelativeIndent  int    `xml:"relativeIndent,attr,omitempty"`
	ShrinkToFit     bool   `xml:"shrinkToFit,attr,omitempty"`
	TextRotation    int    `xml:"textRotation,attr,omitempty"`
	Vertical        string `xml:"vertical,attr,omitempty"`
	WrapText        bool   `xml:"wrapText,attr,omitempty"`
}

type xlsxProtection struct {
	Hidden *bool `xml:"hidden,attr"`
	Locked *bool `xml:"locked,attr"`
}

type xlsxLine struct {
	Style string     `xml:"style,attr,omitempty"`
	Color *xlsxColor `xml:"color"`
}

type xlsxColor struct {
	Auto    bool    `xml:"auto,attr,omitempty"`
	RGB     string  `xml:"rgb,attr,omitempty"`
	Indexed int     `xml:"indexed,attr,omitempty"`
	Theme   *int    `xml:"theme,attr"`
	Tint    float64 `xml:"tint,attr,omitempty"`
}

type xlsxFonts struct {
	Count int         `xml:"count,attr"`
	Font  []*xlsxFont `xml:"font"`
}

type xlsxFont struct {
	Name      *attrValString `xml:"name"`
	Charset   *attrValInt    `xml:"charset"`
	Family    *attrValInt    `xml:"family"`
	B         *attrValBool   `xml:"b"`
	I         *attrValBool   `xml:"i"`
	Strike    *attrValBool   `xml:"strike"`
	Outline   *attrValBool   `xml:"outline"`
	Shadow    *attrValBool   `xml:"shadow"`
	Condense  *attrValBool   `xml:"condense"`
	Extend    *attrValBool   `xml:"extend"`
	Color     *xlsxColor     `xml:"color"`
	Sz        *attrValFloat  `xml:"sz"`
	U         *attrValString `xml:"u"`
	VertAlign *attrValString `xml:"vertAlign"`
	Scheme    *attrValString `xml:"scheme"`
}

type xlsxFills struct {
	Count int         `xml:"count,attr"`
	Fill  []*xlsxFill `xml:"fill"`
}

type xlsxFill struct {
	PatternFill  *xlsxPatternFill  `xml:"patternFill"`
	GradientFill *xlsxGradientFill `xml:"gradientFill"`
}

type xlsxPatternFill struct {
	PatternType string     `xml:"patternType,attr,omitempty"`
	FgColor     *xlsxColor `xml:"fgColor"`
	BgColor     *xlsxColor `xml:"bgColor"`
}

type xlsxGradientFill struct {
	Bottom float64                 `xml:"bottom,attr,omitempty"`
	Degree float64                 `xml:"degree,attr,omitempty"`
	Left   float64                 `xml:"left,attr,omitempty"`
	Right  float64                 `xml:"right,attr,omitempty"`
	Top    float64                 `xml:"top,attr,omitempty"`
	Type   string                  `xml:"type,attr,omitempty"`
	Stop   []*xlsxGradientFillStop `xml:"stop"`
}

type xlsxGradientFillStop struct {
	Position float64   `xml:"position,attr"`
	Color    xlsxColor `xml:"color,omitempty"`
}

type xlsxBorders struct {
	Count  int           `xml:"count,attr"`
	Border []*xlsxBorder `xml:"border"`
}

type xlsxBorder struct {
	DiagonalDown bool      `xml:"diagonalDown,attr,omitempty"`
	DiagonalUp   bool      `xml:"diagonalUp,attr,omitempty"`
	Outline      bool      `xml:"outline,attr,omitempty"`
	Left         *xlsxLine `xml:"left"`
	Right        *xlsxLine `xml:"right"`
	Top          *xlsxLine `xml:"top"`
	Bottom       *xlsxLine `xml:"bottom"`
	Diagonal     *xlsxLine `xml:"diagonal"`
	Vertical     *xlsxLine `xml:"vertical"`
	Horizontal   *xlsxLine `xml:"horizontal"`
}

type xlsxCellStyles struct {
	XMLName   xml.Name         `xml:"cellStyles"`
	Count     int              `xml:"count,attr"`
	CellStyle []*xlsxCellStyle `xml:"cellStyle"`
}

type xlsxCellStyle struct {
	XMLName       xml.Name `xml:"cellStyle"`
	Name          string   `xml:"name,attr"`
	XfID          int      `xml:"xfId,attr"`
	BuiltInID     *int     `xml:"builtinId,attr"`
	ILevel        *int     `xml:"iLevel,attr"`
	Hidden        *bool    `xml:"hidden,attr"`
	CustomBuiltIn *bool    `xml:"customBuiltin,attr"`
}

type xlsxCellStyleXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

type xlsxXf struct {
	NumFmtID          *int            `xml:"numFmtId,attr"`
	FontID            *int            `xml:"fontId,attr"`
	FillID            *int            `xml:"fillId,attr"`
	BorderID          *int            `xml:"borderId,attr"`
	XfID              *int            `xml:"xfId,attr"`
	QuotePrefix       *bool           `xml:"quotePrefix,attr"`
	PivotButton       *bool           `xml:"pivotButton,attr"`
	ApplyNumberFormat *bool           `xml:"applyNumberFormat,attr"`
	ApplyFont         *bool           `xml:"applyFont,attr"`
	ApplyFill         *bool           `xml:"applyFill,attr"`
	ApplyBorder       *bool           `xml:"applyBorder,attr"`
	ApplyAlignment    *bool           `xml:"applyAlignment,attr"`
	ApplyProtection   *bool           `xml:"applyProtection,attr"`
	Alignment         *xlsxAlignment  `xml:"alignment"`
	Protection        *xlsxProtection `xml:"protection"`
}

type xlsxCellXfs struct {
	Count int      `xml:"count,attr"`
	Xf    []xlsxXf `xml:"xf,omitempty"`
}

type xlsxDxfs struct {
	Count int        `xml:"count,attr"`
	Dxfs  []*xlsxDxf `xml:"dxf"`
}

type xlsxDxf struct {
	Font       *xlsxFont           `xml:"font"`
	NumFmt     *xlsxNumFmt         `xml:"numFmt"`
	Fill       *xlsxFill           `xml:"fill"`
	Alignment  *xlsxAlignment      `xml:"alignment"`
	Border     *xlsxBorder         `xml:"border"`
	Protection *xlsxProtection     `xml:"protection"`
	ExtLst     *xlsxPositiveSize2D `xml:"extLst"`
}

type xlsxTableStyles struct {
	Count             int               `xml:"count,attr"`
	DefaultPivotStyle string            `xml:"defaultPivotStyle,attr"`
	DefaultTableStyle string            `xml:"defaultTableStyle,attr"`
	TableStyles       []*xlsxTableStyle `xml:"tableStyle"`
}

type xlsxTableStyle struct {
	Name              string `xml:"name,attr,omitempty"`
	Pivot             int    `xml:"pivot,attr"`
	Count             int    `xml:"count,attr,omitempty"`
	Table             bool   `xml:"table,attr,omitempty"`
	TableStyleElement string `xml:",innerxml"`
}

type xlsxNumFmts struct {
	Count  int           `xml:"count,attr"`
	NumFmt []*xlsxNumFmt `xml:"numFmt"`
}

type xlsxNumFmt struct {
	NumFmtID     int    `xml:"numFmtId,attr"`
	FormatCode   string `xml:"formatCode,attr"`
	FormatCode16 string `xml:"http://schemas.microsoft.com/office/spreadsheetml/2015/02/main formatCode16,attr,omitempty"`
}

type xlsxIndexedColors struct {
	RgbColor []xlsxColor `xml:"rgbColor"`
}

type xlsxStyleColors struct {
	IndexedColors *xlsxIndexedColors `xml:"indexedColors"`
	MruColors     *xlsxInnerXML      `xml:"mruColors"`
}

type Alignment struct {
	Horizontal      string
	Indent          int
	JustifyLastLine bool
	ReadingOrder    uint64
	RelativeIndent  int
	ShrinkToFit     bool
	TextRotation    int
	Vertical        string
	WrapText        bool
}

type Border struct {
	Type  string
	Color string
	Style int
}

type Font struct {
	Bold         bool
	Italic       bool
	Underline    string
	Family       string
	Size         float64
	Strike       bool
	Color        string
	ColorIndexed int
	ColorTheme   *int
	ColorTint    float64
	VertAlign    string
	Charset      *int
}

type Fill struct {
	Type         string
	Pattern      int
	Color        []string
	Shading      int
	Transparency int
}

type Protection struct {
	Hidden bool
	Locked bool
}

type Style struct {
	Border        []Border
	Fill          Fill
	Font          *Font
	Alignment     *Alignment
	Protection    *Protection
	NumFmt        int
	DecimalPlaces *int
	CustomNumFmt  *string
	NegRed        bool
}
