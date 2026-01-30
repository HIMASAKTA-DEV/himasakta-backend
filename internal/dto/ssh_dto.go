package dto

type (
	CreateSSHRequest struct {
		ProposalID     string  `json:"proposal_id" binding:"required"`
		ItemID         string  `json:"item_id" binding:"required"`
		UraianKomponen string  `json:"uraian_komponen" binding:"required"`
		Spesifikasi    string  `json:"spesifikasi" binding:"required"`
		Satuan         string  `json:"satuan" binding:"required"`
		HargaSatuan    float64 `json:"harga_satuan" binding:"required"`
		Rekening       string  `json:"rekening" binding:"required"`
		Kel            int     `json:"kel" binding:"required"`
		StatusKomponen string  `json:"status_komponen" binding:"required"`
		StatusSIPD     string  `json:"status_sipd" binding:"required"`
		DasarHarga     string  `json:"dasar_harga" binding:"required"`
	}

	UpdateSSHRequest struct {
		ID             string  `json:"-"`
		ProposalID     string  `json:"proposal_id"`
		ItemID         string  `json:"item_id"`
		UraianKomponen string  `json:"uraian_komponen"`
		Spesifikasi    string  `json:"spesifikasi"`
		Satuan         string  `json:"satuan"`
		HargaSatuan    float64 `json:"harga_satuan"`
		Rekening       string  `json:"rekening"`
		Kel            int     `json:"kel"`
		StatusKomponen string  `json:"status_komponen"`
		StatusSIPD     string  `json:"status_sipd"`
		DasarHarga     string  `json:"dasar_harga"`
	}
)
