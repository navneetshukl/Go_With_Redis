package Redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func RedisConnection() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func InsertIntoRedis(key string, value interface{}) {
	client := RedisConnection()

	err := client.Set(key, value, 0).Err()

	if err != nil {
		fmt.Printf("Error in inserting ", err)
		return

	}
	fmt.Println("Inserting Successful")
}

func GetFromRedis(key string){

	client:=RedisConnection()

	ans,err:=client.Get(key).Result()

	if err!=nil{
		log.Fatal("Error in Fetching value from redis ",err)
		return 
	}
	fmt.Println("Value is ",ans)
}
