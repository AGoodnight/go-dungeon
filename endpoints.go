package main

import (
	"dungeon/creatures"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCreatures(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, creatures.CreatureQuery)
}
