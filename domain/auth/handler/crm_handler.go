package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	Config "skegsTech/auth-service-go/config"
	"skegsTech/auth-service-go/domain/auth/entity"
	"skegsTech/auth-service-go/domain/auth/request"
	"skegsTech/auth-service-go/domain/auth/service"
	"skegsTech/auth-service-go/logger"
	"skegsTech/auth-service-go/myerror"
	"skegsTech/auth-service-go/util"
	"time"

	"github.com/golang-jwt/jwt"
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
	user, err := c.service.Create(ctx, payload)
	if err != nil {
		logger.Error(ctx, err)
		util.Errorf(w, http.StatusInternalServerError, nil, err)
		return
	}
	
	response := entityToResponse(user)

	util.Success(w, http.StatusOK, response, "")
}

func (c *crmAuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	//  init config
	config, err := Config.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	
	ctx := r.Context()

	// get payload
	payload := &request.CreateLoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		util.Error(w, http.StatusBadRequest, nil, "Invalid request")
		return
	}

	// validate
	if errors := util.ValidateRequest(payload); len(errors) > 0 {
		util.Error(w, http.StatusBadRequest, errors, "Validation error")
		return
	}

	// get user
	user, err := c.service.Get(ctx, payload.Email)
	if err != nil {
		if errors.Is(err, myerror.ErrRecordNotFound) {
			util.Error(w, http.StatusNotFound, nil, err.Error())
			return
		}
		
		logger.Error(ctx, err)
		util.Errorf(w, http.StatusInternalServerError, nil, err)
		return
	}

	// validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password));
	if err != nil {
		util.Error(w, http.StatusBadRequest, err, "Email or Password incorrect")
		return
	}

	claims := &entity.Claims{
		UserId: user.Id,
		UserName: user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    user.Name,
		},
	}
	// generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.AuthJwtSecretCrm))

	response := entityToResponse(user)
	response["token"] = signedToken

	util.Success(w, http.StatusOK, response, "Success Login")
}

func (c *crmAuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// get authenticated userId
	userId, _ := util.GetAuthenticatedUserID(ctx)
	fmt.Println(userId)

	// get user
	user, err := c.service.GetById(ctx, userId)
	if err != nil {
		if errors.Is(err, myerror.ErrRecordNotFound) {
			util.Error(w, http.StatusNotFound, nil, err.Error())
			return
		}
		
		logger.Error(ctx, err)
		util.Errorf(w, http.StatusInternalServerError, nil, err)
		return
	}

	response := entityToResponse(user)

	util.Success(w, http.StatusOK, response, "Success Get Account")
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
