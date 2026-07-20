package excelize

import (
	"container/list"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/xuri/efp"
)

const (
	formulaErrorDIV         = "#DIV/0!"
	formulaErrorNAME        = "#NAME?"
	formulaErrorNA          = "#N/A"
	formulaErrorNUM         = "#NUM!"
	formulaErrorVALUE       = "#VALUE!"
	formulaErrorREF         = "#REF!"
	formulaErrorNULL        = "#NULL!"
	formulaErrorSPILL       = "#SPILL!"
	formulaErrorCALC        = "#CALC!"
	formulaErrorGETTINGDATA = "#GETTING_DATA"

	_ byte = iota
	criteriaEq
	criteriaLe
	criteriaGe
	criteriaNe
	criteriaL
	criteriaG
	criteriaErr
	criteriaRegexp

	categoryWeightAndMass
	categoryDistance
	categoryTime
	categoryPressure
	categoryForce
	categoryEnergy
	categoryPower
	categoryMagnetism
	categoryTemperature
	categoryVolumeAndLiquidMeasure
	categoryArea
	categoryInformation
	categorySpeed

	matchModeExact      = 0
	matchModeMinGreater = 1
	matchModeMaxLess    = -1
	matchModeWildcard   = 2

	searchModeLinear        = 1
	searchModeReverseLinear = -1
	searchModeAscBinary     = 2
	searchModeDescBinary    = -2

	maxFinancialIterations = 128
	financialPrecision     = 1.0e-08

	monthRe    = `((jan|january)|(feb|february)|(mar|march)|(apr|april)|(may)|(jun|june)|(jul|july)|(aug|august)|(sep|september)|(oct|october)|(nov|november)|(dec|december))`
	df1        = `(([0-9])+)/(([0-9])+)/(([0-9])+)`
	df2        = monthRe + ` (([0-9])+), (([0-9])+)`
	df3        = `(([0-9])+)-(([0-9])+)-(([0-9])+)`
	df4        = `(([0-9])+)-` + monthRe + `-(([0-9])+)`
	datePrefix = `^((` + df1 + `|` + df2 + `|` + df3 + `|` + df4 + `) )?`
	tfhh       = `(([0-9])+) (am|pm)`
	tfhhmm     = `(([0-9])+):(([0-9])+)( (am|pm))?`
	tfmmss     = `(([0-9])+):(([0-9])+\.([0-9])+)( (am|pm))?`
	tfhhmmss   = `(([0-9])+):(([0-9])+):(([0-9])+(\.([0-9])+)?)( (am|pm))?`
	timeSuffix = `( (` + tfhh + `|` + tfhhmm + `|` + tfmmss + `|` + tfhhmmss + `))?$`
)

var (
	wildcardTokenRE = regexp.MustCompile(`~[*?~]|[*?]|[\s\S]`)

	wildcardPatternMap = map[string]string{
		"~*": regexp.QuoteMeta("*"),
		"~?": regexp.QuoteMeta("?"),
		"~~": regexp.QuoteMeta("~"),
		"*":  ".*",
		"?":  ".",
	}

	wildcardBareTokens = map[string]bool{"*": true, "?": true}

	tokenPriority = map[string]int{
		"^":  5,
		"*":  4,
		"/":  4,
		"+":  3,
		"-":  3,
		"&":  2,
		"=":  1,
		"<>": 1,
		"<":  1,
		"<=": 1,
		">":  1,
		">=": 1,
	}
	month2num = map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
		"jan":       1,
		"feb":       2,
		"mar":       3,
		"apr":       4,
		"jun":       6,
		"jul":       7,
		"aug":       8,
		"sep":       9,
		"oct":       10,
		"nov":       11,
		"dec":       12,
	}
	dateFormats = map[string]*regexp.Regexp{
		"mm/dd/yy":    regexp.MustCompile(`^` + df1 + timeSuffix),
		"mm dd, yy":   regexp.MustCompile(`^` + df2 + timeSuffix),
		"yy-mm-dd":    regexp.MustCompile(`^` + df3 + timeSuffix),
		"yy-mmStr-dd": regexp.MustCompile(`^` + df4 + timeSuffix),
	}
	timeFormats = map[string]*regexp.Regexp{
		"hh":       regexp.MustCompile(datePrefix + tfhh + `$`),
		"hh:mm":    regexp.MustCompile(datePrefix + tfhhmm + `$`),
		"mm:ss":    regexp.MustCompile(datePrefix + tfmmss + `$`),
		"hh:mm:ss": regexp.MustCompile(datePrefix + tfhhmmss + `$`),
	}
	dateOnlyFormats = []*regexp.Regexp{
		regexp.MustCompile(`^` + df1 + `$`),
		regexp.MustCompile(`^` + df2 + `$`),
		regexp.MustCompile(`^` + df3 + `$`),
		regexp.MustCompile(`^` + df4 + `$`),
	}
	addressFmtMaps = map[string]func(col, row int) (string, error){
		"1_TRUE": func(col, row int) (string, error) {
			return CoordinatesToCellName(col, row, true)
		},
		"1_FALSE": func(col, row int) (string, error) {
			return fmt.Sprintf("R%dC%d", row, col), nil
		},
		"2_TRUE": func(col, row int) (string, error) {
			column, err := ColumnNumberToName(col)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("%s$%d", column, row), nil
		},
		"2_FALSE": func(col, row int) (string, error) {
			return fmt.Sprintf("R%dC[%d]", row, col), nil
		},
		"3_TRUE": func(col, row int) (string, error) {
			column, err := ColumnNumberToName(col)
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("$%s%d", column, row), nil
		},
		"3_FALSE": func(col, row int) (string, error) {
			return fmt.Sprintf("R[%d]C%d", row, col), nil
		},
		"4_TRUE": func(col, row int) (string, error) {
			return CoordinatesToCellName(col, row, false)
		},
		"4_FALSE": func(col, row int) (string, error) {
			return fmt.Sprintf("R[%d]C[%d]", row, col), nil
		},
	}
	formulaFnNameReplacer = strings.NewReplacer("_xlfn.", "", ".", "dot")
	formulaFormats        = []*regexp.Regexp{
		regexp.MustCompile(`^(\d+)$`),
		regexp.MustCompile(`^=(.*)$`),
		regexp.MustCompile(`^<>(.*)$`),
		regexp.MustCompile(`^<=(.*)$`),
		regexp.MustCompile(`^>=(.*)$`),
		regexp.MustCompile(`^<(.*)$`),
		regexp.MustCompile(`^>(.*)$`),
	}
	formulaCriterias = []byte{
		criteriaEq,
		criteriaEq,
		criteriaNe,
		criteriaLe,
		criteriaGe,
		criteriaL,
		criteriaG,
	}

	th0      = "\u0E28\u0E39\u0E19\u0E22\u0E4C"
	th1      = "\u0E2B\u0E19\u0E36\u0E48\u0E07"
	th2      = "\u0E2A\u0E2D\u0E07"
	th3      = "\u0E2A\u0E32\u0E21"
	th4      = "\u0E2A\u0E35\u0E48"
	th5      = "\u0E2B\u0E49\u0E32"
	th6      = "\u0E2B\u0E01"
	th7      = "\u0E40\u0E08\u0E47\u0E14"
	th8      = "\u0E41\u0E1B\u0E14"
	th9      = "\u0E40\u0E01\u0E49\u0E32"
	th10     = "\u0E2A\u0E34\u0E1A"
	th11     = "\u0E40\u0E2D\u0E47\u0E14"
	th20     = "\u0E22\u0E35\u0E48"
	th1e2    = "\u0E23\u0E49\u0E2D\u0E22"
	th1e3    = "\u0E1E\u0E31\u0E19"
	th1e4    = "\u0E2B\u0E21\u0E37\u0E48\u0E19"
	th1e5    = "\u0E41\u0E2A\u0E19"
	th1e6    = "\u0E25\u0E49\u0E32\u0E19"
	thDot0   = "\u0E16\u0E49\u0E27\u0E19"
	thBaht   = "\u0E1A\u0E32\u0E17"
	thSatang = "\u0E2A\u0E15\u0E32\u0E07\u0E04\u0E4C"
	thMinus  = "\u0E25\u0E1A"
)

type calcContext struct {
	mu                sync.Mutex
	entry             string
	maxCalcIterations uint
	iterations        map[string]uint
	iterationsCache   map[string]formulaArg
}

type cellRef struct {
	Col   int
	Row   int
	Sheet string
}

type cellRange struct {
	From cellRef
	To   cellRef
}

type formulaCriteria struct {
	Type      byte
	Condition formulaArg
}

type ArgType byte

const (
	ArgUnknown ArgType = iota
	ArgNumber
	ArgString
	ArgList
	ArgMatrix
	ArgError
	ArgEmpty
)

type formulaArg struct {
	SheetName            string
	Number               float64
	String               string
	List                 []formulaArg
	Matrix               [][]formulaArg
	Boolean              bool
	Error                string
	Type                 ArgType
	cellRefs, cellRanges *list.List
}

func (fa formulaArg) Value() (value string) { _ = "STUB: not implemented"; return "" }

func (fa formulaArg) ToNumber() formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func (fa formulaArg) ToBool() formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func (fa formulaArg) ToList() []formulaArg { _ = "STUB: not implemented"; return nil }

type formulaFuncs struct {
	f           *File
	ctx         *calcContext
	sheet, cell string
}

func (fn *formulaFuncs) implicitIntersect(arg formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (f *File) CalcCellValue(sheet, cell string, opts ...Options) (result string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *File) storeCalcCache(entry, result string, rawCellValue bool) {
	_ = "STUB: not implemented"
	return
}

func (f *File) clearCalcCache() { _ = "STUB: not implemented"; return }

func (f *File) calcCellValue(ctx *calcContext, sheet, cell string) (result formulaArg, err error) {
	_ = "STUB: not implemented"
	return *new(formulaArg), nil
}

func getPriority(token efp.Token) (pri int) { _ = "STUB: not implemented"; return 0 }

func newNumberFormulaArg(n float64) formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func newStringFormulaArg(s string) formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func newMatrixFormulaArg(m [][]formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func newListFormulaArg(l []formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func newBoolFormulaArg(b bool) formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func newErrorFormulaArg(formulaError, msg string) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func newEmptyFormulaArg() formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func (f *File) evalInfixExp(ctx *calcContext, sheet, cell string, tokens []efp.Token) (formulaArg, error) {
	_ = "STUB: not implemented"
	return *new(formulaArg), nil
}

func (f *File) evalInfixExpFunc(ctx *calcContext, sheet, cell string, token, nextToken efp.Token, opfStack, opdStack, opftStack, opfdStack, argsStack *Stack) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func prepareEvalInfixExp(opfStack, opftStack, opfdStack, argsStack *Stack) {
	_ = "STUB: not implemented"
	return
}

func calcPow(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcEq(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcNEq(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcL(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcLe(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcG(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcGe(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcSplice(rOpd, lOpd formulaArg, opdStack *Stack) error {
	_ = "STUB: not implemented"
	return nil
}

func calcAdd(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calcSubtract(rOpd, lOpd formulaArg, opdStack *Stack) error {
	_ = "STUB: not implemented"
	return nil
}

func calcMultiply(rOpd, lOpd formulaArg, opdStack *Stack) error {
	_ = "STUB: not implemented"
	return nil
}

func calcDiv(rOpd, lOpd formulaArg, opdStack *Stack) error { _ = "STUB: not implemented"; return nil }

func calculate(opdStack *Stack, opt efp.Token) error { _ = "STUB: not implemented"; return nil }

func (f *File) parseOperatorPrefixToken(optStack, opdStack *Stack, token efp.Token) {
	_ = "STUB: not implemented"
	return
}

func isFunctionStartToken(token efp.Token) bool { _ = "STUB: not implemented"; return false }

func isFunctionStopToken(token efp.Token) bool { _ = "STUB: not implemented"; return false }

func isBeginParenthesesToken(token efp.Token) bool { _ = "STUB: not implemented"; return false }

func isEndParenthesesToken(token efp.Token) bool { _ = "STUB: not implemented"; return false }

func isOperatorPrefixToken(token efp.Token) bool { _ = "STUB: not implemented"; return false }

func isOperand(token efp.Token) bool { _ = "STUB: not implemented"; return false }

func tokenToFormulaArg(token efp.Token) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func formulaArgToToken(arg formulaArg) efp.Token { _ = "STUB: not implemented"; return *new(efp.Token) }

func (f *File) parseToken(ctx *calcContext, sheet string, token efp.Token, opdStack, optStack *Stack) error {
	_ = "STUB: not implemented"
	return nil
}

func parseRef(ref string) (cellRef, bool, bool, error) {
	_ = "STUB: not implemented"
	return *new(cellRef), false, false, nil
}

func (cr *cellRange) prepareCellRange(col, row bool, cellRef cellRef) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) parseReference(ctx *calcContext, sheet, reference string) (formulaArg, error) {
	_ = "STUB: not implemented"
	return *new(formulaArg), nil
}

func (f *File) parse3DReference(ctx *calcContext, parts []string) (formulaArg, error) {
	_ = "STUB: not implemented"
	return *new(formulaArg), nil
}

func split3DReference(reference string) []string { _ = "STUB: not implemented"; return nil }

func readSheetToken(s string) (name, rest string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func (f *File) expand3DSheetRange(sheet1, sheet2 string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func prepareValueRange(cr cellRange, valueRange []int) { _ = "STUB: not implemented"; return }

func prepareValueRef(cr cellRef, valueRange []int) { _ = "STUB: not implemented"; return }

func (f *File) cellResolver(ctx *calcContext, sheet, cell string) (formulaArg, error) {
	_ = "STUB: not implemented"
	return *new(formulaArg), nil
}

func (f *File) rangeResolver(ctx *calcContext, cellRefs, cellRanges *list.List) (arg formulaArg, err error) {
	_ = "STUB: not implemented"
	return *new(formulaArg), nil
}

func callFuncByName(receiver interface{}, name string, params []reflect.Value) (arg formulaArg) {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func formulaCriteriaParser(exp formulaArg) *formulaCriteria { _ = "STUB: not implemented"; return nil }

func formulaCriteriaEval(val formulaArg, criteria *formulaCriteria) (result bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (fn *formulaFuncs) BESSELI(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BESSELJ(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) bassel(argsList *list.List, modfied bool) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BESSELK(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) besselK0(x formulaArg) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) besselK1(x formulaArg) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) besselK2(x, n formulaArg) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) BESSELY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) besselY0(x formulaArg) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) besselY1(x formulaArg) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) besselY2(x, n formulaArg) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) BIN2DEC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BIN2HEX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BIN2OCT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) bin2dec(number string) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BITAND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BITLSHIFT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BITOR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BITRSHIFT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BITXOR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) bitwise(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COMPLEX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func cmplx2str(num complex128, suffix string) string { _ = "STUB: not implemented"; return "" }

func str2cmplx(c string) string { _ = "STUB: not implemented"; return "" }

type conversionUnit struct {
	group       uint8
	allowPrefix bool
}

var conversionUnits = map[string]conversionUnit{

	"g":        {group: categoryWeightAndMass, allowPrefix: true},
	"sg":       {group: categoryWeightAndMass, allowPrefix: false},
	"lbm":      {group: categoryWeightAndMass, allowPrefix: false},
	"u":        {group: categoryWeightAndMass, allowPrefix: true},
	"ozm":      {group: categoryWeightAndMass, allowPrefix: false},
	"grain":    {group: categoryWeightAndMass, allowPrefix: false},
	"cwt":      {group: categoryWeightAndMass, allowPrefix: false},
	"shweight": {group: categoryWeightAndMass, allowPrefix: false},
	"uk_cwt":   {group: categoryWeightAndMass, allowPrefix: false},
	"lcwt":     {group: categoryWeightAndMass, allowPrefix: false},
	"hweight":  {group: categoryWeightAndMass, allowPrefix: false},
	"stone":    {group: categoryWeightAndMass, allowPrefix: false},
	"ton":      {group: categoryWeightAndMass, allowPrefix: false},
	"uk_ton":   {group: categoryWeightAndMass, allowPrefix: false},
	"LTON":     {group: categoryWeightAndMass, allowPrefix: false},
	"brton":    {group: categoryWeightAndMass, allowPrefix: false},

	"m":         {group: categoryDistance, allowPrefix: true},
	"mi":        {group: categoryDistance, allowPrefix: false},
	"Nmi":       {group: categoryDistance, allowPrefix: false},
	"in":        {group: categoryDistance, allowPrefix: false},
	"ft":        {group: categoryDistance, allowPrefix: false},
	"yd":        {group: categoryDistance, allowPrefix: false},
	"ang":       {group: categoryDistance, allowPrefix: true},
	"ell":       {group: categoryDistance, allowPrefix: false},
	"ly":        {group: categoryDistance, allowPrefix: false},
	"parsec":    {group: categoryDistance, allowPrefix: false},
	"pc":        {group: categoryDistance, allowPrefix: false},
	"Pica":      {group: categoryDistance, allowPrefix: false},
	"Picapt":    {group: categoryDistance, allowPrefix: false},
	"pica":      {group: categoryDistance, allowPrefix: false},
	"survey_mi": {group: categoryDistance, allowPrefix: false},

	"yr":  {group: categoryTime, allowPrefix: false},
	"day": {group: categoryTime, allowPrefix: false},
	"d":   {group: categoryTime, allowPrefix: false},
	"hr":  {group: categoryTime, allowPrefix: false},
	"mn":  {group: categoryTime, allowPrefix: false},
	"min": {group: categoryTime, allowPrefix: false},
	"sec": {group: categoryTime, allowPrefix: true},
	"s":   {group: categoryTime, allowPrefix: true},

	"Pa":   {group: categoryPressure, allowPrefix: true},
	"p":    {group: categoryPressure, allowPrefix: true},
	"atm":  {group: categoryPressure, allowPrefix: true},
	"at":   {group: categoryPressure, allowPrefix: true},
	"mmHg": {group: categoryPressure, allowPrefix: true},
	"psi":  {group: categoryPressure, allowPrefix: true},
	"Torr": {group: categoryPressure, allowPrefix: true},

	"N":    {group: categoryForce, allowPrefix: true},
	"dyn":  {group: categoryForce, allowPrefix: true},
	"dy":   {group: categoryForce, allowPrefix: true},
	"lbf":  {group: categoryForce, allowPrefix: false},
	"pond": {group: categoryForce, allowPrefix: true},

	"J":   {group: categoryEnergy, allowPrefix: true},
	"e":   {group: categoryEnergy, allowPrefix: true},
	"c":   {group: categoryEnergy, allowPrefix: true},
	"cal": {group: categoryEnergy, allowPrefix: true},
	"eV":  {group: categoryEnergy, allowPrefix: true},
	"ev":  {group: categoryEnergy, allowPrefix: true},
	"HPh": {group: categoryEnergy, allowPrefix: false},
	"hh":  {group: categoryEnergy, allowPrefix: false},
	"Wh":  {group: categoryEnergy, allowPrefix: true},
	"wh":  {group: categoryEnergy, allowPrefix: true},
	"flb": {group: categoryEnergy, allowPrefix: false},
	"BTU": {group: categoryEnergy, allowPrefix: false},
	"btu": {group: categoryEnergy, allowPrefix: false},

	"HP": {group: categoryPower, allowPrefix: false},
	"h":  {group: categoryPower, allowPrefix: false},
	"W":  {group: categoryPower, allowPrefix: true},
	"w":  {group: categoryPower, allowPrefix: true},
	"PS": {group: categoryPower, allowPrefix: false},
	"T":  {group: categoryMagnetism, allowPrefix: true},
	"ga": {group: categoryMagnetism, allowPrefix: true},

	"C":    {group: categoryTemperature, allowPrefix: false},
	"cel":  {group: categoryTemperature, allowPrefix: false},
	"F":    {group: categoryTemperature, allowPrefix: false},
	"fah":  {group: categoryTemperature, allowPrefix: false},
	"K":    {group: categoryTemperature, allowPrefix: false},
	"kel":  {group: categoryTemperature, allowPrefix: false},
	"Rank": {group: categoryTemperature, allowPrefix: false},
	"Reau": {group: categoryTemperature, allowPrefix: false},

	"l":        {group: categoryVolumeAndLiquidMeasure, allowPrefix: true},
	"L":        {group: categoryVolumeAndLiquidMeasure, allowPrefix: true},
	"lt":       {group: categoryVolumeAndLiquidMeasure, allowPrefix: true},
	"tsp":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"tspm":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"tbs":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"oz":       {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"cup":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"pt":       {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"us_pt":    {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"uk_pt":    {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"qt":       {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"uk_qt":    {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"gal":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"uk_gal":   {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"ang3":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: true},
	"ang^3":    {group: categoryVolumeAndLiquidMeasure, allowPrefix: true},
	"barrel":   {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"bushel":   {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"in3":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"in^3":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"ft3":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"ft^3":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"ly3":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"ly^3":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"m3":       {group: categoryVolumeAndLiquidMeasure, allowPrefix: true},
	"m^3":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: true},
	"mi3":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"mi^3":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"yd3":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"yd^3":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"Nmi3":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"Nmi^3":    {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"Pica3":    {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"Pica^3":   {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"Picapt3":  {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"Picapt^3": {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"GRT":      {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"regton":   {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},
	"MTON":     {group: categoryVolumeAndLiquidMeasure, allowPrefix: false},

	"ha":       {group: categoryArea, allowPrefix: true},
	"uk_acre":  {group: categoryArea, allowPrefix: false},
	"us_acre":  {group: categoryArea, allowPrefix: false},
	"ang2":     {group: categoryArea, allowPrefix: true},
	"ang^2":    {group: categoryArea, allowPrefix: true},
	"ar":       {group: categoryArea, allowPrefix: true},
	"ft2":      {group: categoryArea, allowPrefix: false},
	"ft^2":     {group: categoryArea, allowPrefix: false},
	"in2":      {group: categoryArea, allowPrefix: false},
	"in^2":     {group: categoryArea, allowPrefix: false},
	"ly2":      {group: categoryArea, allowPrefix: false},
	"ly^2":     {group: categoryArea, allowPrefix: false},
	"m2":       {group: categoryArea, allowPrefix: true},
	"m^2":      {group: categoryArea, allowPrefix: true},
	"Morgen":   {group: categoryArea, allowPrefix: false},
	"mi2":      {group: categoryArea, allowPrefix: false},
	"mi^2":     {group: categoryArea, allowPrefix: false},
	"Nmi2":     {group: categoryArea, allowPrefix: false},
	"Nmi^2":    {group: categoryArea, allowPrefix: false},
	"Pica2":    {group: categoryArea, allowPrefix: false},
	"Pica^2":   {group: categoryArea, allowPrefix: false},
	"Picapt2":  {group: categoryArea, allowPrefix: false},
	"Picapt^2": {group: categoryArea, allowPrefix: false},
	"yd2":      {group: categoryArea, allowPrefix: false},
	"yd^2":     {group: categoryArea, allowPrefix: false},

	"byte": {group: categoryInformation, allowPrefix: true},
	"bit":  {group: categoryInformation, allowPrefix: true},

	"m/s":   {group: categorySpeed, allowPrefix: true},
	"m/sec": {group: categorySpeed, allowPrefix: true},
	"m/h":   {group: categorySpeed, allowPrefix: true},
	"m/hr":  {group: categorySpeed, allowPrefix: true},
	"mph":   {group: categorySpeed, allowPrefix: false},
	"admkn": {group: categorySpeed, allowPrefix: false},
	"kn":    {group: categorySpeed, allowPrefix: false},
}

var unitConversions = map[byte]map[string]float64{

	categoryWeightAndMass: {
		"g":        1,
		"sg":       6.85217658567918e-05,
		"lbm":      2.20462262184878e-03,
		"u":        6.02214179421676e+23,
		"ozm":      3.52739619495804e-02,
		"grain":    1.54323583529414e+01,
		"cwt":      2.20462262184878e-05,
		"shweight": 2.20462262184878e-05,
		"uk_cwt":   1.96841305522212e-05,
		"lcwt":     1.96841305522212e-05,
		"hweight":  1.96841305522212e-05,
		"stone":    1.57473044417770e-04,
		"ton":      1.10231131092439e-06,
		"uk_ton":   9.84206527611061e-07,
		"LTON":     9.84206527611061e-07,
		"brton":    9.84206527611061e-07,
	},

	categoryDistance: {
		"m":         1,
		"mi":        6.21371192237334e-04,
		"Nmi":       5.39956803455724e-04,
		"in":        3.93700787401575e+01,
		"ft":        3.28083989501312e+00,
		"yd":        1.09361329833771e+00,
		"ang":       1.0e+10,
		"ell":       8.74890638670166e-01,
		"ly":        1.05700083402462e-16,
		"parsec":    3.24077928966473e-17,
		"pc":        3.24077928966473e-17,
		"Pica":      2.83464566929134e+03,
		"Picapt":    2.83464566929134e+03,
		"pica":      2.36220472440945e+02,
		"survey_mi": 6.21369949494950e-04,
	},

	categoryTime: {
		"yr":  3.16880878140289e-08,
		"day": 1.15740740740741e-05,
		"d":   1.15740740740741e-05,
		"hr":  2.77777777777778e-04,
		"mn":  1.66666666666667e-02,
		"min": 1.66666666666667e-02,
		"sec": 1,
		"s":   1,
	},

	categoryPressure: {
		"Pa":   1,
		"p":    1,
		"atm":  9.86923266716013e-06,
		"at":   9.86923266716013e-06,
		"mmHg": 7.50063755419211e-03,
		"psi":  1.45037737730209e-04,
		"Torr": 7.50061682704170e-03,
	},

	categoryForce: {
		"N":    1,
		"dyn":  1.0e+5,
		"dy":   1.0e+5,
		"lbf":  2.24808923655339e-01,
		"pond": 1.01971621297793e+02,
	},

	categoryEnergy: {
		"J":   1,
		"e":   9.99999519343231e+06,
		"c":   2.39006249473467e-01,
		"cal": 2.38846190642017e-01,
		"eV":  6.24145700000000e+18,
		"ev":  6.24145700000000e+18,
		"HPh": 3.72506430801000e-07,
		"hh":  3.72506430801000e-07,
		"Wh":  2.77777916238711e-04,
		"wh":  2.77777916238711e-04,
		"flb": 2.37304222192651e+01,
		"BTU": 9.47815067349015e-04,
		"btu": 9.47815067349015e-04,
	},

	categoryPower: {
		"HP": 1,
		"h":  1,
		"W":  7.45699871582270e+02,
		"w":  7.45699871582270e+02,
		"PS": 1.01386966542400e+00,
	},

	categoryMagnetism: {
		"T":  1,
		"ga": 10000,
	},

	categoryVolumeAndLiquidMeasure: {
		"l":        1,
		"L":        1,
		"lt":       1,
		"tsp":      2.02884136211058e+02,
		"tspm":     2.0e+02,
		"tbs":      6.76280454036860e+01,
		"oz":       3.38140227018430e+01,
		"cup":      4.22675283773038e+00,
		"pt":       2.11337641886519e+00,
		"us_pt":    2.11337641886519e+00,
		"uk_pt":    1.75975398639270e+00,
		"qt":       1.05668820943259e+00,
		"uk_qt":    8.79876993196351e-01,
		"gal":      2.64172052358148e-01,
		"uk_gal":   2.19969248299088e-01,
		"ang3":     1.0e+27,
		"ang^3":    1.0e+27,
		"barrel":   6.28981077043211e-03,
		"bushel":   2.83775932584017e-02,
		"in3":      6.10237440947323e+01,
		"in^3":     6.10237440947323e+01,
		"ft3":      3.53146667214886e-02,
		"ft^3":     3.53146667214886e-02,
		"ly3":      1.18093498844171e-51,
		"ly^3":     1.18093498844171e-51,
		"m3":       1.0e-03,
		"m^3":      1.0e-03,
		"mi3":      2.39912758578928e-13,
		"mi^3":     2.39912758578928e-13,
		"yd3":      1.30795061931439e-03,
		"yd^3":     1.30795061931439e-03,
		"Nmi3":     1.57426214685811e-13,
		"Nmi^3":    1.57426214685811e-13,
		"Pica3":    2.27769904358706e+07,
		"Pica^3":   2.27769904358706e+07,
		"Picapt3":  2.27769904358706e+07,
		"Picapt^3": 2.27769904358706e+07,
		"GRT":      3.53146667214886e-04,
		"regton":   3.53146667214886e-04,
		"MTON":     8.82866668037215e-04,
	},

	categoryArea: {
		"ha":       1,
		"uk_acre":  2.47105381467165e+00,
		"us_acre":  2.47104393046628e+00,
		"ang2":     1.0e+24,
		"ang^2":    1.0e+24,
		"ar":       1.0e+02,
		"ft2":      1.07639104167097e+05,
		"ft^2":     1.07639104167097e+05,
		"in2":      1.55000310000620e+07,
		"in^2":     1.55000310000620e+07,
		"ly2":      1.11725076312873e-28,
		"ly^2":     1.11725076312873e-28,
		"m2":       1.0e+04,
		"m^2":      1.0e+04,
		"Morgen":   4.0e+00,
		"mi2":      3.86102158542446e-03,
		"mi^2":     3.86102158542446e-03,
		"Nmi2":     2.91553349598123e-03,
		"Nmi^2":    2.91553349598123e-03,
		"Pica2":    8.03521607043214e+10,
		"Pica^2":   8.03521607043214e+10,
		"Picapt2":  8.03521607043214e+10,
		"Picapt^2": 8.03521607043214e+10,
		"yd2":      1.19599004630108e+04,
		"yd^2":     1.19599004630108e+04,
	},

	categoryInformation: {
		"bit":  1,
		"byte": 0.125,
	},

	categorySpeed: {
		"m/s":   1,
		"m/sec": 1,
		"m/h":   3.60e+03,
		"m/hr":  3.60e+03,
		"mph":   2.23693629205440e+00,
		"admkn": 1.94260256941567e+00,
		"kn":    1.94384449244060e+00,
	},
}

var conversionMultipliers = map[string]float64{
	"Y":  1e24,
	"Z":  1e21,
	"E":  1e18,
	"P":  1e15,
	"T":  1e12,
	"G":  1e9,
	"M":  1e6,
	"k":  1e3,
	"h":  1e2,
	"e":  1e1,
	"da": 1e1,
	"d":  1e-1,
	"c":  1e-2,
	"m":  1e-3,
	"u":  1e-6,
	"n":  1e-9,
	"p":  1e-12,
	"f":  1e-15,
	"a":  1e-18,
	"z":  1e-21,
	"y":  1e-24,
	"Yi": math.Pow(2, 80),
	"Zi": math.Pow(2, 70),
	"Ei": math.Pow(2, 60),
	"Pi": math.Pow(2, 50),
	"Ti": math.Pow(2, 40),
	"Gi": math.Pow(2, 30),
	"Mi": math.Pow(2, 20),
	"ki": math.Pow(2, 10),
}

func getUnitDetails(uom string) (unit string, catgory byte, res float64, ok bool) {
	_ = "STUB: not implemented"
	return "", 0, 0, false
}

func resolveTemperatureSynonyms(uom string) string { _ = "STUB: not implemented"; return "" }

func convertTemperature(fromUOM, toUOM string, value float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) CONVERT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DEC2BIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DEC2HEX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DEC2OCT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) dec2x(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DELTA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ERF(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ERFdotPRECISE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) erfc(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ERFC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ERFCdotPRECISE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GESTEP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HEX2BIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HEX2DEC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HEX2OCT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) hex2dec(number string) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMABS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMAGINARY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMARGUMENT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMCONJUGATE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMCOS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMCOSH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMCOT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMCSC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMCSCH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMDIV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMEXP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMLN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMLOG10(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMLOG2(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMPOWER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMPRODUCT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMREAL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMSEC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMSECH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMSIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMSINH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMSQRT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMSUB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMSUM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IMTAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) OCT2BIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) OCT2DEC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) OCT2HEX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) oct2dec(number string) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ABS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ACOS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ACOSH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ACOT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ACOTH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AGGREGATE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ARABIC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ASIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ASINH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ATAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ATANH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ATAN2(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BASE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CEILING(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CEILINGdotMATH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CEILINGdotPRECISE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COMBIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COMBINA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COSH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COTH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CSC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CSCH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DECIMAL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DEGREES(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EVEN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EXP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func fact(number float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) FACT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FACTDOUBLE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FLOOR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FLOORdotMATH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FLOORdotPRECISE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func gcd(x, y float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) GCD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) INT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISOdotCEILING(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func lcm(a, b float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) LCM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOG(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOG10(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func minor(sqMtx [][]float64, idx int) [][]float64 { _ = "STUB: not implemented"; return nil }

func det(sqMtx [][]float64) float64 { _ = "STUB: not implemented"; return 0 }

func newNumberMatrix(arg formulaArg, phalanx bool) (numMtx [][]float64, ele formulaArg) {
	_ = "STUB: not implemented"
	return nil, *new(formulaArg)
}

func newFormulaArgMatrix(numMtx [][]float64) (arg [][]formulaArg) {
	_ = "STUB: not implemented"
	return nil
}

func (fn *formulaFuncs) MDETERM(argsList *list.List) (result formulaArg) {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func cofactorMatrix(i, j int, A [][]float64) float64 { _ = "STUB: not implemented"; return 0 }

func adjugateMatrix(A [][]float64) (adjA [][]float64) { _ = "STUB: not implemented"; return nil }

func (fn *formulaFuncs) MINVERSE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MMULT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MOD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MROUND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MULTINOMIAL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MUNIT(argsList *list.List) (result formulaArg) {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ODD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PI(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) POWER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PRODUCT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) QUOTIENT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RADIANS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RAND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RANDBETWEEN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

type romanNumerals struct {
	n float64
	s string
}

var romanTable = [][]romanNumerals{
	{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	},
	{
		{1000, "M"},
		{950, "LM"},
		{900, "CM"},
		{500, "D"},
		{450, "LD"},
		{400, "CD"},
		{100, "C"},
		{95, "VC"},
		{90, "XC"},
		{50, "L"},
		{45, "VL"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	},
	{
		{1000, "M"},
		{990, "XM"},
		{950, "LM"},
		{900, "CM"},
		{500, "D"},
		{490, "XD"},
		{450, "LD"},
		{400, "CD"},
		{100, "C"},
		{99, "IC"},
		{90, "XC"},
		{50, "L"},
		{45, "VL"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	},
	{
		{1000, "M"},
		{995, "VM"},
		{990, "XM"},
		{950, "LM"},
		{900, "CM"},
		{500, "D"},
		{495, "VD"},
		{490, "XD"},
		{450, "LD"},
		{400, "CD"},
		{100, "C"},
		{99, "IC"},
		{90, "XC"},
		{50, "L"},
		{45, "VL"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	},
	{
		{1000, "M"},
		{999, "IM"},
		{995, "VM"},
		{990, "XM"},
		{950, "LM"},
		{900, "CM"},
		{500, "D"},
		{499, "ID"},
		{495, "VD"},
		{490, "XD"},
		{450, "LD"},
		{400, "CD"},
		{100, "C"},
		{99, "IC"},
		{90, "XC"},
		{50, "L"},
		{45, "VL"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	},
}

func (fn *formulaFuncs) ROMAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

type roundMode byte

const (
	closest roundMode = iota
	down
	up
)

func (fn *formulaFuncs) round(number, digits float64, mode roundMode) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) ROUND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ROUNDDOWN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ROUNDUP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SEC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SECH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SERIESSUM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SIGN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SINH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SQRT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SQRTPI(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STDEV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STDEVdotS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STDEVA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcStdevPow(result, count float64, n, m formulaArg) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func calcStdev(stdeva bool, result, count float64, mean, token formulaArg) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (fn *formulaFuncs) stdev(stdeva bool, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) POISSONdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) POISSON(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func prepareProbArgs(argsList *list.List) []formulaArg { _ = "STUB: not implemented"; return nil }

func (fn *formulaFuncs) PROB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUBTOTAL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUMIF(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUMIFS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) sumproduct(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUMPRODUCT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUMSQ(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) sumx(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUMX2MY2(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUMX2PY2(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUMXMY2(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TANH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TRUNC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AVEDEV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AVERAGE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AVERAGEA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AVERAGEIF(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AVERAGEIFS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func getBetaHelperContFrac(fX, fA, fB float64) float64 { _ = "STUB: not implemented"; return 0 }

func getLanczosSum(fZ float64) float64 { _ = "STUB: not implemented"; return 0 }

func getBeta(fAlpha, fBeta float64) float64 { _ = "STUB: not implemented"; return 0 }

func getBetaDistPDF(fX, fA, fB float64) float64 { _ = "STUB: not implemented"; return 0 }

func getLogBeta(fAlpha, fBeta float64) float64 { _ = "STUB: not implemented"; return 0 }

func getBetaDist(fXin, fAlpha, fBeta float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) prepareBETAdotDISTArgs(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BETAdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BETADIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func d1mach(i int) float64 { _ = "STUB: not implemented"; return 0 }

func chebyshevInit(nos int, eta float64, dos []float64) int { _ = "STUB: not implemented"; return 0 }

func chebyshevEval(n int, x float64, a []float64) float64 { _ = "STUB: not implemented"; return 0 }

func lgammacor(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func logrelerr(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func logBeta(a, b float64) float64 { _ = "STUB: not implemented"; return 0 }

func pbetaRaw(alnsml, ans, eps, p, pin, q, sml, x, y float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func pbeta(x, pin, qin float64) (ans float64) { _ = "STUB: not implemented"; return 0 }

func betainvProbIterator(alpha1, alpha3, beta1, beta2, beta3, logBeta, maxCumulative, prob1, prob2 float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcBetainv(probability, alpha, beta, lower, upper float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) betainv(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BETAINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BETAdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func incompleteGamma(a, x float64) float64 { _ = "STUB: not implemented"; return 0 }

func binomCoeff(n, k float64) float64 { _ = "STUB: not implemented"; return 0 }

func binomdist(x, n, p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) BINOMdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BINOMDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) BINOMdotDISTdotRANGE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func binominv(n, p, alpha float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) BINOMdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHIDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHIINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHITEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func getGammaSeries(fA, fX float64) float64 { _ = "STUB: not implemented"; return 0 }

func getGammaContFraction(fA, fX float64) float64 { _ = "STUB: not implemented"; return 0 }

func getLogGammaHelper(fZ float64) float64 { _ = "STUB: not implemented"; return 0 }

func getGammaHelper(fZ float64) float64 { _ = "STUB: not implemented"; return 0 }

func getLogGamma(fZ float64) float64 { _ = "STUB: not implemented"; return 0 }

func getLowRegIGamma(fA, fX float64) float64 { _ = "STUB: not implemented"; return 0 }

func getChiSqDistCDF(fX, fDF float64) float64 { _ = "STUB: not implemented"; return 0 }

func getChiSqDistPDF(fX, fDF float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) CHISQdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHISQdotDISTdotRT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHISQdotTEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func hasChangeOfSign(u, w float64) bool { _ = "STUB: not implemented"; return false }

type calcInverseIterator struct {
	name        string
	fp, fDF, nT float64
}

func (iterator *calcInverseIterator) callBack(x float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func inverseQuadraticInterpolation(iterator calcInverseIterator, fAx, fAy, fBx, fBy float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcIterateInverse(iterator calcInverseIterator, fAx, fBx float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) CHISQdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHISQdotINVdotRT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) confidence(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CONFIDENCE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CONFIDENCEdotNORM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CONFIDENCEdotT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) covar(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COVAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COVARIANCEdotP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COVARIANCEdotS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcStringCountSum(countText bool, count, sum float64, num, arg formulaArg) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (fn *formulaFuncs) countSum(countText bool, args []formulaArg) (count, sum float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (fn *formulaFuncs) CORREL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUNT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUNTA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUNTBLANK(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUNTIF(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func formulaIfsMatch(args []formulaArg) (cellRefs []cellRef) { _ = "STUB: not implemented"; return nil }

func (fn *formulaFuncs) COUNTIFS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CRITBINOM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DEVSQ(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FISHER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FISHERINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FORECAST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FORECASTdotLINEAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

type sortedItem struct {
	idx    int
	number float64
}

func matrixToSortedColumnList(arg formulaArg) ([]sortedItem, *formulaArg) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (fn *formulaFuncs) FREQUENCY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GAMMA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GAMMAdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GAMMADIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func gammainv(probability, alpha, beta float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) GAMMAdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GAMMAINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GAMMALN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GAMMALNdotPRECISE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GAUSS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GEOMEAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func getNewMatrix(c, r int) (matrix [][]float64) { _ = "STUB: not implemented"; return nil }

func approxSub(a, b float64) float64 { _ = "STUB: not implemented"; return 0 }

func matrixClone(matrix [][]float64) (cloneMatrix [][]float64) {
	_ = "STUB: not implemented"
	return nil
}

type trendGrowthMatrixInfo struct {
	trendType, nCX, nCY, nRX, nRY, M, N int
	mtxX, mtxY                          [][]float64
}

func prepareTrendGrowthMtxX(mtxX [][]float64) [][]float64 { _ = "STUB: not implemented"; return nil }

func prepareTrendGrowthMtxY(bLOG bool, mtxY [][]float64) [][]float64 {
	_ = "STUB: not implemented"
	return nil
}

func prepareTrendGrowth(bLOG bool, mtxX, mtxY [][]float64) (*trendGrowthMatrixInfo, formulaArg) {
	_ = "STUB: not implemented"
	return nil, *new(formulaArg)
}

func calcPosition(mtx [][]float64, idx int) (row, col int) { _ = "STUB: not implemented"; return 0, 0 }

func getDouble(mtx [][]float64, idx int) float64 { _ = "STUB: not implemented"; return 0 }

func putDouble(mtx [][]float64, idx int, val float64) { _ = "STUB: not implemented"; return }

func calcMeanOverAll(mtx [][]float64, n int) float64 { _ = "STUB: not implemented"; return 0 }

func calcSumProduct(mtxA, mtxB [][]float64, m int) float64 { _ = "STUB: not implemented"; return 0 }

func calcColumnMeans(mtxX, mtxRes [][]float64, c, r int) { _ = "STUB: not implemented"; return }

func calcColumnsDelta(mtx, columnMeans [][]float64, c, r int) { _ = "STUB: not implemented"; return }

func calcSign(val float64) float64 { _ = "STUB: not implemented"; return 0 }

func calcColsMaximumNorm(mtxA [][]float64, c, r, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcFastMult(mtxA, mtxB, mtxR [][]float64, n, m, l int) { _ = "STUB: not implemented"; return }

func calcRowsEuclideanNorm(mtxA [][]float64, c, r, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcRowsSumProduct(mtxA [][]float64, a int, mtxB [][]float64, b, r, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcSolveWithUpperRightTriangle(mtxA [][]float64, vecR []float64, mtxS [][]float64, k int, bIsTransposed bool) {
	_ = "STUB: not implemented"
	return
}

func calcRowQRDecomposition(mtxA [][]float64, vecR []float64, k, n int) bool {
	_ = "STUB: not implemented"
	return false
}

func calcApplyColsHouseholderTransformation(mtxA [][]float64, r int, mtxY [][]float64, n int) {
	_ = "STUB: not implemented"
	return
}

func calcRowMeans(mtxX, mtxRes [][]float64, c, r int) { _ = "STUB: not implemented"; return }

func calcRowsDelta(mtx, rowMeans [][]float64, c, r int) { _ = "STUB: not implemented"; return }

func calcColumnMaximumNorm(mtxA [][]float64, r, c, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcColsEuclideanNorm(mtxA [][]float64, r, c, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcColsSumProduct(mtxA [][]float64, a int, mtxB [][]float64, b, c, n int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func calcColQRDecomposition(mtxA [][]float64, vecR []float64, k, n int) bool {
	_ = "STUB: not implemented"
	return false
}

func calcApplyRowsHouseholderTransformation(mtxA [][]float64, c int, mtxY [][]float64, n int) {
	_ = "STUB: not implemented"
	return
}

func calcTrendGrowthSimpleRegression(bConstant, bGrowth bool, mtxY, mtxX, newX, mtxRes [][]float64, meanY float64, N int) {
	_ = "STUB: not implemented"
	return
}

func calcTrendGrowthMultipleRegressionPart1(bConstant, bGrowth bool, mtxY, mtxX, newX, mtxRes [][]float64, meanY float64, RXN, K, N int) {
	_ = "STUB: not implemented"
	return
}

func calcTrendGrowthMultipleRegressionPart2(bConstant, bGrowth bool, mtxY, mtxX, newX, mtxRes [][]float64, meanY float64, nCXN, K, N int) {
	_ = "STUB: not implemented"
	return
}

func calcTrendGrowthRegression(bConstant, bGrowth bool, trendType, nCXN, nRXN, K, N int, mtxY, mtxX, newX, mtxRes [][]float64) {
	_ = "STUB: not implemented"
	return
}

func calcTrendGrowth(mtxY, mtxX, newX [][]float64, bConstant, bGrowth bool) ([][]float64, formulaArg) {
	_ = "STUB: not implemented"
	return nil, *new(formulaArg)
}

func (fn *formulaFuncs) trendGrowth(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) GROWTH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HARMEAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func checkHYPGEOMDISTArgs(sampleS, numberSample, populationS, numberPop formulaArg) bool {
	_ = "STUB: not implemented"
	return false
}

func (fn *formulaFuncs) prepareHYPGEOMDISTArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HYPGEOMdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HYPGEOMDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) INTERCEPT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) KURT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EXPONdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EXPONDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FdotDISTdotRT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareFinvArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FdotINVdotRT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FdotTEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FTEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOGINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOGNORMdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOGNORMdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOGNORMDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MODE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MODEdotMULT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MODEdotSNGL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NEGBINOMdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NEGBINOMDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMdotSdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMSDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMSINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NORMdotSdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func norminv(p float64) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

func (fn *formulaFuncs) kth(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LARGE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MAX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MAXA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MAXIFS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcListMatrixMax(maxa bool, maxVal float64, arg formulaArg) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) maxValue(maxa bool, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MEDIAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MINA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MINIFS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcListMatrixMin(mina bool, minVal float64, arg formulaArg) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) minValue(mina bool, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) pearsonProduct(name string, n int, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PEARSON(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERCENTILEdotEXC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERCENTILEdotINC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERCENTILE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) percentrank(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERCENTRANKdotEXC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERCENTRANKdotINC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERCENTRANK(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERMUT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PERMUTATIONA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PHI(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) QUARTILE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) QUARTILEdotEXC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) QUARTILEdotINC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) rank(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RANKdotEQ(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RANK(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RSQ(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) skew(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SKEW(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SKEWdotP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SLOPE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SMALL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STANDARDIZE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) stdevp(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STDEVP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STDEVdotP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STDEVPA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) STEYX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func getTDist(T, fDF, nType float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) TdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TdotDISTdot2T(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TdotDISTdotRT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TdotINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TdotINVdot2T(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TINV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TREND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func tTest(bTemplin bool, mtx1, mtx2 [][]formulaArg, c1, c2, r1, r2 int) (float64, float64, bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

func (fn *formulaFuncs) tTest(mtx1, mtx2 [][]formulaArg, fTails, fTyp float64) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TTEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TdotTEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TRIMMEAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) vars(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VARA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VARP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VARdotP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VARdotS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VARPA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) WEIBULL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) WEIBULLdotDIST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ZdotTEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ZTEST(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ERRORdotTYPE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISBLANK(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISERR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISERROR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISEVEN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISFORMULA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISLOGICAL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISNA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISNONTEXT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISNUMBER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISODD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISREF(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISTEXT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) N(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SHEET(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SHEETS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TYPE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) T(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FALSE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IFERROR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IFNA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IFS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NOT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) OR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SWITCH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TRUE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcXor(argsList *list.List) formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func (fn *formulaFuncs) XOR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DATE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcDateDif(unit string, diff float64, seq []int, startArg, endArg formulaArg) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) DATEDIF(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func isDateOnlyFmt(dateString string) bool { _ = "STUB: not implemented"; return false }

func isTimeOnlyFmt(timeString string) bool { _ = "STUB: not implemented"; return false }

func strToTimePatternHandler1(subMatch []string) (h, m int, s float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func strToTimePatternHandler2(subMatch []string) (h, m int, s float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func strToTimePatternHandler3(subMatch []string) (h, m int, s float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func strToTimePatternHandler4(subMatch []string) (h, m int, s float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func strToTime(str string) (int, int, float64, bool, bool, formulaArg) {
	_ = "STUB: not implemented"
	return 0, 0, 0, false, false, *new(formulaArg)
}

func strToDatePatternHandler1(subMatch []string) (int, int, int, bool, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, false, nil
}

func strToDatePatternHandler2(subMatch []string) (int, int, int, bool, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, false, nil
}

func strToDatePatternHandler3(subMatch []string) (int, int, int, bool, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, false, nil
}

func strToDatePatternHandler4(subMatch []string) (int, int, int, bool, error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, false, nil
}

func strToDate(str string) (int, int, int, bool, formulaArg) {
	_ = "STUB: not implemented"
	return 0, 0, 0, false, *new(formulaArg)
}

func (fn *formulaFuncs) DATEVALUE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DAY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DAYS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DAYS360(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISOWEEKNUM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EDATE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EOMONTH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HOUR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MINUTE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MONTH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func genWeekendMask(weekend int) []byte { _ = "STUB: not implemented"; return nil }

func isWorkday(weekendMask []byte, date float64) bool { _ = "STUB: not implemented"; return false }

func prepareWorkday(weekend formulaArg) ([]byte, int) { _ = "STUB: not implemented"; return nil, 0 }

func toExcelDateArg(arg formulaArg) formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func prepareHolidays(args formulaArg) []int { _ = "STUB: not implemented"; return nil }

func workdayIntl(endDate, sign int, holidays []int, weekendMask []byte, startDate float64) int {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) NETWORKDAYS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NETWORKDAYSdotINTL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) WORKDAY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) WORKDAYdotINTL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) YEAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func yearFracBasisCond(sy, sm, sd, ey, em, ed int) bool { _ = "STUB: not implemented"; return false }

func yearFracBasis0(startDate, endDate float64) (dayDiff, daysInYear float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func yearFracBasis1(startDate, endDate float64) (dayDiff, daysInYear float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func yearFracBasis4(startDate, endDate float64) (dayDiff, daysInYear float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func yearFrac(startDate, endDate float64, basis int) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func getYearDays(year, basis int) int { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) YEARFRAC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NOW(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SECOND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TIME(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TIMEVALUE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TODAY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func makeDate(y int, m time.Month, d int) int64 { _ = "STUB: not implemented"; return 0 }

func daysBetween(startDate, endDate int64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) WEEKDAY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) weeknum(snTime time.Time, returnType int) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) WEEKNUM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func prepareToText(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ARRAYTOTEXT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func bahttextAppendDigit(text string, digit int) string { _ = "STUB: not implemented"; return "" }

func bahttextAppendPow10(text string, digit, pow10 int) string {
	_ = "STUB: not implemented"
	return ""
}

func bahttextAppendBlock(text string, val int) string { _ = "STUB: not implemented"; return "" }

func (fn *formulaFuncs) BAHTTEXT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CLEAN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CODE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) code(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CONCAT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CONCATENATE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) concat(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DBCS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EXACT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FIXED(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FIND(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FINDB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareFindArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) find(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LEFT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LEFTB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) leftRight(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LEN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LENB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOWER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MID(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MIDB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) mid(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PROPER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) REPLACE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) REPLACEB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) replace(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) REPT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RIGHT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RIGHTB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SEARCH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SEARCHB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SUBSTITUTE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TEXT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareTextAfterBefore(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func textAfterBeforeSearch(text string, delimiter []string, startPos int, reverseSearch bool) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func textAfterBeforeResult(name, modifiedDelimiter string, text []rune, foundIdx, repeatZero, textLen int, matchEndActive, matchEnd, reverseSearch bool) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) textAfterBefore(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TEXTAFTER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TEXTBEFORE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TEXTJOIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func textJoin(arg *list.Element, arr []string, ignoreEmpty bool) ([]string, formulaArg) {
	_ = "STUB: not implemented"
	return nil, *new(formulaArg)
}

func (fn *formulaFuncs) TRIM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) UNICHAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) UNICODE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) UNIQUE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func transposeFormulaArgsMatrix(args [][]formulaArg) [][]formulaArg {
	_ = "STUB: not implemented"
	return nil
}

func transposeFormulaArgsList(args []formulaArg, cols, rows int) ([]formulaArg, int, int) {
	_ = "STUB: not implemented"
	return nil, 0, 0
}

func concatValues(args []formulaArg) string { _ = "STUB: not implemented"; return "" }

type uniqueArgs struct {
	cellRange   []formulaArg
	cols        int
	rows        int
	byColumn    bool
	exactlyOnce bool
}

func getFormulaUniqueArgs(argsList *list.List) (uniqueArgs, *formulaArg) {
	_ = "STUB: not implemented"
	return *new(uniqueArgs), nil
}

func (fn *formulaFuncs) UPPER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VALUE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VALUETOTEXT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IF(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ADDRESS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ANCHORARRAY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CHOOSE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func matchPatternToRegExp(findText string, dbcs bool) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func matchPattern(findText, withinText string, dbcs bool, startNum int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func compareFormulaArg(lhs, rhs, matchMode formulaArg, caseSensitive bool) byte {
	_ = "STUB: not implemented"
	return 0
}

func compareFormulaArgList(lhs, rhs, matchMode formulaArg, caseSensitive bool) byte {
	_ = "STUB: not implemented"
	return 0
}

func compareFormulaArgMatrix(lhs, rhs, matchMode formulaArg, caseSensitive bool) byte {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) COLUMN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcColsRowsMinMax(cols bool, argsList *list.List) (minVal, maxVal int) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (fn *formulaFuncs) COLUMNS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FORMULATEXT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func checkHVLookupArgs(name string, argsList *list.List) (idx int, lookupValue, tableArray, matchMode, errArg formulaArg) {
	_ = "STUB: not implemented"
	return 0, *new(formulaArg), *new(formulaArg), *new(formulaArg), *new(formulaArg)
}

func (fn *formulaFuncs) HLOOKUP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) HYPERLINK(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcMatchMatrix(vertical bool, matchType int, criteria *formulaCriteria, lookupArray [][]formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcMatch(matchType int, criteria *formulaCriteria, lookupArray []formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MATCH(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TRANSPOSE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func lookupLinearSearch(vertical bool, lookupValue, lookupArray, matchMode, searchMode formulaArg) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (fn *formulaFuncs) VLOOKUP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func lookupBinarySearch(vertical bool, lookupValue, lookupArray, matchMode, searchMode formulaArg) (matchIdx int, wasExact bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func checkLookupArgs(argsList *list.List) (arrayForm bool, lookupValue, lookupVector, errArg formulaArg) {
	_ = "STUB: not implemented"
	return false, *new(formulaArg), *new(formulaArg), *new(formulaArg)
}

func iterateLookupArgs(lookupValue, lookupVector formulaArg) ([]formulaArg, int, bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

func (fn *formulaFuncs) index(array formulaArg, rowIdx, colIdx int) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func validateMatchMode(mode float64) bool { _ = "STUB: not implemented"; return false }

func validateSearchMode(mode float64) bool { _ = "STUB: not implemented"; return false }

func (fn *formulaFuncs) prepareXlookupArgs(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) xlookup(lookupRows, lookupCols, returnArrayRows, returnArrayCols, matchIdx int,
	condition1, condition2, condition3, condition4 bool, returnArray formulaArg,
) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) XLOOKUP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) INDEX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) INDIRECT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) LOOKUP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func lookupCol(arr formulaArg, idx int) []formulaArg { _ = "STUB: not implemented"; return nil }

func (fn *formulaFuncs) ROW(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ROWS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ENCODEURL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func validateFrequency(freq float64) bool { _ = "STUB: not implemented"; return false }

func (fn *formulaFuncs) ACCRINT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ACCRINTM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareAmorArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AMORDEGRC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) AMORLINC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareCouponArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func is30BasisMethod(basis int) bool { _ = "STUB: not implemented"; return false }

func getDaysInMonthRange(fromMonth, toMonth int) int { _ = "STUB: not implemented"; return 0 }

func getDayOnBasis(y, m, d, basis int) int { _ = "STUB: not implemented"; return 0 }

func coupdays(from, to time.Time, basis int) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) COUPDAYBS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUPDAYS(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUPDAYSNC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) coupons(name string, arg formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUPNCD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUPNUM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) COUPPCD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CUMIPMT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) CUMPRINC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) cumip(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcDbArgsCompare(cost, salvage, life, period formulaArg) bool {
	_ = "STUB: not implemented"
	return false
}

func (fn *formulaFuncs) DB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DDB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareDataValueArgs(n int, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) discIntrate(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DISC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DOLLAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DOLLARDE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DOLLARFR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) dollar(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareDurationArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) duration(settlement, maturity, coupon, yld, frequency, basis formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DURATION(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EFFECT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) EUROCONVERT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) FVSCHEDULE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) INTRATE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IPMT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func calcIpmt(name string, typ, per, pmt, pv, rate formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ipmt(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) IRR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ISPMT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MDURATION(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) MIRR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NOMINAL(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NPER(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) NPV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func aggrBetween(startPeriod, endPeriod float64, initialValue []float64, f func(acc []float64, index float64) []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func fold(f func(acc []float64, index float64) []float64, state []float64, source []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func changeMonth(date time.Time, numMonths float64, returnLastMonth bool) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func datesAggregate(startDate, endDate time.Time, numMonths float64, f func(pcd, ncd time.Time) float64, acc float64, returnLastMonth bool) (time.Time, time.Time, float64) {
	_ = "STUB: not implemented"
	return *new(time.Time), *new(time.Time), 0
}

func coupNumber(maturity, settlement, numMonths float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func prepareOddYldOrPrArg(name string, arg formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareOddfArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ODDFPRICE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func getODDFPRICE(f func(yld float64) float64, x, cnt, prec float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (fn *formulaFuncs) ODDFYIELD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareOddlArgs(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) oddl(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ODDLPRICE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) ODDLYIELD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PDURATION(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PMT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PPMT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) price(settlement, maturity, rate, yld, redemption, frequency, basis formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func checkPriceYieldArgs(name string, rate, prYld, redemption, frequency formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) priceYield(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PRICE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PRICEDISC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PRICEMAT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) PV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) rate(nper, pmt, pv, fv, t, guess formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RATE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RECEIVED(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) RRI(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SLN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) SYD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TBILLEQ(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TBILLPRICE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) TBILLYIELD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareVdbArgs(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) vdb(cost, salvage, life, life1, period, factor formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) VDB(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) prepareXArgs(values, dates formulaArg) (valuesArg, datesArg []float64, err formulaArg) {
	_ = "STUB: not implemented"
	return nil, nil, *new(formulaArg)
}

func (fn *formulaFuncs) xirr(values, dates []float64, guess float64) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func xirrPart1(values, dates []float64, rate float64) float64 { _ = "STUB: not implemented"; return 0 }

func xirrPart2(values, dates []float64, rate float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fn *formulaFuncs) XIRR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) XNPV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) yield(settlement, maturity, rate, pr, redemption, frequency, basis formulaArg) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) YIELD(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) YIELDDISC(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) YIELDMAT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

type calcDatabase struct {
	col, row int
	indexMap map[int]int
	database [][]formulaArg
	criteria [][]formulaArg
}

func newCalcDatabase(database, field, criteria formulaArg) *calcDatabase {
	_ = "STUB: not implemented"
	return nil
}

func (db *calcDatabase) columnIndex(database [][]formulaArg, field formulaArg) int {
	_ = "STUB: not implemented"
	return 0
}

func (db *calcDatabase) criteriaEval() bool { _ = "STUB: not implemented"; return false }

func (db *calcDatabase) value() formulaArg { _ = "STUB: not implemented"; return *new(formulaArg) }

func (db *calcDatabase) next() bool { _ = "STUB: not implemented"; return false }

func (fn *formulaFuncs) database(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DAVERAGE(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) dcount(name string, argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DCOUNT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DCOUNTA(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DGET(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DMAX(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DMIN(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DPRODUCT(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DSTDEV(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DSTDEVP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DSUM(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DVAR(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DVARP(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func (fn *formulaFuncs) DISPIMG(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

type sortbyArgs struct {
	array    []formulaArg
	cols     int
	rows     int
	sortKeys []sortbyKey
}

type sortbyKey struct {
	byArray   []formulaArg
	cols      int
	rows      int
	ascending bool
}

type rowWithKeys struct {
	rowData  []formulaArg
	sortKeys [][]formulaArg
}

func (fn *formulaFuncs) SORTBY(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func checkSortbyArgs(argsList *list.List) formulaArg {
	_ = "STUB: not implemented"
	return *new(formulaArg)
}

func prepareSortbyArgs(argsList *list.List) (sortbyArgs, *formulaArg) {
	_ = "STUB: not implemented"
	return *new(sortbyArgs), nil
}

func parseSortOrderArg(byArray *list.Element, key *sortbyKey, keyCount, argsLen int) (*list.Element, *formulaArg) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compareRowsForSortby(i, j rowWithKeys, sortKeys []sortbyKey) bool {
	_ = "STUB: not implemented"
	return false
}
