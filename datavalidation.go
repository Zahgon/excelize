package excelize

import (
	"strings"
)

type DataValidationType byte

const (
	_ DataValidationType = iota
	DataValidationTypeNone
	DataValidationTypeCustom
	DataValidationTypeDate
	DataValidationTypeDecimal
	DataValidationTypeList
	DataValidationTypeTextLength
	DataValidationTypeTime
	DataValidationTypeWhole
)

type DataValidationErrorStyle byte

const (
	_ DataValidationErrorStyle = iota
	DataValidationErrorStyleStop
	DataValidationErrorStyleWarning
	DataValidationErrorStyleInformation
)

const (
	styleStop        = "stop"
	styleWarning     = "warning"
	styleInformation = "information"
)

type DataValidationOperator byte

const (
	_ DataValidationOperator = iota
	DataValidationOperatorBetween
	DataValidationOperatorEqual
	DataValidationOperatorGreaterThan
	DataValidationOperatorGreaterThanOrEqual
	DataValidationOperatorLessThan
	DataValidationOperatorLessThanOrEqual
	DataValidationOperatorNotBetween
	DataValidationOperatorNotEqual
)

var (
	formulaEscaper = strings.NewReplacer(
		`&`, `&amp;`,
		`<`, `&lt;`,
		`>`, `&gt;`,
	)
	formulaUnescaper = strings.NewReplacer(
		`&amp;`, `&`,
		`&lt;`, `<`,
		`&gt;`, `>`,
	)

	dataValidationTypeMap = map[DataValidationType]string{
		DataValidationTypeNone:       "none",
		DataValidationTypeCustom:     "custom",
		DataValidationTypeDate:       "date",
		DataValidationTypeDecimal:    "decimal",
		DataValidationTypeList:       "list",
		DataValidationTypeTextLength: "textLength",
		DataValidationTypeTime:       "time",
		DataValidationTypeWhole:      "whole",
	}

	dataValidationOperatorMap = map[DataValidationOperator]string{
		DataValidationOperatorBetween:            "between",
		DataValidationOperatorEqual:              "equal",
		DataValidationOperatorGreaterThan:        "greaterThan",
		DataValidationOperatorGreaterThanOrEqual: "greaterThanOrEqual",
		DataValidationOperatorLessThan:           "lessThan",
		DataValidationOperatorLessThanOrEqual:    "lessThanOrEqual",
		DataValidationOperatorNotBetween:         "notBetween",
		DataValidationOperatorNotEqual:           "notEqual",
	}
)

func NewDataValidation(allowBlank bool) *DataValidation { _ = "STUB: not implemented"; return nil }

func (dv *DataValidation) SetError(style DataValidationErrorStyle, title, msg string) {
	_ = "STUB: not implemented"
	return
}

func (dv *DataValidation) SetInput(title, msg string) { _ = "STUB: not implemented"; return }

func (dv *DataValidation) SetDropList(keys []string) error { _ = "STUB: not implemented"; return nil }

func (dv *DataValidation) SetRange(f1, f2 interface{}, t DataValidationType, o DataValidationOperator) error {
	_ = "STUB: not implemented"
	return nil
}

func (dv *DataValidation) SetSqrefDropList(sqref string) { _ = "STUB: not implemented"; return }

func (dv *DataValidation) SetSqref(sqref string) { _ = "STUB: not implemented"; return }

func (f *File) AddDataValidation(sheet string, dv *DataValidation) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetDataValidations(sheet string) ([]*DataValidation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDataValidations(dvs *xlsxDataValidations) []*DataValidation {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) DeleteDataValidation(sheet string, sqref ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) deleteDataValidation(ws *xlsxWorksheet, delCells map[int][][]int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) deleteX14DataValidation(ws *xlsxWorksheet, delCells map[int][][]int) error {
	_ = "STUB: not implemented"
	return nil
}

func squashSqref(cells [][]int) []string { _ = "STUB: not implemented"; return nil }

func deleteCellsFromSqref(sqref string, delCells map[int][][]int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (dv *xlsxInnerXML) isFormula() bool { _ = "STUB: not implemented"; return false }

func unescapeDataValidationFormula(val string) string { _ = "STUB: not implemented"; return "" }
