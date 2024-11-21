package dto

type CreateProductInputDto struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInputDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthInputDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateAuthOutputDto struct {
	AccessToken string `json:"access_token"`
}
