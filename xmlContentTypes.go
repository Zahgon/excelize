package excelize

import (
	"encoding/xml"
	"sync"
)

type xlsxTypes struct {
	mu        sync.Mutex
	XMLName   xml.Name       `xml:"http://schemas.openxmlformats.org/package/2006/content-types Types"`
	Defaults  []xlsxDefault  `xml:"Default"`
	Overrides []xlsxOverride `xml:"Override"`
}

type xlsxOverride struct {
	PartName    string `xml:",attr"`
	ContentType string `xml:",attr"`
}

type xlsxDefault struct {
	Extension   string `xml:",attr"`
	ContentType string `xml:",attr"`
}
