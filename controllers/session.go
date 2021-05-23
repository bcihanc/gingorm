package controllers

import (
	"fmt"
	"gingorm/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateVariableInSession godoc
// @Summary create new variable in session
// @Description create new variable in session description
// @Param  data body models.SessionInput true "SessionInput"
// @Router /session [post]
func CreateVariableInSession(c *gin.Context) {
	session := sessions.Default(c)

	var input models.SessionInput

	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
		return
	}

	session.Set(input.Key, input.Value)
	if saveErr := session.Save(); saveErr != nil {
		fmt.Errorf(saveErr.Error())
	}
	c.JSON(http.StatusOK, gin.H{input.Key: input.Value})
}

// GetVariableFromSession godoc
// @Summary create new variable in session
// @Description create new variable in session description
// @Param key path string true "key is stored value"
// @Router /session/{key} [get]
// @Success 200 {object} models.SessionInput
func GetVariableFromSession(c *gin.Context) {
	var output models.SessionInput
	key := c.Param("key")
	session := sessions.Default(c)
	value := session.Get(key)

	output.Key = key
	output.Value = fmt.Sprintf("%v", value)

	c.JSON(http.StatusOK, output)
}
