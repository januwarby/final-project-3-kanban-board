package dto

import "time"

type NewUserRequest struct {
	FullName string `json:"full_name" valid:"required~full_name cannot be empty"`
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty,min=6"`
	Role     string `json:"role" valid:"required~role cannot be empty"`
}

type NewUserResponse struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type RoleUser struct {
	Role string `json:"role"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserUpdateResponse struct {
	Id       int       `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	UpdatedAt time.Time `json:"updatedAt"`
}
