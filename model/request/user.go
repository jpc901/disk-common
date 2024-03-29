package request

type UserSignUpRequest struct {
	Username        string `json:"username" binding:"required"`
	Password        string `josn:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UserSignInRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserInfoRequest struct {
	Username string `form:"username" binding:"required"`
}
