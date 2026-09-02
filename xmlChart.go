package excelize

import "encoding/xml"

type xlsxChartSpace struct {
	XMLName        xml.Name        `xml:"http://schemas.openxmlformats.org/drawingml/2006/chart chartSpace"`
	XMLNSa         string          `xml:"xmlns:a,attr"`
	Date1904       *attrValBool    `xml:"date1904"`
	Lang           *attrValString  `xml:"lang"`
	RoundedCorners *attrValBool    `xml:"roundedCorners"`
	Chart          cChart          `xml:"chart"`
	SpPr           *cSpPr          `xml:"spPr"`
	TxPr           *cTxPr          `xml:"txPr"`
	PrintSettings  *cPrintSettings `xml:"printSettings"`
}

type cThicknessSpPr struct {
	Thickness *attrValInt `xml:"thickness"`
	SpPr      *cSpPr      `xml:"spPr"`
}

type cChart struct {
	Title            *cTitle            `xml:"title"`
	AutoTitleDeleted *cAutoTitleDeleted `xml:"autoTitleDeleted"`
	View3D           *cView3D           `xml:"view3D"`
	Floor            *cThicknessSpPr    `xml:"floor"`
	SideWall         *cThicknessSpPr    `xml:"sideWall"`
	BackWall         *cThicknessSpPr    `xml:"backWall"`
	PlotArea         *cPlotArea         `xml:"plotArea"`
	Legend           *cLegend           `xml:"legend"`
	PlotVisOnly      *attrValBool       `xml:"plotVisOnly"`
	DispBlanksAs     *attrValString     `xml:"dispBlanksAs"`
	ShowDLblsOverMax *attrValBool       `xml:"showDLblsOverMax"`
}

type cTitle struct {
	Tx      cTx          `xml:"tx,omitempty"`
	Layout  *cLayout     `xml:"layout"`
	Overlay *attrValBool `xml:"overlay"`
	SpPr    cSpPr        `xml:"spPr,omitempty"`
	TxPr    cTxPr        `xml:"txPr,omitempty"`
	ExtLst  *xlsxExtLst  `xml:"extLst"`
}

type cLayout struct {
	ManualLayout *cManualLayout `xml:"manualLayout"`
	ExtLst       *xlsxExtLst    `xml:"extLst"`
}

type cManualLayout struct {
	LayoutTarget *attrValString `xml:"layoutTarget"`
	XMode        *attrValString `xml:"xMode"`
	YMode        *attrValString `xml:"yMode"`
	WMode        *attrValString `xml:"wMode"`
	HMode        *attrValString `xml:"hMode"`
	X            *attrValFloat  `xml:"x"`
	Y            *attrValFloat  `xml:"y"`
	W            *attrValFloat  `xml:"w"`
	H            *attrValFloat  `xml:"h"`
	ExtLst       *xlsxExtLst    `xml:"extLst"`
}

type cTx struct {
	StrRef *cStrRef `xml:"strRef"`
	Rich   *cRich   `xml:"rich,omitempty"`
}

type cRich struct {
	BodyPr   aBodyPr `xml:"a:bodyPr,omitempty"`
	LstStyle string  `xml:"a:lstStyle,omitempty"`
	P        []aP    `xml:"a:p"`
}

type aBodyPr struct {
	Anchor           string  `xml:"anchor,attr,omitempty"`
	AnchorCtr        bool    `xml:"anchorCtr,attr"`
	Rot              int     `xml:"rot,attr"`
	BIns             float64 `xml:"bIns,attr,omitempty"`
	CompatLnSpc      bool    `xml:"compatLnSpc,attr,omitempty"`
	ForceAA          bool    `xml:"forceAA,attr,omitempty"`
	FromWordArt      bool    `xml:"fromWordArt,attr,omitempty"`
	HorzOverflow     string  `xml:"horzOverflow,attr,omitempty"`
	LIns             float64 `xml:"lIns,attr,omitempty"`
	NumCol           int     `xml:"numCol,attr,omitempty"`
	RIns             float64 `xml:"rIns,attr,omitempty"`
	RtlCol           bool    `xml:"rtlCol,attr,omitempty"`
	SpcCol           int     `xml:"spcCol,attr,omitempty"`
	SpcFirstLastPara bool    `xml:"spcFirstLastPara,attr"`
	TIns             float64 `xml:"tIns,attr,omitempty"`
	Upright          bool    `xml:"upright,attr,omitempty"`
	Vert             string  `xml:"vert,attr,omitempty"`
	VertOverflow     string  `xml:"vertOverflow,attr,omitempty"`
	Wrap             string  `xml:"wrap,attr,omitempty"`
}

type aP struct {
	PPr        *aPPr        `xml:"a:pPr"`
	R          *aR          `xml:"a:r"`
	EndParaRPr *aEndParaRPr `xml:"a:endParaRPr"`
}

type aPPr struct {
	DefRPr aRPr `xml:"a:defRPr"`
}

type aSrgbClr struct {
	Val      *string     `xml:"val,attr"`
	Tint     *attrValInt `xml:"a:tint"`
	Shade    *attrValInt `xml:"a:shade"`
	Comp     *attrValInt `xml:"a:comp"`
	Inv      *attrValInt `xml:"a:inv"`
	Gray     *attrValInt `xml:"a:gray"`
	Alpha    *attrValInt `xml:"a:alpha"`
	AlphaOff *attrValInt `xml:"a:alphaOff"`
	AlphaMod *attrValInt `xml:"a:alphaMod"`
	Hue      *attrValInt `xml:"a:hue"`
	HueOff   *attrValInt `xml:"a:hueOff"`
	HueMod   *attrValInt `xml:"a:hueMod"`
	Sat      *attrValInt `xml:"a:sat"`
	SatOff   *attrValInt `xml:"a:satOff"`
	SatMod   *attrValInt `xml:"a:satMod"`
	Lum      *attrValInt `xml:"a:lum"`
	LumOff   *attrValInt `xml:"a:lumOff"`
	LumMod   *attrValInt `xml:"a:lumMod"`
	Red      *attrValInt `xml:"a:red"`
	RedOff   *attrValInt `xml:"a:redOff"`
	RedMod   *attrValInt `xml:"a:redMod"`
	Green    *attrValInt `xml:"a:green"`
	GreenOff *attrValInt `xml:"a:greenOff"`
	GreenMod *attrValInt `xml:"a:greenMod"`
	Blue     *attrValInt `xml:"a:blue"`
	BlueOff  *attrValInt `xml:"a:blueOff"`
	BlueMod  *attrValInt `xml:"a:blueMod"`
	Gamma    *attrValInt `xml:"a:gamma"`
	InvGamma *attrValInt `xml:"a:invGamma"`
}

type aSolidFill struct {
	SchemeClr *aSchemeClr    `xml:"a:schemeClr"`
	SrgbClr   *aSrgbClr      `xml:"a:srgbClr"`
	PrstClr   *attrValString `xml:"a:prstClr"`
}

type aSchemeClr struct {
	Val    string      `xml:"val,attr,omitempty"`
	LumMod *attrValInt `xml:"a:lumMod"`
	LumOff *attrValInt `xml:"a:lumOff"`
}

type attrValInt struct {
	Val *int `xml:"val,attr"`
}

type attrValFloat struct {
	Val *float64 `xml:"val,attr"`
}

type attrValBool struct {
	Val *bool `xml:"val,attr"`
}

type attrValString struct {
	Val *string `xml:"val,attr"`
}

type xlsxCTTextFont struct {
	Typeface    string `xml:"typeface,attr"`
	Panose      string `xml:"panose,attr,omitempty"`
	PitchFamily string `xml:"pitchFamily,attr,omitempty"`
	Charset     string `xml:"Charset,attr,omitempty"`
}

type aR struct {
	RPr aRPr   `xml:"a:rPr,omitempty"`
	T   string `xml:"a:t,omitempty"`
}

type aRPr struct {
	AltLang    string          `xml:"altLang,attr,omitempty"`
	B          bool            `xml:"b,attr"`
	Baseline   int             `xml:"baseline,attr"`
	Bmk        string          `xml:"bmk,attr,omitempty"`
	Cap        string          `xml:"cap,attr,omitempty"`
	Dirty      bool            `xml:"dirty,attr,omitempty"`
	Err        bool            `xml:"err,attr,omitempty"`
	I          bool            `xml:"i,attr"`
	Kern       int             `xml:"kern,attr"`
	Kumimoji   bool            `xml:"kumimoji,attr,omitempty"`
	Lang       string          `xml:"lang,attr,omitempty"`
	NoProof    bool            `xml:"noProof,attr,omitempty"`
	NormalizeH bool            `xml:"normalizeH,attr,omitempty"`
	SmtClean   bool            `xml:"smtClean,attr,omitempty"`
	SmtID      uint64          `xml:"smtId,attr,omitempty"`
	Spc        int             `xml:"spc,attr"`
	Strike     string          `xml:"strike,attr,omitempty"`
	Sz         float64         `xml:"sz,attr,omitempty"`
	U          string          `xml:"u,attr,omitempty"`
	SolidFill  *aSolidFill     `xml:"a:solidFill"`
	Latin      *xlsxCTTextFont `xml:"a:latin"`
	Ea         *xlsxCTTextFont `xml:"a:ea"`
	Cs         *xlsxCTTextFont `xml:"a:cs"`
}

type cDTable struct {
	ShowHorzBorder *attrValBool `xml:"showHorzBorder"`
	ShowVertBorder *attrValBool `xml:"showVertBorder"`
	ShowOutline    *attrValBool `xml:"showOutline"`
	ShowKeys       *attrValBool `xml:"showKeys"`
	SpPr           *cSpPr       `xml:"spPr"`
	TxPr           *cTxPr       `xml:"txPr"`
	ExtLst         *xlsxExtLst  `xml:"extLst"`
}

type cSpPr struct {
	NoFill    *string     `xml:"a:noFill"`
	SolidFill *aSolidFill `xml:"a:solidFill"`
	Ln        *aLn        `xml:"a:ln"`
	Sp3D      *aSp3D      `xml:"a:sp3d"`
	EffectLst *string     `xml:"a:effectLst"`
}

type aSp3D struct {
	ContourW   int          `xml:"contourW,attr"`
	ContourClr *aContourClr `xml:"a:contourClr"`
}

type aContourClr struct {
	SchemeClr *aSchemeClr `xml:"a:schemeClr"`
}

type aLn struct {
	Algn      string         `xml:"algn,attr,omitempty"`
	Cap       string         `xml:"cap,attr,omitempty"`
	Cmpd      string         `xml:"cmpd,attr,omitempty"`
	W         int            `xml:"w,attr,omitempty"`
	NoFill    *attrValString `xml:"a:noFill"`
	Round     string         `xml:"a:round,omitempty"`
	SolidFill *aSolidFill    `xml:"a:solidFill"`
	PrstDash  *attrValString `xml:"a:prstDash"`
	PrstClr   *xlsxInnerXML  `xml:"a:prstClr"`
}

type cTxPr struct {
	BodyPr   aBodyPr `xml:"a:bodyPr,omitempty"`
	LstStyle string  `xml:"a:lstStyle,omitempty"`
	P        aP      `xml:"a:p,omitempty"`
}

type aEndParaRPr struct {
	Lang    string `xml:"lang,attr"`
	AltLang string `xml:"altLang,attr,omitempty"`
	Sz      int    `xml:"sz,attr,omitempty"`
}

type cAutoTitleDeleted struct {
	Val bool `xml:"val,attr"`
}

type cView3D struct {
	RotX         *attrValInt `xml:"rotX"`
	RotY         *attrValInt `xml:"rotY"`
	RAngAx       *attrValInt `xml:"rAngAx"`
	DepthPercent *attrValInt `xml:"depthPercent"`
	Perspective  *attrValInt `xml:"perspective"`
	ExtLst       *xlsxExtLst `xml:"extLst"`
}

type cPlotArea struct {
	Layout         *string    `xml:"layout"`
	AreaChart      []*cCharts `xml:"areaChart"`
	Area3DChart    []*cCharts `xml:"area3DChart"`
	BarChart       []*cCharts `xml:"barChart"`
	Bar3DChart     []*cCharts `xml:"bar3DChart"`
	BubbleChart    []*cCharts `xml:"bubbleChart"`
	DoughnutChart  []*cCharts `xml:"doughnutChart"`
	LineChart      []*cCharts `xml:"lineChart"`
	Line3DChart    []*cCharts `xml:"line3DChart"`
	StockChart     []*cCharts `xml:"stockChart"`
	PieChart       []*cCharts `xml:"pieChart"`
	Pie3DChart     []*cCharts `xml:"pie3DChart"`
	OfPieChart     []*cCharts `xml:"ofPieChart"`
	RadarChart     []*cCharts `xml:"radarChart"`
	ScatterChart   []*cCharts `xml:"scatterChart"`
	Surface3DChart []*cCharts `xml:"surface3DChart"`
	SurfaceChart   []*cCharts `xml:"surfaceChart"`
	CatAx          []*cAxs    `xml:"catAx"`
	ValAx          []*cAxs    `xml:"valAx"`
	DateAx         []*cAxs    `xml:"dateAx"`
	SerAx          []*cAxs    `xml:"serAx"`
	DTable         *cDTable   `xml:"dTable"`
	SpPr           *cSpPr     `xml:"spPr"`
}

type cCharts struct {
	BarDir       *attrValString `xml:"barDir"`
	BubbleScale  *attrValFloat  `xml:"bubbleScale"`
	Grouping     *attrValString `xml:"grouping"`
	RadarStyle   *attrValString `xml:"radarStyle"`
	ScatterStyle *attrValString `xml:"scatterStyle"`
	OfPieType    *attrValString `xml:"ofPieType"`
	VaryColors   *attrValBool   `xml:"varyColors"`
	Wireframe    *attrValBool   `xml:"wireframe"`
	Ser          *[]cSer        `xml:"ser"`
	SplitPos     *attrValInt    `xml:"splitPos"`
	SerLines     *attrValString `xml:"serLines"`
	DLbls        *cDLbls        `xml:"dLbls"`
	DropLines    *cLines        `xml:"dropLines"`
	HiLowLines   *cLines        `xml:"hiLowLines"`
	UpDownBars   *cUpDownBars   `xml:"upDownBars"`
	GapWidth     *attrValInt    `xml:"gapWidth"`
	Shape        *attrValString `xml:"shape"`
	HoleSize     *attrValInt    `xml:"holeSize"`
	Smooth       *attrValBool   `xml:"smooth"`
	Overlap      *attrValInt    `xml:"overlap"`
	AxID         []*attrValInt  `xml:"axId"`
}

type cAxs struct {
	AxID           *attrValInt    `xml:"axId"`
	Scaling        *cScaling      `xml:"scaling"`
	Delete         *attrValBool   `xml:"delete"`
	AxPos          *attrValString `xml:"axPos"`
	MajorGridlines *cLines        `xml:"majorGridlines"`
	MinorGridlines *cLines        `xml:"minorGridlines"`
	Title          *cTitle        `xml:"title"`
	NumFmt         *cNumFmt       `xml:"numFmt"`
	MajorTickMark  *attrValString `xml:"majorTickMark"`
	MinorTickMark  *attrValString `xml:"minorTickMark"`
	TickLblPos     *attrValString `xml:"tickLblPos"`
	SpPr           *cSpPr         `xml:"spPr"`
	TxPr           *cTxPr         `xml:"txPr"`
	CrossAx        *attrValInt    `xml:"crossAx"`
	Crosses        *attrValString `xml:"crosses"`
	CrossBetween   *attrValString `xml:"crossBetween"`
	MajorUnit      *attrValFloat  `xml:"majorUnit"`
	MinorUnit      *attrValFloat  `xml:"minorUnit"`
	Auto           *attrValBool   `xml:"auto"`
	LblAlgn        *attrValString `xml:"lblAlgn"`
	LblOffset      *attrValInt    `xml:"lblOffset"`
	TickLblSkip    *attrValInt    `xml:"tickLblSkip"`
	TickMarkSkip   *attrValInt    `xml:"tickMarkSkip"`
	NoMultiLvlLbl  *attrValBool   `xml:"noMultiLvlLbl"`
}

type cUpDownBars struct {
	GapWidth *attrValString `xml:"gapWidth"`
	UpBars   *cLines        `xml:"upBars"`
	DownBars *cLines        `xml:"downBars"`
	ExtLst   *xlsxExtLst    `xml:"extLst"`
}

type cLines struct {
	SpPr *cSpPr `xml:"spPr"`
}

type cScaling struct {
	LogBase     *attrValFloat  `xml:"logBase"`
	Orientation *attrValString `xml:"orientation"`
	Max         *attrValFloat  `xml:"max"`
	Min         *attrValFloat  `xml:"min"`
}

type cNumFmt struct {
	FormatCode   string `xml:"formatCode,attr"`
	SourceLinked bool   `xml:"sourceLinked,attr"`
}

type cSer struct {
	IDx              *attrValInt  `xml:"idx"`
	Order            *attrValInt  `xml:"order"`
	Tx               *cTx         `xml:"tx"`
	SpPr             *cSpPr       `xml:"spPr"`
	InvertIfNegative *attrValBool `xml:"invertIfNegative"`
	Marker           *cMarker     `xml:"marker"`
	DPt              []*cDPt      `xml:"dPt"`
	DLbls            *cDLbls      `xml:"dLbls"`
	Cat              *cCat        `xml:"cat"`
	Val              *cVal        `xml:"val"`
	XVal             *cCat        `xml:"xVal"`
	YVal             *cVal        `xml:"yVal"`
	Smooth           *attrValBool `xml:"smooth"`
	BubbleSize       *cVal        `xml:"bubbleSize"`
	Bubble3D         *attrValBool `xml:"bubble3D"`
}

type cMarker struct {
	Symbol *attrValString `xml:"symbol"`
	Size   *attrValInt    `xml:"size"`
	SpPr   *cSpPr         `xml:"spPr"`
}

type cDPt struct {
	IDx      *attrValInt  `xml:"idx"`
	Bubble3D *attrValBool `xml:"bubble3D"`
	SpPr     *cSpPr       `xml:"spPr"`
}

type cCat struct {
	StrRef *cStrRef `xml:"strRef"`
}

type cStrRef struct {
	F        string     `xml:"f"`
	StrCache *cStrCache `xml:"strCache"`
}

type cStrCache struct {
	Pt      []*cPt      `xml:"pt"`
	PtCount *attrValInt `xml:"ptCount"`
}

type cPt struct {
	IDx int     `xml:"idx,attr"`
	V   *string `xml:"v"`
}

type cVal struct {
	NumRef *cNumRef `xml:"numRef"`
}

type cNumRef struct {
	F        string     `xml:"f"`
	NumCache *cNumCache `xml:"numCache"`
}

type cNumCache struct {
	FormatCode string      `xml:"formatCode"`
	Pt         []*cPt      `xml:"pt"`
	PtCount    *attrValInt `xml:"ptCount"`
}

type cDLbls struct {
	NumFmt          *cNumFmt       `xml:"numFmt"`
	SpPr            *cSpPr         `xml:"spPr"`
	TxPr            *cTxPr         `xml:"txPr"`
	DLblPos         *attrValString `xml:"dLblPos"`
	ShowLegendKey   *attrValBool   `xml:"showLegendKey"`
	ShowVal         *attrValBool   `xml:"showVal"`
	ShowCatName     *attrValBool   `xml:"showCatName"`
	ShowSerName     *attrValBool   `xml:"showSerName"`
	ShowPercent     *attrValBool   `xml:"showPercent"`
	ShowBubbleSize  *attrValBool   `xml:"showBubbleSize"`
	ShowLeaderLines *attrValBool   `xml:"showLeaderLines"`
}

type cLegendEntry struct {
	IDx  *attrValInt `xml:"idx"`
	TxPr *cTxPr      `xml:"txPr"`
}

type cLegend struct {
	Layout      *string        `xml:"layout"`
	LegendPos   *attrValString `xml:"legendPos"`
	LegendEntry []cLegendEntry `xml:"legendEntry"`
	Overlay     *attrValBool   `xml:"overlay"`
	SpPr        *cSpPr         `xml:"spPr"`
	TxPr        *cTxPr         `xml:"txPr"`
}

type cPrintSettings struct {
	HeaderFooter *string       `xml:"headerFooter"`
	PageMargins  *cPageMargins `xml:"pageMargins"`
	PageSetup    *string       `xml:"pageSetup"`
}

type cPageMargins struct {
	B      float64 `xml:"b,attr"`
	Footer float64 `xml:"footer,attr"`
	Header float64 `xml:"header,attr"`
	L      float64 `xml:"l,attr"`
	R      float64 `xml:"r,attr"`
	T      float64 `xml:"t,attr"`
}

type LineOptions struct {
	Type   LineType
	Dash   LineDashType
	Fill   Fill
	Smooth bool
	Width  float64
}

type ChartNumFmt struct {
	CustomNumFmt string
	SourceLinked bool
}

type ChartAxis struct {
	None              bool
	DropLines         bool
	HighLowLines      bool
	MajorGridLines    bool
	MinorGridLines    bool
	MajorUnit         float64
	TickLabelPosition ChartTickLabelPositionType
	TickLabelSkip     int
	ReverseOrder      bool
	Secondary         bool
	Maximum           *float64
	Minimum           *float64
	Alignment         Alignment
	Font              Font
	LogBase           float64
	NumFmt            ChartNumFmt
	Title             ChartTitle
	axID              int
}

type ChartDimension struct {
	Width  uint
	Height uint
}

type ChartUpDownBar struct {
	Fill   Fill
	Border LineOptions
}

type ChartPlotArea struct {
	SecondPlotValues  int
	ShowBubbleSize    bool
	ShowCatName       bool
	ShowDataTable     bool
	ShowDataTableKeys bool
	ShowLeaderLines   bool
	ShowPercent       bool
	ShowSerName       bool
	ShowVal           bool
	Fill              Fill
	UpBars            ChartUpDownBar
	DownBars          ChartUpDownBar
	NumFmt            ChartNumFmt
}

type Chart struct {
	Type         ChartType
	Series       []ChartSeries
	Format       GraphicOptions
	Dimension    ChartDimension
	Legend       ChartLegend
	Title        ChartTitle
	VaryColors   *bool
	XAxis        ChartAxis
	YAxis        ChartAxis
	PlotArea     ChartPlotArea
	Fill         Fill
	Border       LineOptions
	ShowBlanksAs string
	BubbleSize   int
	HoleSize     int
	GapWidth     *uint
	Overlap      *int
	order        int
}

type ChartTitle struct {
	Fill      Fill
	Border    LineOptions
	Paragraph []RichTextRun
	Font      *Font
	Formula   string
	OffsetX   int
	OffsetY   int
	Width     int
	Height    int
	Overlay   bool
}

type ChartLegend struct {
	Position      string
	ShowLegendKey bool
	Font          *Font
}

type ChartMarker struct {
	Border LineOptions
	Fill   Fill
	Symbol string
	Size   int
}

type ChartDataLabel struct {
	Alignment Alignment
	Font      Font
	Fill      Fill
}

type ChartDataPoint struct {
	Index int
	Fill  Fill
}

type ChartSeries struct {
	Name              string
	Categories        string
	Values            string
	Sizes             string
	Fill              Fill
	Legend            ChartLegend
	Line              LineOptions
	Marker            ChartMarker
	DataLabel         ChartDataLabel
	DataLabelPosition ChartDataLabelPositionType
	DataPoint         []ChartDataPoint
}
