package dto

type (
	LoginRequest struct {
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required"`
		RememberMe bool   `json:"remember_me"`
		UserAgent  string `json:"-"`
		IP         string `json:"-"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	RegisterRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}

	VerifyEmailRequest struct {
		Token string `json:"token" binding:"required"`
	}

	LoginResponse struct {
		AccessToken  string  `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
		Role         string  `json:"role"`
	}

	LoginWithGoogleResponse struct {
		NeedRegistration bool   `json:"need_registration"`
		Token            string `json:"token"`
		RegisterToken    string `json:"register_token"`
		Role             string `json:"role"`
	}

	ForgetPasswordRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	ChangePasswordRequest struct {
		Email       string
		NewPassword string `json:"new_password" binding:"required"`
	}

	GetMe struct {
		PersonalInfo PersonalInfo `json:"personal_info"`
	}

	PersonalInfo struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}
)
