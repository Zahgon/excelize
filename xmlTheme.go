package excelize

import "encoding/xml"

type xlsxTheme struct {
	XMLName           xml.Name              `xml:"a:theme"`
	XMLNSa            string                `xml:"xmlns:a,attr"`
	XMLNSr            string                `xml:"xmlns:r,attr"`
	Name              string                `xml:"name,attr"`
	ThemeElements     xlsxBaseStyles        `xml:"a:themeElements"`
	ObjectDefaults    xlsxObjectDefaults    `xml:"a:objectDefaults"`
	ExtraClrSchemeLst xlsxExtraClrSchemeLst `xml:"a:extraClrSchemeLst"`
	CustClrLst        *xlsxInnerXML         `xml:"a:custClrLst"`
	ExtLst            *xlsxExtLst           `xml:"a:extLst"`
}

type xlsxBaseStyles struct {
	ClrScheme  xlsxColorScheme `xml:"a:clrScheme"`
	FontScheme xlsxFontScheme  `xml:"a:fontScheme"`
	FmtScheme  xlsxStyleMatrix `xml:"a:fmtScheme"`
	ExtLst     *xlsxExtLst     `xml:"a:extLst"`
}

type xlsxCTColor struct {
	ScrgbClr  *xlsxInnerXML  `xml:"a:scrgbClr"`
	SrgbClr   *attrValString `xml:"a:srgbClr"`
	HslClr    *xlsxInnerXML  `xml:"a:hslClr"`
	SysClr    *xlsxSysClr    `xml:"a:sysClr"`
	SchemeClr *xlsxInnerXML  `xml:"a:schemeClr"`
	PrstClr   *xlsxInnerXML  `xml:"a:prstClr"`
}

type xlsxColorScheme struct {
	Name     string      `xml:"name,attr"`
	Dk1      xlsxCTColor `xml:"a:dk1"`
	Lt1      xlsxCTColor `xml:"a:lt1"`
	Dk2      xlsxCTColor `xml:"a:dk2"`
	Lt2      xlsxCTColor `xml:"a:lt2"`
	Accent1  xlsxCTColor `xml:"a:accent1"`
	Accent2  xlsxCTColor `xml:"a:accent2"`
	Accent3  xlsxCTColor `xml:"a:accent3"`
	Accent4  xlsxCTColor `xml:"a:accent4"`
	Accent5  xlsxCTColor `xml:"a:accent5"`
	Accent6  xlsxCTColor `xml:"a:accent6"`
	Hlink    xlsxCTColor `xml:"a:hlink"`
	FolHlink xlsxCTColor `xml:"a:folHlink"`
	ExtLst   *xlsxExtLst `xml:"a:extLst"`
}

type xlsxObjectDefaults struct {
	ObjectDefaults string `xml:",innerxml"`
}

type xlsxExtraClrSchemeLst struct {
	ExtraClrSchemeLst string `xml:",innerxml"`
}

type xlsxCTSupplementalFont struct {
	Script   string `xml:"script,attr"`
	Typeface string `xml:"typeface,attr"`
}

type xlsxFontCollection struct {
	Latin  *xlsxCTTextFont          `xml:"a:latin"`
	Ea     *xlsxCTTextFont          `xml:"a:ea"`
	Cs     *xlsxCTTextFont          `xml:"a:cs"`
	Font   []xlsxCTSupplementalFont `xml:"a:font"`
	ExtLst *xlsxExtLst              `xml:"a:extLst"`
}

type xlsxFontScheme struct {
	Name      string             `xml:"name,attr"`
	MajorFont xlsxFontCollection `xml:"a:majorFont"`
	MinorFont xlsxFontCollection `xml:"a:minorFont"`
	ExtLst    *xlsxExtLst        `xml:"a:extLst"`
}

type xlsxStyleMatrix struct {
	Name           string             `xml:"name,attr,omitempty"`
	FillStyleLst   xlsxFillStyleLst   `xml:"a:fillStyleLst"`
	LnStyleLst     xlsxLnStyleLst     `xml:"a:lnStyleLst"`
	EffectStyleLst xlsxEffectStyleLst `xml:"a:effectStyleLst"`
	BgFillStyleLst xlsxBgFillStyleLst `xml:"a:bgFillStyleLst"`
}

type xlsxFillStyleLst struct {
	FillStyleLst string `xml:",innerxml"`
}

type xlsxLnStyleLst struct {
	LnStyleLst string `xml:",innerxml"`
}

type xlsxEffectStyleLst struct {
	EffectStyleLst string `xml:",innerxml"`
}

type xlsxBgFillStyleLst struct {
	BgFillStyleLst string `xml:",innerxml"`
}

type xlsxSysClr struct {
	Val     string `xml:"val,attr"`
	LastClr string `xml:"lastClr,attr"`
}

type decodeTheme struct {
	XMLName           xml.Name              `xml:"http://schemas.openxmlformats.org/drawingml/2006/main theme"`
	Name              string                `xml:"name,attr"`
	ThemeElements     decodeBaseStyles      `xml:"themeElements"`
	ObjectDefaults    xlsxObjectDefaults    `xml:"objectDefaults"`
	ExtraClrSchemeLst xlsxExtraClrSchemeLst `xml:"extraClrSchemeLst"`
	CustClrLst        *xlsxInnerXML         `xml:"custClrLst"`
	ExtLst            *xlsxExtLst           `xml:"extLst"`
}

type decodeBaseStyles struct {
	ClrScheme  decodeColorScheme `xml:"clrScheme"`
	FontScheme decodeFontScheme  `xml:"fontScheme"`
	FmtScheme  decodeStyleMatrix `xml:"fmtScheme"`
	ExtLst     *xlsxExtLst       `xml:"extLst"`
}

type decodeColorScheme struct {
	Name     string        `xml:"name,attr"`
	Dk1      decodeCTColor `xml:"dk1"`
	Lt1      decodeCTColor `xml:"lt1"`
	Dk2      decodeCTColor `xml:"dk2"`
	Lt2      decodeCTColor `xml:"lt2"`
	Accent1  decodeCTColor `xml:"accent1"`
	Accent2  decodeCTColor `xml:"accent2"`
	Accent3  decodeCTColor `xml:"accent3"`
	Accent4  decodeCTColor `xml:"accent4"`
	Accent5  decodeCTColor `xml:"accent5"`
	Accent6  decodeCTColor `xml:"accent6"`
	Hlink    decodeCTColor `xml:"hlink"`
	FolHlink decodeCTColor `xml:"folHlink"`
	ExtLst   *xlsxExtLst   `xml:"extLst"`
}

type decodeFontScheme struct {
	Name      string               `xml:"name,attr"`
	MajorFont decodeFontCollection `xml:"majorFont"`
	MinorFont decodeFontCollection `xml:"minorFont"`
	ExtLst    *xlsxExtLst          `xml:"extLst"`
}

type decodeFontCollection struct {
	Latin  *xlsxCTTextFont          `xml:"latin"`
	Ea     *xlsxCTTextFont          `xml:"ea"`
	Cs     *xlsxCTTextFont          `xml:"cs"`
	Font   []xlsxCTSupplementalFont `xml:"font"`
	ExtLst *xlsxExtLst              `xml:"extLst"`
}

type decodeCTColor struct {
	ScrgbClr  *xlsxInnerXML  `xml:"scrgbClr"`
	SrgbClr   *attrValString `xml:"srgbClr"`
	HslClr    *xlsxInnerXML  `xml:"hslClr"`
	SysClr    *xlsxSysClr    `xml:"sysClr"`
	SchemeClr *xlsxInnerXML  `xml:"schemeClr"`
	PrstClr   *xlsxInnerXML  `xml:"prstClr"`
}

type decodeStyleMatrix struct {
	Name           string             `xml:"name,attr,omitempty"`
	FillStyleLst   xlsxFillStyleLst   `xml:"fillStyleLst"`
	LnStyleLst     xlsxLnStyleLst     `xml:"lnStyleLst"`
	EffectStyleLst xlsxEffectStyleLst `xml:"effectStyleLst"`
	BgFillStyleLst xlsxBgFillStyleLst `xml:"bgFillStyleLst"`
}
