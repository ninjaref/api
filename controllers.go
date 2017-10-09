package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DBController manages access to our database.
type DBController struct {
	db *gorm.DB
}

// NewDBController creates a controller associated with the given `gorm.DB`.
func NewDBController(db *gorm.DB) *DBController {
	return &DBController{db: db}
}

// Ninjas returns all competitors in the database.
func (dbc *DBController) Ninjas(c *gin.Context) {
	ninjas := []Ninja{}

	if err := dbc.db.Find(&ninjas).Error; err != nil {
		c.JSON(406, gin.H{"error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": ninjas})
}

// Ninja returns the competitor with the given ID.
func (dbc *DBController) Ninja(c *gin.Context) {
	ninja := Ninja{}

	if err := dbc.db.Find(&ninja, c.Param("id")).Error; err != nil {
		c.JSON(406, gin.H{"error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": ninja})
}

// Leaderboard returns the top-15 competitors by Ninja Rating.
//
// TODO: implement separate divisions for men, women, and overall.
func (dbc *DBController) Leaderboard(c *gin.Context) {
	summaries := []CareerSummary{}

	// Select the top-15 competitors ordered by to Ninja Rating:
	sel := "*, speed + consistency + success AS rating"
	ord := "rating desc"
	query := dbc.db.Preload("Ninja").Limit(15).Select(sel).Order(ord)

	if err := query.Find(&summaries).Error; err != nil {
		c.JSON(406, gin.H{"error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": summaries})
}
