package bannerapp

type Admin struct {
	Id            int    `json:"id" db:"id"`
	AdminUsername string `json:"adminUsername" binding:"required"`
	AdminPassword string `json:"adminPassword" binding:"required"`
	AdminStatus   bool   `json:"adminStatus" binding:"required"`
	SystemPasword string `json:"systemPasword" binding:"required"`
}
