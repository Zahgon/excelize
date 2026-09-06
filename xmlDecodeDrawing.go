package excelize

import "encoding/xml"

type decodeCellAnchor struct {
	EditAs           string                  `xml:"editAs,attr,omitempty"`
	From             *decodeFrom             `xml:"from"`
	To               *decodeTo               `xml:"to"`
	Ext              *decodePositiveSize2D   `xml:"ext"`
	Sp               *decodeSp               `xml:"sp"`
	Pic              *decodePic              `xml:"pic"`
	ClientData       *decodeClientData       `xml:"clientData"`
	AlternateContent []*xlsxAlternateContent `xml:"AlternateContent"`
	Content          string                  `xml:",innerxml"`
}

type decodeCellAnchorPos struct {
	EditAs           string                  `xml:"editAs,attr,omitempty"`
	From             *xlsxFrom               `xml:"from"`
	To               *xlsxTo                 `xml:"to"`
	Pos              *xlsxInnerXML           `xml:"pos"`
	Ext              *xlsxPositiveSize2D     `xml:"ext"`
	Sp               *xlsxSp                 `xml:"sp"`
	GrpSp            *xlsxInnerXML           `xml:"grpSp"`
	GraphicFrame     *xlsxInnerXML           `xml:"graphicFrame"`
	CxnSp            *xlsxInnerXML           `xml:"cxnSp"`
	Pic              *xlsxInnerXML           `xml:"pic"`
	ContentPart      *xlsxInnerXML           `xml:"contentPart"`
	AlternateContent []*xlsxAlternateContent `xml:"AlternateContent"`
	ClientData       *xlsxInnerXML           `xml:"clientData"`
}

type decodeChoice struct {
	XMLName      xml.Name           `xml:"Choice"`
	XMLNSA14     string             `xml:"a14,attr"`
	XMLNSSle15   string             `xml:"sle15,attr"`
	Requires     string             `xml:"Requires,attr"`
	GraphicFrame decodeGraphicFrame `xml:"graphicFrame"`
}

type decodeGraphicFrame struct {
	Macro            string                 `xml:"macro,attr"`
	NvGraphicFramePr decodeNvGraphicFramePr `xml:"nvGraphicFramePr"`
}

type decodeNvGraphicFramePr struct {
	CNvPr decodeCNvPr `xml:"cNvPr"`
}

type decodeSp struct {
	Macro      string        `xml:"macro,attr,omitempty"`
	TextLink   string        `xml:"textlink,attr,omitempty"`
	FLocksText bool          `xml:"fLocksText,attr,omitempty"`
	FPublished *bool         `xml:"fPublished,attr"`
	NvSpPr     *decodeNvSpPr `xml:"nvSpPr"`
	SpPr       *decodeSpPr   `xml:"spPr"`
}

type decodeNvSpPr struct {
	CNvPr   *decodeCNvPr          `xml:"cNvPr"`
	ExtLst  *decodePositiveSize2D `xml:"extLst"`
	CNvSpPr *decodeCNvSpPr        `xml:"cNvSpPr"`
}

type decodeCNvSpPr struct {
	TxBox bool `xml:"txBox,attr"`
}

type decodeWsDr struct {
	XMLName          xml.Name            `xml:"http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing wsDr"`
	A                string              `xml:"xmlns a,attr"`
	Xdr              string              `xml:"xmlns xdr,attr"`
	R                string              `xml:"xmlns r,attr"`
	AlternateContent []*xlsxInnerXML     `xml:"http://schemas.openxmlformats.org/markup-compatibility/2006 AlternateContent"`
	OneCellAnchor    []*decodeCellAnchor `xml:"oneCellAnchor,omitempty"`
	TwoCellAnchor    []*decodeCellAnchor `xml:"twoCellAnchor,omitempty"`
}

type decodeCNvPr struct {
	XMLName    xml.Name          `xml:"cNvPr"`
	ID         int               `xml:"id,attr"`
	Name       string            `xml:"name,attr"`
	Descr      string            `xml:"descr,attr"`
	Title      string            `xml:"title,attr,omitempty"`
	HlinkClick *decodeHlinkClick `xml:"hlinkClick,omitempty"`
}

type decodeHlinkClick struct {
	RID            string `xml:"id,attr,omitempty"`
	InvalidURL     string `xml:"invalidUrl,attr,omitempty"`
	Action         string `xml:"action,attr,omitempty"`
	TgtFrame       string `xml:"tgtFrame,attr,omitempty"`
	Tooltip        string `xml:"tooltip,attr,omitempty"`
	History        bool   `xml:"history,attr,omitempty"`
	HighlightClick bool   `xml:"highlightClick,attr,omitempty"`
	EndSnd         bool   `xml:"endSnd,attr,omitempty"`
}

type decodePicLocks struct {
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

type decodeBlip struct {
	Embed  string `xml:"embed,attr"`
	Cstate string `xml:"cstate,attr,omitempty"`
	R      string `xml:"r,attr"`
}

type decodeStretch struct {
	FillRect string `xml:"fillRect"`
}

type decodeOff struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}

type decodePositiveSize2D struct {
	Cx int `xml:"cx,attr"`
	Cy int `xml:"cy,attr"`
}

type decodePrstGeom struct {
	Prst string `xml:"prst,attr"`
}

type decodeXfrm struct {
	Off decodeOff            `xml:"off"`
	Ext decodePositiveSize2D `xml:"ext"`
}

type decodeCNvPicPr struct {
	PicLocks decodePicLocks `xml:"picLocks"`
}

type decodeNvPicPr struct {
	CNvPr    decodeCNvPr    `xml:"cNvPr"`
	CNvPicPr decodeCNvPicPr `xml:"cNvPicPr"`
}

type decodeBlipFill struct {
	Blip    decodeBlip    `xml:"blip"`
	Stretch decodeStretch `xml:"stretch"`
}

type decodeSpPr struct {
	Xfrm     decodeXfrm     `xml:"xfrm"`
	PrstGeom decodePrstGeom `xml:"prstGeom"`
}

type decodePic struct {
	NvPicPr  decodeNvPicPr  `xml:"nvPicPr"`
	BlipFill decodeBlipFill `xml:"blipFill"`
	SpPr     decodeSpPr     `xml:"spPr"`
}

type decodeFrom struct {
	Col    int `xml:"col"`
	ColOff int `xml:"colOff"`
	Row    int `xml:"row"`
	RowOff int `xml:"rowOff"`
}

type decodeTo struct {
	Col    int `xml:"col"`
	ColOff int `xml:"colOff"`
	Row    int `xml:"row"`
	RowOff int `xml:"rowOff"`
}

type decodeClientData struct {
	FLocksWithSheet  bool `xml:"fLocksWithSheet,attr"`
	FPrintsWithSheet bool `xml:"fPrintsWithSheet,attr"`
}

type decodeCellImages struct {
	XMLName   xml.Name          `xml:"http://www.wps.cn/officeDocument/2017/etCustomData cellImages"`
	CellImage []decodeCellImage `xml:"cellImage"`
}

type decodeCellImage struct {
	Pic decodePic `xml:"pic"`
}
