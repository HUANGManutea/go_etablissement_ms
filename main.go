package main

import (
	"log"
	"reflect"
	"time"

	Search "go_etablissement_ms/search"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

// CacheMiddleware will add the cache to the context
func CacheMiddleware(cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("cache", cache)
		c.Next()
	}
}

func main() {
	router := gin.Default()

	// redis server
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	// cache
	ctxCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	router.Use(CacheMiddleware(ctxCache))

	router.GET("/etablissement/:id", getEtablissement)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getEtablissement(ctx *gin.Context) {
	id := ctx.Param("id")

	ctxCache, ok := ctx.MustGet("cache").(*cache.Cache)
	if !ok {
		log.Println("Error getting cache in context")
	}

	var line string

	err := ctxCache.Get(ctx, id, &line)

	if err != nil {
		log.Printf("Error getting the line in cache: line [%v] err [%v]\n", line, err)
	}

	// rien dans le cache
	if line == "" {
		log.Printf("COUCOU\n")
		searchline, err := Search.BinarySearchFile("./data.txt", id, 0, 7)
		log.Printf("line: [%v] type of line [%v]\n", line, reflect.TypeOf(line))
		if err != nil {
			log.Printf("Error getting the etablissement line: %v\n", err)
		}

		// assign
		line = searchline

		// on met dans le cache (dans cache local + redis)
		err = ctxCache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   id,
			TTL:   time.Hour,
			Value: line,
		})
		if err != nil {
			log.Printf("Error setting the cache: %v\n", err)
		}
	}

	ctx.JSON(200, gin.H{
		"message": line,
	})
}
