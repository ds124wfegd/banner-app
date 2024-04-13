package handler

import (
	"net/http"
	"strconv"

	bannerapp "github.com/ds124wfegd/banner-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBanner(c *gin.Context) {

	adminId, err := getAdminId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if adminId < 2000000000 {
		newErrorResponse(c, http.StatusForbidden, "the user does not have access")
		return
	} else {

		var input bannerapp.Banner
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		id, err := h.services.BannerList.Create(input)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})

	}
}

type getAllBannersResponse struct {
	Data []bannerapp.Banner `json:"data"`
}

func (h *Handler) getAllBanners(c *gin.Context) {

	adminId, err := getAdminId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, err := h.services.BannerList.GetAll(adminId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBannersResponse{
		Data: lists,
	})

}

func (h *Handler) getBannersByFeature(c *gin.Context) { // для пользователя

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parametr")
		return
	}

	lists, err := h.services.BannerList.GetByFeature(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBannersResponse{
		Data: lists,
	})

}

func (h *Handler) getBannersByTag(c *gin.Context) { // для пользователя

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parametr")
		return
	}

	lists, err := h.services.BannerList.GetByTag(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBannersResponse{
		Data: lists,
	})

}

func (h *Handler) updateBanner(c *gin.Context) { // для админа
	adminId, err := getAdminId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if adminId < 2000000000 {
		newErrorResponse(c, http.StatusForbidden, "the user does not have access")
		return
	} else {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id parametr")
			return
		}

		var input bannerapp.UpdateBanner
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := h.services.BannerList.UpdateBanner(id, input); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

	}
}

func (h *Handler) deleteBannerByFeature(c *gin.Context) { // для админа

	adminId, err := getAdminId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if adminId < 2000000000 {
		newErrorResponse(c, http.StatusForbidden, "the user does not have access")
		return
	} else {
		adminId, err := getAdminId(c)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		if adminId < 1000000000 {
			newErrorResponse(c, http.StatusForbidden, "the user does not have access")
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id parametr")
			return
		}

		err = h.services.BannerList.DeleteByFeature(id)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}
}
