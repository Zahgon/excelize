package excelize

import "encoding/xml"

type xlsxCalcChain struct {
	XMLName xml.Name         `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main calcChain"`
	C       []xlsxCalcChainC `xml:"c"`
}

type xlsxCalcChainC struct {
	R string `xml:"r,attr"`
	I int    `xml:"i,attr,omitempty"`
	L bool   `xml:"l,attr,omitempty"`
	S bool   `xml:"s,attr,omitempty"`
	T bool   `xml:"t,attr,omitempty"`
	A bool   `xml:"a,attr,omitempty"`
}

type xlsxVolTypes struct {
	XMLName xml.Name      `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main volTypes"`
	VolType []xlsxVolType `xml:"volType"`
	ExtLst  *xlsxExtLst   `xml:"extLst"`
}

type xlsxVolType struct {
	Type string        `xml:"type,attr"`
	Main []xlsxVolMain `xml:"main"`
}

type xlsxVolMain struct {
	First string         `xml:"first,attr"`
	Tp    []xlsxVolTopic `xml:"tp"`
}

type xlsxVolTopic struct {
	T   string            `xml:"t,attr,omitempty"`
	V   string            `xml:"v"`
	Stp []string          `xml:"stp"`
	Tr  []xlsxVolTopicRef `xml:"tr"`
}

type xlsxVolTopicRef struct {
	R string `xml:"r,attr"`
	S int    `xml:"s,attr"`
}
