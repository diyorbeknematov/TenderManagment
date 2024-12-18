package handler

import (
	"fmt"
	"net/http"
	"regexp"
	"tender/api/token"
	"tender/model"

	"github.com/gin-gonic/gin"
)

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// @Summary 	Registarion
// @Description Registration
// @Tags		Register And Login
// @Accept		json
// @Produce		json
// @Param 		user body model.UserRegisterReq true "User info"
// @Success		200 {object} model.UserRegisterResp
// @Failure 	400 {object} model.APIError
// @Failure 	404 {object} model.APIError
// @Failure 	500 {object} model.APIError
// @Router 		/register [post]
func (h *Handler) RegistrationHandler(ctx *gin.Context) {
	var user model.UserRegisterReq

	if err := ctx.ShouldBindJSON(&user); err != nil {
		h.Log.Error(fmt.Sprintf("Requestni structga o'qishda xatolik bor: %v", err))
		ctx.JSON(model.ErrInvalidInput.Code, model.ErrInvalidInput)
		return
	}

	if IsValidEmail(user.Email) {
		ctx.JSON(400, "email is incorrect format")
		return
	}

	exists, err := h.Service.IsUserExists(model.IsUserExists{
		Username: user.Username,
		Email:    user.Email,
	})

	if err != nil {
		h.Log.Error(fmt.Sprintf("Userni bor yoki yo'qligini tekshirishda xatolik: %v", err))
		ctx.JSON(model.ErrInternalServerError.Code, model.ErrInternalServerError)
		return
	}

	if exists {
		h.Log.Error("User oldin ro'yxatdan o'tgan")
		ctx.JSON(model.ErrEmailAlreadyExists.Code, model.ErrEmailAlreadyExists)
		return
	}

	resp, err := h.Service.Registration(user)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Userni bor yoki yo'qligini tekshirishda xatolik: %v", err))
		ctx.JSON(model.ErrInternalServerError.Code, model.ErrInternalServerError)
		return
	}

	token, err := token.GenerateToken(model.Token{
		ID:       resp.ID,
		Username: resp.Username,
		Role:     resp.Role,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error accesstoken generate qilishda xatolik: %v", err))
		ctx.JSON(model.ErrInternalServerError.Code, model.ErrInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// @Summary 	Login
// @Description Login
// @Tags		Register And Login
// @Accept		json
// @Produce		json
// @Param 		Info body model.LoginUser true "User info"
// @Success		200 {object} model.LoginResp
// @Failure 	400 {object} model.APIError
// @Failure 	404 {object} model.APIError
// @Failure 	500 {object} model.APIError
// @Router 		/login [post]
func (h *Handler) LoginHandler(ctx *gin.Context) {
	var login model.LoginUser

	if err := ctx.ShouldBindJSON(&login); err != nil {
		h.Log.Error(fmt.Sprintf("Requestni structga o'qishda xatolik bor: %v", err))
		ctx.JSON(model.ErrInvalidInput.Code, model.ErrInvalidInput)
		return
	}

	resp, err := h.Service.GetUserByUsername(login)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Userni username va password boyicha login qilishda xatolik: %v", err))
		ctx.JSON(model.ErrInternalServerError.Code, model.ErrInternalServerError)
		return
	}

	token, err := token.GenerateToken(model.Token{
		ID:       resp.ID,
		Username: resp.Username,
		Role:     resp.Role,
	})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error accesstoken generate qilishda xatolik: %v", err))
		ctx.JSON(model.ErrInternalServerError.Code, model.ErrInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
