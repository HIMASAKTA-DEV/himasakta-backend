package dto

type (
	CreateProposalRequest struct {
		No            string  `json:"no" binding:"required"`
		ComponentName string  `json:"component_name" binding:"required"`
		Specification string  `json:"specification" binding:"required"`
		Unit          string  `json:"unit" binding:"required"`
		UnitPrice     float64 `json:"unit_price" binding:"required"`
		AccountCode   string  `json:"account_code" binding:"required"`
	}

	UpdateProposalRequest struct {
		ID            string  `json:"-"`
		No            string  `json:"no"`
		ComponentName string  `json:"component_name"`
		Specification string  `json:"specification"`
		Unit          string  `json:"unit"`
		UnitPrice     float64 `json:"unit_price"`
		AccountCode   string  `json:"account_code"`
	}
)
