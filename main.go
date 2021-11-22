package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// // redis server
	// ring := redis.NewRing(&redis.RingOptions{
	// 	Addrs: map[string]string{
	// 		"server1": ":6379",
	// 	},
	// })

	// // cache
	// mycache := cache.New(&cache.Options{
	// 	Redis:      ring,
	// 	LocalCache: cache.NewTinyLFU(1000, time.Minute),
	// })
	router.GET("/etablissement/:id", getEtablissement)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEtablissement(ctx *gin.Context) {
	id := ctx.Param("id")

	BinarySearchFile("./data.txt", id, 0, 7)

	line, err := BinarySearchFile("../data.txt", "0000010", 0, 7)
	if err != nil {
		log.Printf("Error getting the etablissement line: %v\n", err)
	}
	ctx.JSON(200, gin.H{
		"message": line,
	})
}
