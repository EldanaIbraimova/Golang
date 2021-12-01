package main

import (
	redis_cache "HW7/internal/cache/redis-cache"
	"HW7/internal/http"
	"HW7/internal/store/inmemory"
	"context"
	"log"
)

func main() {
	store := inmemory.Init()

	//srv := http.NewServer(context.Background(), ":8080", store)
	//if err := srv.Run(); err != nil {
	//	log.Println(err)
	//}

	cache := redis_cache.NewRedisCache(":8080", 1, 1800)

	srv := http.NewServer(context.Background(),
		http.WithAddress(":8080"),
		http.WithStore(store),
		http.WithCache(cache),
	)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
}