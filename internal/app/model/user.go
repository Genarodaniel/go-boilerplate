package model

type PostUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type PostUserRequest struct {
	Email    string `validate:"required" json:"email"`
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required" json:"password"`
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
