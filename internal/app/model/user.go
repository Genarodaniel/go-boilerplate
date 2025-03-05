package model

type PostUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type PostUserRequest struct {
	Email string `validate:"required" json:"email"`
	Name  string `validate:"required" json:"name"`
}

type PostUserLogin struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserQueue struct {
	ClientID string `json:"client_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
