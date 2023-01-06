package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tianmarillio/instant-note-backend/config"
	"github.com/tianmarillio/instant-note-backend/src/models"
)

func INoteCreate(c *gin.Context) {
	// get user from auth
	var reqUser, _ = c.Get("user")
	user := reqUser.(*models.User)

	// extract body
	var body struct {
		Title   string `json:"title"`
		Body    string `json:"body"`
		FileUrl string `json:"fileUrl"`
		Type    string `json:"type"`
	}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// create db row
	iNote := models.INote{
		ID:      uuid.NewString(),
		UserID:  user.ID,
		Title:   body.Title,
		Body:    body.Body,
		FileUrl: body.FileUrl,
		Type:    body.Type,
	}
	result := config.DB.Create(&iNote)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	// return json
	c.JSON(http.StatusCreated, gin.H{
		"id": iNote.ID,
	})
}

func INoteFindAll(c *gin.Context) {
	// get user from auth
	var reqUser, _ = c.Get("user")
	user := reqUser.(*models.User)

	// query to db
	var iNotes []*models.INote
	result := config.DB.
		Where("user_id = ?", user.ID).
		Find(&iNotes)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}

	// return json
	c.JSON(http.StatusAccepted, gin.H{
		"iNotes": iNotes,
	})
}

func INoteFindById(c *gin.Context) {
	// get user from auth
	var reqUser, _ = c.Get("user")
	user := reqUser.(*models.User)

	// get params from context
	id := c.Param("id")

	// query to db
	var iNote *models.INote
	result := config.DB.
		Where("user_id = ? AND id = ?", user.ID, id).
		Find(&iNote)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// return json
	c.JSON(http.StatusAccepted, gin.H{
		"iNote": iNote,
	})
}

func INoteUpdate(c *gin.Context) {

	// get user from auth
	var reqUser, _ = c.Get("user")
	user := reqUser.(*models.User)

	// get params from context
	id := c.Param("id")

	// extract body
	var body struct {
		Title   string `json:"title"`
		Body    string `json:"body"`
		FileUrl string `json:"fileUrl"`
		Type    string `json:"type"`
	}
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// update db
	iNote := models.INote{
		Title:   body.Title,
		Body:    body.Body,
		FileUrl: body.FileUrl,
		Type:    body.Type,
	}
	result := config.DB.
		Where("user_id = ? AND id = ?", user.ID, id).
		Updates(&iNote)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// return json
	c.JSON(http.StatusAccepted, gin.H{
		"id": id,
	})
}

func INoteDelete(c *gin.Context) {
	// get user from auth
	var reqUser, _ = c.Get("user")
	user := reqUser.(*models.User)

	// get params from context
	id := c.Param("id")

	// query to db
	var iNote *models.INote
	result := config.DB.
		Where("user_id = ? AND id = ?", user.ID, id).
		Delete(&iNote)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	// return json
	c.JSON(http.StatusAccepted, gin.H{
		"id": id,
	})
}
