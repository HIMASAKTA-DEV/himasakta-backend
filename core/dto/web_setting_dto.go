package dto

type WebSettings struct {
	ExternalSOPLink    string         `json:"ExternalSOPLink"`
	InternalSOPLink    string         `json:"InternalSOPLink"`
	DeskripsiHimpunan  string         `json:"DeskripsiHimpunan"`
	FotoHimpunan       string         `json:"FotoHimpunan"`
	SocialMedia       SocialMediaDTO `json:"SocialMedia"`
	InMaintenance      bool           `json:"InMaintenance"`
}

type SocialMediaDTO struct {
	Instagram string `json:"instagram"`
	TikTok    string `json:"tiktok"`
	YouTube   string `json:"youtube"`
	LinkedIn  string `json:"linkedin"`
	Linktree  string `json:"linktree"`
}
