package excelize

import (
	"reflect"
)

func (f *File) stylesReader() (*xlsxStyleSheet, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) styleSheetWriter() { _ = "STUB: not implemented"; return }

func (f *File) themeWriter() { _ = "STUB: not implemented"; return }

func (f *File) sharedStringsWriter() { _ = "STUB: not implemented"; return }

func parseFormatStyleSet(style *Style) (*Style, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) NewStyle(style *Style) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var (
	styleBorders = []string{
		"none",
		"thin",
		"medium",
		"dashed",
		"dotted",
		"thick",
		"double",
		"hair",
		"mediumDashed",
		"dashDot",
		"mediumDashDot",
		"dashDotDot",
		"mediumDashDotDot",
		"slantDashDot",
	}

	styleBorderTypes = []string{
		"left", "right", "top", "bottom", "diagonalUp", "diagonalDown",
	}

	styleFillPatterns = []string{
		"none",
		"solid",
		"mediumGray",
		"darkGray",
		"lightGray",
		"darkHorizontal",
		"darkVertical",
		"darkDown",
		"darkUp",
		"darkGrid",
		"darkTrellis",
		"lightHorizontal",
		"lightVertical",
		"lightDown",
		"lightUp",
		"lightGrid",
		"lightTrellis",
		"gray125",
		"gray0625",
	}

	styleFillVariants = func() []xlsxGradientFill {
		return []xlsxGradientFill{
			{Degree: 90, Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Degree: 270, Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Degree: 90, Stop: []*xlsxGradientFillStop{{}, {Position: 0.5}, {Position: 1}}},
			{Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Degree: 180, Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Stop: []*xlsxGradientFillStop{{}, {Position: 0.5}, {Position: 1}}},
			{Degree: 45, Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Degree: 255, Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Degree: 45, Stop: []*xlsxGradientFillStop{{}, {Position: 0.5}, {Position: 1}}},
			{Degree: 135, Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Degree: 315, Stop: []*xlsxGradientFillStop{{}, {Position: 1}}},
			{Degree: 135, Stop: []*xlsxGradientFillStop{{}, {Position: 0.5}, {Position: 1}}},
			{Stop: []*xlsxGradientFillStop{{}, {Position: 1}}, Type: "path"},
			{Stop: []*xlsxGradientFillStop{{}, {Position: 1}}, Type: "path", Left: 1, Right: 1},
			{Stop: []*xlsxGradientFillStop{{}, {Position: 1}}, Type: "path", Bottom: 1, Top: 1},
			{Stop: []*xlsxGradientFillStop{{}, {Position: 1}}, Type: "path", Bottom: 1, Left: 1, Right: 1, Top: 1},
			{Stop: []*xlsxGradientFillStop{{}, {Position: 1}}, Type: "path", Bottom: 0.5, Left: 0.5, Right: 0.5, Top: 0.5},
		}
	}

	getXfIDFuncs = map[string]func(int, xlsxXf, *Style) bool{
		"numFmt": func(numFmtID int, xf xlsxXf, style *Style) bool {
			if style.CustomNumFmt == nil && numFmtID == -1 {
				return xf.NumFmtID != nil && *xf.NumFmtID == 0
			}
			if style.NegRed || (style.DecimalPlaces != nil && *style.DecimalPlaces != 2) {
				return false
			}
			return xf.NumFmtID != nil && *xf.NumFmtID == numFmtID
		},
		"font": func(fontID int, xf xlsxXf, style *Style) bool {
			if style.Font == nil || fontID == 0 {
				return (xf.FontID == nil || *xf.FontID == 0) && (xf.ApplyFont == nil || !*xf.ApplyFont)
			}
			return xf.FontID != nil && *xf.FontID == fontID && xf.ApplyFont != nil && *xf.ApplyFont
		},
		"fill": func(fillID int, xf xlsxXf, style *Style) bool {
			if style.Fill.Type == "" || fillID == 0 {
				return (xf.FillID == nil || *xf.FillID == 0) && (xf.ApplyFill == nil || !*xf.ApplyFill)
			}
			return xf.FillID != nil && *xf.FillID == fillID && xf.ApplyFill != nil && *xf.ApplyFill
		},
		"border": func(borderID int, xf xlsxXf, style *Style) bool {
			if len(style.Border) == 0 {
				return (xf.BorderID == nil || *xf.BorderID == 0) && (xf.ApplyBorder == nil || !*xf.ApplyBorder)
			}
			return xf.BorderID != nil && *xf.BorderID == borderID && xf.ApplyBorder != nil && *xf.ApplyBorder
		},
		"alignment": func(ID int, xf xlsxXf, style *Style) bool {
			if style.Alignment == nil {
				return xf.ApplyAlignment == nil || !*xf.ApplyAlignment
			}
			return reflect.DeepEqual(xf.Alignment, newAlignment(style))
		},
		"protection": func(ID int, xf xlsxXf, style *Style) bool {
			if style.Protection == nil {
				return xf.ApplyProtection == nil || !*xf.ApplyProtection
			}
			return reflect.DeepEqual(xf.Protection, newProtection(style)) && xf.ApplyProtection != nil && *xf.ApplyProtection
		},
	}

	extractStyleCondFuncs = map[string]func(xlsxXf, *xlsxStyleSheet) bool{
		"fill": func(xf xlsxXf, s *xlsxStyleSheet) bool {
			return (xf.ApplyFill == nil || (xf.ApplyFill != nil && *xf.ApplyFill)) &&
				xf.FillID != nil && s.Fills != nil &&
				*xf.FillID < len(s.Fills.Fill)
		},
		"border": func(xf xlsxXf, s *xlsxStyleSheet) bool {
			return (xf.ApplyBorder == nil || (xf.ApplyBorder != nil && *xf.ApplyBorder)) &&
				xf.BorderID != nil && s.Borders != nil &&
				*xf.BorderID < len(s.Borders.Border)
		},
		"font": func(xf xlsxXf, s *xlsxStyleSheet) bool {
			return (xf.ApplyFont == nil || (xf.ApplyFont != nil && *xf.ApplyFont)) &&
				xf.FontID != nil && s.Fonts != nil &&
				*xf.FontID < len(s.Fonts.Font)
		},
		"alignment": func(xf xlsxXf, s *xlsxStyleSheet) bool {
			return xf.ApplyAlignment == nil || (xf.ApplyAlignment != nil && *xf.ApplyAlignment)
		},
		"protection": func(xf xlsxXf, s *xlsxStyleSheet) bool {
			return xf.ApplyProtection == nil || (xf.ApplyProtection != nil && *xf.ApplyProtection)
		},
	}

	drawContFmtFunc = map[string]func(p int, ct, ref, GUID string, fmtCond *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule){
		"cellIs":            drawCondFmtCellIs,
		"timePeriod":        drawCondFmtTimePeriod,
		"text":              drawCondFmtText,
		"top10":             drawCondFmtTop10,
		"aboveAverage":      drawCondFmtAboveAverage,
		"duplicateValues":   drawCondFmtDuplicateUniqueValues,
		"uniqueValues":      drawCondFmtDuplicateUniqueValues,
		"containsBlanks":    drawCondFmtBlanks,
		"notContainsBlanks": drawCondFmtNoBlanks,
		"containsErrors":    drawCondFmtErrors,
		"notContainsErrors": drawCondFmtNoErrors,
		"2_color_scale":     drawCondFmtColorScale,
		"3_color_scale":     drawCondFmtColorScale,
		"dataBar":           drawCondFmtDataBar,
		"expression":        drawCondFmtExp,
		"iconSet":           drawCondFmtIconSet,
	}

	extractContFmtFunc = map[string]func(*File, string, *xlsxCfRule, *xlsxExtLst) ConditionalFormatOptions{
		"cellIs": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtCellIs(c, extLst)
		},
		"timePeriod": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtTimePeriod(c, ref, extLst)
		},
		"containsText": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtText(c, extLst)
		},
		"notContainsText": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtText(c, extLst)
		},
		"beginsWith": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtText(c, extLst)
		},
		"endsWith": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtText(c, extLst)
		},
		"top10": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtTop10(c, extLst)
		},
		"aboveAverage": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtAboveAverage(c, extLst)
		},
		"duplicateValues": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtDuplicateUniqueValues(c, extLst)
		},
		"uniqueValues": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtDuplicateUniqueValues(c, extLst)
		},
		"containsBlanks": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtBlanks(c, extLst)
		},
		"notContainsBlanks": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtNoBlanks(c, extLst)
		},
		"containsErrors": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtErrors(c, extLst)
		},
		"notContainsErrors": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtNoErrors(c, extLst)
		},
		"colorScale": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtColorScale(c, extLst)
		},
		"dataBar": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtDataBar(c, extLst)
		},
		"expression": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtExp(c, extLst)
		},
		"iconSet": func(f *File, ref string, c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
			return f.extractCondFmtIconSet(c, extLst)
		},
	}

	validType = map[string]string{
		"cell":          "cellIs",
		"average":       "aboveAverage",
		"duplicate":     "duplicateValues",
		"unique":        "uniqueValues",
		"top":           "top10",
		"bottom":        "top10",
		"text":          "text",
		"time_period":   "timePeriod",
		"blanks":        "containsBlanks",
		"no_blanks":     "notContainsBlanks",
		"errors":        "containsErrors",
		"no_errors":     "notContainsErrors",
		"2_color_scale": "2_color_scale",
		"3_color_scale": "3_color_scale",
		"data_bar":      "dataBar",
		"formula":       "expression",
		"icon_set":      "iconSet",
	}

	criteriaType = map[string]string{
		"!=":                       "notEqual",
		"<":                        "lessThan",
		"<=":                       "lessThanOrEqual",
		"<>":                       "notEqual",
		"=":                        "equal",
		"==":                       "equal",
		">":                        "greaterThan",
		">=":                       "greaterThanOrEqual",
		"begins with":              "beginsWith",
		"between":                  "between",
		"containing":               "containsText",
		"continue month":           "nextMonth",
		"continue week":            "nextWeek",
		"ends with":                "endsWith",
		"equal to":                 "equal",
		"greater than or equal to": "greaterThanOrEqual",
		"greater than":             "greaterThan",
		"last 7 days":              "last7Days",
		"last month":               "lastMonth",
		"last week":                "lastWeek",
		"less than or equal to":    "lessThanOrEqual",
		"less than":                "lessThan",
		"not between":              "notBetween",
		"not containing":           "notContains",
		"not equal to":             "notEqual",
		"this month":               "thisMonth",
		"this week":                "thisWeek",
		"today":                    "today",
		"tomorrow":                 "tomorrow",
		"yesterday":                "yesterday",
	}

	operatorType = map[string]string{
		"beginsWith":         "begins with",
		"between":            "between",
		"containsText":       "containing",
		"endsWith":           "ends with",
		"equal":              "equal to",
		"greaterThan":        "greater than",
		"greaterThanOrEqual": "greater than or equal to",
		"last7Days":          "last 7 days",
		"lastMonth":          "last month",
		"lastWeek":           "last week",
		"lessThan":           "less than",
		"lessThanOrEqual":    "less than or equal to",
		"nextMonth":          "continue month",
		"nextWeek":           "continue week",
		"notBetween":         "not between",
		"notContains":        "not containing",
		"notEqual":           "not equal to",
		"thisMonth":          "this month",
		"thisWeek":           "this week",
		"today":              "today",
		"tomorrow":           "tomorrow",
		"yesterday":          "yesterday",
	}

	cellIsCriteriaType = []string{
		"equal",
		"notEqual",
		"greaterThan",
		"lessThan",
		"greaterThanOrEqual",
		"lessThanOrEqual",
		"containsText",
		"notContains",
		"beginsWith",
		"endsWith",
	}

	cfvo3 = &xlsxCfRule{IconSet: &xlsxIconSet{Cfvo: []*xlsxCfvo{
		{Type: "percent", Val: "0"},
		{Type: "percent", Val: "33"},
		{Type: "percent", Val: "67"},
	}}}

	cfvo4 = &xlsxCfRule{IconSet: &xlsxIconSet{Cfvo: []*xlsxCfvo{
		{Type: "percent", Val: "0"},
		{Type: "percent", Val: "25"},
		{Type: "percent", Val: "50"},
		{Type: "percent", Val: "75"},
	}}}

	cfvo5 = &xlsxCfRule{IconSet: &xlsxIconSet{Cfvo: []*xlsxCfvo{
		{Type: "percent", Val: "0"},
		{Type: "percent", Val: "20"},
		{Type: "percent", Val: "40"},
		{Type: "percent", Val: "60"},
		{Type: "percent", Val: "80"},
	}}}

	condFmtIconSetPresets = map[string]*xlsxCfRule{
		"3Arrows":         cfvo3,
		"3ArrowsGray":     cfvo3,
		"3Flags":          cfvo3,
		"3Signs":          cfvo3,
		"3Symbols":        cfvo3,
		"3Symbols2":       cfvo3,
		"3TrafficLights1": cfvo3,
		"3TrafficLights2": cfvo3,
		"4Arrows":         cfvo4,
		"4ArrowsGray":     cfvo4,
		"4Rating":         cfvo4,
		"4RedToBlack":     cfvo4,
		"4TrafficLights":  cfvo4,
		"5Arrows":         cfvo5,
		"5ArrowsGray":     cfvo5,
		"5Quarters":       cfvo5,
		"5Rating":         cfvo5,
	}

	x14cfvo3 = &xlsxX14CfRule{IconSet: &xlsx14IconSet{Cfvo: []*xlsx14Cfvo{
		{Type: "percent", F: "0"},
		{Type: "percent", F: "33"},
		{Type: "percent", F: "67"},
	}}}

	x14cfvo5 = &xlsxX14CfRule{IconSet: &xlsx14IconSet{Cfvo: []*xlsx14Cfvo{
		{Type: "percent", F: "0"},
		{Type: "percent", F: "20"},
		{Type: "percent", F: "40"},
		{Type: "percent", F: "60"},
		{Type: "percent", F: "80"},
	}}}

	condFmtX14IconSetPresets = map[string]*xlsxX14CfRule{
		"3Stars":     x14cfvo3,
		"3Triangles": x14cfvo3,
		"5Boxes":     x14cfvo5,
	}
)

func (clr *decodeCTColor) colorChoice() *string { _ = "STUB: not implemented"; return nil }

func (f *File) GetBaseColor(hexColor string, indexedColor int, themeColor *int) string {
	_ = "STUB: not implemented"
	return ""
}

func (f *File) getThemeColor(clr *xlsxColor) string { _ = "STUB: not implemented"; return "" }

func (f *File) extractBorders(bdr *xlsxBorder, s *xlsxStyleSheet, style *Style) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractFills(fl *xlsxFill, s *xlsxStyleSheet, style *Style) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractGradientFill(gf *xlsxGradientFill, fill *Fill) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractPatternFill(pf *xlsxPatternFill, fill *Fill) {
	_ = "STUB: not implemented"
	return
}

func extractFont(fnt *xlsxFont) *Font { _ = "STUB: not implemented"; return nil }

func (f *File) extractNumFmt(n *int, s *xlsxStyleSheet, style *Style) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractAlignment(a *xlsxAlignment, s *xlsxStyleSheet, style *Style) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractProtection(p *xlsxProtection, s *xlsxStyleSheet, style *Style) {
	_ = "STUB: not implemented"
	return
}

func (f *File) GetStyle(idx int) (*Style, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) getStyleID(ss *xlsxStyleSheet, style *Style) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) NewConditionalStyle(style *Style) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) GetConditionalStyle(idx int) (*Style, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newDxfNumFmt(styleSheet *xlsxStyleSheet, style *Style, dxf *xlsxDxf) *xlsxNumFmt {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetDefaultFont() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (f *File) SetDefaultFont(fontName string) error { _ = "STUB: not implemented"; return nil }

func (f *File) readDefaultFont() (*xlsxFont, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *File) getFontID(styleSheet *xlsxStyleSheet, style *Style) int {
	_ = "STUB: not implemented"
	return 0
}

func newFontColor(font *Font) *xlsxColor { _ = "STUB: not implemented"; return nil }

func (fnt *Font) newFont() *xlsxFont { _ = "STUB: not implemented"; return nil }

func getNumFmtID(styleSheet *xlsxStyleSheet, style *Style) int { _ = "STUB: not implemented"; return 0 }

func newNumFmt(styleSheet *xlsxStyleSheet, style *Style) int { _ = "STUB: not implemented"; return 0 }

func setCustomNumFmt(styleSheet *xlsxStyleSheet, style *Style) int {
	_ = "STUB: not implemented"
	return 0
}

func getCustomNumFmtID(styleSheet *xlsxStyleSheet, style *Style) (customNumFmtID int) {
	_ = "STUB: not implemented"
	return 0
}

func isLangNumFmt(ID int) bool { _ = "STUB: not implemented"; return false }

func setLangNumFmt(style *Style) int { _ = "STUB: not implemented"; return 0 }

func getFillID(styleSheet *xlsxStyleSheet, style *Style) (fillID int) {
	_ = "STUB: not implemented"
	return 0
}

func newFills(style *Style, fg bool) *xlsxFill { _ = "STUB: not implemented"; return nil }

func newAlignment(style *Style) *xlsxAlignment { _ = "STUB: not implemented"; return nil }

func newProtection(style *Style) *xlsxProtection { _ = "STUB: not implemented"; return nil }

func getBorderID(styleSheet *xlsxStyleSheet, style *Style) (borderID int) {
	_ = "STUB: not implemented"
	return 0
}

func newBorders(style *Style) *xlsxBorder { _ = "STUB: not implemented"; return nil }

func setCellXfs(style *xlsxStyleSheet, fontID, numFmtID, fillID, borderID int, applyAlignment, applyProtection bool, alignment *xlsxAlignment, protection *xlsxProtection) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) GetCellStyle(sheet, cell string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (f *File) SetCellStyle(sheet, topLeftCell, bottomRightCell string, styleID int) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) SetConditionalFormat(sheet, rangeRef string, opts []ConditionalFormatOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func prepareConditionalFormatRange(rangeRef string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func (f *File) appendCfRule(ws *xlsxWorksheet, rule *xlsxX14CfRule, sqref string) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) extractCondFmtCellIs(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtTimePeriod(c *xlsxCfRule, ref string, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtText(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtTop10(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtAboveAverage(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtDuplicateUniqueValues(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtBlanks(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtNoBlanks(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtErrors(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtNoErrors(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtColorScale(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtDataBarRule(ID string, format *ConditionalFormatOptions, condFmts []decodeX14ConditionalFormatting) {
	_ = "STUB: not implemented"
	return
}

func (f *File) extractCondFmtDataBar(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtExp(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractCondFmtIconSet(c *xlsxCfRule, extLst *xlsxExtLst) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) extractX14CondFmtIconSetRule(c *decodeX14CfRule) ConditionalFormatOptions {
	_ = "STUB: not implemented"
	return *new(ConditionalFormatOptions)
}

func (f *File) GetConditionalFormats(sheet string) (map[string][]ConditionalFormatOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) UnsetConditionalFormat(sheet, rangeRef string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *decodeX14ConditionalFormattingRules) deleteCfRule(delCells map[int][][]int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtCellIs(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtTimePeriod(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtText(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtTop10(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtAboveAverage(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtDuplicateUniqueValues(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtColorScale(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtDataBar(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtExp(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtErrors(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtNoErrors(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtBlanks(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtNoBlanks(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func drawCondFmtIconSet(p int, ct, ref, GUID string, format *ConditionalFormatOptions) (*xlsxCfRule, *xlsxX14CfRule) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getPaletteColor(color string) string { _ = "STUB: not implemented"; return "" }

func (f *File) themeReader() (*decodeTheme, error) { _ = "STUB: not implemented"; return nil, nil }

func ThemeColor(baseColor string, tint float64) string { _ = "STUB: not implemented"; return "" }
