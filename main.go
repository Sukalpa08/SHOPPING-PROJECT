package main

import (
	"shops/router"

	"github.com/gin-gonic/gin"
)

func main() {

	routers := gin.Default()
	routers = router.Setup()
	routers.Run("localhost:8080")
}

/*{
		addr := envstring("port", "8080")
	routers := gin.New()
	routers.Use(gin.Logger())
	router.Setup()
	router.Run("" + addr)
}
func envstring(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}*/
