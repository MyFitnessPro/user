package main

import (
	"net/http"

	goFirebase "github.com/MyFitnessPro/firebase"
	_ "github.com/MyFitnessPro/user/docs"
	utils "github.com/MyFitnessPro/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get user
// @Description Get user by ID and role
// @Tags User
// @Accept  json
// @Produce  json
// @Param uid query string true "User ID"
// @Param role query string true "User Role"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "Invalid request parameters"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Failed to operate on user"
// @Router /get [get]
func HandleGetUserRequest(c *gin.Context, client *goFirebase.FirebaseClient) {
	uid := c.GetString("uid")
	role := c.GetString("role")

	userData, err := client.GetDocument(role, uid)
	if utils.HandleHTTPError(c, err, http.StatusInternalServerError, "Failed to operate on user") {
		return
	}

	c.JSON(http.StatusOK, userData)
	utils.HandleHTTPError(c, nil, http.StatusNotFound, "User not found")
}

// @Summary Delete user
// @Description Delete user by ID and role
// @Tags User
// @Accept  json
// @Produce  json
// @Param uid query string true "User ID"
// @Param role query string true "User Role"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {string} string "Invalid request parameters"
// @Failure 500 {string} string "Failed to operate on user"
// @Router /delete [delete]
func HandleDeleteUserRequest(c *gin.Context, client *goFirebase.FirebaseClient) {
	uid := c.GetString("uid")
	role := c.GetString("role")

	err := client.DeleteDocument(role, uid)
	if utils.HandleHTTPError(c, err, http.StatusInternalServerError, "Failed to operate on user") {
		return
	}

	c.JSON(http.StatusOK, "User deleted successfully")
	utils.HandleHTTPError(c, nil, http.StatusNotFound, "User not found")
}

// @Summary Upsert user
// @Description Upsert user by ID and role
// @Tags User
// @Accept  json
// @Produce  json
// @Param uid query string true "User ID"
// @Param role query string true "User Role"
// @Param userData body interface{} true "User data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "Invalid request parameters"
// @Failure 400 {string} string "Body not found in context"
// @Failure 500 {string} string "Body is of an incorrect type"
// @Failure 500 {string} string "Failed to operate on user"
// @Router /upsert [post]
func HandleUpsertUserRequest(c *gin.Context, client *goFirebase.FirebaseClient) {
	uid := c.GetString("uid")
	role := c.GetString("role")

	userDataInterface, exists := c.Get("userData")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body not found in context"})
		return
	}
	userData, ok := userDataInterface.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Body is of an incorrect type"})
		return
	}

	err := client.UpsertDocument(role, uid, userData)
	if utils.HandleHTTPError(c, err, http.StatusInternalServerError, "Failed to create user in Firestore") {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User upserted successfully"})
}
