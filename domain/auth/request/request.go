package request

type CreateAuthRequest struct {
    Name        string `json:"name" validate:"required"`
    Email       string `json:"email" validate:"required,email"`
    Password    string `json:"password" validate:"required,min=6,max=16"`
}

type UpdateAuthRequest struct {
    Name        string `json:"name" validate:"required"`
    Email       string `json:"email" validate:"required,email"`
    Password    string `json:"password" validate:"required,min=6,max=16"`
}
