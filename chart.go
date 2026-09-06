package excelize

type ChartType byte

const (
	Area ChartType = iota
	AreaStacked
	AreaPercentStacked
	Area3D
	Area3DStacked
	Area3DPercentStacked
	Bar
	BarStacked
	BarPercentStacked
	Bar3DClustered
	Bar3DStacked
	Bar3DPercentStacked
	Bar3DConeClustered
	Bar3DConeStacked
	Bar3DConePercentStacked
	Bar3DPyramidClustered
	Bar3DPyramidStacked
	Bar3DPyramidPercentStacked
	Bar3DCylinderClustered
	Bar3DCylinderStacked
	Bar3DCylinderPercentStacked
	Col
	ColStacked
	ColPercentStacked
	Col3D
	Col3DClustered
	Col3DStacked
	Col3DPercentStacked
	Col3DCone
	Col3DConeClustered
	Col3DConeStacked
	Col3DConePercentStacked
	Col3DPyramid
	Col3DPyramidClustered
	Col3DPyramidStacked
	Col3DPyramidPercentStacked
	Col3DCylinder
	Col3DCylinderClustered
	Col3DCylinderStacked
	Col3DCylinderPercentStacked
	Doughnut
	Line
	Line3D
	Pie
	Pie3D
	PieOfPie
	BarOfPie
	Radar
	Scatter
	Surface3D
	WireframeSurface3D
	Contour
	WireframeContour
	Bubble
	Bubble3D
	StockHighLowClose
	StockOpenHighLowClose
)

type LineDashType byte

const (
	LineDashUnset LineDashType = iota
	LineDashSolid
	LineDashDot
	LineDashDash
	LineDashLgDash
	LineDashSashDot
	LineDashLgDashDot
	LineDashLgDashDotDot
	LineDashSysDash
	LineDashSysDot
	LineDashSysDashDot
	LineDashSysDashDotDot
)

type LineType byte

const (
	LineUnset LineType = iota
	LineSolid
	LineNone
	LineAutomatic
)

type ChartTickLabelPositionType byte

const (
	ChartTickLabelNextToAxis ChartTickLabelPositionType = iota
	ChartTickLabelHigh
	ChartTickLabelLow
	ChartTickLabelNone
)

var (
	chartView3DRotX = map[ChartType]int{
		Area:                        0,
		AreaStacked:                 0,
		AreaPercentStacked:          0,
		Area3D:                      15,
		Area3DStacked:               15,
		Area3DPercentStacked:        15,
		Bar:                         0,
		BarStacked:                  0,
		BarPercentStacked:           0,
		Bar3DClustered:              15,
		Bar3DStacked:                15,
		Bar3DPercentStacked:         15,
		Bar3DConeClustered:          15,
		Bar3DConeStacked:            15,
		Bar3DConePercentStacked:     15,
		Bar3DPyramidClustered:       15,
		Bar3DPyramidStacked:         15,
		Bar3DPyramidPercentStacked:  15,
		Bar3DCylinderClustered:      15,
		Bar3DCylinderStacked:        15,
		Bar3DCylinderPercentStacked: 15,
		Col:                         0,
		ColStacked:                  0,
		ColPercentStacked:           0,
		Col3D:                       15,
		Col3DClustered:              15,
		Col3DStacked:                15,
		Col3DPercentStacked:         15,
		Col3DCone:                   15,
		Col3DConeClustered:          15,
		Col3DConeStacked:            15,
		Col3DConePercentStacked:     15,
		Col3DPyramid:                15,
		Col3DPyramidClustered:       15,
		Col3DPyramidStacked:         15,
		Col3DPyramidPercentStacked:  15,
		Col3DCylinder:               15,
		Col3DCylinderClustered:      15,
		Col3DCylinderStacked:        15,
		Col3DCylinderPercentStacked: 15,
		Doughnut:                    0,
		Line:                        0,
		Line3D:                      20,
		Pie:                         0,
		Pie3D:                       30,
		PieOfPie:                    0,
		BarOfPie:                    0,
		Radar:                       0,
		Scatter:                     0,
		Surface3D:                   15,
		WireframeSurface3D:          15,
		Contour:                     90,
		WireframeContour:            90,
	}
	chartView3DRotY = map[ChartType]int{
		Area:                        0,
		AreaStacked:                 0,
		AreaPercentStacked:          0,
		Area3D:                      20,
		Area3DStacked:               20,
		Area3DPercentStacked:        20,
		Bar:                         0,
		BarStacked:                  0,
		BarPercentStacked:           0,
		Bar3DClustered:              20,
		Bar3DStacked:                20,
		Bar3DPercentStacked:         20,
		Bar3DConeClustered:          20,
		Bar3DConeStacked:            20,
		Bar3DConePercentStacked:     20,
		Bar3DPyramidClustered:       20,
		Bar3DPyramidStacked:         20,
		Bar3DPyramidPercentStacked:  20,
		Bar3DCylinderClustered:      20,
		Bar3DCylinderStacked:        20,
		Bar3DCylinderPercentStacked: 20,
		Col:                         0,
		ColStacked:                  0,
		ColPercentStacked:           0,
		Col3D:                       20,
		Col3DClustered:              20,
		Col3DStacked:                20,
		Col3DPercentStacked:         20,
		Col3DCone:                   20,
		Col3DConeClustered:          20,
		Col3DConeStacked:            20,
		Col3DConePercentStacked:     20,
		Col3DPyramid:                20,
		Col3DPyramidClustered:       20,
		Col3DPyramidStacked:         20,
		Col3DPyramidPercentStacked:  20,
		Col3DCylinder:               20,
		Col3DCylinderClustered:      20,
		Col3DCylinderStacked:        20,
		Col3DCylinderPercentStacked: 20,
		Doughnut:                    0,
		Line:                        0,
		Line3D:                      15,
		Pie:                         0,
		Pie3D:                       0,
		PieOfPie:                    0,
		BarOfPie:                    0,
		Radar:                       0,
		Scatter:                     0,
		Surface3D:                   20,
		WireframeSurface3D:          20,
		Contour:                     0,
		WireframeContour:            0,
	}
	plotAreaChartOverlap = map[ChartType]int{
		BarStacked:        100,
		BarPercentStacked: 100,
		ColStacked:        100,
		ColPercentStacked: 100,
	}
	chartView3DPerspective = map[ChartType]int{
		Line3D:           30,
		Contour:          0,
		WireframeContour: 0,
	}
	chartView3DRAngAx = map[ChartType]int{
		Area:                        0,
		AreaStacked:                 0,
		AreaPercentStacked:          0,
		Area3D:                      1,
		Area3DStacked:               1,
		Area3DPercentStacked:        1,
		Bar:                         0,
		BarStacked:                  0,
		BarPercentStacked:           0,
		Bar3DClustered:              1,
		Bar3DStacked:                1,
		Bar3DPercentStacked:         1,
		Bar3DConeClustered:          1,
		Bar3DConeStacked:            1,
		Bar3DConePercentStacked:     1,
		Bar3DPyramidClustered:       1,
		Bar3DPyramidStacked:         1,
		Bar3DPyramidPercentStacked:  1,
		Bar3DCylinderClustered:      1,
		Bar3DCylinderStacked:        1,
		Bar3DCylinderPercentStacked: 1,
		Col:                         0,
		ColStacked:                  0,
		ColPercentStacked:           0,
		Col3D:                       1,
		Col3DClustered:              1,
		Col3DStacked:                1,
		Col3DPercentStacked:         1,
		Col3DCone:                   1,
		Col3DConeClustered:          1,
		Col3DConeStacked:            1,
		Col3DConePercentStacked:     1,
		Col3DPyramid:                1,
		Col3DPyramidClustered:       1,
		Col3DPyramidStacked:         1,
		Col3DPyramidPercentStacked:  1,
		Col3DCylinder:               1,
		Col3DCylinderClustered:      1,
		Col3DCylinderStacked:        1,
		Col3DCylinderPercentStacked: 1,
		Doughnut:                    0,
		Line:                        0,
		Line3D:                      0,
		Pie:                         0,
		Pie3D:                       0,
		PieOfPie:                    0,
		BarOfPie:                    0,
		Radar:                       0,
		Scatter:                     0,
		Surface3D:                   0,
		WireframeSurface3D:          0,
		Contour:                     0,
		Bubble:                      0,
		Bubble3D:                    0,
	}
	chartLegendPosition = map[string]string{
		"bottom":    "b",
		"left":      "l",
		"right":     "r",
		"top":       "t",
		"top_right": "tr",
	}
	chartValAxNumFmtFormatCode = map[ChartType]string{
		Area:                        "General",
		AreaStacked:                 "General",
		AreaPercentStacked:          "0%",
		Area3D:                      "General",
		Area3DStacked:               "General",
		Area3DPercentStacked:        "0%",
		Bar:                         "General",
		BarStacked:                  "General",
		BarPercentStacked:           "0%",
		Bar3DClustered:              "General",
		Bar3DStacked:                "General",
		Bar3DPercentStacked:         "0%",
		Bar3DConeClustered:          "General",
		Bar3DConeStacked:            "General",
		Bar3DConePercentStacked:     "0%",
		Bar3DPyramidClustered:       "General",
		Bar3DPyramidStacked:         "General",
		Bar3DPyramidPercentStacked:  "0%",
		Bar3DCylinderClustered:      "General",
		Bar3DCylinderStacked:        "General",
		Bar3DCylinderPercentStacked: "0%",
		Col:                         "General",
		ColStacked:                  "General",
		ColPercentStacked:           "0%",
		Col3D:                       "General",
		Col3DClustered:              "General",
		Col3DStacked:                "General",
		Col3DPercentStacked:         "0%",
		Col3DCone:                   "General",
		Col3DConeClustered:          "General",
		Col3DConeStacked:            "General",
		Col3DConePercentStacked:     "0%",
		Col3DPyramid:                "General",
		Col3DPyramidClustered:       "General",
		Col3DPyramidStacked:         "General",
		Col3DPyramidPercentStacked:  "0%",
		Col3DCylinder:               "General",
		Col3DCylinderClustered:      "General",
		Col3DCylinderStacked:        "General",
		Col3DCylinderPercentStacked: "0%",
		Doughnut:                    "General",
		Line:                        "General",
		Line3D:                      "General",
		Pie:                         "General",
		Pie3D:                       "General",
		PieOfPie:                    "General",
		BarOfPie:                    "General",
		Radar:                       "General",
		Scatter:                     "General",
		Surface3D:                   "General",
		WireframeSurface3D:          "General",
		Contour:                     "General",
		WireframeContour:            "General",
		Bubble:                      "General",
		Bubble3D:                    "General",
		StockHighLowClose:           "General",
		StockOpenHighLowClose:       "General",
	}
	chartValAxCrossBetween = map[ChartType]string{
		Area:                        "midCat",
		AreaStacked:                 "midCat",
		AreaPercentStacked:          "midCat",
		Area3D:                      "midCat",
		Area3DStacked:               "midCat",
		Area3DPercentStacked:        "midCat",
		Bar:                         "between",
		BarStacked:                  "between",
		BarPercentStacked:           "between",
		Bar3DClustered:              "between",
		Bar3DStacked:                "between",
		Bar3DPercentStacked:         "between",
		Bar3DConeClustered:          "between",
		Bar3DConeStacked:            "between",
		Bar3DConePercentStacked:     "between",
		Bar3DPyramidClustered:       "between",
		Bar3DPyramidStacked:         "between",
		Bar3DPyramidPercentStacked:  "between",
		Bar3DCylinderClustered:      "between",
		Bar3DCylinderStacked:        "between",
		Bar3DCylinderPercentStacked: "between",
		Col:                         "between",
		ColStacked:                  "between",
		ColPercentStacked:           "between",
		Col3D:                       "between",
		Col3DClustered:              "between",
		Col3DStacked:                "between",
		Col3DPercentStacked:         "between",
		Col3DCone:                   "between",
		Col3DConeClustered:          "between",
		Col3DConeStacked:            "between",
		Col3DConePercentStacked:     "between",
		Col3DPyramid:                "between",
		Col3DPyramidClustered:       "between",
		Col3DPyramidStacked:         "between",
		Col3DPyramidPercentStacked:  "between",
		Col3DCylinder:               "between",
		Col3DCylinderClustered:      "between",
		Col3DCylinderStacked:        "between",
		Col3DCylinderPercentStacked: "between",
		Doughnut:                    "between",
		Line:                        "between",
		Line3D:                      "between",
		Pie:                         "between",
		Pie3D:                       "between",
		PieOfPie:                    "between",
		BarOfPie:                    "between",
		Radar:                       "between",
		Scatter:                     "between",
		Surface3D:                   "midCat",
		WireframeSurface3D:          "midCat",
		Contour:                     "midCat",
		WireframeContour:            "midCat",
		Bubble:                      "midCat",
		Bubble3D:                    "midCat",
		StockHighLowClose:           "between",
		StockOpenHighLowClose:       "between",
	}
	plotAreaChartGrouping = map[ChartType]string{
		Area:                        "standard",
		AreaStacked:                 "stacked",
		AreaPercentStacked:          "percentStacked",
		Area3D:                      "standard",
		Area3DStacked:               "stacked",
		Area3DPercentStacked:        "percentStacked",
		Bar:                         "clustered",
		BarStacked:                  "stacked",
		BarPercentStacked:           "percentStacked",
		Bar3DClustered:              "clustered",
		Bar3DStacked:                "stacked",
		Bar3DPercentStacked:         "percentStacked",
		Bar3DConeClustered:          "clustered",
		Bar3DConeStacked:            "stacked",
		Bar3DConePercentStacked:     "percentStacked",
		Bar3DPyramidClustered:       "clustered",
		Bar3DPyramidStacked:         "stacked",
		Bar3DPyramidPercentStacked:  "percentStacked",
		Bar3DCylinderClustered:      "clustered",
		Bar3DCylinderStacked:        "stacked",
		Bar3DCylinderPercentStacked: "percentStacked",
		Col:                         "clustered",
		ColStacked:                  "stacked",
		ColPercentStacked:           "percentStacked",
		Col3D:                       "standard",
		Col3DClustered:              "clustered",
		Col3DStacked:                "stacked",
		Col3DPercentStacked:         "percentStacked",
		Col3DCone:                   "standard",
		Col3DConeClustered:          "clustered",
		Col3DConeStacked:            "stacked",
		Col3DConePercentStacked:     "percentStacked",
		Col3DPyramid:                "standard",
		Col3DPyramidClustered:       "clustered",
		Col3DPyramidStacked:         "stacked",
		Col3DPyramidPercentStacked:  "percentStacked",
		Col3DCylinder:               "standard",
		Col3DCylinderClustered:      "clustered",
		Col3DCylinderStacked:        "stacked",
		Col3DCylinderPercentStacked: "percentStacked",
		Line:                        "standard",
		Line3D:                      "standard",
	}
	plotAreaChartBarDir = map[ChartType]string{
		Bar:                         "bar",
		BarStacked:                  "bar",
		BarPercentStacked:           "bar",
		Bar3DClustered:              "bar",
		Bar3DStacked:                "bar",
		Bar3DPercentStacked:         "bar",
		Bar3DConeClustered:          "bar",
		Bar3DConeStacked:            "bar",
		Bar3DConePercentStacked:     "bar",
		Bar3DPyramidClustered:       "bar",
		Bar3DPyramidStacked:         "bar",
		Bar3DPyramidPercentStacked:  "bar",
		Bar3DCylinderClustered:      "bar",
		Bar3DCylinderStacked:        "bar",
		Bar3DCylinderPercentStacked: "bar",
		Col:                         "col",
		ColStacked:                  "col",
		ColPercentStacked:           "col",
		Col3D:                       "col",
		Col3DClustered:              "col",
		Col3DStacked:                "col",
		Col3DPercentStacked:         "col",
		Col3DCone:                   "col",
		Col3DConeStacked:            "col",
		Col3DConeClustered:          "col",
		Col3DConePercentStacked:     "col",
		Col3DPyramid:                "col",
		Col3DPyramidClustered:       "col",
		Col3DPyramidStacked:         "col",
		Col3DPyramidPercentStacked:  "col",
		Col3DCylinder:               "col",
		Col3DCylinderClustered:      "col",
		Col3DCylinderStacked:        "col",
		Col3DCylinderPercentStacked: "col",
		Line:                        "standard",
		Line3D:                      "standard",
	}
	barColChartTypes = []ChartType{
		Bar,
		BarStacked,
		BarPercentStacked,
		Bar3DClustered,
		Bar3DStacked,
		Bar3DPercentStacked,
		Bar3DConeClustered,
		Bar3DConeStacked,
		Bar3DConePercentStacked,
		Bar3DPyramidClustered,
		Bar3DPyramidStacked,
		Bar3DPyramidPercentStacked,
		Bar3DCylinderClustered,
		Bar3DCylinderStacked,
		Bar3DCylinderPercentStacked,
		Col,
		ColStacked,
		ColPercentStacked,
		Col3D,
		Col3DClustered,
		Col3DStacked,
		Col3DPercentStacked,
		Col3DCone,
		Col3DConeStacked,
		Col3DConeClustered,
		Col3DConePercentStacked,
		Col3DPyramid,
		Col3DPyramidClustered,
		Col3DPyramidStacked,
		Col3DPyramidPercentStacked,
		Col3DCylinder,
		Col3DCylinderClustered,
		Col3DCylinderStacked,
		Col3DCylinderPercentStacked,
	}
	orientation = map[bool]string{
		true:  "maxMin",
		false: "minMax",
	}
	catAxPos = map[bool]string{
		true:  "t",
		false: "b",
	}
	valAxPos = map[bool]string{
		true:  "r",
		false: "l",
	}
	tickLblPosVal = map[ChartTickLabelPositionType]string{
		ChartTickLabelNextToAxis: "nextTo",
		ChartTickLabelHigh:       "high",
		ChartTickLabelLow:        "low",
		ChartTickLabelNone:       "none",
	}
	tickLblPosNone = map[ChartType]string{
		Contour:          "none",
		WireframeContour: "none",
	}
)

func parseChartOptions(opts *Chart) (*Chart, error) { _ = "STUB: not implemented"; return nil, nil }

func (opts *Chart) parseSeries() error { _ = "STUB: not implemented"; return nil }

func (opts *Chart) parseTitle() error { _ = "STUB: not implemented"; return nil }

func (f *File) AddChart(sheet, cell string, chart *Chart, combo ...*Chart) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) AddChartSheet(sheet string, chart *Chart, combo ...*Chart) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) getChartOptions(opts *Chart, combo []*Chart) (*Chart, []*Chart, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (f *File) DeleteChart(sheet, cell string) error { _ = "STUB: not implemented"; return nil }

func (f *File) countCharts() int { _ = "STUB: not implemented"; return 0 }

func ptToEMUs(pt float64) int { _ = "STUB: not implemented"; return 0 }
