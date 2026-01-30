package mypdf

type (
	GenerateRequestData struct {
		PackageInfoData       PackageInfoData
		DisciplineSectionData DisciplineSectionData
		CommentRow            []CommentRow
	}

	PackageInfoData struct {
		Package           string
		ContractorInitial string
	}

	DisciplineSectionData struct {
		Discipline string
		// AreaOfConcernID          string
		// AreaOfConcernDescription string
		Consolidator string
	}

	CommentRow struct {
		No              string
		Page            string
		SMEInitial      string
		SMEComment      string
		RefDocNo        string
		RefDocTitle     string
		DocStatus       string
		Status          string
		SMECloseComment string
	}
)
