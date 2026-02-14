package dto

type CreateNrpWhitelistRequest struct {
	Nrp  string `json:"nrp" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateNrpWhitelistRequest struct {
	Nrp  string `json:"nrp" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CheckNrpWhitelistRequest struct {
	Nrp string `json:"nrp" binding:"required"`
}
