package excelize

type FormControlType byte

const (
	FormControlNote FormControlType = iota
	FormControlButton
	FormControlOptionButton
	FormControlSpinButton
	FormControlCheckBox
	FormControlGroupBox
	FormControlLabel
	FormControlScrollBar
)

type HeaderFooterImagePositionType byte

const (
	HeaderFooterImagePositionLeft HeaderFooterImagePositionType = iota
	HeaderFooterImagePositionCenter
	HeaderFooterImagePositionRight
)

func (f *File) GetComments(sheet string) ([]Comment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) getSheetComments(sheetFile string) string { _ = "STUB: not implemented"; return "" }

func (f *File) AddComment(sheet string, opts Comment) error { _ = "STUB: not implemented"; return nil }

func (f *File) DeleteComment(sheet, cell string) error { _ = "STUB: not implemented"; return nil }

func (f *File) deleteFormControl(sheetRelationshipsDrawingVML, cell string, isComment bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addComment(commentsXML string, opts vmlOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) countComments() int { _ = "STUB: not implemented"; return 0 }

func (f *File) commentsReader(path string) (*xlsxComments, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) commentsWriter() { _ = "STUB: not implemented"; return }

func (f *File) AddFormControl(sheet string, opts FormControl) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) DeleteFormControl(sheet, cell string) error { _ = "STUB: not implemented"; return nil }

func (f *File) countVMLDrawing() int { _ = "STUB: not implemented"; return 0 }

func (f *File) decodeVMLDrawingReader(path string) (*decodeVmlDrawing, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) vmlDrawingWriter() { _ = "STUB: not implemented"; return }

func (f *File) addVMLObject(opts vmlOptions) error { _ = "STUB: not implemented"; return nil }

func prepareFormCtrlOptions(opts *vmlOptions) *vmlOptions { _ = "STUB: not implemented"; return nil }

func formCtrlText(opts *vmlOptions) []vmlFont { _ = "STUB: not implemented"; return nil }

var formCtrlPresets = map[FormControlType]formCtrlPreset{
	FormControlNote: {
		objectType:   "Note",
		autoFill:     "True",
		filled:       "",
		fillColor:    "#FBF6D6",
		stroked:      "",
		strokeColor:  "#EDEAA1",
		strokeButton: "",
		fill: &vFill{
			Color2: "#FBFE82",
			Angle:  -180,
			Type:   "gradient",
			Fill:   &oFill{Ext: "view", Type: "gradientUnscaled"},
		},
		textHAlign:  "",
		textVAlign:  "",
		noThreeD:    nil,
		firstButton: nil,
		shadow:      &vShadow{On: "t", Color: "black", Obscured: "t"},
	},
	FormControlButton: {
		objectType:   "Button",
		autoFill:     "True",
		filled:       "",
		fillColor:    "buttonFace [67]",
		stroked:      "",
		strokeColor:  "windowText [64]",
		strokeButton: "t",
		fill: &vFill{
			Color2: "buttonFace [67]",
			Angle:  -180,
			Type:   "gradient",
			Fill:   &oFill{Ext: "view", Type: "gradientUnscaled"},
		},
		textHAlign:  "Center",
		textVAlign:  "Center",
		noThreeD:    nil,
		firstButton: nil,
		shadow:      nil,
	},
	FormControlCheckBox: {
		objectType:   "Checkbox",
		autoFill:     "True",
		filled:       "f",
		fillColor:    "window [65]",
		stroked:      "f",
		strokeColor:  "windowText [64]",
		strokeButton: "",
		fill:         nil,
		textHAlign:   "",
		textVAlign:   "Center",
		noThreeD:     stringPtr(""),
		firstButton:  nil,
		shadow:       nil,
	},
	FormControlGroupBox: {
		objectType:   "GBox",
		autoFill:     "False",
		filled:       "f",
		fillColor:    "",
		stroked:      "f",
		strokeColor:  "windowText [64]",
		strokeButton: "",
		fill:         nil,
		textHAlign:   "",
		textVAlign:   "",
		noThreeD:     stringPtr(""),
		firstButton:  nil,
		shadow:       nil,
	},
	FormControlLabel: {
		objectType:   "Label",
		autoFill:     "False",
		filled:       "f",
		fillColor:    "window [65]",
		stroked:      "f",
		strokeColor:  "windowText [64]",
		strokeButton: "",
		fill:         nil,
		textHAlign:   "",
		textVAlign:   "",
		noThreeD:     nil,
		firstButton:  nil,
		shadow:       nil,
	},
	FormControlOptionButton: {
		objectType:   "Radio",
		autoFill:     "False",
		filled:       "f",
		fillColor:    "window [65]",
		stroked:      "f",
		strokeColor:  "windowText [64]",
		strokeButton: "",
		fill:         nil,
		textHAlign:   "",
		textVAlign:   "Center",
		noThreeD:     stringPtr(""),
		firstButton:  stringPtr(""),
		shadow:       nil,
	},
	FormControlScrollBar: {
		objectType:   "Scroll",
		autoFill:     "",
		filled:       "",
		fillColor:    "",
		stroked:      "f",
		strokeColor:  "windowText [64]",
		strokeButton: "",
		fill:         nil,
		textHAlign:   "",
		textVAlign:   "",
		noThreeD:     nil,
		firstButton:  nil,
		shadow:       nil,
	},
	FormControlSpinButton: {
		objectType:   "Spin",
		autoFill:     "False",
		filled:       "",
		fillColor:    "",
		stroked:      "f",
		strokeColor:  "windowText [64]",
		strokeButton: "",
		fill:         nil,
		textHAlign:   "",
		textVAlign:   "",
		noThreeD:     nil,
		firstButton:  nil,
		shadow:       nil,
	},
}

func (sp *encodeShape) addFormCtrl(opts *vmlOptions) error { _ = "STUB: not implemented"; return nil }

func (f *File) addFormCtrlShape(preset formCtrlPreset, col, row int, anchor string, opts *vmlOptions) (*encodeShape, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *File) addDrawingVML(sheetID int, drawingVML string, opts *vmlOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) GetFormControls(sheet string) ([]FormControl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractFormControl(clientData string) (FormControl, error) {
	_ = "STUB: not implemented"
	return *new(FormControl), nil
}

func extractAnchorCell(anchor string) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func extractVMLFont(font []decodeVMLFont) []RichTextRun { _ = "STUB: not implemented"; return nil }

func (f *File) AddHeaderFooterImage(sheet string, opts *HeaderFooterImageOptions) error {
	_ = "STUB: not implemented"
	return nil
}
