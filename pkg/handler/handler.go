package handler

import (
	"github.com/ds124wfegd/banner-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up-admin", h.signUpAdmin)
		auth.POST("/sign-in-admin", h.signInAdmin)

	}

	api := router.Group("/api", h.userIdentity)
	{
		banner := api.Group("/banner")
		{
			banner.GET("/user_banner/feature=:id", h.getBannersByFeature)
			banner.GET("/user_banner/tag=:id", h.getBannersByTag)
			banner.GET("/user_banner/banner", h.getAllBanners)
		}
	}

	api_admin := router.Group("/api_admin", h.adminIdentity)
	{
		bannerAdmin := api_admin.Group("/bannerAdmin")
		{
			bannerAdmin.POST("/", h.createBanner)
			bannerAdmin.GET("/admin_banner/feature=:id", h.getBannersByFeature)
			bannerAdmin.GET("/admin_banner/tag=:id", h.getBannersByTag)
			bannerAdmin.GET("/getAll/banner", h.getAllBanners)
			bannerAdmin.GET("/updateBanner/:id", h.updateBanner)
			bannerAdmin.DELETE("/deleteByFeature/:id", h.deleteBannerByFeature)

			bannerAdmin.POST("content/createContent=:id", h.createContent)
			bannerAdmin.GET("content/getAllContnetById=:id", h.getAllContent)
			bannerAdmin.DELETE("content/deleteContnetById=:id", h.deleteContentById)

		}

	}
	return router
}
