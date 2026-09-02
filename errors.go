package excelize

import (
	"errors"
	"fmt"
)

var (
	ErrAddVBAProject = errors.New("unsupported VBA project")

	ErrAttrValBool = errors.New("unexpected child of attrValBool")

	ErrCellCharsLength = fmt.Errorf("cell value must be 0-%d characters", TotalCellChars)

	ErrCellStyles = fmt.Errorf("the cell styles exceeds the %d limit", MaxCellStyles)

	ErrChartTitle = errors.New("cannot set both 'Formula' and 'Paragraph' for chart title")

	ErrColumnNumber = fmt.Errorf("the column number must be greater than or equal to %d and less than or equal to %d", MinColumns, MaxColumns)

	ErrColumnWidth = fmt.Errorf("the width of the column must be less than or equal to %d characters", MaxColumnWidth)

	ErrCoordinates = errors.New("coordinates length must be 4")

	ErrCustomNumFmt = errors.New("custom number format can not be empty")

	ErrDataValidationFormulaLength = fmt.Errorf("data validation must be 0-%d characters", MaxFieldLength)

	ErrDataValidationRange = errors.New("data validation range exceeds limit")

	ErrDefinedNameDuplicate = errors.New("the same name already exists on the scope")

	ErrDefinedNameScope = errors.New("no defined name on the scope")

	ErrExistsSheet = errors.New("the same name sheet already exists")

	ErrExistsTableName = errors.New("the same name table already exists")

	ErrFillType = errors.New("fill type value must be one of 'gradient' or 'pattern'")

	ErrFillGradientColor = errors.New("fill color value must be an array of two colors for 'gradient' type")

	ErrFillGradientShading = errors.New("fill shading value must be between 0 and 16 for 'gradient' type")

	ErrFillPatternColor = errors.New("fill color value must be empty or an array of one color for 'pattern' type")

	ErrFillPattern = errors.New("fill pattern value must be between 0 and 18")

	ErrFontLength = fmt.Errorf("the length of the font family name must be less than or equal to %d", MaxFontFamilyLength)

	ErrFontSize = fmt.Errorf("font size must be an integer from %d to %d points", MinFontSize, MaxFontSize)

	ErrFormControlValue = fmt.Errorf("scroll value must be an integer from 0 to %d", MaxFormControlValue)

	ErrGroupSheets = errors.New("group worksheet must contain an active worksheet")

	ErrImgExt = errors.New("unsupported image extension")

	ErrInvalidFormula = errors.New("formula not valid")

	ErrMaxFilePathLength = fmt.Errorf("file path length exceeds maximum limit %d characters", MaxFilePathLength)

	ErrMaxRowHeight = fmt.Errorf("the height of the row must be less than or equal to %d points", MaxRowHeight)

	ErrMaxRows = errors.New("row number exceeds maximum limit")

	ErrNameLength = fmt.Errorf("the name length exceeds the %d characters limit", MaxFieldLength)

	ErrMaxGraphicAltTextLength = fmt.Errorf("the alt text length exceeds the %d characters limit", MaxGraphicAltTextLength)

	ErrMaxGraphicNameLength = fmt.Errorf("the name length exceeds the %d characters limit", MaxGraphicNameLength)

	ErrOptionsUnzipSizeLimit = errors.New("the value of UnzipSizeLimit should be greater than or equal to UnzipXMLSizeLimit")

	ErrOutlineLevel = errors.New("invalid outline level")

	ErrPageSetupAdjustTo = errors.New("adjust to value must be an integer from 0 to 400")

	ErrParameterInvalid = errors.New("parameter is invalid")

	ErrParameterRequired = errors.New("parameter is required")

	ErrPasswordLengthInvalid = errors.New("password length invalid")

	ErrPivotTableShowValuesAsBaseField = errors.New("this kind of show values as type requires a base field")

	ErrPivotTableShowValuesAsBaseItem = errors.New("this kind of show values as type and base field requires a base item")

	ErrPivotTableClassicLayout = errors.New("cannot enable ClassicLayout and CompactData in the same time")

	ErrSave = errors.New("no path defined for file, consider File.WriteTo or File.Write")

	ErrSheetIdx = errors.New("invalid worksheet index")

	ErrSheetNameBlank = errors.New("the sheet name can not be blank")

	ErrSheetNameInvalid = errors.New("the sheet can not contain any of the characters :\\/?*[or]")

	ErrSheetNameLength = fmt.Errorf("the sheet name length exceeds the %d characters limit", MaxSheetNameLength)

	ErrSheetNameSingleQuote = errors.New("the first or last character of the sheet name can not be a single quote")

	ErrSparkline = errors.New("must have the same number of 'Location' and 'Range' parameters")

	ErrSparklineLocation = errors.New("parameter 'Location' is required")

	ErrSparklineRange = errors.New("parameter 'Range' is required")

	ErrSparklineStyle = errors.New("parameter 'Style' value must be an integer from 0 to 35")

	ErrSparklineType = errors.New("parameter 'Type' value must be one of 'line', 'column' or 'win_loss'")

	ErrTotalSheetHyperlinks = errors.New("over maximum limit hyperlinks in a worksheet")

	ErrTransparency = errors.New("transparency value must be an integer from 0 to 100")

	ErrUnknownEncryptMechanism = errors.New("unknown encryption mechanism")

	ErrUnprotectSheet = errors.New("worksheet has set no protect")

	ErrUnprotectSheetPassword = errors.New("worksheet protect password not match")

	ErrUnprotectWorkbook = errors.New("workbook has set no protect")

	ErrUnprotectWorkbookPassword = errors.New("workbook protect password not match")

	ErrUnsupportedEncryptMechanism = errors.New("unsupported encryption mechanism")

	ErrUnsupportedHashAlgorithm = errors.New("unsupported hash algorithm")

	ErrUnsupportedNumberFormat = errors.New("unsupported number format token")

	ErrUnsupportedPivotTableShowValuesAsType = errors.New("unsupported pivot table show values as type")

	ErrWorkbookFileFormat = errors.New("unsupported workbook file format")

	ErrWorkbookPassword = errors.New("the supplied open workbook password is not correct")
)

type ErrSheetNotExist struct {
	SheetName string
}

func (err ErrSheetNotExist) Error() string { _ = "STUB: not implemented"; return "" }

func newAddCommentError(cell string) error { _ = "STUB: not implemented"; return nil }

func newCellNameToCoordinatesError(cell string, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func newChartTitleError(name string) error { _ = "STUB: not implemented"; return nil }

func newCoordinatesToCellNameError(col, row int) error { _ = "STUB: not implemented"; return nil }

func newFieldLengthError(name string) error { _ = "STUB: not implemented"; return nil }

func newInvalidAutoFilterColumnError(col string) error { _ = "STUB: not implemented"; return nil }

func newInvalidAutoFilterExpError(exp string) error { _ = "STUB: not implemented"; return nil }

func newInvalidAutoFilterOperatorError(op, exp string) error { _ = "STUB: not implemented"; return nil }

func newInvalidCellNameError(cell string) error { _ = "STUB: not implemented"; return nil }

func newInvalidColumnNameError(col string) error { _ = "STUB: not implemented"; return nil }

func newInvalidExcelDateError(dateValue float64) error { _ = "STUB: not implemented"; return nil }

func newInvalidLinkTypeError(linkType string) error { _ = "STUB: not implemented"; return nil }

func newInvalidNameError(name string) error { _ = "STUB: not implemented"; return nil }

func newInvalidOptionalValue(name, value string, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

func newInvalidRowNumberError(row int) error { _ = "STUB: not implemented"; return nil }

func newInvalidSharedStringIndex(idx int) error { _ = "STUB: not implemented"; return nil }

func newInvalidSlicerNameError(name string) error { _ = "STUB: not implemented"; return nil }

func newInvalidStyleID(styleID int) error { _ = "STUB: not implemented"; return nil }

func newNoExistSlicerError(name string) error { _ = "STUB: not implemented"; return nil }

func newNoExistTableError(name string) error { _ = "STUB: not implemented"; return nil }

func newNotWorksheetError(name string) error { _ = "STUB: not implemented"; return nil }

func newPivotTableColFieldsError(data []string) error { _ = "STUB: not implemented"; return nil }

func newPivotTableRowFieldsError(data []string) error { _ = "STUB: not implemented"; return nil }

func newPivotTableDataRangeError(msg string) error { _ = "STUB: not implemented"; return nil }

func newPivotTableSelectedItemError(item, field string) error {
	_ = "STUB: not implemented"
	return nil
}

func newPivotTableRangeError(msg string) error { _ = "STUB: not implemented"; return nil }

func newPivotTableShowValuesAsBaseFieldError(field string) error {
	_ = "STUB: not implemented"
	return nil
}

func newStreamSetRowError(row int) error { _ = "STUB: not implemented"; return nil }

func newStreamSetRowOrderError(name string) error { _ = "STUB: not implemented"; return nil }

func newUnknownFilterTokenError(token string) error { _ = "STUB: not implemented"; return nil }

func newUnsupportedChartType(chartType ChartType) error { _ = "STUB: not implemented"; return nil }

func newUnsupportedPivotCacheSourceType(sourceType string) error {
	_ = "STUB: not implemented"
	return nil
}

func newUnzipSizeLimitError(unzipSizeLimit int64) error { _ = "STUB: not implemented"; return nil }

func newViewIdxError(viewIndex int) error { _ = "STUB: not implemented"; return nil }
