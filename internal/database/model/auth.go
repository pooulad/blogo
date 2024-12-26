package model

type RegisterInput struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Skill     string `json:"skill"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
