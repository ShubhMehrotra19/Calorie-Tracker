package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Replace with your actual MongoDB client initialization logic
var entryCollection *mongo.Collection = openCollection(Client, "calories")

func AddEntry(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel() // Ensure context is always released

	var entries []bson.M
	cursor, err := entryCollection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(entries) // Debugging print
	c.JSON(http.StatusOK, entries)
}

func DeleteEntry(c *gin.Context) {
	entryID := c.Param("id")
	docID, err := primitive.ObjectIDFromHex(entryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entry ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	result, err := entryCollection.DeleteOne(ctx, bson.M{"_id": docID})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deletedCount": result.DeletedCount})
}

// Placeholder functions
func GetEntries(c *gin.Context)             {}
func GetEntriesByIngredient(c *gin.Context) {}
func GetEntryById(c *gin.Context)           {}
func UpdateEntry(c *gin.Context)            {}
func UpdateIngredient(c *gin.Context)       {}
