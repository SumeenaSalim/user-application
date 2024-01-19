package middleware

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"user-app/internal/consts"
	"user-app/internal/dto/input"
	"user-app/internal/utils"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// ValidateCreateUserRequest validates the Create User request.
func ValidateCreateUserRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var request input.CreateUserRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to convert create user request to byte, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Println("failed to unmarshal data for create user request, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		// Validate the request using the validator library
		if err := validateCreateUserRequest(request); err != nil {
			log.Println("failed to validate request body, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx = context.WithValue(ctx, consts.ContextKeyCreateUserRequest, request)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ValidateUpdateUserRequest validates the Update User request.
func ValidateUpdateUserRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var request input.UpdateUserRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to convert update user request to byte, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Println("failed to unmarshal data for update user request, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if request.Name == "" && request.Email == "" {
			log.Println("either name or email is required, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		// Validate the request using the validator library
		if err := validateUpdateUserRequest(request); err != nil {
			log.Println("failed to validate request body, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedId, err := uuid.Parse(chi.URLParam(r, consts.UserId))
		if err != nil {
			log.Println("invalid user id, error: ", err)
			utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		ctx = context.WithValue(ctx, consts.ContextKeyUpdateUserRequest, request)
		ctx = context.WithValue(ctx, consts.UserId, parsedId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// validateUpdateUserRequest validates the Update User request fields.
func validateUpdateUserRequest(request input.UpdateUserRequest) error {
	validate := validator.New()
	return validate.Struct(request)
}

// validateCreateUserRequest validates the Create User request fields.
func validateCreateUserRequest(request input.CreateUserRequest) error {
	validate := validator.New()
	return validate.Struct(request)
}
