package handler

import (
	"net/http"
	"strconv"

	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createContent(c *gin.Context) {

	adminId, err := getAdminId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if adminId < 2000000000 {
		newErrorResponse(c, http.StatusForbidden, "the user does not have access")
		return
	} else {

		bannerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
			return
		}

		var input bannerapp.Content
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		id, err := h.services.BannerContent.CreateContent(bannerId, input)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})
	}
}

func (h *Handler) getAllContent(c *gin.Context) {
	adminId, err := getAdminId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if adminId < 2000000000 {
		newErrorResponse(c, http.StatusForbidden, "the user does not have access")
		return
	} else {

		bannerId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
			return
		}

		content, err := h.services.BannerContent.GetAllContent(bannerId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, content)
	}
}

func (h *Handler) deleteContentById(c *gin.Context) {
	adminId, err := getAdminId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if adminId < 2000000000 {
		newErrorResponse(c, http.StatusForbidden, "the user does not have access")
		return
	} else {

		contentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
			return
		}

		err = h.services.BannerContent.DeleteContent(contentId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, statusResponse{"ok"})
	}
}
