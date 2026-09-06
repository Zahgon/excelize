package excelize

import "encoding/xml"

type xlsxComments struct {
	XMLName     xml.Name        `xml:"http://schemas.openxmlformats.org/spreadsheetml/2006/main comments"`
	Authors     xlsxAuthor      `xml:"authors"`
	CommentList xlsxCommentList `xml:"commentList"`
	cells       []string
}

type xlsxAuthor struct {
	Author []string `xml:"author"`
}

type xlsxCommentList struct {
	Comment []xlsxComment `xml:"comment"`
}

type xlsxComment struct {
	Ref      string   `xml:"ref,attr"`
	AuthorID int      `xml:"authorId,attr"`
	Text     xlsxText `xml:"text"`
}

type xlsxText struct {
	T          *string          `xml:"t"`
	R          []xlsxR          `xml:"r"`
	RPh        *xlsxPhoneticRun `xml:"rPh"`
	PhoneticPr *xlsxPhoneticPr  `xml:"phoneticPr"`
}

type xlsxPhoneticRun struct {
	Sb uint32 `xml:"sb,attr"`
	Eb uint32 `xml:"eb,attr"`
	T  string `xml:"t"`
}

type Comment struct {
	Author    string
	AuthorID  int
	Cell      string
	Text      string
	Width     uint
	Height    uint
	Paragraph []RichTextRun
}
