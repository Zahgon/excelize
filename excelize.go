package excelize

import (
	"encoding/xml"
	"io"
	"io/fs"
	"os"
	"sync"
)

type File struct {
	mu               sync.Mutex
	checked          sync.Map
	formulaChecked   bool
	zip64Entries     []string
	options          *Options
	sharedStringItem [][]uint
	sharedStringsMap map[string]int
	sharedStringTemp *os.File
	sheetMap         map[string]string
	streams          map[string]*StreamWriter
	tempFiles        sync.Map
	xmlAttr          sync.Map
	calcCache        sync.Map
	calcRawCache     sync.Map
	formulaArgCache  sync.Map
	CalcChain        *xlsxCalcChain
	CharsetReader    func(charset string, input io.Reader) (rdr io.Reader, err error)
	Comments         map[string]*xlsxComments
	ContentTypes     *xlsxTypes
	DecodeVMLDrawing map[string]*decodeVmlDrawing
	DecodeCellImages *decodeCellImages
	Drawings         sync.Map
	Path             string
	Pkg              sync.Map
	Relationships    sync.Map
	SharedStrings    *xlsxSST
	Sheet            sync.Map
	SheetCount       int
	Styles           *xlsxStyleSheet
	Theme            *decodeTheme
	VMLDrawing       map[string]*vmlDrawing
	VolatileDeps     *xlsxVolTypes
	WorkBook         *xlsxWorkbook
	ZipWriter        func(io.Writer) ZipWriter
}

type ZipWriter interface {
	Create(name string) (io.Writer, error)
	AddFS(fsys fs.FS) error
	Close() error
}

type Options struct {
	MaxCalcIterations uint
	Password          string
	RawCellValue      bool
	UnzipSizeLimit    int64
	UnzipXMLSizeLimit int64
	TmpDir            string
	ShortDatePattern  string
	LongDatePattern   string
	LongTimePattern   string
	CultureInfo       CultureName
}

func OpenFile(filename string, opts ...Options) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newFile() *File { _ = "STUB: not implemented"; return nil }

func (f *File) checkOpenReaderOptions() error { _ = "STUB: not implemented"; return nil }

func OpenReader(r io.Reader, opts ...Options) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openReaderAt(r io.ReaderAt, size int64, opts ...Options) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getOptions(opts ...Options) *Options { _ = "STUB: not implemented"; return nil }

func (f *File) CharsetTranscoder(fn func(charset string, input io.Reader) (rdr io.Reader, err error)) *File {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) SetZipWriter(fn func(io.Writer) ZipWriter) *File {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) xmlNewDecoder(rdr io.Reader) (ret *xml.Decoder) {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) setDefaultTimeStyle(sheet, cell string, format int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) workSheetReader(sheet string) (ws *xlsxWorksheet, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func checkRowNum(r int) error { _ = "STUB: not implemented"; return nil }

func lastRowNum(r xlsxRow) int { _ = "STUB: not implemented"; return 0 }

func (ws *xlsxWorksheet) checkSheet() error { _ = "STUB: not implemented"; return nil }

func (ws *xlsxWorksheet) checkSheetR0(sheetData *xlsxSheetData, rowData *xlsxRow, r0 bool) {
	_ = "STUB: not implemented"
	return
}

func (f *File) setRels(rID, relPath, relType, target, targetMode string) int {
	_ = "STUB: not implemented"
	return 0
}

func (f *File) addRels(relPath, relType, target, targetMode string) int {
	_ = "STUB: not implemented"
	return 0
}

func (f *File) UpdateLinkedValue() error { _ = "STUB: not implemented"; return nil }

func (f *File) AddVBAProject(file []byte) error { _ = "STUB: not implemented"; return nil }

func (f *File) setContentTypePartProjectExtensions(contentType string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) metadataReader() (*xlsxMetadata, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) richValueReader() (*xlsxRichValueData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) richValueRelReader() (*xlsxRichValueRels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) richValueStructuresReader() (*xlsxRichValueStructures, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) richValueWebImageReader() (*xlsxWebImagesSupportingRichData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getRichDataRichValueRelRelationships(rID string) *xlsxRelationship {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getRichValueWebImageRelationships(rID string) *xlsxRelationship {
	_ = "STUB: not implemented"
	return nil
}
