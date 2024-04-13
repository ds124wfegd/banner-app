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
		}
	}

	api_admin := router.Group("/api_admin", h.adminIdentity)
	{
		bannerAdmin := api_admin.Group("/bannerAdmin")
		{
			bannerAdmin.GET("/admin_banner/feature=:id", h.getBannersByFeature)
			bannerAdmin.GET("/admin_banner/tag=:id", h.getBannersByTag)
			bannerAdmin.POST("/", h.createBanner)
			bannerAdmin.POST("/createContent=:id", h.createContent)
			bannerAdmin.GET("/getAll/banner", h.getAllBanners)
			bannerAdmin.GET("/getAllContnetById=:id", h.getAllContent)
			bannerAdmin.DELETE("/deleteContnetById=:id", h.deleteContentById)
			bannerAdmin.GET("/updateBanner/:id", h.updateBanner)
			bannerAdmin.DELETE("/deleteByFeature/:id", h.deleteBannerByFeature)
		}
	}
	return router
}
