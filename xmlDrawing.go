package excelize

import (
	"encoding/xml"
	"sync"
)

type xlsxCNvPr struct {
	ID         int             `xml:"id,attr"`
	Name       string          `xml:"name,attr"`
	Descr      string          `xml:"descr,attr"`
	Title      string          `xml:"title,attr,omitempty"`
	HlinkClick *xlsxHlinkClick `xml:"a:hlinkClick"`
}

type xlsxHlinkClick struct {
	R              string `xml:"xmlns:r,attr,omitempty"`
	RID            string `xml:"r:id,attr,omitempty"`
	InvalidURL     string `xml:"invalidUrl,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	TgtFrame       string `xml:"tgtFrame,attr,omitempty"`
	Tooltip        string `xml:"tooltip,attr,omitempty"`
	History        bool   `xml:"history,attr,omitempty"`
	HighlightClick bool   `xml:"highlightClick,attr,omitempty"`
	EndSnd         bool   `xml:"endSnd,attr,omitempty"`
}

type xlsxPicLocks struct {
	NoAdjustHandles    bool `xml:"noAdjustHandles,attr,omitempty"`
	NoChangeArrowheads bool `xml:"noChangeArrowheads,attr,omitempty"`
	NoChangeAspect     bool `xml:"noChangeAspect,attr"`
	NoChangeShapeType  bool `xml:"noChangeShapeType,attr,omitempty"`
	NoCrop             bool `xml:"noCrop,attr,omitempty"`
	NoEditPoints       bool `xml:"noEditPoints,attr,omitempty"`
	NoGrp              bool `xml:"noGrp,attr,omitempty"`
	NoMove             bool `xml:"noMove,attr,omitempty"`
	NoResize           bool `xml:"noResize,attr,omitempty"`
	NoRot              bool `xml:"noRot,attr,omitempty"`
	NoSelect           bool `xml:"noSelect,attr,omitempty"`
}

type xlsxBlip struct {
	Embed   string                        `xml:"r:embed,attr"`
	Cstate  string                        `xml:"cstate,attr,omitempty"`
	R       string                        `xml:"xmlns:r,attr"`
	ExtList *xlsxEGOfficeArtExtensionList `xml:"a:extLst"`
}

type xlsxStretch struct {
	FillRect string `xml:"a:fillRect"`
}

type xlsxOff struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

type xlsxPositiveSize2D struct {
	Cx int `xml:"cx,attr"`
	Cy int `xml:"cy,attr"`
}

type xlsxPrstGeom struct {
	Prst string `xml:"prst,attr"`
}

type xlsxXfrm struct {
	Off xlsxOff            `xml:"a:off"`
	Ext xlsxPositiveSize2D `xml:"a:ext"`
}

type xlsxCNvPicPr struct {
	PicLocks xlsxPicLocks `xml:"a:picLocks"`
}

type xlsxNvPicPr struct {
	CNvPr    xlsxCNvPr    `xml:"xdr:cNvPr"`
	CNvPicPr xlsxCNvPicPr `xml:"xdr:cNvPicPr"`
}

type xlsxCTSVGBlip struct {
	XMLNSaAVG string `xml:"xmlns:asvg,attr"`
	Embed     string `xml:"r:embed,attr"`
	Link      string `xml:"r:link,attr,omitempty"`
}

type xlsxCTOfficeArtExtension struct {
	XMLName xml.Name      `xml:"a:ext"`
	URI     string        `xml:"uri,attr"`
	SVGBlip xlsxCTSVGBlip `xml:"asvg:svgBlip"`
}

type xlsxEGOfficeArtExtensionList struct {
	Ext []xlsxCTOfficeArtExtension `xml:"ext"`
}

type xlsxBlipFill struct {
	Blip    xlsxBlip    `xml:"a:blip"`
	Stretch xlsxStretch `xml:"a:stretch"`
}

type xlsxSpPr struct {
	Xfrm      xlsxXfrm     `xml:"a:xfrm"`
	PrstGeom  xlsxPrstGeom `xml:"a:prstGeom"`
	SolidFill *aSolidFill  `xml:"a:solidFill"`
	Ln        *aLn         `xml:"a:ln"`
}

type xlsxPic struct {
	NvPicPr  xlsxNvPicPr  `xml:"xdr:nvPicPr"`
	BlipFill xlsxBlipFill `xml:"xdr:blipFill"`
	SpPr     xlsxSpPr     `xml:"xdr:spPr"`
}

type xlsxFrom struct {
	Col    int `xml:"xdr:col"`
	ColOff int `xml:"xdr:colOff"`
	Row    int `xml:"xdr:row"`
	RowOff int `xml:"xdr:rowOff"`
}

type xlsxTo struct {
	Col    int `xml:"xdr:col"`
	ColOff int `xml:"xdr:colOff"`
	Row    int `xml:"xdr:row"`
	RowOff int `xml:"xdr:rowOff"`
}

type xdrClientData struct {
	FLocksWithSheet  bool `xml:"fLocksWithSheet,attr"`
	FPrintsWithSheet bool `xml:"fPrintsWithSheet,attr"`
}

type xdrCellAnchor struct {
	EditAs           string                  `xml:"editAs,attr,omitempty"`
	Pos              *xlsxPoint2D            `xml:"xdr:pos"`
	From             *xlsxFrom               `xml:"xdr:from"`
	To               *xlsxTo                 `xml:"xdr:to"`
	Ext              *xlsxPositiveSize2D     `xml:"xdr:ext"`
	Sp               *xdrSp                  `xml:"xdr:sp"`
	Pic              *xlsxPic                `xml:"xdr:pic,omitempty"`
	GraphicFrame     string                  `xml:",innerxml"`
	AlternateContent []*xlsxAlternateContent `xml:"mc:AlternateContent"`
	ClientData       *xdrClientData          `xml:"xdr:clientData"`
}

type xlsxCellAnchorPos struct {
	EditAs           string                  `xml:"editAs,attr,omitempty"`
	From             *xlsxFrom               `xml:"xdr:from"`
	To               *xlsxTo                 `xml:"xdr:to"`
	Pos              *xlsxInnerXML           `xml:"xdr:pos"`
	Ext              *xlsxPositiveSize2D     `xml:"xdr:ext"`
	Sp               *xlsxSp                 `xml:"xdr:sp"`
	GrpSp            *xlsxInnerXML           `xml:"xdr:grpSp"`
	GraphicFrame     *xlsxInnerXML           `xml:"xdr:graphicFrame"`
	CxnSp            *xlsxInnerXML           `xml:"xdr:cxnSp"`
	Pic              *xlsxInnerXML           `xml:"xdr:pic"`
	ContentPart      *xlsxInnerXML           `xml:"xdr:contentPart"`
	AlternateContent []*xlsxAlternateContent `xml:"mc:AlternateContent"`
	ClientData       *xlsxInnerXML           `xml:"xdr:clientData"`
}

type xlsxSp struct {
	Macro      string `xml:"macro,attr,omitempty"`
	TextLink   string `xml:"textlink,attr,omitempty"`
	FLocksText bool   `xml:"fLocksText,attr,omitempty"`
	FPublished *bool  `xml:"fPublished,attr"`
	Content    string `xml:",innerxml"`
}

type xlsxPoint2D struct {
	XMLName xml.Name `xml:"xdr:pos"`
	X       int      `xml:"x,attr"`
	Y       int      `xml:"y,attr"`
}

type xlsxWsDr struct {
	mu               sync.Mutex
	XMLName          xml.Name                `xml:"xdr:wsDr"`
	NS               string                  `xml:"xmlns,attr,omitempty"`
	A                string                  `xml:"xmlns:a,attr,omitempty"`
	Xdr              string                  `xml:"xmlns:xdr,attr,omitempty"`
	R                string                  `xml:"xmlns:r,attr,omitempty"`
	AlternateContent []*xlsxAlternateContent `xml:"mc:AlternateContent"`
	AbsoluteAnchor   []*xdrCellAnchor        `xml:"xdr:absoluteAnchor"`
	OneCellAnchor    []*xdrCellAnchor        `xml:"xdr:oneCellAnchor"`
	TwoCellAnchor    []*xdrCellAnchor        `xml:"xdr:twoCellAnchor"`
}

type xlsxGraphicFrame struct {
	XMLName          xml.Name             `xml:"xdr:graphicFrame"`
	Macro            string               `xml:"macro,attr"`
	NvGraphicFramePr xlsxNvGraphicFramePr `xml:"xdr:nvGraphicFramePr"`
	Xfrm             xlsxXfrm             `xml:"xdr:xfrm"`
	Graphic          *xlsxGraphic         `xml:"a:graphic"`
}

type xlsxNvGraphicFramePr struct {
	CNvPr                *xlsxCNvPr `xml:"xdr:cNvPr"`
	ChicNvGraphicFramePr string     `xml:"xdr:cNvGraphicFramePr"`
}

type xlsxGraphic struct {
	GraphicData *xlsxGraphicData `xml:"a:graphicData"`
}

type xlsxGraphicData struct {
	URI   string     `xml:"uri,attr"`
	Chart *xlsxChart `xml:"c:chart,omitempty"`
	Sle   *xlsxSle   `xml:"sle:slicer"`
}

type xlsxSle struct {
	XMLNS string `xml:"xmlns:sle,attr"`
	Name  string `xml:"name,attr"`
}

type xlsxChart struct {
	C   string `xml:"xmlns:c,attr"`
	RID string `xml:"r:id,attr"`
	R   string `xml:"xmlns:r,attr"`
}

type xdrSp struct {
	XMLName  xml.Name   `xml:"xdr:sp"`
	Macro    string     `xml:"macro,attr"`
	Textlink string     `xml:"textlink,attr"`
	NvSpPr   *xdrNvSpPr `xml:"xdr:nvSpPr"`
	SpPr     *xlsxSpPr  `xml:"xdr:spPr"`
	Style    *xdrStyle  `xml:"xdr:style"`
	TxBody   *xdrTxBody `xml:"xdr:txBody"`
}

type xdrNvSpPr struct {
	CNvPr   *xlsxCNvPr  `xml:"xdr:cNvPr"`
	CNvSpPr *xdrCNvSpPr `xml:"xdr:cNvSpPr"`
}

type xdrCNvSpPr struct {
	TxBox bool `xml:"txBox,attr"`
}

type xdrStyle struct {
	LnRef     *aRef     `xml:"a:lnRef"`
	FillRef   *aRef     `xml:"a:fillRef"`
	EffectRef *aRef     `xml:"a:effectRef"`
	FontRef   *aFontRef `xml:"a:fontRef"`
}

type aRef struct {
	Idx       int            `xml:"idx,attr"`
	ScrgbClr  *aScrgbClr     `xml:"a:scrgbClr"`
	SchemeClr *attrValString `xml:"a:schemeClr"`
	SrgbClr   *attrValString `xml:"a:srgbClr"`
}

type aScrgbClr struct {
	R float64 `xml:"r,attr"`
	G float64 `xml:"g,attr"`
	B float64 `xml:"b,attr"`
}

type aFontRef struct {
	Idx       string         `xml:"idx,attr"`
	SchemeClr *attrValString `xml:"a:schemeClr"`
}

type xdrTxBody struct {
	BodyPr *aBodyPr `xml:"a:bodyPr"`
	P      []*aP    `xml:"a:p"`
}

type Picture struct {
	Extension  string
	File       []byte
	Format     *GraphicOptions
	InsertType PictureInsertType
}

type GraphicOptions struct {
	AltText             string
	Name                string
	PrintObject         *bool
	Locked              *bool
	LockAspectRatio     bool
	AutoFit             bool
	AutoFitIgnoreAspect bool
	OffsetX             int
	OffsetY             int
	ScaleX              float64
	ScaleY              float64
	Hyperlink           string
	HyperlinkType       string
	Positioning         string
}

type Shape struct {
	Cell      string
	Type      string
	Macro     string
	Width     uint
	Height    uint
	Format    GraphicOptions
	Fill      Fill
	Line      LineOptions
	Paragraph []RichTextRun
}
