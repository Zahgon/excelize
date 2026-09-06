package excelize

import "encoding/xml"

type xlsxMetadata struct {
	XMLName         xml.Name             `xml:"metadata"`
	MetadataTypes   *xlsxInnerXML        `xml:"metadataTypes"`
	MetadataStrings *xlsxInnerXML        `xml:"metadataStrings"`
	MdxMetadata     *xlsxInnerXML        `xml:"mdxMetadata"`
	FutureMetadata  []xlsxFutureMetadata `xml:"futureMetadata"`
	CellMetadata    *xlsxMetadataBlocks  `xml:"cellMetadata"`
	ValueMetadata   *xlsxMetadataBlocks  `xml:"valueMetadata"`
	ExtLst          *xlsxInnerXML        `xml:"extLst"`
}

type xlsxFutureMetadata struct {
	Bk     []xlsxFutureMetadataBlock `xml:"bk"`
	ExtLst *xlsxInnerXML             `xml:"extLst"`
}

type xlsxFutureMetadataBlock struct {
	ExtLst *xlsxInnerXML `xml:"extLst"`
}

type xlsxMetadataBlocks struct {
	Count int                 `xml:"count,attr,omitempty"`
	Bk    []xlsxMetadataBlock `xml:"bk"`
}

type xlsxMetadataBlock struct {
	Rc []xlsxMetadataRecord `xml:"rc"`
}

type xlsxMetadataRecord struct {
	T int `xml:"t,attr"`
	V int `xml:"v,attr"`
}

type xlsxRichValueData struct {
	XMLName xml.Name        `xml:"rvData"`
	Count   int             `xml:"count,attr,omitempty"`
	Rv      []xlsxRichValue `xml:"rv"`
	ExtLst  *xlsxInnerXML   `xml:"extLst"`
}

type xlsxRichValue struct {
	S  int           `xml:"s,attr"`
	V  []string      `xml:"v"`
	Fb *xlsxInnerXML `xml:"fb"`
}

type xlsxRichValueRels struct {
	XMLName xml.Name                       `xml:"richValueRels"`
	Rels    []xlsxRichValueRelRelationship `xml:"rel"`
	ExtLst  *xlsxInnerXML                  `xml:"extLst"`
}

type xlsxRichValueRelRelationship struct {
	ID string `xml:"id,attr"`
}

type xlsxRichValueStructures struct {
	XMLName xml.Name                 `xml:"rvStructures"`
	Count   int                      `xml:"count,attr,omitempty"`
	S       []xlsxRichValueStructure `xml:"s"`
	ExtLst  *xlsxInnerXML            `xml:"extLst"`
}

type xlsxRichValueStructure struct {
	T string             `xml:"t,attr"`
	K []xlsxRichValueKey `xml:"k"`
}

type xlsxRichValueKey struct {
	N string `xml:"n,attr"`
	T string `xml:"t,attr,omitempty"`
}

type xlsxWebImagesSupportingRichData struct {
	XMLName     xml.Name                         `xml:"webImagesSrd"`
	WebImageSrd []xlsxWebImageSupportingRichData `xml:"webImageSrd"`
	ExtLst      *xlsxInnerXML                    `xml:"extLst"`
}

type xlsxWebImageSupportingRichData struct {
	Address           xlsxExternalReference `xml:"address"`
	MoreImagesAddress xlsxExternalReference `xml:"moreImagesAddress"`
	Blip              xlsxExternalReference `xml:"blip"`
}
