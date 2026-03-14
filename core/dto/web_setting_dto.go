package dto

import "fmt"

type WebSettings struct {
	ExternalSOPLink   string           `json:"ExternalSOPLink"`
	InternalSOPLink   string           `json:"InternalSOPLink"`
	DeskripsiHimpunan string           `json:"DeskripsiHimpunan"`
	FotoHimpunan      string           `json:"FotoHimpunan"`
	SocialMedia       []SocialMediaDTO `json:"SocialMedia"`
	InMaintenance     bool             `json:"InMaintenance"`
}

type SocialMediaDTO struct {
	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}

func (w WebSettings) Validate() error {
	if len(w.SocialMedia) > 20 {
		return fmt.Errorf("social media links cannot exceed 20")
	}
	return nil
}
