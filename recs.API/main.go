package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type song struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Link   string `json:"link"`
	Season string `json:"season"`
}

var songs = []song{
	{ID: "1", Title: "Alcohol Free", Link: "https://www.youtube.com/watch?v=XA2YEHn-A8Q&ab_channel=JYPEntertainment", Season: "Summer"},
	{ID: "2", Title: "Running Up That Hill", Link: "https://www.youtube.com/watch?v=wp43OdtAAkM&ab_channel=KateBushMusic", Season: "Spring"},
}

func main() {
	router := gin.Default()

	router.GET("/songs", getSongs)
	router.GET("/songs/:id", getSongById)
	router.POST("/songs", postSongs)

	router.Run("localhost:3001")
}

// get songs as JSON
func getSongs(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.IndentedJSON(http.StatusOK, songs)
}

//post song
func postSongs(c *gin.Context) {
	var newSong song

	//call Bind JSON to bind received JSON to the new song
	if err := c.BindJSON(&newSong); err != nil {
		return
	}

	// add new song to slice
	songs = append(songs, newSong)
	c.IndentedJSON(http.StatusCreated, newSong)
}

func getSongById(c *gin.Context) {
	id := c.Param("id")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	for _, a := range songs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
}
