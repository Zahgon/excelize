package excelize

func (f *File) prepareDrawing(ws *xlsxWorksheet, drawingID int, sheet, drawingXML string) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (f *File) prepareChartSheetDrawing(cs *xlsxChartsheet, drawingID int, sheet string) {
	_ = "STUB: not implemented"
	return
}

func (f *File) addChart(opts *Chart, comboCharts []*Chart) { _ = "STUB: not implemented"; return }

func (f *File) drawBaseChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawDoughnutChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawLineChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawLine3DChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawPieChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawPie3DChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawPieOfPieChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawBarOfPieChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawRadarChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawScatterChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawSurface3DChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawSurfaceChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawBubbleChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawStockChart(pa *cPlotArea, opts *Chart) *cPlotArea {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartGapWidth(opts *Chart) *attrValInt { _ = "STUB: not implemented"; return nil }

func (f *File) drawChartOverlap(opts *Chart) *attrValInt { _ = "STUB: not implemented"; return nil }

func (f *File) drawChartShape(opts *Chart) *attrValString { _ = "STUB: not implemented"; return nil }

func (f *File) drawChartSeries(opts *Chart) *[]cSer { _ = "STUB: not implemented"; return nil }

func (fill *Fill) drawShapeFill(spPr *cSpPr) *cSpPr { _ = "STUB: not implemented"; return nil }

func (f *File) drawChartSeriesSpPr(i int, opts *Chart) *cSpPr {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartSeriesDPt(i int, opts *Chart) []*cDPt {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartSeriesCat(v ChartSeries, opts *Chart) *cCat {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartSeriesVal(v ChartSeries, opts *Chart) *cVal {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartSeriesMarker(i int, opts *Chart) *cMarker {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartSeriesXVal(v ChartSeries, opts *Chart) *cCat {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartSeriesYVal(v ChartSeries, opts *Chart) *cVal {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawCharSeriesBubbleSize(v ChartSeries, opts *Chart) *cVal {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawCharSeriesBubble3D(opts *Chart) *attrValBool {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawChartNumFmt(labels ChartNumFmt) *cNumFmt { _ = "STUB: not implemented"; return nil }

func (f *File) drawChartDLbls(opts *Chart) *cDLbls { _ = "STUB: not implemented"; return nil }

func inSupportedChartDataLabelsPositionType(a []ChartDataLabelPositionType, x ChartDataLabelPositionType) int {
	_ = "STUB: not implemented"
	return 0
}

func (f *File) drawChartSeriesDLbls(i int, opts *Chart) *cDLbls {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawPlotAreaCatAx(pa *cPlotArea, opts *Chart) []*cAxs {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawPlotAreaValAx(pa *cPlotArea, opts *Chart) []*cAxs {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) drawPlotAreaSerAx(opts *Chart) []*cAxs { _ = "STUB: not implemented"; return nil }

func drawChartFont(fnt *Font, r *aRPr) { _ = "STUB: not implemented"; return }

func (ct *ChartTitle) drawPlotAreaTitles(vert string) *cTitle {
	_ = "STUB: not implemented"
	return nil
}

func (ct *ChartTitle) drawTitlesManualLayout() *cLayout { _ = "STUB: not implemented"; return nil }

func (f *File) drawPlotAreaDTable(opts *Chart) *cDTable { _ = "STUB: not implemented"; return nil }

func (f *File) drawPlotAreaSpPr() *cSpPr { _ = "STUB: not implemented"; return nil }

func (f *File) drawPlotAreaTxPr(opts *ChartAxis) *cTxPr { _ = "STUB: not implemented"; return nil }

func (l *LineOptions) drawChartLn() *aLn { _ = "STUB: not implemented"; return nil }

func (c *cChart) drawChartLegend(opts *Chart) { _ = "STUB: not implemented"; return }

func (f *File) drawingParser(path string) (*xlsxWsDr, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (f *File) addDrawingChart(sheet, drawingXML, cell string, width, height, rID int, opts *GraphicOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) addSheetDrawingChart(drawingXML string, rID int, opts *GraphicOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) deleteDrawing(col, row int, drawingXML, drawingType string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractEmbedRID(pic *xlsxPic, decodePic *decodePic) string {
	_ = "STUB: not implemented"
	return ""
}

func getUnusedCellAnchorRID(delRID, refRID []string, rIDMaps map[string]int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (f *File) deleteDrawingRels(rels, rID string) { _ = "STUB: not implemented"; return }

func (f *File) genAxID(opts *Chart) []*attrValInt { _ = "STUB: not implemented"; return nil }
