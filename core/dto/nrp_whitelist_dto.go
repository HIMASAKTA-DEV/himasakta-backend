package dto

type CreateNrpWhitelistRequest struct {
	Nrp  string `json:"nrp" binding:"required"`
	Name string `json:"name"`
}

type UpdateNrpWhitelistRequest struct {
	Nrp  *string `json:"nrp"`
	Name *string `json:"name"`
}

type CheckNrpWhitelistRequest struct {
	Nrp string `json:"nrp" binding:"required"`
}
