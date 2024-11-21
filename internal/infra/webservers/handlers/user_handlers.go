package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/brunoofgod/go-simple-api/internal/dto"
	"github.com/brunoofgod/go-simple-api/internal/entity"
	"github.com/brunoofgod/go-simple-api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	userDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, Jwt *jwtauth.JWTAuth, JwtExpiresIn int) *UserHandler {
	return &UserHandler{userDB: userDB, Jwt: Jwt, JwtExpiresIn: JwtExpiresIn}
}

// Create User godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserInputDto true "Create User"
// @Success 201
// @Failure 500 {object} Error
// @Router /users [post]
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	userDto := dto.CreateUserInputDto{}

	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = u.userDB.Create(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// AuthUser godoc
// @Summary Auth User
// @Description Auth User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body dto.AuthInputDto true "Auth User"
// @Success 200
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /users/auth [post]
func (u *UserHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var user dto.AuthInputDto

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userFound, err := u.userDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)

		return
	}

	if !userFound.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)

		return
	}

	_, jwtToken, _ := u.Jwt.Encode(map[string]interface{}{
		"sub": userFound.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(u.JwtExpiresIn)).Unix(),
	})

	accessToken := dto.CreateAuthOutputDto{
		AccessToken: jwtToken,
	}

	err = json.NewEncoder(w).Encode(accessToken)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)

		return
	}

	w.WriteHeader(http.StatusOK)
}
