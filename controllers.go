package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
)

// NinjaController ...
type NinjaController struct {
	model *Ninja
	db    *gorp.DbMap
}

// NewNinjaController ...
func NewNinjaController(db *gorp.DbMap) *NinjaController {
	return &NinjaController{model: new(Ninja), db: db}
}

// All returns all competitors in the database.
func (ninja *NinjaController) All(c *gin.Context) {
	data, err := ninja.model.All(ninja.db)
	if err != nil {
		c.JSON(406, gin.H{"Message": "No results.", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"data": data})
}

/*
//One ...
func (ctrl ArticleController) One(c *gin.Context) {
	userID := getUserID(c)

	if userID == 0 {
		c.JSON(403, gin.H{"message": "Please login first"})
		c.Abort()
		return
	}

	id := c.Param("id")

	if id, err := strconv.ParseInt(id, 10, 64); err == nil {

		data, err := articleModel.One(userID, id)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}*/
