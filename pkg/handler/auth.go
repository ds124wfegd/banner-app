package handler

import (
	"net/http"

	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/gin-gonic/gin"
)

const (
	userToken     = "userId"
	adminToken    = "adminId"
	tokenAdminCtx = "tokenAdmin"
)

// /////////////////////////////////////for users////////////////////////////////////////////////
func (h *Handler) signUp(c *gin.Context) {
	var input bannerapp.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// /////////////////////////////////////for admins////////////////////////////////////////////////

func (h *Handler) signUpAdmin(c *gin.Context) {
	var inputAdmin bannerapp.Admin

	if err := c.BindJSON(&inputAdmin); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateAdmin(inputAdmin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type signInInputAdmin struct {
	AdminUsername string `json:"adminUsername" binding:"required"`
	AdminPassword string `json:"adminPassword" binding:"required"`
}

func (h *Handler) signInAdmin(c *gin.Context) {

	var input signInInputAdmin

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokenAdmin, err := h.services.Authorization.GenerateTokenAdmin(input.AdminUsername, input.AdminPassword)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"tokenAdmin": tokenAdmin,
	})
	c.Set(tokenAdminCtx, tokenAdmin)
}
