package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ninjaref/api"
)

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	r.Use(sessions.Sessions("ninjaref-session", store))
	r.Use(CORSMiddleware())

	db, err := api.NewDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true)
	db.AutoMigrate(&api.Ninja{}, &api.CareerSummary{})

	v1 := r.Group("/v1")
	{
		controller := api.NewDBController(db)

		// Ninja endpoint:
		v1.GET("/ninjas", controller.Ninjas)
		v1.GET("/ninjas/:id", controller.Ninja)

		// Leaderboard endpoint:
		v1.GET("/leaderboard/:division", controller.Leaderboard)
	}

	r.LoadHTMLGlob("./views/*.html")
	r.Static("/views", "./views")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
