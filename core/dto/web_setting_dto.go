package dto

import "fmt"

type WebSettings struct {
	ExternalSOPLink     string           `json:"ExternalSOPLink"`
	InternalSOPLink     string           `json:"InternalSOPLink"`
	DeskripsiHimpunan   string           `json:"DeskripsiHimpunan"`
	VisiHimpunan        string           `json:"VisiHimpunan"`
	MisiHimpunan        string           `json:"MisiHimpunan"`
	FotoHimpunan        string           `json:"FotoHimpunan"`
	FotoSejarahHimpunan string           `json:"FotoSejarahHimpunan"`
	SocialMedia         []SocialMediaDTO `json:"SocialMedia"`
	InMaintenance       bool             `json:"InMaintenance"`
}

type AuthSettings struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
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
