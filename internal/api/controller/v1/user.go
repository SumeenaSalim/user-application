package v1

import (
	"net/http"
	"user-app/internal/service"
	"user-app/internal/utils"
)

type UserControllerInterface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response, err := u.userService.CreateUser(ctx)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.SendSuccessResponse(w, http.StatusCreated, response)
}

func (u *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response, err := u.userService.UpdateUser(ctx)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.SendSuccessResponse(w, http.StatusOK, response)
}
