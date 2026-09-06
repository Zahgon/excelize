package excelize

import "encoding/xml"

type xlsxCustomProperties struct {
	XMLName  xml.Name       `xml:"http://schemas.openxmlformats.org/officeDocument/2006/custom-properties Properties"`
	Vt       string         `xml:"xmlns:vt,attr"`
	Property []xlsxProperty `xml:"property"`
}

type xlsxProperty struct {
	XMLName    xml.Name `xml:"property"`
	FmtID      string   `xml:"fmtid,attr"`
	PID        int      `xml:"pid,attr"`
	Name       string   `xml:"name,attr,omitempty"`
	LinkTarget string   `xml:"linkTarget,attr,omitempty"`
	Vector     *string  `xml:"vt:vector"`
	Array      *string  `xml:"vt:array"`
	Blob       *string  `xml:"vt:blob"`
	Oblob      *string  `xml:"vt:oblob"`
	Empty      *string  `xml:"vt:empty"`
	Null       *string  `xml:"vt:null"`
	I1         *int8    `xml:"vt:i1"`
	I2         *int16   `xml:"vt:i2"`
	I4         *int32   `xml:"vt:i4"`
	I8         *int64   `xml:"vt:i8"`
	Int        *int     `xml:"vt:int"`
	Ui1        *uint8   `xml:"vt:ui1"`
	Ui2        *uint16  `xml:"vt:ui2"`
	Ui4        *uint32  `xml:"vt:ui4"`
	Ui8        *uint64  `xml:"vt:ui8"`
	Uint       *uint    `xml:"vt:uint"`
	R4         *float32 `xml:"vt:r4"`
	R8         *float64 `xml:"vt:r8"`
	Decimal    *string  `xml:"vt:decimal"`
	Lpstr      *string  `xml:"vt:lpstr"`
	Lpwstr     *string  `xml:"vt:lpwstr"`
	Bstr       *string  `xml:"vt:bstr"`
	Date       *string  `xml:"vt:date"`
	FileTime   *string  `xml:"vt:filetime"`
	Bool       *bool    `xml:"vt:bool"`
	Cy         *string  `xml:"vt:cy"`
	Error      *string  `xml:"vt:error"`
	Stream     *string  `xml:"vt:stream"`
	Ostream    *string  `xml:"vt:ostream"`
	Storage    *string  `xml:"vt:storage"`
	Ostorage   *string  `xml:"vt:ostorage"`
	Vstream    *string  `xml:"vt:vstream"`
	ClsID      *string  `xml:"vt:clsid"`
}

type decodeCustomProperties struct {
	XMLName  xml.Name         `xml:"http://schemas.openxmlformats.org/officeDocument/2006/custom-properties Properties"`
	Vt       string           `xml:"xmlns:vt,attr"`
	Property []decodeProperty `xml:"property"`
}

type decodeProperty struct {
	XMLName    xml.Name `xml:"property"`
	FmtID      string   `xml:"fmtid,attr"`
	PID        int      `xml:"pid,attr"`
	Name       string   `xml:"name,attr,omitempty"`
	LinkTarget string   `xml:"linkTarget,attr,omitempty"`
	Vector     *string  `xml:"vector"`
	Array      *string  `xml:"array"`
	Blob       *string  `xml:"blob"`
	Oblob      *string  `xml:"oblob"`
	Empty      *string  `xml:"empty"`
	Null       *string  `xml:"null"`
	I1         *int8    `xml:"i1"`
	I2         *int16   `xml:"i2"`
	I4         *int32   `xml:"i4"`
	I8         *int64   `xml:"i8"`
	Int        *int     `xml:"int"`
	Ui1        *uint8   `xml:"ui1"`
	Ui2        *uint16  `xml:"ui2"`
	Ui4        *uint32  `xml:"ui4"`
	Ui8        *uint64  `xml:"ui8"`
	Uint       *uint    `xml:"uint"`
	R4         *float32 `xml:"r4"`
	R8         *float64 `xml:"r8"`
	Decimal    *string  `xml:"decimal"`
	Lpstr      *string  `xml:"lpstr"`
	Lpwstr     *string  `xml:"lpwstr"`
	Bstr       *string  `xml:"bstr"`
	Date       *string  `xml:"date"`
	FileTime   *string  `xml:"filetime"`
	Bool       *bool    `xml:"bool"`
	Cy         *string  `xml:"cy"`
	Error      *string  `xml:"error"`
	Stream     *string  `xml:"stream"`
	Ostream    *string  `xml:"ostream"`
	Storage    *string  `xml:"storage"`
	Ostorage   *string  `xml:"ostorage"`
	Vstream    *string  `xml:"vstream"`
	ClsID      *string  `xml:"clsid"`
}

type CustomProperty struct {
	Name  string
	Value interface{}
}
