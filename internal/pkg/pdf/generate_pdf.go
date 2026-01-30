package mypdf

import (
	"bytes"
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

var (
	ColorBlack  = []int{0, 0, 0}
	ColorGray   = []int{231, 230, 230}
	ColorYellow = []int{255, 255, 0}
	ColorWhite  = []int{255, 255, 255}
)

const (
	marginBottom    = 15.0
	pageHeight      = 210.0
	minRowHeight    = 6.0
	headerHeight    = 8.0
	tableStartYPage = 10.0
)

func Generate(req []GenerateRequestData) (*bytes.Buffer, string, error) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetAutoPageBreak(false, 0)

	for _, r := range req {
		pdf.AddPage()
		setupPDFDefaults(pdf)
		drawHeader(pdf)
		drawPackageInfo(pdf, r.PackageInfoData)
		drawDisciplineSection(pdf, r.DisciplineSectionData)

		drawMainTableWithPageBreak(pdf, r.CommentRow)
	}

	filename := "comment_resolution_sheet.pdf"

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, "", err
	}

	return &buf, filename, nil
}

func setupPDFDefaults(pdf *gofpdf.Fpdf) {
	pdf.SetTextColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])
	pdf.SetDrawColor(ColorBlack[0], ColorBlack[1], ColorBlack[2])
	pdf.SetLineWidth(0.3)
}

func drawHeader(pdf *gofpdf.Fpdf) {
	logoOpt := gofpdf.ImageOptions{
		ImageType: "PNG",
		ReadDpi:   true,
	}

	pdf.ImageOptions("./assets/image/Logo-CRS.png", 10, 8, 55, 12, false, logoOpt, 0, "")

	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(100, 10)
	pdf.Cell(30, 8, "COMMENT RESOLUTION SHEET")
}

func drawPackageInfo(pdf *gofpdf.Fpdf, data PackageInfoData) {
	pdf.SetFont("Arial", "", 8)
	startY := 38.0

	leftInfo := [][2]string{
		{"Package", fmt.Sprintf(": %s", data.Package)},
		{"FEED Contractor", fmt.Sprintf(": %s", data.ContractorInitial)},
	}

	for i, info := range leftInfo {
		y := startY + float64(i*5)
		pdf.SetXY(10, y)
		pdf.Cell(35, 4, info[0])
		pdf.SetXY(45, y)
		pdf.Cell(60, 4, info[1])
	}

	rightInfo := []string{
		"Inc. Transmittal",
		"Out. Transmittal",
		"Out. Transmittal Date",
	}

	for i, label := range rightInfo {
		y := startY + float64(i*5)
		pdf.SetXY(200, y)
		pdf.Cell(40, 4, label)
		pdf.SetXY(245, y)
		pdf.Cell(5, 4, ":")
	}
}

// drawDisciplineSection draws the discipline information table
func drawDisciplineSection(pdf *gofpdf.Fpdf, data DisciplineSectionData) {
	disciplineY := 56.0
	rowHeight := 6.0

	columns := []struct {
		x     float64
		width float64
		label string
	}{
		{10, 30, "Discipline"},
		// {40, 30, "Area of Concern ID"},
		// {70, 90, "Area of Concern Description"},
		{40, 50, "Consolidator"},
	}

	pdf.SetFont("Arial", "B", 8)
	setFillColor(pdf, ColorGray)

	for _, col := range columns {
		pdf.Rect(col.x, disciplineY, col.width, rowHeight, "FD")
		pdf.SetXY(col.x, disciplineY+1.5)
		pdf.Cell(col.width, 4, col.label)
	}

	pdf.SetFont("Arial", "", 7)
	contentY := disciplineY + rowHeight
	contentHeight := 10.0

	disciplineData := []struct {
		x     float64
		width float64
		text  string
		vPos  float64
	}{
		{10, 30, data.Discipline, 1},
		// {40, 30, data.AreaOfConcernID, 3},
		// {70, 90, data.AreaOfConcernDescription, 3},
		{40, 50, data.Consolidator, 3},
	}

	for _, data := range disciplineData {
		pdf.Rect(data.x, contentY, data.width, contentHeight, "D")
		pdf.SetXY(data.x, contentY+data.vPos)
		pdf.MultiCell(data.width, 3, data.text, "", "L", false)
	}

	pdf.SetFont("Arial", "I", 7)
	pdf.SetXY(10, contentY+12)
	pdf.Cell(100, 4, "*Please manually sort page number in ascending order")
}

// calculateRowHeight calculates the height needed for a row based on its content
func calculateRowHeight(pdf *gofpdf.Fpdf, row CommentRow, colWidths []float64) float64 {
	pdf.SetFont("Arial", "", 7)
	lineHeight := 3.0
	maxLines := 1

	rowData := []string{
		row.No,
		row.Page,
		row.SMEInitial,
		row.SMEComment,
		row.RefDocNo,
		row.RefDocTitle,
		row.DocStatus,
		row.Status,
		row.SMECloseComment,
	}

	// Calculate lines needed for each column
	for i, text := range rowData {
		if text == "" {
			continue
		}
		lines := pdf.SplitLines([]byte(text), colWidths[i]-2)
		numLines := len(lines)
		if numLines > maxLines {
			maxLines = numLines
		}
	}

	// Calculate height: padding + (lines * lineHeight)
	height := 2.0 + float64(maxLines)*lineHeight
	if height < minRowHeight {
		height = minRowHeight
	}

	return height
}

// drawMainTableWithPageBreak draws the main table with manual page break handling
func drawMainTableWithPageBreak(pdf *gofpdf.Fpdf, rows []CommentRow) {
	tableStartY := 86.0
	currentY := tableStartY

	colWidths := []float64{10, 20, 20, 40, 40, 40, 30, 25, 40}
	headers := []string{
		"No.",
		"Page *",
		"SME Initial",
		"SME\nComment",
		"Ref. Document No.",
		"Ref. Document Title",
		"Doc. Status",
		"Status",
		"SME Close Out\nComments",
	}

	// Draw header pertama kali
	drawTableHeaders(pdf, headers, colWidths, currentY)
	currentY += headerHeight

	pdf.SetFont("Arial", "", 7)

	// Loop through all rows
	for _, row := range rows {
		// Calculate height for this row
		rowHeight := calculateRowHeight(pdf, row, colWidths)

		// Cek apakah masih cukup ruang untuk row ini
		if currentY+rowHeight > pageHeight-marginBottom {
			// Pindah ke halaman baru
			pdf.AddPage()
			currentY = tableStartYPage

			// Draw header lagi di halaman baru
			drawTableHeaders(pdf, headers, colWidths, currentY)
			currentY += headerHeight
			pdf.SetFont("Arial", "", 7)
		}

		// Draw row with calculated height
		drawTableRowMultiline(pdf, row, currentY, colWidths, rowHeight)
		currentY += rowHeight
	}
}

// drawTableHeaders draws the table header row
func drawTableHeaders(pdf *gofpdf.Fpdf, headers []string, colWidths []float64, y float64) {
	pdf.SetFont("Arial", "B", 7)
	x := 10.0

	for i, header := range headers {
		if i >= 4 {
			setFillColor(pdf, ColorYellow)
		} else {
			setFillColor(pdf, ColorGray)
		}

		pdf.Rect(x, y, colWidths[i], headerHeight, "FD")
		pdf.SetXY(x+1, y+1)
		pdf.MultiCell(colWidths[i]-2, 3, header, "", "C", false)
		x += colWidths[i]
	}
}

// drawTableRowMultiline draws a single data row with multiline support
func drawTableRowMultiline(pdf *gofpdf.Fpdf, row CommentRow, y float64, colWidths []float64, height float64) {
	x := 10.0
	rowData := []string{
		row.No,
		row.Page,
		row.SMEInitial,
		row.SMEComment,
		row.RefDocNo,
		row.RefDocTitle,
		row.DocStatus,
		row.Status,
		row.SMECloseComment,
	}

	for i, data := range rowData {
		// Draw cell border
		pdf.Rect(x, y, colWidths[i], height, "D")

		// Draw text with wrapping
		pdf.SetXY(x+1, y+1)

		// Save current position for MultiCell
		startX := x + 1
		startY := y + 1

		// Use MultiCell for text wrapping
		pdf.SetXY(startX, startY)
		pdf.MultiCell(colWidths[i]-2, 3, data, "", "L", false)

		x += colWidths[i]
	}
}

// setFillColor sets the fill color from RGB array
func setFillColor(pdf *gofpdf.Fpdf, color []int) {
	pdf.SetFillColor(color[0], color[1], color[2])
}
