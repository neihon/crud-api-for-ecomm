package main

import (
	"github.com/gin-gonic/gin"
	"github.com/neihon/crud-api-with-authentication/sqliteDb"
	"github.com/neihon/crud-api-with-authentication/user_class"
	"log"
	"net/http"
	"strconv"
)

// routes
func createUser(c *gin.Context) {
	// Parse request, validate data, create user in database, return response
	var newUser user_class.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database := sqliteDb.CreateDatabase()
	createdUser := sqliteDb.CreateUser(database, newUser)

	c.JSON(http.StatusCreated, createdUser)
}

func getUser(c *gin.Context) {
	userIdParam := c.Param("Id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	database := sqliteDb.CreateDatabase()

	user, err := sqliteDb.GetUserById(database, uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
	// Fetch user by ID, update fields, save changes, return updated user
	userIdParam := c.Param("Id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	database := sqliteDb.CreateDatabase()

	user, err := sqliteDb.GetUserById(database, uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := database.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func deleteUser(c *gin.Context) {
	userIdParam := c.Param("Id")
	userId, err := strconv.ParseUint(userIdParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	database := sqliteDb.CreateDatabase()

	result := database.Delete(&user_class.User{}, uint(userId))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "user": userId})
}

func main() {
	database := sqliteDb.CreateDatabase()
	err := sqliteDb.DbMigrateModels(database)
	if err != nil {
		log.Fatal(err)
	}

	// routing
	router := gin.Default()
	router.POST("/user", createUser)
	router.GET("/users/:id", getUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
}
