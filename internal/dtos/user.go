package dtos

type User struct {
	ID       uint   `jsom:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Role     string `json:"role"`
}
