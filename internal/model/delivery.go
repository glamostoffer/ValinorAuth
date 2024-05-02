package model

type (
	AdminSignUpRequest struct {
		Login       string
		Password    string
		InviteToken string
	}

	SignUpRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	SignInRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	SignInResponse struct {
		Token string `json:"token"`
	}
)
