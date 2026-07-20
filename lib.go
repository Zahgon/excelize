package excelize

import (
	"archive/zip"
	"container/list"
	"encoding/xml"
	"os"
	"reflect"
	"regexp"
)

func (f *File) ReadZipReader(r *zip.Reader) (map[string][]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (f *File) unzipToTemp(zipFile *zip.File) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) readXML(name string) []byte { _ = "STUB: not implemented"; return nil }

func (f *File) readBytes(name string) []byte { _ = "STUB: not implemented"; return nil }

func (f *File) readTemp(name string) (file *os.File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) saveFileList(name string, content []byte) { _ = "STUB: not implemented"; return }

func readFile(file *zip.File) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func SplitCellName(cell string) (string, int, error) { _ = "STUB: not implemented"; return "", 0, nil }

func JoinCellName(col string, row int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ColumnNameToNumber(name string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var columnNames = func() []string {
	names := make([]string, MaxColumns+1)
	for i := 1; i <= MaxColumns; i++ {
		num := i
		l := 0
		for n := num; n > 0; n = (n - 1) / 26 {
			l++
		}
		buf := make([]byte, l)
		for num > 0 {
			l--
			buf[l] = byte((num-1)%26 + 'A')
			num = (num - 1) / 26
		}
		names[i] = string(buf)
	}
	return names
}()

func ColumnNumberToName(num int) (string, error) { _ = "STUB: not implemented"; return "", nil }

func CellNameToCoordinates(cell string) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func CoordinatesToCellName(col, row int, abs ...bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func rangeRefToCoordinates(ref string) ([]int, error) { _ = "STUB: not implemented"; return nil, nil }

func cellRefsToCoordinates(firstCell, lastCell string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func sortCoordinates(coordinates []int) error { _ = "STUB: not implemented"; return nil }

func coordinatesToRangeRef(coordinates []int, abs ...bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) getDefinedNameRefTo(definedNameName, currentSheet string) (refTo string) {
	_ = "STUB: not implemented"
	return ""
}

func flatSqref(sqref string) (cells map[int][][]int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func inCoordinates(a [][]int, x []int) int { _ = "STUB: not implemented"; return 0 }

func inStrSlice(a []string, x string, caseSensitive bool) int { _ = "STUB: not implemented"; return 0 }

func inFloat64Slice(a []float64, x float64) int { _ = "STUB: not implemented"; return 0 }

func boolPtr(b bool) *bool { _ = "STUB: not implemented"; return nil }

func intPtr(i int) *int { _ = "STUB: not implemented"; return nil }

func uintPtr(u uint) *uint { _ = "STUB: not implemented"; return nil }

func float64Ptr(f float64) *float64 { _ = "STUB: not implemented"; return nil }

func stringPtr(s string) *string { _ = "STUB: not implemented"; return nil }

func (attr *attrValFloat) Value() float64 { _ = "STUB: not implemented"; return 0 }

func (avb *attrValBool) Value() bool { _ = "STUB: not implemented"; return false }

func (avb *attrValString) Value() string { _ = "STUB: not implemented"; return "" }

func (avb attrValBool) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (avb *attrValBool) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (ext xlsxExt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func (ext *xlsxExt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	_ = "STUB: not implemented"
	return nil
}

func namespaceStrictToTransitional(content []byte) []byte { _ = "STUB: not implemented"; return nil }

func bytesReplace(s, source, target []byte, n int) []byte { _ = "STUB: not implemented"; return nil }

func genSheetPasswd(plaintext string) string { _ = "STUB: not implemented"; return "" }

func getRootElement(d *xml.Decoder) []xml.Attr { _ = "STUB: not implemented"; return nil }

func genXMLNamespace(attr []xml.Attr) string { _ = "STUB: not implemented"; return "" }

func getXMLNamespace(space string, attr []xml.Attr) string { _ = "STUB: not implemented"; return "" }

func (f *File) replaceNameSpaceBytes(path string, contentMarshal []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addNameSpaces(path string, ns xml.Attr) { _ = "STUB: not implemented"; return }

func (f *File) setIgnorableNameSpace(path string, index int, ns xml.Attr) {
	_ = "STUB: not implemented"
	return
}

func (f *File) addSheetNameSpace(sheet string, ns xml.Attr) { _ = "STUB: not implemented"; return }

func isNumeric(s string) (bool, int, float64) { _ = "STUB: not implemented"; return false, 0, 0 }

var (
	bstrExp       = regexp.MustCompile(`_x[a-fA-F\d]{4}_`)
	bstrEscapeExp = regexp.MustCompile(`x[a-fA-F\d]{4}_`)
)

func bstrUnmarshal(s string) (result string) { _ = "STUB: not implemented"; return "" }

func bstrMarshal(s string) (result string) { _ = "STUB: not implemented"; return "" }

func floatToFraction(x float64, numeratorPlaceHolder, denominatorPlaceHolder int) string {
	_ = "STUB: not implemented"
	return ""
}

func floatToFracUseContinuedFraction(r float64, denominatorLimit int64) (num, den int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func assignFieldValue(field string, immutable, mutable reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func setPtrFields(immutable, mutable reflect.Value) { _ = "STUB: not implemented"; return }

func setNoPtrFieldsVal(fields []string, immutable, mutable reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func setPtrFieldsVal(fields []string, immutable, mutable reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func countUTF16String(s string) int { _ = "STUB: not implemented"; return 0 }

func truncateUTF16Units(s string, length int) string { _ = "STUB: not implemented"; return "" }

type Stack struct {
	list *list.List
}

func NewStack() *Stack { _ = "STUB: not implemented"; return nil }

func (stack *Stack) Push(value interface{}) { _ = "STUB: not implemented"; return }

func (stack *Stack) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (stack *Stack) Peek() interface{} { _ = "STUB: not implemented"; return nil }

func (stack *Stack) Len() int { _ = "STUB: not implemented"; return 0 }

func (stack *Stack) Empty() bool { _ = "STUB: not implemented"; return false }
