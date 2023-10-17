package handlers

import (
	"github.com/fur1ouswolf/crud/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewPerson godoc
// @Summary 	Create a person
// @Description Create a person
// @ID			create-person
// @Accept		json
// @Produce		json
// @Param		person 	body 		Person	true	"Person"
// @Success		200		{object}	Person
// @Failure		400		{object}	error
// @Router		/person [post]
func (h *Handler) NewPerson(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		h.Logger.Error("Failed to bind JSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.db.Create(&person)
	h.Logger.Info("Person created:", person)
	c.JSON(http.StatusOK, person)
}

// AllPersons godoc
// @Summary 	Get all persons
// @Description Get all persons
// @ID 			get-all-persons
// @Success 	200 	{object} 	Person
// @Produce  	json
// @Router 		/person [get]
func (h *Handler) AllPersons(c *gin.Context) {
	var persons []models.Person
	h.db.Find(&persons)
	c.JSON(http.StatusOK, persons)
}

// Person godoc
// @Summary 	Get a person
// @Description Get a person
// @ID			get-person
// @Produce 	json
// @Param		id 		path 		int		true	"Person ID"
// @Success		200		{object}	Person
// @Failure		400		{object}	error
// @Router		/person/{id} [get]
func (h *Handler) Person(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	h.db.First(&person, id)
	if person.ID == 0 {
		h.Logger.Error("Person not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

// UpdatePerson godoc
// @Summary 	Update a person
// @Description Update a person
// @ID			update-person
// @Accept		json
// @Produce		json
// @Param		id 		path 		int		true	"Person ID"
// @Param		person 	body 		Person	true	"Person"
// @Success		200		{object}	Person
// @Failure		400		{object}	error
// @Router		/person/{id} [put]
func (h *Handler) UpdatePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	h.db.Find(&person, id)
	if err := c.ShouldBindJSON(&person); err != nil {
		h.Logger.Error("Failed to bind JSON:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.db.Save(&person)
	h.Logger.Info("Person updated:", person)
	c.JSON(http.StatusOK, person)
}

// DeletePerson godoc
// @Summary 	Delete a person
// @Description Delete a person
// @ID			delete-person
// @Produce		json
// @Param		id 		path 		int		true	"Person ID"
// @Success		200		{object}	Person
// @Failure		400		{object}	error
// @Router		/person/{id} [delete]
func (h *Handler) DeletePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")
	h.db.First(&person, id)
	if person.ID == 0 {
		h.Logger.Error("Person not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found"})
		return
	}
	h.db.Delete(&person, id)
	h.Logger.Info("Person deleted:", person)
	c.JSON(http.StatusOK, person)
}
