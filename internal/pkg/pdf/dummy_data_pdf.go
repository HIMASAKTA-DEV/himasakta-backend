package mypdf

func GetSampleIFRRows() []CommentRow {
	return []CommentRow{
		{"1", "Page 20", "ABC", "comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1 comment 1", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"2", "Page 30", "ABC", "comment 2", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"3", "Page 35", "DEF", "comment 3", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"4", "Page 40", "RST", "comment 4", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"5", "Page 55", "KLM", "comment 5", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"6", "Page 56", "KLM", "comment 6", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"7", "Page 20", "ABC", "comment 1", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"8", "Page 30", "ABC", "comment 2", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"9", "Page 35", "DEF", "comment 3", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"10", "Page 40", "RST", "comment 4", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"11", "Page 55", "KLM", "comment 5", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"12", "Page 56", "KLM", "comment 6", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"13", "Page 20", "ABC", "comment 1", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"14", "Page 30", "ABC", "comment 2", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"15", "Page 35", "DEF", "comment 3", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"16", "Page 40", "RST", "comment 4", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"17", "Page 55", "KLM", "comment 5", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
		{"18", "Page 56", "KLM", "comment 6", "Doc No 123456", "Doc Title ABCDEFG", "IFR Comment", "NA", "NA"},
	}
}

func GetSampleIFURows() []CommentRow {
	return []CommentRow{
		{"1", "Page 20", "ABC", "comment 1", "Doc No 123456", "Doc Title ABCDEFG", "IFU Comment", "Accept", ""},
		{"2", "Page 25", "ABC", "comment 2", "Doc No 123456", "Doc Title ABCDEFG", "IFU Comment", "Accept", ""},
		{"3", "Page 35", "DEF", "comment 3", "Doc No 123456", "Doc Title ABCDEFG", "IFU Comment", "Reject", "Close out comment"},
		{"4", "Page 40", "RST", "comment 4", "Doc No 123456", "Doc Title ABCDEFG", "IFU Comment", "Reject", "Close out comment"},
		{"5", "Page 55", "KLM", "comment 5", "Doc No 123456", "Doc Title ABCDEFG", "IFU Comment", "Accept", ""},
		{"6", "Page 56", "KLM", "comment 6", "Doc No 123456", "Doc Title ABCDEFG", "IFU Comment", "Accept", ""},
	}
}
