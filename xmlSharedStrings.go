package excelize

import (
	"encoding/xml"
	"sync"
)

type xlsxSST struct {
	mu          sync.Mutex
	XMLName     xml.Name `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main sst"`
	Count       int      `xml:"count,attr"`
	UniqueCount int      `xml:"uniqueCount,attr"`
	SI          []xlsxSI `xml:"si"`
}

type xlsxSI struct {
	T          *xlsxT             `xml:"t,omitempty"`
	R          []xlsxR            `xml:"r"`
	RPh        []*xlsxPhoneticRun `xml:"rPh"`
	PhoneticPr *xlsxPhoneticPr    `xml:"phoneticPr"`
}

type xlsxR struct {
	XMLName xml.Name `xml:"r"`
	RPr     *xlsxRPr `xml:"rPr"`
	T       *xlsxT   `xml:"t"`
}

type xlsxT struct {
	XMLName xml.Name `xml:"t"`
	Space   xml.Attr `xml:"space,attr,omitempty"`
	Val     string   `xml:",chardata"`
}

type xlsxRPr struct {
	RFont     *attrValString `xml:"rFont"`
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

type RichTextRun struct {
	Font *Font
	Text string
}
