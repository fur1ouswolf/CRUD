package handlers

import (
	"github.com/fur1ouswolf/crud/internal/models"
	"github.com/gin-gonic/gin"
)

// AllRegions godoc
// @Summary 	Get all regions
// @Description Get all regions
// @ID 			get-all-regions
// @Success 	200 	{object} 	Region
// @Produce  	json
// @Router 		/region [get]
func (h *Handler) AllRegions(c *gin.Context) {
	var regions []models.Region
	h.db.Find(&regions)
	c.JSON(200, regions)
}

// Residents godoc
// @Summary 	Get a residents
// @Description Get a residents
// @ID			get-region
// @Produce 	json
// @Param		id 		path 		int		true	"Region ID"
// @Success		200		{object}	Region
// @Router		person/region/{id} [get]
func (h *Handler) Residents(c *gin.Context) {
	var persons []models.Person
	id := c.Params.ByName("id")
	h.db.Where("region_id = ?", id).Find(&persons)
	c.JSON(200, persons)
}

// CreateRegion godoc
// @Summary		Create a region
// @Description Create a region
// @ID			create-region
// @Accept		json
// @Produce		json
// @Param		region 	body 		Region	true	"Region"
// @Success		200		{object}	Region
// @Router		/region [post]
func (h *Handler) CreateRegion(c *gin.Context) {
	var region models.Region
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.db.Create(&region)
	h.Logger.Info("Region created:", region)
	c.JSON(200, region)
}

// DeleteRegion godoc
//
// @Summary		Delete a region
// @Description Delete a region
// @ID 			delete-region
// @Produce  	json
// @Param 		id 		path 	int 	true 	"Region ID"
// @Success 	200 	{object} 	Region
// @Router 		/region/{id} [delete]
func (h *Handler) DeleteRegion(c *gin.Context) {
	var region models.Region
	id := c.Params.ByName("id")
	h.db.First(&region, id)
	if region.ID == 0 {
		c.JSON(400, gin.H{"error": "Region not found"})
		return
	}
	h.db.Delete(&region, id)
	h.Logger.Info("Region deleted:", region)
	c.JSON(200, region)
}
