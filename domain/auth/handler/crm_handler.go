package handler

import (
	"encoding/json"
	"net/http"

	"skegsTech/auth-service-go/domain/auth/entity"
	"skegsTech/auth-service-go/domain/auth/request"
	"skegsTech/auth-service-go/domain/auth/service"
	"skegsTech/auth-service-go/logger"
	"skegsTech/auth-service-go/util"

	"golang.org/x/crypto/bcrypt"
)

type crmAuthHandler struct {
	service service.AuthService
}

func NewCrmAuthHandler(sv service.AuthService) CrmAuthHandler {
	return &crmAuthHandler{
		service: sv,
	}
}

func (c *crmAuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get payload
	payload := &request.CreateAuthRequest{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		util.Error(w, http.StatusBadRequest, nil, "Invalid request")
		return
	}

	// validate
	if errors := util.ValidateRequest(payload); len(errors) > 0 {
		util.Error(w, http.StatusBadRequest, errors, "Validation error")
		return
	}

	// hash possword
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		logger.Error(ctx, err)
		util.Errorf(w, http.StatusBadRequest, nil, err)
		return
	}
	payload.Password = string(hashPassword)

	// create
	user, err := c.service.Register(ctx, payload)
	if err != nil {
		logger.Error(ctx, err)
		util.Errorf(w, http.StatusInternalServerError, nil, err)
		return
	}

	util.Success(w, http.StatusOK, entityToResponse(user), "")
}

func entityToResponse(user *entity.User) map[string]interface{} {
	return map[string]interface{}{
		"id":          	user.Id,
		"uniqueId":   	user.UniqueId,
		"name":        	user.Name,
		"email":       	user.Email,
		"createdAt":   	user.CreatedAt,
		"updatedAt":   	user.UpdatedAt,
	}
}
