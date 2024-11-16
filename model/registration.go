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
	ID        string `json:"id"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

type GetUser struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Password string `db:"password"`
}
