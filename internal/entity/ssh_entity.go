package entity

import (
	"github.com/google/uuid"
)

type SSH struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProposalID uuid.UUID `json:"proposal_id"`
	ItemID     uuid.UUID `json:"item_id"`

	UraianKomponen string  `json:"uraian_komponen"`
	Spesifikasi    string  `json:"spesifikasi"`
	Satuan         string  `json:"satuan"`
	HargaSatuan    float64 `json:"harga_satuan"`
	Rekening       string  `json:"rekening"`
	Kel            int     `json:"kel"`
	StatusKomponen string  `json:"status_komponen"`
	StatusSIPD     string  `json:"status_sipd"`
	DasarHarga     string  `json:"dasar_harga"`

	Timestamp

	Proposal *Proposal `json:"proposal,omitempty" gorm:"foreignKey:ProposalID;references:ID"`
	Item     *Item     `json:"item,omitempty" gorm:"foreignKey:ItemID;references:ID"`
}

func (s *SSH) TableName() string {
	return "sshs"
}
