package dto

type (
	CreateSPTJMRequest struct {
		SKPDName  string `json:"skpd_name" binding:"required"`
		Address   string `json:"address" binding:"required"`
		PIC1Name  string `json:"pic1_name" binding:"required"`
		PIC1Email string `json:"pic1_email" binding:"required,email"`
		PIC1Phone string `json:"pic1_phone" binding:"required"`
		PIC2Name  string `json:"pic2_name" binding:"required"`
		PIC2Email string `json:"pic2_email" binding:"required,email"`
		PIC2Phone string `json:"pic2_phone" binding:"required"`
		FileURL   string `json:"file_url" binding:"required"`
	}

	UpdateSPTJMRequest struct {
		ID        string `json:"id"`
		SKPDName  string `json:"skpd_name" binding:"required"`
		Address   string `json:"address" binding:"required"`
		PIC1Name  string `json:"pic1_name" binding:"required"`
		PIC1Email string `json:"pic1_email" binding:"required,email"`
		PIC1Phone string `json:"pic1_phone" binding:"required"`
		PIC2Name  string `json:"pic2_name" binding:"required"`
		PIC2Email string `json:"pic2_email" binding:"required,email"`
		PIC2Phone string `json:"pic2_phone" binding:"required"`
		FileURL   string `json:"file_url" binding:"required"`
	}

	SPTJMResponse struct {
		ID        string `json:"id"`
		SKPDName  string `json:"skpd_name"`
		Address   string `json:"address"`
		PIC1Name  string `json:"pic1_name"`
		PIC1Email string `json:"pic1_email"`
		PIC1Phone string `json:"pic1_phone"`
		PIC2Name  string `json:"pic2_name"`
		PIC2Email string `json:"pic2_email"`
		PIC2Phone string `json:"pic2_phone"`
		FileURL   string `json:"file_url"`
	}
)
