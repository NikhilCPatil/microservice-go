// routes/people.go
package routes

import (
	"github.com/gin-gonic/gin"
	"/db"
	"net/http"
)

// Person represents a person in the system
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreatePerson creates a new person
func CreatePerson(c *gin.Context) {
	var person Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection()
	_, err := collection.InsertOne(c.Request.Context(), person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Person created successfully"})
}

// GetPeople retrieves all people
func GetPeople(c *gin.Context) {
	collection := db.GetCollection()
	cur, err := collection.Find(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(c.Request.Context())

	var people []Person
	for cur.Next(c.Request.Context()) {
		var person Person
		err := cur.Decode(&person)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		people = append(people, person)
	}

	c.JSON(http.StatusOK, people)
}

// RegisterPeopleRoutes registers people-related routes
func RegisterPeopleRoutes(router *gin.Engine) {
	router.POST("/people", CreatePerson)
	router.GET("/people", GetPeople)
}
