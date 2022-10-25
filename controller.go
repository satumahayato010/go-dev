package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func QiitaGet(c *gin.Context) {
	resp, err := http.Get("http://qita")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var data []Item

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"item": data})
}
