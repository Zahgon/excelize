package excelize

import "encoding/xml"

type DocProperties struct {
	Category       string
	ContentStatus  string
	Created        string
	Creator        string
	Description    string
	Identifier     string
	Keywords       string
	LastModifiedBy string
	Modified       string
	Revision       string
	Subject        string
	Title          string
	Language       string
	Version        string
}

type decodeDcTerms struct {
	Text string `xml:",chardata"`
	Type string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`
}

type decodeCoreProperties struct {
	XMLName        xml.Name       `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Title          string         `xml:"http://purl.org/dc/elements/1.1/ title,omitempty"`
	Subject        string         `xml:"http://purl.org/dc/elements/1.1/ subject,omitempty"`
	Creator        string         `xml:"http://purl.org/dc/elements/1.1/ creator"`
	Keywords       string         `xml:"keywords,omitempty"`
	Description    string         `xml:"http://purl.org/dc/elements/1.1/ description,omitempty"`
	LastModifiedBy string         `xml:"lastModifiedBy"`
	Language       string         `xml:"http://purl.org/dc/elements/1.1/ language,omitempty"`
	Identifier     string         `xml:"http://purl.org/dc/elements/1.1/ identifier,omitempty"`
	Revision       string         `xml:"revision,omitempty"`
	Created        *decodeDcTerms `xml:"http://purl.org/dc/terms/ created"`
	Modified       *decodeDcTerms `xml:"http://purl.org/dc/terms/ modified"`
	ContentStatus  string         `xml:"contentStatus,omitempty"`
	Category       string         `xml:"category,omitempty"`
	Version        string         `xml:"version,omitempty"`
}

type xlsxDcTerms struct {
	Text string `xml:",chardata"`
	Type string `xml:"xsi:type,attr"`
}

type xlsxCoreProperties struct {
	XMLName        xml.Name     `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties coreProperties"`
	Dc             string       `xml:"xmlns:dc,attr"`
	Dcterms        string       `xml:"xmlns:dcterms,attr"`
	Dcmitype       string       `xml:"xmlns:dcmitype,attr"`
	XSI            string       `xml:"xmlns:xsi,attr"`
	Title          string       `xml:"dc:title,omitempty"`
	Subject        string       `xml:"dc:subject,omitempty"`
	Creator        string       `xml:"dc:creator"`
	Keywords       string       `xml:"keywords,omitempty"`
	Description    string       `xml:"dc:description,omitempty"`
	LastModifiedBy string       `xml:"lastModifiedBy"`
	Language       string       `xml:"dc:language,omitempty"`
	Identifier     string       `xml:"dc:identifier,omitempty"`
	Revision       string       `xml:"revision,omitempty"`
	Created        *xlsxDcTerms `xml:"dcterms:created"`
	Modified       *xlsxDcTerms `xml:"dcterms:modified"`
	ContentStatus  string       `xml:"contentStatus,omitempty"`
	Category       string       `xml:"category,omitempty"`
	Version        string       `xml:"version,omitempty"`
}
