package model

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type IsUserExists struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRegisterReq struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password string `db:"password"`
}

type UserRegisterResp struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
}

type GetUser struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	Role     string `db:"role"`
	Password string `db:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
